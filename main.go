package main

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"math/rand"
)

func main() {

	// Params
	filterSize := 4096
	numSamples := 2000
	numSites := 1000
	numHashes := 1
	hstsProp := 0.2
	httpProp := 0.2

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
	sites := lines[0:numSites]

	// We will be destroying the main list, copy it for use later
	var total_sites []string = make([]string, len(sites))
	_ = copy(total_sites, sites)

	
	// Generate n sites which have hsts
	hsts := make([]string, 0)
	for i:=0; i<len(sites); i++ {
		p := rand.Float64()
		if p<hstsProp {
			hsts = append(hsts, sites[i]);
			// Remove the chosen site from the main list
			// So it can't be chosen as a http site or a https_no_hsts site
			sites[i] = sites[len(sites)-1]
			sites = sites[:len(sites)-1]
		}
	}
	// Generate n sites which don't have https
	http := make([]string, 0)
	for i:=0; i<len(sites); i++ {
		p := rand.Float64()
		if p<httpProp {
			http = append(http, sites[i]);
			sites[i] = sites[len(sites)-1]
			sites = sites[:len(sites)-1]
		}
	}

	// All remaining sites have https but not hsts
	https_no_hsts := sites
	
	fmt.Printf("HSTS: %d\n",len(hsts))
	fmt.Printf("http only: %d\n",len(http))
	fmt.Printf("https_no_hsts: %d\n",len(https_no_hsts))

	c := NewCrews(filterSize, numHashes);
	
	source := rand.NewSource(time.Now().UnixNano()) 
	hsts_zipf := rand.NewZipf(rand.New(source), 1.1, 2.0, uint64(len(hsts)-1))

	// Sample n sites to report to crews, these can and will be duplicates
	for i:=0; i<numSamples; i++ {
		n := hsts_zipf.Uint64()
		go c.ReportHsts(hsts[n])
	}

	http_zipf := rand.NewZipf(rand.New(source), 1.1, 2.0, uint64(len(http)-1))
	for i:=0; i<numSamples; i++ {
		n := http_zipf.Uint64()
		go c.ReportHttp(http[n]);
	}
	
	var disasters uint = 0
	var final_benefit uint =0
	var no_benefit uint =0
	var initial_true_hsts uint =0
	var disasters_averted uint =0
	var accidental_upgrades uint =0
	var accidental_upgrades_averted uint =0

	// Iterate through our 3 lists of sites. Each one has 2 potential oucomes
	
	for i:=0; i<len(hsts);i++ {
		if c.PrimaryTest(hsts[i]) {
			initial_true_hsts += 1
			if c.SecondaryTest(hsts[i]) {
				no_benefit += 1
			} else {
				final_benefit +=1
			}
		}
	}

	for i:=0; i<len(http);i++ {
		if c.PrimaryTest(http[i]) {
			if c.SecondaryTest(http[i]) {
				disasters_averted += 1
			} else {
				disasters += 1
			}
		}
	}

	for i:=0; i<len(https_no_hsts);i++ {
		if c.PrimaryTest(https_no_hsts[i]) {
			if c.SecondaryTest(https_no_hsts[i]) {
				accidental_upgrades += 1
			} else {
				accidental_upgrades_averted += 1
			}
		}
	}
	
	fmt.Printf("Initial True HSTS: %d\n",initial_true_hsts)
	fmt.Printf("Disasters Averted: %d\n", disasters_averted)
	fmt.Printf("No Benefit: %d\n", no_benefit)   

	fmt.Printf("Accidental upgrades averted: %d\n", accidental_upgrades_averted)
	fmt.Printf("Accidental Upgrades: %d\n", accidental_upgrades)

	fmt.Println("-------")
	fmt.Printf("Final Benefit: %d\n", final_benefit)
	fmt.Printf("Disasters: %d\n", disasters)

}

