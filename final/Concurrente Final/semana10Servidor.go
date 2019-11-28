package main

import (
	"fmt"
	"net"
	"bufio"
)

func main(){

	ln, _ := net.Listen("tcp","10.11.98.212:8000")
	defer ln.Close() //se ejecuta al final (defer)
	con, _ := ln.Accept()
	defer ln.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Println(msg)
}