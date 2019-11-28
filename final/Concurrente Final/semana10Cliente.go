package main

import (
	"fmt"
	"net"
)

func main(){
	con, _ := net.Dial("tcp","10.11.98.229:8000")
	defer con.Close()
	fmt.Println(con,"Galarza D:")

}