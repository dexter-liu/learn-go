package main

import "fmt"

func Dexter1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("程序有错误\n")
		}
	}()

	var sex string

	fmt.Scanln(&sex)

	if sex != "男" && sex != "女" {
		return fmt.Errorf("录入性别既不是 男， 也不是女")
	}

	var weight int
	fmt.Scanln(&weight)
	if weight <= 10 || weight > 200 {
		return fmt.Errorf("输入体重 %d 超出有效范围", weight)
	}
}
