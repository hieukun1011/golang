package main

import (
	"fmt"
	"strings"
)

func task1() {
	str_input := "This is a sample text. This text is just an example"
	str := strings.ReplaceAll(str_input, ".", "")
	list_input := strings.Split(str, " ")
	dict_input := make(map[string]int)
	for _, rec := range list_input {
		dict_input[rec] = dict_input[rec] + 1
	}
	fmt.Println(dict_input)
}

func checkPrimeNumber(n int) int {
	flag := 1
	if n < 2 {
		flag = 0
		return flag
	}
	for p := 2; p < n; p++ {
		if n%p == 0 {
			flag = 0
			break
		}
	}
	return flag
}

func task2() {
	var n int
	fmt.Print(">> nhap mot so n: ")
	fmt.Scanln(&n)

	str_output := ""
	for i := 0; i < n; i++ {
		check := checkPrimeNumber(i)
		if check == 1 {
			str_output += fmt.Sprintf("%d ", i)
		}
	}
	fmt.Println(str_output)
}