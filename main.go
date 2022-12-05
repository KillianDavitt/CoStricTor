package main

import (
	"os"
	"bufio"
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
	hstsProp := 0.2
	httpProp := 0.2
	sitesToCheck := 100000
	filterSizes := [14]int{5000,7500,2500,2000,1800,1700,1600,1500,1400,1300,1200,1100,1000}
	//secondaryThresholds := [6]float64{0.01,0.02,0.03,0.04,0.05,0.06}

	numSites := 1000
	numJobs := 1

	perms := make([]interface{},0)
	for i:=0; i<len(filterSizes); i++ {
		p := []interface{}{filterSizes[i],300000,0.0,1.0,0.004,0.00335}
		perms = append(perms,p)
	}
	// Divide the parameters in chunks for the array job
	
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
	checkSites := lines[numSites:sitesToCheck]
	hsts, http, https_no_hsts := generateSites(sites, hstsProp, httpProp);
	checkHsts, checkHttp, checkHttpsNoHsts := generateSites(checkSites, hstsProp, httpProp);
        var wg sync.WaitGroup
	for _,params := range jobs[jobNumber-1] {
		wg.Add(1)
		go runSim(params.([]interface{}), hsts, http, https_no_hsts, checkHsts, checkHttp, checkHttpsNoHsts, &wg, hstsProp, httpProp, numSites)
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
