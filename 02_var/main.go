package main

import "fmt"

func main(){
	var name string
	name="Brad"
	surname,email:="Phut","mee@gmail.com"

	var age int32 = 37
	const isCool=true
	nickname:="mee"
	size1:=3.7
	size2:=30
	var size3 float32=3.7

	fmt.Println(surname," and ",email)
	fmt.Println(name,age)
	fmt.Printf("name type is %T\n",name)
	fmt.Printf("age type is %T\n",age)
	fmt.Printf("isCool type is %T\n",isCool)
	fmt.Printf("nickname type is %T\n",nickname)
	fmt.Printf("size1 type is %T\n",size1)
	fmt.Printf("size2 type is %T\n",size2)
	fmt.Printf("size3 type is %T\n",size3)


}

/*
Phut  and  mee@gmail.com
Brad 37
name type is string
age type is int32
isCool type is bool
nickname type is string
size1 type is float64
size2 type is int
size3 type is float32
*/