package main

import "fmt"

func main(){
	//Arrays
	var fruitArr[2]string
	fruitArr[0]="apple"
	fruitArr[1]="orange"
	fmt.Println(fruitArr)

	fruitArr2:=[2]string{"apple","orange"}
	fmt.Println(fruitArr2)

	fruitSlice:=[]string{"mango","apple","orange","grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0:2])
}