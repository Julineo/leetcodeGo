/*
http://www.golangbootcamp.com/book/control_flow

You have 50 bitcoins to distribute to 10 users: Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth The coins will be distributed based on the vowels contained in each name where:

a: 1 coin e: 1 coin i: 2 coins o: 3 coins u: 4 coins

and a user can’t get more than 10 coins. Print a map with each user’s name and the amount of coins distributed. After distributing all the coins, you should have 2 coins left.

The output should look something like that:

map[Matthew:2 Peter:2 Giana:4 Adriano:7 Elizabeth:5 Sarah:2 Augustus:10 Heidi:5 Emilie:6 Aaron:5]
Coins left: 2

Note that Go doesn’t keep the order of the keys in a map, so your results might not look exactly the same but the key/value mapping should be the same.
*/

package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	dist = make(map[string]int, len(users))
)

func main() {
	for _, name := range users {
		for _, l := range name {
			switch l {
			case 'a', 'e', 'A', 'E':
				dist[name] += 1
			case 'i', 'I':
				dist[name] += 2
			case 'o', 'O':
				dist[name] += 3
			case 'u', 'U':
				dist[name] += 4
			}
			if dist[name] > 10 {
				dist[name] = 10
			}
		}
	}
	for _, v := range dist {
		coins -= v
	}

	fmt.Println(dist)
	fmt.Println("Coins left:", coins)
}
