package main

import (
	"os"
	"bufio"
	"github.com/schwarmco/go-cartesian-product"
	"sync"
	"fmt"
	"strconv"
	"flag"
	"log"
	"runtime"
"runtime/pprof"
	"net/http/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")


func main() {

flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }
	
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
	filterSizes := []interface{}{4792}
	sampleSizes := []interface{}{300000}
	numSites := 50000
	primMod := []interface{}{0.02}
	secMod := []interface{}{0.1}
	ps := []interface{}{0.000001,0.00001,0.0001}
	qs := []interface{}{0.9}
	numsHashes := []interface{}{7}

	// Get the cartesian product, i.e. all possible combinations of the parameters
	prm := cartesian.Iter(filterSizes, sampleSizes, ps, qs, numsHashes, primMod, secMod)

	// Result is a channel, draw all items from it to make it a slice
	perms := make([]interface{},len(prm))
	for x := range prm{
		perms = append(perms,x)
	}

	// Divide the parameters in chunks for the array job
	numJobs := 1
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


	if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}

func chunkSlice(slice []interface{}, chunkSize int) ([][]interface{}, error) {

	chunks := make([][]interface{}, 0, (len(slice)+chunkSize-1)/chunkSize)
	for chunkSize < len(slice) {
		slice, chunks = slice[chunkSize:], append(chunks, slice[0:chunkSize:chunkSize])
	}
	chunks = append(chunks, slice)

	return chunks, nil
}
