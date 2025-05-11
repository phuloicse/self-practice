#! /bin/bash



curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{"username":"loiluong", "password":"1234"}'


curl --location 'http://localhost:8080/protected/hello' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJsb2lsdW9uZyIsImV4cCI6MTc0Njk1Nzg2OH0.2EdPOAGZU8hFWS56RiW6H_sm8ADPmM4IDhCCGsB9tSc'

