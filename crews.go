package main

type Crews struct {
	primary *BloomFilter    // filter data
	secondary *BloomFilter
	primaryThreshold float64
	secondaryThreshold float64
	p float64
	q float64
}

func NewCrews(filterSize int, numHashes int, primaryThreshold float64, secondaryThreshold float64, p float64, q float64) *Crews {
	return &Crews{
		primary: NewBloomFilter(uint(filterSize), uint(numHashes)),
		secondary:   NewBloomFilter(uint(filterSize), uint(numHashes)),
		primaryThreshold: primaryThreshold,
		secondaryThreshold: secondaryThreshold,
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
	return c.primary.Test([]byte(s))[0] >= uint(c.primaryThreshold * float64(c.primary.count))
}

func (c *Crews) SecondaryTest(s string) bool {
	return c.secondary.Test([]byte(s))[0] >= uint((c.secondaryThreshold * float64(c.secondary.count)))
}

