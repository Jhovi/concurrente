package main

import(
	"fmt"
	"sort"
)
func (a StringSlice) Len() int           { return len(a) }
func (a StringSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StringSlice) Less(i, j int) bool { return a[i] < a[j] }


sort.Sort(StringSlice.names)

var names = []*Name{
	{"Pablo", "Galarza"},
	{"Jose","Perez"},
	{"Alan", "Garcia"}
}

func main(){
	values := []int{3,1,2}
	fmt.Println(sort.IntsAreSorted(values))
}