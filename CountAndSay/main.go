package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", countAndSay(5))
}

func countAndSay(n int) string {
	ns := []string{"0","1"}
    letter, letterP := "", ""
	for i := 1; i < n; i++ {

		count := 1
        res := ""
		for j, a := range ns[i] {
			letter = string(a)
            fmt.Println("j",j)
            fmt.Println("string(a):", string(a))
			if letterP == letter {
				count++
                fmt.Println("count:", count)
			} 
            if letterP != letter || j == len(ns[i])-1{
                fmt.Println("second:", res, strconv.Itoa(count), letter)
				res = res + strconv.Itoa(count) + letter
			}
            letterP = letter
            fmt.Println("boom")
		}
        fmt.Println("letterP:", letterP)
        fmt.Println("letter:", letter)
        fmt.Println("res:", res)
        ns = append(ns, res)
	}
    fmt.Println("ns",ns)
	return ns[n]
}

func countAndSayReversed(n int) string {
	ns := strconv.Itoa(n)
	ns = "13112221"
	cFlag := true
	count, letter, res := "", "", ""
	for _, a := range ns {
		if cFlag {
			count = string(a)
			cFlag = false
			fmt.Println(count)
		} else {
			letter = string(a)
			cFlag = true
			fmt.Println(letter)
			n, _ := strconv.Atoi(count)
			for i := 0; i < n; i++ {
				res = res + letter
			}
		}
	}
	return res
}
