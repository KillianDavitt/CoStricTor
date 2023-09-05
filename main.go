package main

import (
	"os"
	"bufio"
	"sync"
	"fmt"
	"strconv"
	"strings"
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

	dat, err := os.ReadFile("params.csv")
	if err  != nil {
		fmt.Println(err)
	}
	config_lines := strings.Split(string(dat), "\n")
	config := make(map[string]float64)
	for i:=0; i<len(config_lines)-1;i++ {
		line := strings.Split(config_lines[i],",")
		config[line[0]], err = strconv.ParseFloat(line[1],64)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error reading config file")
		}
	}

	
	// Sort out parameters
	hstsProp := config["hstsProp"]
	httpProp := config["httpProp"]
	sitesToCheck := int(config["sitesToCheck"])
	primaryThresh := [1]float64{config["primaryThresholds"]}
	secondaryThreshs := [1]float64{config["secondaryThreshs"]}
	filterSizeStart := int(config["filterSizeStart"])
	filterSizeEnd := int(config["filterSizeEnd"])
	filterSizeStep := int(config["filterSizeStep"])
	numFilterSizes := (filterSizeEnd/filterSizeStep) - (filterSizeStart/filterSizeStep) + 1
	var filterSizes = make([]int, numFilterSizes)
	for i:=0; i< numFilterSizes; i++ {
		filterSizes[i] = filterSizeStep * (i+1)
	}

	

	ps := [1]float64{config["ps"]}
	qs := [1]float64{config["qs"]}

	numSites := 50000
	numJobs := int(config["numJobs"])
	numSubmissions := int(config["numSubmissions"]) //90000000

	var perms []interface{};
	for i:=0; i<len(filterSizes); i++ {
		for j:=0; j<len(ps); j++ {
			p := []interface{}{filterSizes[i],numSubmissions,ps[j],qs[0],primaryThresh[0],secondaryThreshs[0]}
			perms = append(perms,p)
		}
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
		jobNumber=1
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
