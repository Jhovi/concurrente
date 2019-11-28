package main

import "fmt"

func philosopher(name string,rightFork, leftFork chan bool){

	for{
		fmt.Println("%s is thinking\n",name)
		<-rightFork
		<-leftFork
		fmt.Println("%s is eating\n",name)
		rightFork <- true
		leftFork <- true
	}
}

func fork(forky chan bool){
	for {
		forky <- true
		<- forky
	}
}

func main(){
	names := []string{"Platon", "Descartes", "Nietzsche","Aristoteles"}

	forks := make([] chan bool,5)

	forks[0] = make(chan bool,1)
	for i,name:= range names {
		forks[i+1] = make(chan bool)
		go philosopher(name,forks[i], forks[i+1])
		go fork(forks[i])
	}

	go fork(forks[4])
	philosopher(" Susy", forks[4],forks[0])
}