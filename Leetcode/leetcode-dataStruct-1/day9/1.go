package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pars := []rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			jugeCh := stack[len(stack)-1]
			if c != pars[jugeCh] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	// fmt.Println('(', ')')
	// fmt.Println('{', '}')
	// fmt.Println('[', ']')
	res := isValid("}{")
	fmt.Println(res)
}

// func pop(s string) string {

// }
