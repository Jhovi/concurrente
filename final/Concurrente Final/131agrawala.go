package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
)

type Info struct {
  nextIp  string
  nextNum int
  cont    int
  first   bool
}
type RemoteInfo struct {
  Num int         `json:"num"`
  Ip  string      `json:"ip"`
}

var chInfo chan Info
var myNum int
var addrs []string

const myip = "10.11.98.212"

func main() {
	chInfo = make(chan Info)
	myNum = rand.Intn(1000000)
	go func() {
		chInfo<- Info{"", 1000001, 0, true}
	}()

	fmt.Printf("Soy %s\n", myip)
	go registerServer(myip)
	go startServer(myip)
	go agrawalaServer(myip)

	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese direccion remota: ")
	remoteIp, _ := gin.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)


	go func() {
		fmt.Printf("Mi numero es: %d, press enter to start...", myNum)
		gin.ReadString('\n')
		sendMyNumToAll()
	}()

	if remoteIp != "" {
		registerSend(remoteIp, myip)
	}

	notifyServer(myip)
}
func agrawalaServer(hostAddr string) {
	host := fmt.Sprintf("%s:8003", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleAgrawala(conn)
	}
}
func handleAgrawala(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	var remoteInfo RemoteInfo
	json.Unmarshal([]byte(msg), &remoteInfo)
	info := <-chInfo
	info.cont++
	if remoteInfo.Num < myNum {
		info.first = false
	} else if remoteInfo.Num < info.nextNum {
		info.nextNum = remoteInfo.Num
		info.nextIp = remoteInfo.Ip
	}
	go func() {
		chInfo<- info
	}()
	if info.cont == len(addrs) && info.first {
		go startTurn()
	}
}
func startTurn() {
		fmt.Println("Toca iniciar!")
		notifyNext()
}
func notifyNext() {
	info := <-chInfo
	remote := fmt.Sprintf("%s:8004", info.nextIp)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()

	fmt.Fprintln(conn, "Dale!")
}
func sendMyNumToAll() {
	for _, addr := range addrs {
		go sendMyNum(addr)
	}
}
func sendMyNum(remoteAddr string) {
	remote := fmt.Sprintf("%s:8003", remoteAddr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	remoteInfo := RemoteInfo{myNum, myip}
	msg, _ := json.Marshal(remoteInfo)
	fmt.Fprintln(conn, string(msg))
}
func startServer(hostAddr string) {
	host := fmt.Sprintf("%s:8004", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleStart(conn)
	}
}
func handleStart(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	r.ReadString('\n')
	go startTurn()
}
func registerSend(remoteAddr, hostAddr string) {
	remote := fmt.Sprintf("%s:8000", remoteAddr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()

	// Enviar direccion
	fmt.Fprintln(conn, hostAddr)

	// Recibir lista de direcciones
	r := bufio.NewReader(conn)
	strAddrs, _ := r.ReadString('\n')
	var respAddrs []string
	json.Unmarshal([]byte(strAddrs), &respAddrs)

	// agregamos direcciones de nodos a propia libreta
	for _, addr := range respAddrs {
		if addr == remoteAddr {
			return
		}
	}
	addrs = append(respAddrs, remoteAddr)
	fmt.Println(addrs)
}
func registerServer(hostAddr string) {
	host := fmt.Sprintf("%s:8000", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegister(conn)
	}
}
func handleRegister(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIp, _ := r.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)

	// respondemos enviando lista de direcciones de nodos actuales
	byteAddrs, _ := json.Marshal(addrs)
	fmt.Fprintf(conn, "%s\n", string(byteAddrs))

	// notificar a nodos actuales de llegada de nuevo nodo
	for _, addr := range addrs {
		notifySend(addr, remoteIp)
	}

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == remoteIp {
			return
		}
	}
	addrs = append(addrs, remoteIp)
	fmt.Println(addrs)
}
func notifySend(addr, remoteIp string) {
	remote := fmt.Sprintf("%s:8001", addr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintln(conn, remoteIp)
}
func notifyServer(hostAddr string) {
	host := fmt.Sprintf("%s:8001", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleNotify(conn)
	}
}
func handleNotify(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIp, _ := r.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == remoteIp {
			return
		}
	}
	addrs = append(addrs, remoteIp)
	fmt.Println(addrs)
}
