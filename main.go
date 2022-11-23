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
	filterSizes := []interface{}{50000,55000,60000,65000,70000}
	sampleSizes := []interface{}{3000000}
	numSites := 10000
	primaryThresholds := []interface{}{0.00005,0.00001,0.000005,0.000001}
	secondaryThresholds := []interface{}{0.00005,0.00001,0.000005,0.000001}
	ps := []interface{}{0.000001,0.01,0.1}
	qs := []interface{}{0.9}
	numsHashes := []interface{}{1,4,6,10,16,20}

	// Get the cartesian product, i.e. all possible combinations of the parameters
	prm := cartesian.Iter(filterSizes, sampleSizes, primaryThresholds, secondaryThresholds, ps, qs, numsHashes)

	// Result is a channel, draw all items from it to make it a slice
	perms := make([]interface{},len(prm))
	for x := range prm{
		perms = append(perms,x)
	}

	// Divide the parameters in chunks for the array job
	numJobs := 1000
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
