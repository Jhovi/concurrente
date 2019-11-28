package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	ln, _ := net.Listen("tcp","10.11.98.212:8000")
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		
		go handle(con)

	}
}

func handle(con net.Conn){
	defer con.Close()
	r := bufio.NewReader(con)
	for{
		msg, err := r.ReadString('\n')
		if err != nil{
			break
		}
		fmt.Printf("Recibido: %s",msg)
		fmt.Fprint(con,msg)



		//Solo agregue esto para comunicarme con otra pc 
		con2, _ := net.Dial("tcp","10.11.98.211:8000")
		defer con2.Close()

		fmt.Fprint(con2,msg)

		
	}

	
}