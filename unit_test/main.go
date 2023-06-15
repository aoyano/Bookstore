package main

import (
	"strings"
)

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func Min(numbers []int) int {
	var min int

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func GetName() string {
	return "Ayan"
}

func UbdateName(name string) string {
	return name
}

func login(email string) string {
	return email
}

func CalculateDiscount(price, discount float64) float64 {
	return price * discount
}

func EmptyField(f string) bool {
	return len(strings.TrimSpace(f)) == 0
}

func TwoPlusOne(a, b, c, min int) int {
	return (a + b + c) - min
}

func getAccountsByOrder(o string) string {
	return o
}

func Volume(a, b, c int) int {
	return a * b * c
}
