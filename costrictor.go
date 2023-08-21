package main


type Costrictor struct {
	primary *BloomFilter    // filter data
	secondary *BloomFilter
	numWebsites uint
	primaryThresholdModifier float64
	secondaryThresholdModifier float64
	p float64
	q float64
}

func NewCostrictor(filterSize int, numHashes int, numWebsites uint, primaryThresholdModifier float64, secondaryThresholdModifier float64, p float64, q float64) *Costrictor {
	return &Costrictor{
		primary: NewBloomFilter(uint(filterSize), uint(numHashes)),
		secondary:   NewBloomFilter(uint(filterSize), uint(numHashes)),
		numWebsites: numWebsites,
		primaryThresholdModifier: primaryThresholdModifier,
		secondaryThresholdModifier: secondaryThresholdModifier,
		p: p,
		q: q,
	}
}

func (c *Costrictor) ReportHsts(s string) *Costrictor {
	c.primary.Add([]byte(s), c.p, c.q)
	return c
}

func (c *Costrictor) ReportHttp(s string) *Costrictor {
	c.secondary.Add([]byte(s), c.p, c.q)
	return c
}

func (c *Costrictor) PrimaryTest(s string) bool {
	count := float64(c.primary.Test([]byte(s)))
	numInsertions := float64(c.primary.count)
	thresholdMod := c.primaryThresholdModifier
	numWebsites := float64(c.numWebsites)
	p := c.p
	q := c.q
	
	adjustedCount := (count - p * numInsertions)/(q-p)
	threshold := float64(numInsertions/numWebsites) * thresholdMod
	return  adjustedCount >= threshold
}

func (c *Costrictor) SecondaryTest(s string) bool {
	count := float64(c.secondary.Test([]byte(s)))
	numInsertions := float64(c.secondary.count)
	thresholdMod := c.secondaryThresholdModifier
	numWebsites := float64(c.numWebsites)
	p := c.p
	q := c.q
	
	adjustedCount := (count - p * numInsertions)/(q-p)
	threshold := float64(numInsertions/numWebsites) * thresholdMod
	return  adjustedCount >= threshold
}

