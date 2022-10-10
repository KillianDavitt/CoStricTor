package main

import (
	"os"
	"bufio"
	"github.com/schwarmco/go-cartesian-product"
	"sync"
)

func main() {

	// Load in the big list of websites 
	const path string = "websites.txt"
	file, err := os.Open(path)
	if err != nil {
		return 
	}
	defer file.Close()
	
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	filterSizes := []interface{}{2048}
	sampleSizes := []interface{}{100,2000}
	numsSites := []interface{}{100}
	hstsProps := []interface{}{0.2}
	httpProps := []interface{}{0.2}
	primaryThresholds := []interface{}{0.001,0.01,0.0001,0.00001}
	secondaryThresholds := []interface{}{0.001,0.01}
	ps := []interface{}{0.01}
	qs := []interface{}{0.99}
	numsHashes := []interface{}{1,2}
	
	prm := cartesian.Iter(filterSizes, sampleSizes,numsSites,hstsProps, httpProps, primaryThresholds, secondaryThresholds, ps, qs, numsHashes)
	
        var wg sync.WaitGroup
	for params := range prm {
		wg.Add(1)
		go runSim(params, lines, &wg)
	}
	wg.Wait()
}

