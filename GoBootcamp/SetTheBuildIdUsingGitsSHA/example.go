// http://www.golangbootcamp.com/book/tricks_and_tips#sec-git_sha_in_go_binary

package main

import "fmt"

// compile passing -ldflags "-X main.Build <build sha1>"
var Build string

func main() {
	fmt.Printf("Using build: %s\n", Build)
}
