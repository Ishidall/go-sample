package main

import (
	"fmt"
	"errors"
	"github.com/Ishidall/go-sample/box"
	"github.com/Ishidall/go-sample/reflection"
	"github.com/Ishidall/go-sample/goroutine"
)

var varName int
var var1, var2, var3 string
var globalVar string = string("hogehoge")
var numbers map[string]int

func main() {
	varName = 1
	bigInteger := 9000000000000000000
	// bigInteger := 9999999999999999999
	fmt.Printf("%d\n", bigInteger)
	var1, var2, var3 = "hoge", "fuga", `moge
moga`
	arr := [3]int{0, 1, 2}
	slice := arr[0:1]
	sliceMax := arr[:]
	numbers = make(map[string]int)
	numPointer := new(map[string]int)
	num1, num2 := 1, 1

	fmt.Printf("before num1: %v\n", num1)
	fmt.Printf("before num2: %v\n", num2)

	addedNum1 := addCopy(num1)
	addedNum2 := addEntity(&num2) // &num2: num2変数に割り当てられているメモリ上のアドレス

	fmt.Printf("after num1: %v\n", num1)
	fmt.Printf("after num2: %v\n", num2)

	fmt.Printf("addedNum1: %v\n", addedNum1)
	fmt.Printf("addedNum2: %v\n", addedNum2)

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%v\n", numPointer)

	numbers["one"] = 1

	fmt.Printf("%v\n", numbers)

	fmt.Printf("%v\n", arr)

	fmt.Print(slice)
	fmt.Print(append(slice, 5))

	fmt.Print(sliceMax)
	fmt.Print(append(sliceMax, 5))

	fmt.Printf("%v\n", arr)

	incremented := incrementArray(arr)

	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", incremented)

	localVar1, localVar2, localVar3 := 1, "local", "local, im sayin."

	err := errors.New("I am an error.")

	if err != nil {
		fmt.Printf("%s\n", err)
		fmt.Print(err)
	}

	if v := varName; v > 0 {
		fmt.Printf("v is greater than 0. v = %v\n", v)
	} else {
		fmt.Printf("v is less than 0. v = %v\n", v)
	}

	fmt.Printf("%v\n", varName)
	fmt.Print(varName)
	fmt.Printf("%s, %s, %s\n", var1, var2, var3)
	fmt.Printf("%v, %s, %s\n", localVar1, localVar2, localVar3)

	fmt.Printf("%s\n", string(globalVar[2]))
	fmt.Printf("%s\n", globalVar[1:])
	fmt.Printf("%s\n", globalVar[:])

	// gotoFunc()
	// deferFunc()

	box.BoxImpl()
	reflection.ReflectionImpl()
	goroutine.GoroutineImpl()
}

func incrementArray(arr [3]int) [3]int {
	arr[0] = arr[0] + 1
	arr[1] = arr[1] + 1
	arr[2] = arr[2] + 1

	return arr
}

func gotoFunc() {
	i := 0

	Here:

	println(i)

	i++

	if i >= 100 {
		return
	}
	
	goto Here
}

func addCopy(cp int) int {
	cp = cp + 1

	return cp
}

func addEntity(pointer *int) int { // pointer *int: int型の数値を指すpointer
	*pointer = *pointer + 1 // *pointer: pointerが指し示すアドレスの中身

	return *pointer
}

func deferFunc() bool {
	a := 1

	fmt.Printf("a is %v\n", a)

	defer fmt.Print(a)
	defer fmt.Print("******************************")
	defer fmt.Print(1)

	// a++
	// fmt.Printf("a is %v\n", a)

	defer fmt.Print(a)
	defer fmt.Print("******************************")
	defer fmt.Print(2)

	// a++
	// fmt.Printf("a is %v\n", a)

	defer fmt.Print(a)
	defer fmt.Print("******************************")
	defer fmt.Print(3)

	a++
	fmt.Printf("a is %v\n", a)

	a++
	fmt.Printf("a is %v\n", a)

	return true
}
