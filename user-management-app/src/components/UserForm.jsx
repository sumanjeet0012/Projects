import { useState } from 'react';
import { Box, Button, TextField, Typography, Paper } from '@mui/material';
import axios from 'axios';

const UserForm = () => {
  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    email: '',
    phone: ''
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:5000/api/users', formData);
      setFormData({
        firstName: '',
        lastName: '',
        email: '',
        phone: ''
      });
      alert('User added successfully!');
    } catch (error) {
      console.error('Error adding user:', error);
      alert('Error adding user. Please try again.');
    }
  };

  return (
    <Paper elevation={3} sx={{ p: 3, mb: 4 }}>
      <Typography variant="h5" component="h2" gutterBottom>
        Add New User
      </Typography>
      <Box component="form" onSubmit={handleSubmit} sx={{ '& .MuiTextField-root': { mb: 2 } }}>
        <TextField
          fullWidth
          label="First Name"
          name="firstName"
          value={formData.firstName}
          onChange={handleChange}
          required
        />
        <TextField
          fullWidth
          label="Last Name"
          name="lastName"
          value={formData.lastName}
          onChange={handleChange}
          required
        />
        <TextField
          fullWidth
          label="Email"
          name="email"
          type="email"
          value={formData.email}
          onChange={handleChange}
          required
        />
        <TextField
          fullWidth
          label="Phone"
          name="phone"
          value={formData.phone}
          onChange={handleChange}
        />
        <Button
          type="submit"
          variant="contained"
          color="primary"
          fullWidth
          sx={{ mt: 2 }}
        >
          Add User
        </Button>
      </Box>
    </Paper>
  );
};

export default UserForm;
