package main

type Crews struct {
	primary *BloomFilter    // filter data
	secondary *BloomFilter
	numWebsites uint
	primaryThresholdModifier float64
	secondaryThresholdModifier float64
	p float64
	q float64
}

func NewCrews(filterSize int, numHashes int, numWebsites uint, primaryThresholdModifier float64,	secondaryThresholdModifier float64, p float64, q float64) *Crews {
	return &Crews{
		primary: NewBloomFilter(uint(filterSize), uint(numHashes)),
		secondary:   NewBloomFilter(uint(filterSize), uint(numHashes)),
		numWebsites: numWebsites,
		primaryThresholdModifier: primaryThresholdModifier,
		secondaryThresholdModifier: secondaryThresholdModifier,
		p: p,
		q: q,
	}
}

func (c *Crews) ReportHsts(s string) *Crews {
	c.primary.Add([]byte(s), c.p, c.q)
	return c
}

func (c *Crews) ReportHttp(s string) *Crews {
	c.secondary.Add([]byte(s), c.p, c.q)
	return c
}

func (c *Crews) PrimaryTest(s string) bool {
	count := c.primary.Test([]byte(s))
	adjustedCount := uint((float64(count) - c.p * float64(c.primary.count))/(c.q-c.p))
	return  adjustedCount >= uint(( c.primary.count/c.numWebsites)*c.primaryThresholdModifier)
}

func (c *Crews) SecondaryTest(s string) bool {
	count := c.secondary.Test([]byte(s))
	adjustedCount := uint((float64(count) - c.p * float64(c.secondary.count))/(c.q-c.p))
	return  adjustedCount >= uint((c.secondary.count/c.numWebsites)*c.secondaryThresholdModifer)
}

