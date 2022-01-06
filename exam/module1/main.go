package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type MyName struct {
	Name string `json:"name"`
}

func printName(n *MyName) {
	fmt.Println(&n.Name)
}

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "qi", "miao")
	go func(ctx context.Context) {
		println(ctx.Value("qi"))
	}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				println("child")
				return
			default:
				println("default")
			}
		}
	}(timeoutCtx)
	time.Sleep(time.Second)
	select {
	case <-timeoutCtx.Done():
		time.Sleep(time.Second)
		println("main")

	}
	ch := make(chan int)
	go func() {
		println("channel")
		ch <- 2
	}()
	two := <-ch // block 等待
	println(two)

	println("a")
	mmn := MyName{Name: "miaomiao"}
	fmt.Println(mmn.Name)
	myType := reflect.TypeOf(mmn)
	name := myType.Field(0)
	fmt.Println(name)
	tag := name.Tag.Get("json")
	fmt.Println(tag)
	const a = 6
	var aa, bb = 1, 2
	var f = float64(a)
	fmt.Printf("a %d\n %f\n", a, f)
	fmt.Printf("a %d\n %d\n", aa, bb)
	fmt.Println("hello go")
	fmt.Println("hello go")
	fmt.Println("hello go")

	mn := MyName{Name: "qiqi"}
	fmt.Println(mn)
	printName(&mn)

	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	mySlice := make([]int, 10)
	mySlice2 := []int{}
	mySlice2 = append(mySlice2, 2)
	fmt.Println(mySlice2)
	for i, _ := range mySlice {
		mySlice[i] = 3
	}
	fmt.Println(mySlice)
	myMap := make(map[int]int, 3)
	myMap[1] = 1
	myFuncMap := map[string]func() int{
		"a": func() int {
			return 1
		},
	}
	fmt.Println(myMap)
	fmt.Println(myFuncMap)
	fu := myFuncMap["a"]
	fmt.Println("func print: ", fu())
	for _, s := range arr {
		fmt.Println(s)
	}

	for i := 0; i < 5; i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		}
		if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
	fmt.Println(arr)
	n := 0
	reply := &n
	fmt.Println("Multiply:", reply) // Multiply: 50
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50

	/*str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}*/
	/*// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}*/

}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
