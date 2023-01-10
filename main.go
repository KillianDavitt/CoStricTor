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
	//primaryThresh := [16]float64{0.1,0.2,0.3,0.09,0.08,0.07,0.06,0.05,0.04,0.03,0.02,0.1,0.001,0.005,0.0005,0.0001}
	filterSizes := [1]int{60000}//500,1000,1500,2000,2500,3000,3500,4000,4500,5000,5500,6000,6500,7000,7500,8000}
	
	//secondaryThresholds := [6]float64{0.01,0.02,0.03,0.04,0.05,0.06}
	//secondaryThreshs := [16]float64{0.1,0.2,0.3,0.09,0.08,0.07,0.06,0.05,0.04,0.03,0.02,0.1,0.001,0.005,0.0005,0.0001}
	numSites := 15000
	numJobs := 1

	var perms []interface{};
	for i:=0; i<len(filterSizes); i++ {
		//for j:=0; j<len(secondaryThreshs); j++ {
			p := []interface{}{filterSizes[i],3000000,0.000001,0.9,0.045,0.07}
			perms = append(perms,p)
			//}
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
	// The CS TSG cluster sets this environment variable depending on how many jobs are being run simultaneously
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
