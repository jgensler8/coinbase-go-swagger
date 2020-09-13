# coinbase-go-swagger

Another go implementation of Coinbase's Pro API

The API is defined using the Swagger/OpenAPI spec

## How to Update This Library

1. Copy existing `swagger.yaml` and edit in https://editor.swagger.io
2. Generate Go client
3. Extract code to `./swagger`
4. Add shim in respective `client` file to take care of Auth headers