const express = require('express');
const dotenv = require('dotenv');
const cors = require('cors');
const bodyParser = require('body-parser');

dotenv.config({ path: './login.env' });
const config = {
  username: process.env.USERNAME,
  password: process.env.PASSWORD,
  port: process.env.PORT
};

const app = express();
const port = 5000;

app.use(cors());
app.use(bodyParser.json());


app.post('/api/login', (req, res) => {
    const { username, password } = req.body;

    if (username === config.username && password === config.password) {
        res.json({ message: 'Login successful!' });
    } else {
        res.status(401).json({ message: 'Invalid username or password.' });
    }
});

app.listen(port, () => {
    console.log(`Server is running on http://localhost:${config.port}`);
});
