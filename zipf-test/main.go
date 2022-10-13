package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){

	source := rand.NewSource(time.Now().UnixNano()) 
	hsts_zipf := rand.NewZipf(rand.New(source), 1.1, 9999.0, 100)
	
	numReports := 1000
	var reports []int = make([]int, 100) 
	
	for i:=0; i<numReports; i++ {
		n := hsts_zipf.Uint64()
		reports[n]+=1
	}

	fmt.Println(reports)
}
