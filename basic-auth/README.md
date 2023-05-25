## Basic authentication Prerequisite

- Step 1 : Set Temporary Environment Variable 

This step involves add temporary instance of username and password to validate the authentication process. Either this can be added with export command or defined before starting the server.
           
    export username = your_username
    export password = your_password

                OR 

    username=your_username password=your_password go run .

- Step 2 : Setup mkcertTool

mkcert is a simple tool for making locally-trusted development certificates. It requires no configuration. This will setup locally trusted TLS certificates

    brew install mkcert
    mkcert localhost(creates a locally trusted certificate and key)

## cURL To test

- curl -i -u your_username:your_password https://localhost:4000/unauthenticated 

- curl -i -u your_username:your_password https://localhost:4000/authenticated 

- curl -i https://localhost:4000/authenticated  

- curl -i https://localhost:4000/unauthenticated  
