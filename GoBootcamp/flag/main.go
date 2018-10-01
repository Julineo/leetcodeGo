/*
$ go run -ldflags "-X main.xyz=abc" main.go
*/

package main

import "fmt"

var xyz string

func main() {
    fmt.Println(xyz)
}
