package main

import "fmt"

func t(k int) {

}

func demoFunc1() string {
	return "demoFunc1"
}

func demoFunc2() string {
	return "demoFunc2"
}

func demoFunc3(param1, param2 string) {
	fmt.Println(param1, param2)
}

func main() {

	for i := 1; i < 10; i++ {
		k := i
		t(k)
	}

	demoFunc3(demoFunc1(), demoFunc2())

}
