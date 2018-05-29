# hangmango-web-api
 Just for practice：A simple api server for hangman game based on golang

## GO Version

- go version >= 1.10.2

## DB Migrations

- need [goose](https://bitbucket.org/liamstask/goose)
  ```
  goose -env=$GOENV up
  ```
## Install dependency

- need [glide](https://github.com/Masterminds/glide)
  ```
  glide install
  ```

## Run Unit Test

- run all unit test
  ```
  GOENV=test go test -coverprofile coverage.out ./...
  ```

- get total coverage
  ```
  go tool cover -func=coverage.out
  ```
