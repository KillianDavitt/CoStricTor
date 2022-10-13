package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){

	source := rand.NewSource(time.Now().UnixNano()) 
	hsts_zipf := rand.NewZipf(rand.New(source), 1.1, 9999.0, 499)
	
	numReports := 10000
	var reports []uint = make([]uint, 500) 
	
	for i:=0; i<numReports; i++ {
		n := hsts_zipf.Uint64()
		reports[n]+=1
	}

	fmt.Println(reports)
}
