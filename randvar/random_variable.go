package randvar

import "fmt"

//RandVar is a random variable
type RandVar struct {
	Event        string
	isEqualDist  bool
	probabilites map[string]float32
}

// New creates a new random variable
func New(event string, domain []string) *RandVar {
	ln := len(domain)
	domProb := make(map[string]float32, ln)
	prob := float32(1 / ln)

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
func NewWithDist(event string, domain []string, dist []float32) (*RandVar, error) {
	if err := VerifyDist(event, dist); err != nil {
		return nil, err
	}

	ln := len(domain)
	domProb := make(map[string]float32, ln)
	for i, d := range domain {
		domProb[d] = dist[i]
	}

	return &RandVar{
		Event:        event,
		isEqualDist:  false,
		probabilites: domProb,
	}, nil
}

//Probability calculates the probanility of a given event
func (r RandVar) Probability(event string) (float32, error) {
	if prob, ok := r.probabilites[event]; ok {
		return prob, nil
	}
	return 0, fmt.Errorf("event: %s not in domain", event)
}
