# risk-api

This project provides 3 apis for risk management:

| Method | Endpoint      | Description         |
| ------ | ------------- | ------------------- |
| GET    | /v1/risk/{id} | get a risk by id    |
| GET    | /v1/risks     | get a list of risks |
| POST   | /v1/risks     | create a risk       |

## Content

- [Getting started](#getting-stared)
- [Invoke apis](#invoke-apis)
- [Run auto tests](#run-auto-tests)

## Getting stared

1. Clone code from this repository `git clone https://github.com/dongdongjssy/risk-api.git`
2. Make sure you have latest [go](https://go.dev/) installed
3. In root directory, run `go run .`

#### Optional: build an executable file

you can build the project by running `go build .` under root directory which will generate an executable program based on your operating system in the same folder.

## Invoke apis

While the server is running, the api docs can be found here: [Swagger docs](http://localhost:8080/swagger/index.html#/), it contains details of each api definition including endpoint url, parameters, and responses etc.

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

## Run auto tests

To run tests, in root folder run command:

```sh
go test ./... -cover
```

It will recursively find all files end with `_test.go` and gives a report of coverage.

The project contains one test files:

- `risk_handlers_test.go` integration tests for each apis, including both success and failure scenarios.

Following is an example of report after running test:

```
?       github.com/dongdongjssy/risk-api/constants      [no test files]
        github.com/dongdongjssy/risk-api/docs           coverage: 0.0% of statements
        github.com/dongdongjssy/risk-api/models         coverage: 0.0% of statements
        github.com/dongdongjssy/risk-api                coverage: 0.0% of statements
        github.com/dongdongjssy/risk-api/routes         coverage: 0.0% of statements
ok      github.com/dongdongjssy/risk-api/handlers       0.456s  coverage: 83.3% of statements
```
