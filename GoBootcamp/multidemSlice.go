/*
Given a list of names, you need to organize each name within a slice based on its length.

After you implement your solution, you should get the following output (slice of slice of strings):

[[] [] [Ava Mia] [Evan Neil Adam Matt Emma] [Emily Chloe]
[Martin Olivia Sophia Alexis] [Katrina Madison Abigail Addison Natalie]
[Isabella Samantha] [Elizabeth]]

*/

package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	maxL := 0
	for _, v := range names {
		if len(v) > maxL {
			maxL = len(v)
		}
	}
	res := make([][]string, maxL)
	for _, v := range names {
		res[len(v) - 1] = append(res[len(v) - 1], v)
	}
	fmt.Println(res)
}
