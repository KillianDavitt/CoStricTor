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
	sitesToCheck := 1000000
	//primaryThresh := [16]float64{0.1,0.2,0.3,0.09,0.08,0.07,0.06,0.05,0.04,0.03,0.02,0.1,0.001,0.005,0.0005,0.0001}
	filterSizes := [1]int{20000} //[60]int{500,1000,1500,2000,2500,3000,3500,4000,4500,5000,5500,6000,6500,7000,7500,8000,8500,9000,9500,10000,10500,11000,11500,12000,12500,13000,13500,14000,14500,15000,15500,16000,16500,17000,17500,18000,18500,19000,19500,20000,20500,21000,21500,22000,22500,23000,23500,24000,24500,25000,25500,26000,26500,27000,27500,28000,28500,29000,29500,30000}//,30500,31000,31500,32000,32500,33000,33500,34000,34500,35000,35500,36000,36500,37000,37500,38000,38500,39000,39500,40000,40500,41000,41500,42000,42500,43000,43500,44000,44500,45000,45500,46000,46500,47000,47500,48000,48500,49000,49500,50000}
	
	//secondaryThresholds := [6]float64{0.01,0.02,0.03,0.04,0.05,0.06}
	//secondaryThreshs := [16]float64{0.1,0.2,0.3,0.09,0.08,0.07,0.06,0.05,0.04,0.03,0.02,0.1,0.001,0.005,0.0005,0.0001}
	ps := [18]float64{0.1,0.2,0.3,0.4,0.5,0.6,0.7,0.8,0.9,0.0023393011164896805, 0.003010075739861694, 0.0051407670922724216, 0.009817666458252727, 0.01981333734650907, 0.040984938304053545, 0.08463610077414542, 0.16874682568157592, 0.31016900541211584}
	qs := [18]float64{0.9988905440229229, 0.9986599476383371, 0.9978767930134134, 0.9962956450020221, 0.9933071490757152, 0.9879368700291202, 0.9791084544732513, 0.9672734436346139, 0.9607296994499493,0.95,0.9,0.85,0.8,0.75,0.7,0.65,0.6,0.55,0.5}
	//for i:=1; i<200; i++ {
	//	ps[i]=0.0000001 * float64(i)
	//}
	
	numSites := 50000
	numJobs := 18

	var perms []interface{};
	for i:=0; i<len(filterSizes); i++ {
		for j:=0; j<len(ps); j++ {
			p := []interface{}{filterSizes[i],3000000,ps[j],qs[j],0.045,0.07}
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
