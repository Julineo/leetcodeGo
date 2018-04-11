package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(fizzBuzz(15))
}

func fizzBuzz(n int) []string {
    res := []string{}
    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%5 == 0 {
            res = append(res, "FizzBuzz")
            continue
        }
        if i%3 == 0 {
            res = append(res, "Fizz")
            continue
        }
        if i%5 == 0 {
            res = append(res, "Buzz")
            continue
        }
        res = append(res, strconv.Itoa(i))
    }
    return res
}
