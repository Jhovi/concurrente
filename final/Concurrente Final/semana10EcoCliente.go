package main

import ( 
	"fmt"
	"net"
	"os"
	"bufio"
)

func main(){
	con, _ := net.Dial("tcp","10.11.98.211:8000")
	defer con.Close()
	r := bufio.NewReader(con)
	gin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Ingrese mensaje: ")
		msg, _ := gin.ReadString('\n')
		fmt.Fprint(con,msg)
		resp, _ := r.ReadString('\n')
		fmt.Printf("Respuesta: %s",resp)
	}
}