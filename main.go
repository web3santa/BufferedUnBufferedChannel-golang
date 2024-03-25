package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// func (s *Student) setAge(age int) {
// 	s.age = age
// }

func (s *Student) getAverageGrade() float32 {
	sum := 0

	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	sum := 0

	for _, v := range s.grades {
		if sum < v {
			sum = v
		}
	}

	return sum
}

func main() {
	st := Student{"Time", []int{70, 50, 20, 15}, 18}
	s2 := Student{"Joe", []int{70, 50, 20, 15, 99, 43, 34}, 18}
	max1 := st.getMaxGrade()
	max2 := s2.getMaxGrade()
	fmt.Println(max1)
	fmt.Println(max2)
	fmt.Println(st.grades)
	st.grades = append(st.grades, 40, 99, 22)
	fmt.Println(st.grades)

	st.grades = append(st.grades, 234234, 123123, 123123)
	fmt.Println(st.grades)

	indexToRemove := 1
	st.grades = append(st.grades[:indexToRemove], st.grades[indexToRemove+1:]...)
	fmt.Println(st.grades)

}
