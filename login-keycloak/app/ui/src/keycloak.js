// src/KeycloakService.js
import Keycloak from 'keycloak-js';

// const config = {
//   url: process.env.REACT_APP_KEYCLOAK_URL,  
//   realm: process.env.REACT_APP_KEYCLOAK_REALM,  
//   clientId: process.env.REACT_APP_KEYCLOAK_CLIENT_ID, 
// };
const config = {
  // url: "http://localhost:8080/auth",  
  url: "http://localhost:8080/", 
  realm: "login",  
  clientId: "login-page", 
};


const keycloak = new Keycloak({
    url: config.url,
    realm: config.realm,
    clientId: config.clientId,
});

const KeycloakService = {
    init: () => {
        return new Promise((resolve, reject) => {
            keycloak.init({ onLoad: 'login-required' }).then((authenticated) => {
                if (authenticated) {
                    resolve();
                } else {
                    reject('Keycloak authentication failed');
                }
            }).catch((error) => {
                reject(error);
            });
        });
    },

    login: () => {
        keycloak.login();
    },

    logout: () => {
        keycloak.logout();
    },

    getToken: () => {
        return keycloak.token;
    },

    isAuthenticated: () => {
        return keycloak.authenticated;
    },

    getKeycloak: () => keycloak,
};

export default KeycloakService;
