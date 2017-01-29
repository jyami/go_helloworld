package main

import "fmt"

func mainSliceSample() {
	sample1()
	sample2()
	sample3()
	sample4()
}

func sample1() {
	num := [5]int{1, 2, 3, 4, 5}
	var slice1 []int

	slice1 = num[:]
	fmt.Println(slice1)

	slice2 := num[1:4]
	fmt.Println(slice2)

	slice3 := num[4:]
	fmt.Println(slice3)

	slice4 := num[:4]
	fmt.Println(slice4)

	fmt.Println("slice2=", slice2)
	fmt.Println("len=", len(slice2))
	fmt.Println("cap=", cap(slice2))

	slice5 := slice2[1:4]
	fmt.Println("slice5=", slice5)
	fmt.Println("len=", len(slice5))
	fmt.Println("cap=", cap(slice5))

	plusOne(num[:])
	fmt.Println(num)
}

func plusOne(vals []int) {
	for i := 0; i < len(vals); i++ {
		vals[i] += 1
	}
}

func sample2() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("s1=", s1)

	s2 := append(s1, 6, 7)
	fmt.Println("s2=", s2)

	s3 := append(s1, s2...)
	fmt.Println("s3=", s3)
}

func sample3() {
	src1 := []int{1, 2}
	dst := []int{97, 98, 99}

	count := copy(dst, src1)
	fmt.Println("copy count=", count)
	fmt.Println(dst)
	src2 := []int{3}
	count = copy(dst[2:], src2)
	fmt.Println("copy count=", count)
	fmt.Println(dst)
}

func sample4() {
	s1 := make([]string, 5, 10)
	fmt.Println("len=", len(s1))
	fmt.Println("cap=", cap(s1))
	s2 := make([]string, 5)
	fmt.Println("len=", len(s2))
	fmt.Println("cap=", cap(s2))
}
