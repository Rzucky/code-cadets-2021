package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/01_interfaces/stacklibfer"
	"code-cadets-2021/lecture_2/01_interfaces/stacklibfoi"
)

func main() {
	s := stacklibfer.New()

	pushPopPrint(s, []int{1, 2, 3, 4})

	s2 := stacklibfoi.New()

	pushPopPrint(s2, []int{1, 2, 3, 4})
}

func pushPopPrint(stack Stack, numbers []int) {
	for _, num := range numbers {
		stack.Push(num)
	}

	for {
		num, ok := stack.Pop()
		if !ok {
			fmt.Println("Done")
			break
		}

		fmt.Println("Removed from stack:", num)
	}
}

type Stack interface {
	Push(a int)
	Pop() (int, bool)
}
