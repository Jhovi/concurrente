package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	
	Name string
	Estatura float32
	MyAddr Address
}

type Address struct{
	Calle string
	Numero int
}


func main(){
	personas := []Person{
		{"Jon",1.8, Address{"Primavera",2350}},
		{"Dash",1.6, Address{"Cerro Azul",231}},
		{"Fernando",1.7, Address{"La Planicie",666}},
		{"Hans",1.8, Address{"UPC",777}}}

jsonBytes, _ := json.MarshalIndent(personas, "", "\t" )  //json.Marshal(personas)   para mostrar el formato json normal
jsonStr := string(jsonBytes)

fmt.Println(jsonStr)


var otras []Person
json.Unmarshal(jsonBytes, &otras)
fmt.Println(otras)
}


