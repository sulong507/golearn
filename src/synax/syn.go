package main

import "fmt"
import "errors"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var v1 int
	var v2 string
	var v3 [10]int //array
	var v4 []int   //slice
	var v5 struct {
		f int
	}
	var v6 *int           //point
	var v7 map[string]int //map, key:string, value:int
	var v8 func(a int) int

	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8)

	_, _, nickName := GetName()
	fmt.Println(nickName)

	//int uint int32 int64 float32 float64 complex64 complex 128

	const zero = 0.0

	const mask = 1 << 3

	fmt.Println(zero, mask)

	const (
		c0     = iota
		c1     = iota
		Sunday = iota
	)
	fmt.Println(c0, c1, Sunday)

	//data type
	//bool
	//int8,byte,int16,int,uint,uintptr
	//float32 float64
	//complex64,complex128
	//string
	//rune
	//error
	//pointer array slice map chan struct interface

	fmt.Printf("print name: \"%s\" is %d \n", "sulong", 30)

	//str := "hello world."
	//for i, ch := range str {
	//fmt.Println(i, ch)
	//}

	mySlicel1 := make([]int, 5, 10)
	mySlicel2 := []int{1, 2, 3, 4, 5}

	fmt.Println(cap(mySlicel1))
	fmt.Println(cap(mySlicel2))

	//float compare
	f1 := 1.5246676534323
	f2 := 1.5246676534323
	if f1 == f2 {
		fmt.Println("exception")
	}

	//array
	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(nums))

	mySlice3 := mySlicel2[3:]
	fmt.Println(mySlice3)

	mySlice := make([]int, 5, 10)
	fmt.Println("le mySlice: ", len(mySlice))
	fmt.Println("cap mySlice: ", cap(mySlice))
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Println("le mySlice: ", len(mySlice))
	fmt.Println("cap mySlice: ", cap(mySlice))
	mySlice4 := []int{8, 9, 10}
	mySlice = append(mySlice, mySlice4...)
	fmt.Println("le mySlice: ", len(mySlice))
	fmt.Println("cap mySlice: ", cap(mySlice))

	mySlice5 := []int{1, 2, 3, 4, 5}
	mySlice6 := []int{5, 4, 3}
	copy(mySlice5, mySlice6)
	fmt.Println(mySlice5, mySlice6)

	//map[string]PersonInfo key:string value:PersonInfo
	var personMap map[string]PersonInfo
	personMap = make(map[string]PersonInfo, 20) //init size 20

	/*
		personMap = map[string]PersonInfo{
			"1234": PersonInfo{"1234", "Tom", "room 203"},
		}
	*/

	personMap["1234"] = PersonInfo{"1234", "Tom", "room 203"}
	personMap["2345"] = PersonInfo{"2345", "Json", "road linken, No.43. room 504"}

	person, ok := personMap["1234"]
	if ok {
		fmt.Println("fetch success.", person.Name)
	} else {
		fmt.Println("fetch failure.")
	}

	fmt.Println(len(personMap))
	delete(personMap, "1234")
	fmt.Println(len(personMap))

	fmt.Println(FetchNum(1))

	index := 2
	switch index {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	Count()

	ret, err := Add(-2, 1)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("RESULT: ", ret)
	}

	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))

	//defer panic recover
	defer func() {
		fmt.Println("xxxxxxxxxxxxx")
	}()

	defer func() {
		fmt.Println("yyyyyyyyyyyyy")
	}()

}

func GetName() (first, last, nick string) {
	return "May", "Chan", "Chibi Maruko"
}

func FetchNum(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func Count() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func Add(a, b int) (ret int, err error) {
	if a < 0 || b > 0 {
		err = errors.New("should be non-negative numbers!")
		return
	}
	return a + b, nil
}
