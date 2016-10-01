package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	name, course string
	module       float64
)

var (
	name1 = "vimal"
	val   = 2.33
)

func main() {

	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)

	a := 10.000
	b := 3
	c := int(a) + b
	fmt.Println("Hello from", runtime.GOOS)
	fmt.Println("Name is ", name, "and is type of ", reflect.TypeOf(name))
	fmt.Println("module is ", module, "and is type of ", reflect.TypeOf(module))
	fmt.Println(c)
	fmt.Println("Memory address of c", &c)
	course := "Docker deep Dive"
	name2 := "Nigel"
	fmt.Println(name2)
	changeCourse(course)

}

func changeCourse(course string) string {
	course = "First look:  native docker cluster"
	fmt.Println("\n", course)
	return course

}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)
	return module, author
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}

func openfile() {
	_, err := os.Open("c:\\test.txt")

	if err != nil {
		fmt.Println("Error ", err)
	}
}

func timeCount() {
	courseInProg := []string{"string1", "string2", "string3"}

	for _, i := range courseInProg {
		fmt.Println(i)
	}

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			break
		}
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}

func slicearray() {
	myCourses := make([]string, 5, 10)
	//len - 5 capacity 10
	fmt.Printf("%d , %d", len(myCourses), cap(myCourses))

}

func mapUseCase() {
	//map are unordered
	// maps <key> : <value> pairs
	// maps dynamically resizable
	//Maps are refernces
	// go strongly typed
	//key typed should be comparable type
	leaquetitles := make(map[string]int)
	leaquetitles["Sunderland"] = 6
	leaquetitles["Newcastle"] = 4
	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  4,
	}
	println(recentHead2Head)

	// printf %v %v ,key,value
	//%t for bool
	// delete(map,key)
	//unsafe for concurrency
	//cheap to pass passed to function by reference
}

func trystruct() {
	//structs - use to define type
	type courseMeta struct {
		Author string
		Level  string
		Ratin  float64
	}
	//var dockerdeepdive courseMeta
	//dockerdeepdive := new(courseMeta)
	dockerdeepdive := courseMeta{
		Author: "A",
		Level:  "b",
		Ratin:  5,
	}
	fmt.Println(dockerdeepdive)
}

func concurrency() {

	runtime.GOMAXPROCS(2)
	var waitgrp sync.WaitGroup
	waitgrp.Add(2)

	go func() {
		defer waitgrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitgrp.Done()
		fmt.Println("hello")
	}()

	waitgrp.Wait()
}

//go routine - go to safe comms channels to other go routine
/*

Go variable are statically typed
package level variable are globally
shorthad declare and initialize := only work inside the function

Go pass arguments by value not by reference

constant are immutable
const x= val

switch case fallthrough

idiomatic err

for keyword
expresion
  blank expression = infinite loop
  can be boolean expression
  can be a range

arrays vs Slices
Array - is like list with fixed length
Slice can be resized - slice built on top of array
Slice are refernces = passed by refernce

*/
