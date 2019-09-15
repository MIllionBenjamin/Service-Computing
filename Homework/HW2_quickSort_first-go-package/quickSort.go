package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(num []int) {
	if len(num) < 2 {
		return
	}
	compare, i, front, back := num[0], 1, 0, len(num)-1
	for front < back {
		if num[i] > compare {
			num[i], num[back] = num[back], num[i]
			back--
			continue
		}
		num[i], num[front] = num[front], num[i]
		i++
		front++
	}
	num[front] = compare
	quickSort(num[:front])
	quickSort(num[front+1:])
}

func main() {
	myRand := rand.New(rand.NewSource(time.Now().Unix()))
	var randomArray [10]int
	for i := 0; i < 10; i++ {
		randomArray[i] = myRand.Intn(100)
	}
	fmt.Println(randomArray)
	quickSort(randomArray[0:10])
	fmt.Println(randomArray)
}
