FROM node:18-alpine

WORKDIR /app

# Copy only frontend package files
COPY package*.json ./

RUN npm install

# Copy only frontend source files, excluding backend
COPY src/ ./src/
COPY public/ ./public/
COPY index.html ./
COPY vite.config.js ./
COPY .env ./

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host"]
