package randvar

import "fmt"

// RandVar is a random variable
type RandVar struct {
	Event        string
	isEqualDist  bool
	probabilites map[string]float64
}

// New creates a new random variable
func New(event string, domain []string) *RandVar {
	ln := len(domain)
	domProb := make(map[string]float64, ln)
	prob := 1 / float64(ln)

	for _, d := range domain {
		domProb[d] = prob
	}

	return &RandVar{
		Event:        event,
		isEqualDist:  true,
		probabilites: domProb,
	}
}

// NewWithDist creates a random variable with unequal distribution
func NewWithDist(event string, domDist map[string]float64) (*RandVar, error) {
	if err := VerifyDist(event, domDist); err != nil {
		return nil, err
	}

	ln := len(domDist)
	domProb := make(map[string]float64, ln)
	for e, d := range domDist {
		domProb[e] = d
	}

	return &RandVar{
		Event:        event,
		isEqualDist:  false,
		probabilites: domProb,
	}, nil
}

// Probability calculates the probanility of a given event
func (r RandVar) Probability(event string) (float64, error) {
	if prob, ok := r.probabilites[event]; ok {
		return prob, nil
	}
	return -1, fmt.Errorf("event: %s not in domain", event)
}
