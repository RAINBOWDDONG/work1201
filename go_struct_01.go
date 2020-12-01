package main

import "fmt"

//사용자 정의 => type를 쓴 함수 생성
type subscriber struct {
	name   string
	rate   int
	active bool
}

func main() {
	var s1, s2 subscriber 
	//s1, s2, s3 subscriber 이렇게 필요할 때마다 갖다가 
	//원하는 만큼 쓸 수 있음
	fmt.Printf("%#v\n", s1)

	s1.name = "cho"
	s1.rate = 62500
	s1.active = true
	s2.name = "park"

	fmt.Printf("%s\n", s1.name)
	fmt.Println(s1.rate)
	fmt.Println(s1.active)
	fmt.Println(s2.name)

	/*
		var subscriber struct { //구조체 선언
			name   string
			rate   int
			active bool
		}
		fmt.Printf("%#v\n", subscriber)

		subscriber.name = "cho"
		subscriber.rate = 62500
		subscriber.active = true
		fmt.Printf("%s\n", subscriber.name)
		fmt.Println(subscriber.rate)
		fmt.Println(subscriber.active)
	*/

	/*
		subscriber := map[string]float64{}
		//map으로 만들어줌 앞에 key값은 string 뒤key값은  float64로 해줌
		subscriber ["name"] = "cho"
		subscriber ["rate"] = 6250
		subscriber ["active"] = true
		//[]안에는 string형으로 "" 해주면 되지만 뒤에는 float64이기에
		//"cho"랑 true가 오류가 생긴다
		//그래서 여러 타입을 한번에 사용할 수 있도록 하기 위해
		//struct 구조체를 사용한다 !!
		fmt.Println("test")
	*/

	/*
		subscriber := []string{}
		subscriber = append(subscriber, "cho")
		subscriber = append(subscriber, 5000)
		subscriber = append(subscriber, true)
		fmt.Println("test")
	*/
}
