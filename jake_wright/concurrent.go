package main

import(
  "fmt"
  "time"
  "strconv"
  "sync"
)
//main()also is a goroutine
func main(){
  workerapproach()
}


//==========================================
//To run another routine (count) with main routine
func run10goroutine(){
  for i:=1;i<10;i++{
    t := strconv.Itoa(i)
    go count(t)
  }
  count("main")//to prevent main() routine to complete
  //fmt.Scanln()//to prevent main() routine to complet ,Wait for any key enter
}

//To display run another go routine with main()routine
func withonegoroutine(){
  go count("sheep") //this is a goroutine
  count("fish")
}

//To run function with a single routine (main())
func withoutconcurrent(){
  count("sheep:never count fish")
  count("fish")
}

//infinite count function
func count(thing string){
  for i:=1;true;i++{
    fmt.Println(i,thing)
    time.Sleep(time.Millisecond*500)
  }
}

//============================
//ask main() to wait for another goroutine to complete
func wgsync(){
  var wg sync.WaitGroup
      wg.Add(1)//Tell main() to wait for a goroutine

      go func(){
        count2("sheep")
        wg.Done() //tell main() that a goroutine is completed
      }()

      wg.Wait()//When main() come to this line it will wait (block) until wg.Done completed
}
func count2(thing string){
  for i:=1;i<=10;i++{
    fmt.Println(i,thing)
    time.Sleep(time.Millisecond*500)
  }
}
//========================
//Channel is use to communicate between routine
func mainchannel(){
  c:=make(chan string)
  go countchannel("sheep",c)//this function sen input via c

  for{
    msg,open:=<-c//wait until is received (blocking nature)
    if!open{ //to check whether channel is close, if so break for loop and main routine will come to an end
    break
  }
  fmt.Println(msg)

  /*or write another way
  for msg:=range c{
    if!open{
      break
    }
    fmt.Println(msg)
  }
  */
    
  }
}

//input is thing(string)
//use c as channel(string)
//1<i<=10   i+thing  >> tt >> c >> msg (printlin) 
//msg must check channel status(open/close)
func countchannel(thing string,c chan string ){
  for i:=1;i<=10;i++{
    tt := strconv.Itoa(i)+thing
    c<-tt//wait until ready to be sent (blocking nature)
    time.Sleep(time.Millisecond*500)
  }
  close(c)//need to close channel 
}
//===============================
//buffer channel (channel open until buffer is full, it will close)
func bufferchannel(){
  c:=make(chan string,3)//there are 3 string for this buffer
  c<-"hello"
  c<-"world"
  c<-"bye"
  msg1:=<-c
  msg2:=<-c
  msg3:=<-c
  fmt.Println(msg1,msg2,msg3)

}
// channel with buffer 3 strings
//"hello">c>msg1 
//"wold">c>msg2
//"bye">c>msg3

//================================
//select is to receive from any channel who completed
//no need to wait(block) to complete in sequence
func selectchannel(){
  c1:=make(chan string)
  c2:=make(chan string)

  go func(){
    for{
      c1<-"Every 500ms"
      time.Sleep(time.Millisecond*500)
    }
  }()

  go func(){
    for{
      c2<-"Every 2 second"
      time.Sleep(time.Second*2)
    }
  }()

  for{
    select{
    case msg1:=<-c1:
      fmt.Println(msg1)
    case msg2:=<-c2:
      fmt.Println(msg2)
    }
  }  
}

//================================================
//patter of "work approach"

func workerapproach(){

  //create 2 channels for worker()
  jobs:=make(chan int,100) //define channel&buffer value
  results:=make(chan int,100)//define channel&buffer value

  go worker(jobs,results)

  
  for i:=0;i<15;i++{
    jobs<-i
    fmt.Println("jobs",i)
    
    time.Sleep(time.Millisecond*5)
    fmt.Println("results",<-results)

    fmt.Println("------------End------------")
  }

  fmt.Println("============Over===============")
  close(jobs)
  /*
  for j:=0;j<5;j++{
    fmt.Println(<-results)
  }
  */
}


//this function need 2 channels


func worker(jobs<-chan int,results chan<-int){
  for n:=range jobs{
    fmt.Println("worker",n)
    results<-fib(n) 
  }

}

//this is fibinachi (a math-school function)
//fib(0)=0, fib(1)=1, fib(2)=4, fib(3)=7,fib(4)=9, fib(5)=11
func fib(n int)int{ 
  if n<=1{
    fmt.Println("fib n<=1..................")
    return n
  }
  fmt.Println("fib n>1....................")
  return fib(n-1)+fib(n-2)
}