package main

import "fmt"

//사용자 정의 => type를 쓴 함수 생성
type subscriber struct {
	name   string
	rate   int
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name : ", s.name)
	fmt.Println("Monthly rate : ", s.rate)
	fmt.Println("Active? ", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5000
	s.active = true
	return &s  //(&)주소가 가르키는 값을 (*)리턴 시킨다 
}

func applyDiscount(s *subscriber) {
	s.rate = 4000
}

func main() {
	//가입중이던 구독자
	s1 := defaultSubscriber("cho") 
	//s가 가르키는 값을 포인터 형으로 리턴 받고 반환된 값을 applyDiscount함수에 포인트형 s1을 보내면 5000이라는 값을 가진  s.rate가 4000을 가르키게 된다
	//s1.rate = 4500
	applyDiscount(s1)
	printInfo(s1)

	//오늘 가입한 구독자
	s2 := defaultSubscriber("park")
	printInfo(s2)

	var s3 subscriber
	applyDiscount(&s3)
	fmt.Println(s3.rate)
}
