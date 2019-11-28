package main

import "fmt"


var end 
func Source(row [] int, south chan int){
	for elem := range row{
		south <- elem
	}
	close(south)

}

func zero(n int, west chan int){
	for i := 0; i < n; i++{
		west <- 0
	}
	close(west)
}

func sink(north chan int){
	for range north{
		
	}
}

func result(c [][]int,i int, east chan int){
	j := 0
	for c[i][j] := range east{
		j++
	}
}

func multiplier(firstElement int, north,east,west,south chan int){
	for secondElement := range north{
		west <-(<-east + firstElement*secondElement)
		south <- secondElement
	}

	close(west)
	close(south)
}

func main(){
	a := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	b := [][]int{{1,0,2},{0,1,2},{1,0,0}}
	c := [][]int{{0,0,0},{0,0,0},{0,0,0}}

	nra := len(a)
	nca := len(a[0])

	ew := make([][]chan int,nra)
	for i:=range ew{
		ew[i] = make([]chan int, nca + 1)
		for j := range ew[i]{
			ew[i][j] = make(chan int)
		}
	}
	ns := make([][]chan int, nra + 1)
	for i := range ns{
		ns[i] = make([]chan int, nca)
		for j:= range ns[i]{
			ns[i][j] = make(chan int)

		}
	}

	for i := range b {
		go source(b[i], ns[0][i])
		go sink(ns[nra][i])

	}
	for i := range a {
		go zero(nra,ew[i][nca])
		go result(c,i, ew[i][0])
	}

	for i := range a{
		for j:= range b{
			go multiplier(a[i][j],ns[i][j], ew[i][j+1], ns[i+1][j],ew[i][j])
		}
	}

	for _= range a {
		<-end
	}
	fmt.Println(c)
}