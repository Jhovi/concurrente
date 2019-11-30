package main 

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"strings"
	"strconv"
)

const (
	IP = "localhost"
	TCP = "tcp"
)

var nodes map[string]bool = make(map[string]bool)

func getMsg(cn net.Conn) string {
	r := bufio.NewReader(cn)
	msg, _ := r.ReadString('\n')
	return strings.TrimSpace(msg)
}

func host(ip string, port int) string {
	return fmt.Sprintf("%s:%d", ip , port)
}

func ipAdder(chId <-chan int){
	for {
		nodes[<-chId] = true
		fmt.Println(nodes)
	}
}

func EnviarSinRespuesta(msg string, target int){
	cn, _ := net.Dial(TCP, host(IP, target))
	defer cn.Close()
	fmt.Fprintln(cn, msg)
}


func EnviarConRespuesta(msg string, target int)string{
	cn, _ := net.Dial(TCP, host(IP, target))
	defer cn.Close()
	fmt.Fprintln(cn, msg)
	return getMsg(cn)
}


func ServidorAgregador(id int, chId chan<- int){
	ln, _ := net.Listen(TCP, host(IP, id))
	defer ln.Close()
	for {
		cn, _ := ln.Accept()
		go func(cn net.Conn){
			val, _ := strconv.Atoi(getMsg(cn))
			chId <- val
			cn.Close()
		}(cn)
	}
}

func ClienteAgregador(newId int, nodes map[string]bool){
	msg := fmt.Sprintf("%d", newId)
	for target := range nodes {
		go EnviarSinRespuesta(msg, target + 1)
	}
}

func ServidorRegistrador(id int, chId, end chan<- int){
	ln, _ := net.Listen(TCP, host(IP, id))
	defer ln.Close()
	for {
		cn, _:= ln.Accept()
		go func(cn net.Conn){
			newId, _:= strconv.Atoi(getMsg(cn))
			ClienteAgregador(newId, nodes)
			buf, _ := json.Marshal(nodes)
			chId <- newId
			cn.Close()
		}(cn)
	}end <- 0
}

func ClienteRegistrador(id, targetId int, chId chan<- int){
	resp := EnviarConRespuesta(fmt.Sprintf("%d", id), targetId)
	var slc map[string]bool
	_ := json.Unmarshal([]byte(resp), &slc)
	for newId := range slc {
		chId <- newId
	}
}