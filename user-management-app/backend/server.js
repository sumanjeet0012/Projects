import express from 'express';
import mysql from 'mysql2';
import cors from 'cors';
import dotenv from 'dotenv';

dotenv.config();

const app = express();

// Middleware
app.use(cors());
app.use(express.json());

// Create MySQL connection
const db = mysql.createConnection({
  host: 'localhost',
  user: 'root',
  password: 'your_password',
  database: 'user_management'
});

// Connect to database
db.connect((err) => {
  if (err) {
    console.error('Error connecting to MySQL:', err);
    return;
  }
  console.log('Connected to MySQL database');
  
  // Create users table if it doesn't exist
  const createTableQuery = `
    CREATE TABLE IF NOT EXISTS users (
      id INT AUTO_INCREMENT PRIMARY KEY,
      firstName VARCHAR(50) NOT NULL,
      lastName VARCHAR(50) NOT NULL,
      email VARCHAR(100) UNIQUE NOT NULL,
      phone VARCHAR(15),
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
  `;
  
  db.query(createTableQuery, (err) => {
    if (err) {
      console.error('Error creating table:', err);
    } else {
      console.log('Users table ready');
    }
  });
});

// Routes
app.post('/api/users', (req, res) => {
  const { firstName, lastName, email, phone } = req.body;
  const query = 'INSERT INTO users (firstName, lastName, email, phone) VALUES (?, ?, ?, ?)';
  
  db.query(query, [firstName, lastName, email, phone], (err, result) => {
    if (err) {
      res.status(500).json({ error: err.message });
      return;
    }
    res.status(201).json({ id: result.insertId, ...req.body });
  });
});

app.get('/api/users', (req, res) => {
  const query = 'SELECT * FROM users';
  
  db.query(query, (err, results) => {
    if (err) {
      res.status(500).json({ error: err.message });
      return;
    }
    res.json(results);
  });
});

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});
