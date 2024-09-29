# risk-api

This project provides 3 apis for risk management:

- Get a list of risks
- Get a risk by its uuid
- Create a risk

## Getting stared

1. Clone code from this repository `git clone https://github.com/dongdongjssy/risk-api.git`
2. Make sure you have latest [go](https://go.dev/) installed
3. In root directory, run `go run .`

#### Optional: build an executable file

you can build the project by running `go build .` under root directory which will generate an executable program based on your operating system in the same folder.

### Call apis

While the server is running, the api docs can be found here: `http://localhost:8080/swagger/index.html#/`, it contains details of each api definition including endpoint url, parameters, and responses etc.

Options to call apis:

1. There is a `http_client.http` file in root folder, you can send requests inside this file when you have [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension installed in vs code.
2. In [Swagger UI](http://localhost:8080/swagger/index.html#/)
3. Use curl from command line, e.g.

- get risks:
  ```sh
  curl localhost:8080/v1/risks
  ```
- get a risk:
  ```sh
  curl localhost:8080/v1/risk/850cfac4-c958-4e26-9828-c45ca4126d99
  ```
- create a risk:
  ```sh
  curl --json '{"state":"open","title":"some risk","description":"some risk"}' localhost:8080/v1/risks
  ```

4. Other api test tools like postman etc.

### Run auto tests

To run tests, in root folder run command:

```sh
go test ./... -cover
```

It will recursively find all files end with `_test.go` and gives a report of coverage.

The project contains two test files:

- `validate_state_test.go` unit tests for utils.
- `risk_handlers_test.go` integration tests for each apis, including both success and failure scenarios.
