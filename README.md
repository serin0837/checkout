# Checkout-kata

This repo is to solve checkout kata.

## Run locally 
`go run main.go` : just to do get one checkout result from kata.

## Testing
`go test -v ./...` : run all the test verbose

`go test -v -coverprofile cover.out ./...`: run all the tests with coverage

`go tool cover -html cover.out -o cover.html`: get html file of test coverage