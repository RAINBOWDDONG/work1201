package main

import "fmt"

//구조체 두번째 실습
type part struct {
	description string
	count       int
}

func showInfo(p part) { //part 타입의 매개변수를 받아서 출력하는 함수
	fmt.Println("Description : ", p.description)
	fmt.Println("Count : ", p.count)
}

func main() {
	var bolts part //구조체 변수 선언
	bolts.description = "Hex bolts"
	bolts.count = 100
	showInfo(bolts)
}
