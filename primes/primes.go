package primes

import (
	"math"
	"runtime"
	"sort"
	"sync"

	"github.com/jorgenhanssen/primes/cli"
)

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

func (p *Instance) findPrimes(start, end int) (result []int) {
	threshold := calculateThreshold(end)
	if (threshold > p.base[len(p.base)-1]) {
    	p.extendBase(threshold);
	}
	
	if (start < 3) {
		result = []int{2};
		start = 3; 
	} else if start % 2 == 0 {
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
	
	cap := int(math.Ceil(math.Sqrt(float64(number))));
	for i := 0; i < len(p.base); i++ {
		compliment := p.base[i]
		divided := float64(number) / float64(compliment)
		if divided == float64(int64(divided)) {
			return false
		}
		if compliment > cap {
			return true
		}
	}

	return true;
}

func (p *Instance) extendBase(end int) {
	start := p.base[len(p.base)-1];
	extension := p.findPrimes(start + 2, end + 2)
	p.base = append(p.base, extension...)
}


func calculateThreshold(number int) int {
	return int(math.Ceil(math.Sqrt(float64(number))))
}