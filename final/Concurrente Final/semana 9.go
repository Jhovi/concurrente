package main

import "fmt"

func Producer (id int, ch chan string){
	c := 0
	for{
		c++
		ch <- fmt.Sprintf("Product %d made by %d",c,id)
	}
}

func Consumer(id int, ch chan string){
	for{
		fmt.Printf("Consumer %d using %s\n",id, <- ch)
	}
}

func main(){
	ch := make(chan string)
	for i := 0; i < 3; i++{
		go Producer(i,ch)
		go Consumer(i,ch)
	}

	Consumer(3,ch)
}