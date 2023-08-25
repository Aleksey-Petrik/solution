package main

import "strings"

func isOpen(bracket string) bool {
	return strings.Contains("{([", bracket)
}

func isEquals(bracketFirst, bracketSecond string) bool {
	if bracketFirst == "{" && bracketSecond == "}" {
		return true
	} else if bracketFirst == "(" && bracketSecond == ")" {
		return true
	} else if bracketFirst == "[" && bracketSecond == "]" {
		return true
	}
	return false
}

func isBalanced(brackets []string) bool {
	queue := NewQueue()
	for _, bracket := range brackets {
		if isOpen(bracket) {
			queue.Add(bracket)
		} else {
			if queue.Size() == 0 || !isEquals(queue.Remove(), bracket) {
				return false
			}
		}
	}
	if queue.Size() > 0 {
		return false
	}
	return true
}
