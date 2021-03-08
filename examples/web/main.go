package main

import "fmt"

type StatusCode int

const (
	StatusBadRequest StatusCode = iota + 400
	StatusUnauthorized
	StatusPaymentRequired
	StatusForbidden
)

const (
	StatusInternal StatusCode = iota + 500
	StatusNotImplemented
	StatusBadGateway
)

func main() {
	fmt.Println(StatusBadRequest.Error())
}
