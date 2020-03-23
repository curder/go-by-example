package main

import "fmt"

/*
有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distributions = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {
	for _, user := range users {
		for _, char := range user { // 获取名字单个字符
			switch char {
			case 'e', 'E': // 通过单引号比较字符是否相等
				distributions[user] += 1
				coins -= 1
			case 'i', 'I':
				distributions[user] += 2
				coins -= 2
			case 'o', 'O':
				distributions[user] += 3
				coins -= 3
			case 'u', 'U':
				distributions[user] += 4
				coins -= 4
			}
		}
	}
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	fmt.Println("最终剩下", left, "个金币")
	for name, number := range distributions {
		fmt.Printf("用户%s，获得%d个金币 \n", name, number)
	}
}
