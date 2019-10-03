package main

import (
	"fmt"
	"sync"
)

func philosophers(name string, leftFork, rightFork sync.Mutex){
	for{
		fmt.Println(name,"pensando")
		leftFork.Lock()
		rightFork.Lock()
		fmt.Println(name, "comiendo")
		leftFork.Unlock()
		rightFork.Unlock()
	}
}

func main(){
	n := 5
	fork := make([]sync.Mutex,n)
	names := []string{"Socrates", "Platon", "Descartes", 
	"Aristoteles", "Chupetin"}
	for i := 0; i < n ; i++{
		go philosophers(names[i],fork[i],fork[i+1])
	}

	philosophers(names[n-1],fork[n-1],fork[0])
}