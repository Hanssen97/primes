package primes

import (
	"math"
	"runtime"
	"sort"
	"sync"

	"github.com/jorgenhanssen/primes/cli"
)

// MaxSafeInt (2^53) is the largest integer accuratly
// represented by a float64.
const MaxSafeInt = 9007199254740992

type Instance struct {
	base []int
	Program *cli.Instance
	numCPUs int
}

func New() *Instance {
	return &Instance{
		base: []int{2, 3, 4, 5, 7},
		numCPUs: runtime.NumCPU(),
	}
}

func (p *Instance) FindPrimes(start, end int) (result []int) {
	threshold := calculateThreshold(end)
	if (threshold > p.base[len(p.base)-1]) {
    	p.extendBase(threshold);
	}
	
	if (start < 3) {
		result = []int{2};
		start = 3; 
	} else if isDivisible(start, 2) {
		start++;
	}

	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	
	wg.Add(p.numCPUs)
	for i := 0; i < p.numCPUs; i++ {
		go func (i int)  {
			defer wg.Done()
			
			_result := []int{}
			// workers are distributed from start until num workers on all odd numbers.
			// the stride is 2 * num workers (odd numbers).
			for num := start + (2*i); num < end; num += 2 * p.numCPUs {
				if p.isPrime(num) {
					_result = append(_result, num)
				}
			}

			mu.Lock()
			result = append(result, _result...)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	sort.Ints(result)

	return
}

func (p *Instance) isPrime(number int) bool {
	if number == 3 {
		return true
	}
	
	threshold := calculateThreshold(number)
	for _, compliment := range p.base {
		if isDivisible(number, compliment) {
			return false
		}
		if compliment > threshold {
			return true
		}
	}

	return true;
}

func (p *Instance) extendBase(end int) {
	start := p.base[len(p.base)-1];
	extension := p.FindPrimes(start + 2, end + 2)
	p.base = append(p.base, extension...)
}


func calculateThreshold(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}

func isDivisible(a, b int) bool {
	if a < MaxSafeInt {
		// Use this faster method when the value can be
		// represented accuratly as float
		quotient := float64(a) / float64(b)
		return quotient == math.Trunc(quotient)
	}

	// Accurate but slower approach for larger integers
	return a % b == 0
}