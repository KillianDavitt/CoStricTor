package main

import (
	"os"
	"bufio"
	"github.com/schwarmco/go-cartesian-product"
	"sync"
	"fmt"
	"strconv"
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

	// Sort out parameters
	// Some need to be in an interface for the library to get cartesian product of all parameters
	hstsProp := 0.2
	httpProp := 0.2
	filterSizes := []interface{}{2048,4092,8164,16328}
	sampleSizes := []interface{}{10000}
	numSites := 100
	primaryThresholds := []interface{}{0.1,0.05,0.01,0.001,0.0005}
	secondaryThresholds := []interface{}{0.1,0.5,0.001,0.0005}
	ps := []interface{}{0.2}
	qs := []interface{}{0.8}
	numsHashes := []interface{}{1,32,128}

	// Get the cartesian product, i.e. all possible combinations of the parameters
	prm := cartesian.Iter(filterSizes, sampleSizes, primaryThresholds, secondaryThresholds, ps, qs, numsHashes)

	// Result is a channel, draw all items from it to make it a slice
	perms := make([]interface{},len(prm))
	for x := range prm{
		perms = append(perms,x)
	}

	// Divide the parameters in chunks for the array job
	numJobs := 10
	sizeChunks := int(len(perms)/numJobs)
	var jobs [][]interface{};
	jobs, err = chunkSlice(perms, sizeChunks)
	if err != nil {
		fmt.Println("Error dividing up the job chunks")
		return
	}

	// Get current job number
	jobString := os.Getenv("SGE_TASK_ID")
	jobNumber, err := strconv.Atoi(jobString)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error reading the job number from environment")
		return
	}

	var sites []string = make([]string, numSites)
	sites = lines[0:numSites]
	hsts, http, https_no_hsts := generateSites(sites, hstsProp, httpProp);
	
        var wg sync.WaitGroup
	for _,params := range jobs[jobNumber-1] {
		wg.Add(1)
		go runSim(params.([]interface{}), hsts, http, https_no_hsts, &wg, hstsProp, httpProp, numSites)
	}
	wg.Wait()
}

func chunkSlice(slice []interface{}, chunkSize int) ([][]interface{}, error) {

	chunks := make([][]interface{}, 0, (len(slice)+chunkSize-1)/chunkSize)
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, slice)

	return chunks, nil
}
