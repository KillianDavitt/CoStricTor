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
	
	filterSizes := []interface{}{1024,2048,4092,8192}
	sampleSizes := []interface{}{100,500,2000}
	numsSites := []interface{}{100}
	hstsProps := []interface{}{0.2}
	httpProps := []interface{}{0.2}
	primaryThresholds := []interface{}{0.05,0.01,0.005,0.001}
	secondaryThresholds := []interface{}{0.1,0.5}
	ps := []interface{}{0.2}
	qs := []interface{}{0.8}
	numsHashes := []interface{}{1,2,4,6,8}
	
	prm := cartesian.Iter(filterSizes, sampleSizes,numsSites,hstsProps, httpProps, primaryThresholds, secondaryThresholds, ps, qs, numsHashes)
	var sites []string = make([]string, 100)
	sites = lines[0:100]
	hsts, http, https_no_hsts := generateSites(sites, 0.2, 0.2);
	
        var wg sync.WaitGroup
	for params := range prm {
		wg.Add(1)
		go runSim(params, hsts, http, https_no_hsts, &wg)
	}
	wg.Wait()
}

