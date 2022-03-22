package main

import(
	"fmt"
)

func loop1(cnt, step int)(int){
	var total int
	for i := 0; i < cnt; i += 1{
		total += i * step
		defer func(myi, mytotal int){
			fmt.Printf("myi %v, mytotal is %v\n", myi, mytotal)
			fmt.Printf("i %v, total is %v\n", i, total)
		}(i, total)
	}
	return total
}

func main(){
	loop1(10, 1)
}