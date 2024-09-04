package main

import "fmt"

//func main() {
//	s := []int{1, 1, 1}
//	f(s)
//	fmt.Println(s)
//}
//func f(s []int) {
//	// i只是一个副本，不能改变s中元素的值
//	/*for _, i := range s {
//		i++
//	}
//	*/
//
//	for i := range s {
//		s[i] += 1
//	}
//}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}
