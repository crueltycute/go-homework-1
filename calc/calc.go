package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)
func calculate(scanner *bufio.Scanner) (int, error) {
	stack := New()

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())

		if err == nil {  // is int
			stack.Push(number)
		} else {
			switch scanner.Text()[0] {
			case  ' ':
			case '\n':
			case '=':
				result, err := stack.Pop()

				if err != nil {
					return 0, err
				}

				return result, nil
			case '+':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(first + second)
			case '-':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(-first + second)
			case '*':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(first * second)
			case '/':
				first, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				second, err := stack.Pop()
				if err != nil {
					return 0, err
				}

				stack.Push(second / first)
			default:
				return 0, errors.New("undefined symbol")
			}
		}
	}

	result, err := stack.Pop()

	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	result, err := calculate(input)

	if err != nil {
		log.Println("Error")
	} else {
		fmt.Println("Result:", result)
	}
}