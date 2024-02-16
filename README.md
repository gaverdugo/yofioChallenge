# Yofio Technical Challenge

## Install the dependencies
```
go get .
```
## Start the server
```
go run cmd/yofioChallenge/main.go
```
The server starts at localhost:8080
## Usage
Make a POST call to `/credit-assignment`
The body must be a JSON with the following structure
```
{
    "investment": 6700
}
```
The server will return a response with how many of each credits can be assigned to that investment.
```
{
    "credit_type_300": 2,
    "credit_type_500": 1,
    "credit_type_700": 8
}
```