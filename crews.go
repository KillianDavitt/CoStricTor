package main

type Crews struct {
	primary *BloomFilter    // filter data
	secondary *BloomFilter
	primary_threshold float64
	secondary_threshold float64
}

func NewCrews(filter_size int, numHashes int) *Crews {
	return &Crews{
		primary: NewBloomFilter(uint(filter_size), uint(numHashes)),
		secondary:   NewBloomFilter(uint(filter_size), uint(numHashes)),
		primary_threshold: 0.001,
		secondary_threshold: 0.001,
	}
}

func (c *Crews) ReportHsts(s string) *Crews {
	c.primary.Add([]byte(s))
	return c
}

func (c *Crews) ReportHttp(s string) *Crews {
	c.secondary.Add([]byte(s))
	return c
}

func (c *Crews) PrimaryTest(s string) bool {
	return c.primary.Test([]byte(s))[0] >= uint(c.primary_threshold * float64(c.primary.count))
}

func (c *Crews) SecondaryTest(s string) bool {
	return c.secondary.Test([]byte(s))[0] >= uint((c.secondary_threshold * float64(c.secondary.count)))
}

