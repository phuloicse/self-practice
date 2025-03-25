import React, { useState, useEffect } from 'react';
import axios from 'axios';
import KeycloakService from './keycloak';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [authenticated, setAuthenticated] = useState(false);
  const [keycloak, setKeycloak] = useState(null);

  useEffect(() => {
    KeycloakService.init()
      .then(() => {
        setKeycloak(KeycloakService.getKeycloak());
        setAuthenticated(true);
      })
      .catch((error) => {
        console.error('Keycloak initialization failed', error);
        setAuthenticated(false);
      });
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await axios.post('http://localhost:5000/api/login', {
        username,
        password,
      });

      setMessage(response.data.message);
    } catch (error) {
      setMessage(error.response ? error.response.data.message : 'Something went wrong!');
    }
  };

  const handleKeycloakLogin = () => {
    KeycloakService.login();
  };

  const handleKeycloakLogout = () => {
    KeycloakService.logout();
    setAuthenticated(false);
  };

  return (
    <div>
      <h2>Login</h2>

      {!authenticated ? (
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="username">Username:</label>
            <input
              type="text"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit">Login</button>
        </form>
      ) : (
        <div>
          <h3>Welcome, {keycloak?.tokenParsed?.preferred_username}!</h3>
          <button onClick={handleKeycloakLogout}>Logout from Keycloak</button>
        </div>
      )}

      <p>{message}</p>

      {!authenticated && (
        <div>
          <h2>Login Page with Keycloak</h2>
          <button onClick={handleKeycloakLogin}>Login with Keycloak</button>
        </div>
      )}
    </div>
  );
}

export default App;

