package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dexter2() {
	var name string
	var age int
	fmt.Println("enter your name and age: ")
	fmt.Scan(&name, &age)
	fmt.Printf("hello %s, you are %d years old\n", name, age)
}

func Dexter3() {
	var name string
	var age int
	fmt.Println("enter your name and age: ")
	fmt.Scanln(&name, &age)
	fmt.Printf("hello %s, you are %d years old\n", name, age)
}

func Dexter4() {
	var name string
	var age int
	fmt.Println("enter your name and age (formated): ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("hello %s, you are %d years old\n", name, age)
}

func Dexter5() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter your name and age: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Printf("type %T\nyou entered: %s\n", input, input)
}
