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
    mkcert -install





