package main

import (
	"math/rand"
	"sync"
	"time"
	"fmt"
)

func generateSites(sites []string, hstsProp float64, httpProp float64) ([]string,[]string,[]string) {
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
	return hsts,http,https_no_hsts
}

func runSim(prms []interface{}, hsts []string, http []string, https_no_hsts []string,  wg * sync.WaitGroup){
	filterSize := prms[0].(int)
	numSamples := prms[1].(int)
	numSites := prms[2].(int)
	numHashes := prms[9].(int)
	//hstsProp := prms[3].(float64)
	//httpProp := prms[4].(float64)
	var primaryThreshold float64 = prms[5].(float64)
	var secondaryThreshold float64 = prms[6].(float64)
	var p float64 = prms[7].(float64)
	var q float64 = prms[8].(float64)

	
	


	c := NewCrews(filterSize, numHashes, primaryThreshold, secondaryThreshold, p, q);
	
	source := rand.NewSource(time.Now().UnixNano()) 
	hsts_zipf := rand.NewZipf(rand.New(source), 1.1, 9999.0, uint64(len(hsts)-1))

	// Sample n sites to report to crews, these can and will be duplicates
	for i:=0; i<numSamples; i++ {
		_ = hsts_zipf.Uint64()
		n := int(rand.Uint32()) % len(hsts)
		c.ReportHsts(hsts[n])
	}

	http_zipf := rand.NewZipf(rand.New(source), 1.1, 4.0, uint64(len(http)-1))
	for i:=0; i<numSamples; i++ {
		n := http_zipf.Uint64()
		c.ReportHttp(http[n]);
	}
	
	var disasters uint = 0
	var final_benefit uint = 0
	var no_benefit uint = 0
	var initial_true_hsts uint = 0
	var disasters_averted uint = 0
	var accidental_upgrades uint = 0
	var accidental_upgrades_averted uint = 0

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
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d,%f,%f,%f,%f,%d\n",len(hsts), final_benefit,disasters, initial_true_hsts, filterSize, numSamples, numSites, primaryThreshold, secondaryThreshold, p,q, numHashes)
	defer wg.Done()
}

