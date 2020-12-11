package primes

import "testing"

func testNumPrimesInRange(p *Instance, t *testing.T, start, end, expected int) {
	numPrimes := len(p.FindPrimes(start, end))
	if numPrimes != expected {
		t.Errorf("[%d - %d] expected %d; found %d", start, end, expected, numPrimes)
	} else {
		t.Logf("ok [%d - %d]", start, end)
	}
}

func TestCorrectNumberOfPrimesInRanges(t *testing.T) {
	// from https://primes.utm.edu/howmany.html

	p := New()
	testNumPrimesInRange(p, t, 1, 10, 4)
	testNumPrimesInRange(p, t, 1, 100, 25)
	testNumPrimesInRange(p, t, 1, 1000, 168)
	testNumPrimesInRange(p, t, 1, 10000, 1229)
	testNumPrimesInRange(p, t, 1, 100000, 9592)
	testNumPrimesInRange(p, t, 1, 1000000, 78498)
	testNumPrimesInRange(p, t, 1, 10000000, 664579)
	testNumPrimesInRange(p, t, 1, 100000000, 5761455)
}