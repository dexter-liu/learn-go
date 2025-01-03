package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器",
		Long:  "read1",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", -1, "身高")
	cmd.Flags().Float64Var(&weight, "weight", -1, "体重")
	cmd.Flags().IntVar(&age, "age", -1, "年龄")

	cmd.Execute()

}
