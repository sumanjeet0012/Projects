/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/juju/errors"
	localsjgroupv1alpha1 "github.com/projects/configmapsync/api/v1alpha1"
)

// ConfigmapsyncReconciler reconciles a Configmapsync object
type ConfigmapsyncReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	log    logr.Logger
}

// +kubebuilder:rbac:groups=localsjgroup.sumanjeet.com,resources=configmapsyncs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=localsjgroup.sumanjeet.com,resources=configmapsyncs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=localsjgroup.sumanjeet.com,resources=configmapsyncs/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Configmapsync object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.1/pkg/reconcile
func (r *ConfigmapsyncReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the configmapsync instance
	configMapSync := &localsjgroupv1alpha1.Configmapsync{}
	if err := r.Get(ctx, req.NamespacedName, configMapSync); err != nil {
		logger.Error(err, "unable to fetch configMapSync")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Fetch the Source ConfigMap
	sourceConfigmap := &corev1.ConfigMap{}
	sourceConfigmapName := types.NamespacedName{
		Namespace: configMapSync.Spec.SourceNamespace,
		Name:      configMapSync.Spec.ConfigmapName,
	}
	if err := r.Get(ctx, sourceConfigmapName, sourceConfigmap); err != nil {
		logger.Error(err, "unable to fetch source ConfigMap")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Create or Update the destination ConfigMap in the target namespace
	destinationConfigmap := &corev1.ConfigMap{}
	destinationConfigmapName := types.NamespacedName{
		Namespace: configMapSync.Spec.DestinationNamespace,
		Name:      configMapSync.Spec.ConfigmapName,
	}
	if err := r.Get(ctx, destinationConfigmapName, destinationConfigmap); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Creating ConfigMap in destination namespace", "Namespace", configMapSync.Spec.DestinationNamespace)
			destinationConfigmap = &corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      configMapSync.Spec.ConfigmapName,
					Namespace: configMapSync.Spec.DestinationNamespace,
				},
				Data: sourceConfigmap.Data,
			}

			if err := r.Create(ctx, destinationConfigmap); err != nil {
				return ctrl.Result{}, err
			}

		} else {
			return ctrl.Result{}, err
		}
	} else {
		logger.Info("Updating ConfigMap in destination namespace", "Namespace", configMapSync.Spec.DestinationNamespace)
		destinationConfigmap.Data = sourceConfigmap.Data // Update data from source to destination
		if err := r.Update(ctx, destinationConfigmap); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConfigmapsyncReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&localsjgroupv1alpha1.Configmapsync{}).
		Owns(&corev1.ConfigMap{}).
		Named("configmapsync").
		Complete(r)
}
