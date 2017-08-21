package main

import "fmt"

func main() {
	var fibNumList = []int{1, 2}
	i := 1
	//generating fibonacci numbers list
	for {
		if (fibNumList[i] + fibNumList[i-1]) <= 4000000 {
			fibNumList =
				append(fibNumList, fibNumList[i]+fibNumList[i-1])
			i++
		} else {
			break
		}
	}
	fmt.Println(fibNumList)
	var sum int
	//summing up even numbers within the above list
	for _, v := range fibNumList {
		if v%2 == 0 {
			sum += v
		}
	}
	fmt.Println("The sum of Fibonacci even terms, "+
		"whose individual values <= 4 million: ", sum)
}

/*
web page for reference:
	https://projecteuler.net/problem=2

Problem statement:
	Each new term in the Fibonacci sequence is generated by adding
	the previous two terms. By starting with 1 and 2, the first 10 terms
	will be:
		1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
	By considering the terms in the Fibonacci sequence whose values
	do not exceed four million, find the sum of the even-valued terms.
*/
