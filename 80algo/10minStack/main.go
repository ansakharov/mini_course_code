package main

import (
	"errors"
	"fmt"
)

/*Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

// -8
Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.push(-3);

minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2*/

// 1 Точно ли понял задачу?
// 2 Какова сложность алгоритма? Проговори это вслух
// 3 Какие расходы по памяти?
// 4 не уходи в сторону
// 5 Имеет ли смысл реализовавывать мое решение?
// 6 !!! Проверка готового решения

type Stack struct {
	storage    []int
	minStorage []int
}

func (s *Stack) Push(in int) {
	s.storage = append(s.storage, in)

	if len(s.minStorage) == 0 || (in <= s.minStorage[len(s.minStorage)-1]) {
		s.minStorage = append(s.minStorage, in)
	}
}

func (s *Stack) Top() (int, error) {
	if len(s.storage) == 0 {
		return 0, errors.New("empty stack")
	}

	return s.storage[len(s.storage)-1], nil
}

func (s *Stack) Pop() {
	if len(s.storage) == 0 {
		return
	}

	lastValue := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]

	if lastValue == s.minStorage[len(s.minStorage)-1] {
		s.minStorage = s.minStorage[:len(s.minStorage)-1]
	}
}

func (s *Stack) GetMin() (int, error) {
	if len(s.minStorage) == 0 {
		return 0, errors.New("empty stack")
	}

	return s.minStorage[len(s.minStorage)-1], nil
}

func main() {
	minStack := Stack{}
	minStack.Push(1)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2*/
}
