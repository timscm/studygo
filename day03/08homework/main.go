package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：
Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth.
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d. 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	var total int
	for _, user := range users {
		if strings.Contains(user, "e") || strings.Contains(user, "E") {
			distribution[user] += 1
			total += 1
		}
		if strings.Contains(user, "i") || strings.Contains(user, "I") {
			distribution[user] += 2
			total += 2
		}
		if strings.Contains(user, "o") || strings.Contains(user, "O") {

			distribution[user] += 3
			total += 3
		}
		if strings.Contains(user, "u") || strings.Contains(user, "U") {
			distribution[user] += 4
			total += 4
		}
	}
	left = coins - total
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println(distribution)
	fmt.Println("剩下: ", left)
}
