package main 

import (
	"fmt"
	"time"
)

func task1()  {
	time.Sleep(time.Second * 2)
	fmt.Println("task1 おわったよ！")
}

func task2()  {
	fmt.Println("task2 おわったよ！")	
}

func main() {
	go task1()
	go task2()
	
	time.Sleep(time.Second * 3)
}