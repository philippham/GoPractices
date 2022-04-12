package main

import "fmt"

func main() {
	fmt.Println("Hello world !!")
	inputs := []string{
		"]}{]](]}{))}",
		"{}()[()]",
		"()",
		"([]({}{})){}()",
		"[){}(]}]}]))](())(",
		"({{)"}

	for _, input := range inputs {
		fmt.Println(input, " valid braces: ", validateBraces(input))
	}
}

func validateBraces(input string) bool {

	return true
}
