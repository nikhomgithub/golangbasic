package main


import (
	"fmt"
	"errors"
	"math"
)

func main(){
	//variable setting
	var x int
	x=5

	var y int=5

	z:=6

	fmt.Println(x,y,z)

	//condition
	if x>6{
		fmt.Println("more than 6")
	} else if x>4{
		fmt.Println("more than 4")
	} else{
		fmt.Println("less or equal 4")
	}

	//array
	var a[5]int
	a[2]=7
	fmt.Println(a)

	//array's number must be fixed 
	b:= [5]int{1,2,3,4,5} 
	fmt.Println(b)

	//slic is an non-fixed-number array 
	c:=[]int{1,2,3}//slice non-fix member
	c=append(c,4)
	fmt.Println(c)

	//key:value variable
	vt:=make(map[string]int)//key : value
	vt["One"]=1
	vt["Two"]=2
	vt["Three"]=3

	fmt.Println(vt)

	//for loop
	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	//for loop with array
	arr:=[]string{"a","b","c"}
	for index,value := range arr{
		fmt.Println("key",index,"value",value)
	}

	//How to use defined function
	result, err:=sqrt(16)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	//How to use object
	p:=person{name:"Kang",age:23}
	fmt.Println(p)

	//pointer
	i:=7
	inc(&i)//put address of variable as input
	fmt.Println(i)
	
}
//Object
type person struct{
	name string
	age int
}

//Function
func sqrt(x float64)(float64,error){
	if x<0{
		return 0,errors.New("Undefined for negative number")
	}
	return math.Sqrt(x),nil
}

//Pointer
func inc(x *int){ //input is address of input
	*x++
}
