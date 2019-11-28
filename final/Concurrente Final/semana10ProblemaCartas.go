package main

import "fmt"

func pipesort(inCh,outCh chan int){
	min := int(1e9)
	for num := range inCh{
		if num < min{
			outCh <- min
			min = num
		}else {
			outCh <- num
		}
	}
	fmt.Println(min)
	close(outCh)
}

func main(){
	n := 10
	ch := make([] chan int, n+1)
	ch[0] = make(chan int)
	for i := 0; i < n; i++{
		ch[i+1] = make(chan int)
		go pipesort(ch[i], ch[i+1])
	}
	go func(){
		nums := []int{8,3,6,1,9,2,7,5,10,4}
		for _,num := range nums{
			ch[0] <- num //enviamos los numeros al canal
		}
		close(ch[0])
	}()
	for range ch[n]{

	}
}