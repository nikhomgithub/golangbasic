package main

import "fmt"

func greeting(name string)string{
	return "Hello "+name 
}

func getSum1(num1 int,num2 int)int{
	return num1+num2
}

func getSum2(num1,num2 int)int{
	return num1+num2
}

func main(){
	fmt.Println(greeting("Nikhom"))
	fmt.Println(getSum1(1,2))
	fmt.Println(getSum2(1,2))
}


