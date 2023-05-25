## API Keys authentication Prerequisite

- Step 1 : Set Temporary Environment Variable 

This step involves add temporary instance of apiKey to validate the authentication process. Either this can be added with export command or defined before starting the server.
           
    export api-key = your_apiKey

                OR 

    api-key=your_apiKey go run .

- Step 2 : Setup mkcertTool

mkcert is a simple tool for making locally-trusted development certificates. It requires no configuration. This will setup locally trusted TLS certificates

    brew install mkcert
    mkcert -install
