package randvar

import "fmt"

//RandVar is a random variable
type RandVar struct {
	Description         string
	isEqualDistribution bool
	domainProbability   map[string]float32
}

//New creates a new random variable
func New(description string, domain []string) *RandVar {
	domainLen := len(domain)
	domainProbability := make(map[string]float32, domainLen)
	probability := float32(1 / domainLen)

	for i := 0; i < domainLen; i++ {
		domainProbability[domain[i]] = probability
	}

	return &RandVar{
		Description:         description,
		isEqualDistribution: true,
		domainProbability:   domainProbability,
	}
}

//NewWithDistribution creates a random variable with unequal distribution
func NewWithDistribution(description string, domain []string, distribution []float32) (*RandVar, error) {
	if err := veriftyDistributionIsLogical(description, distribution); err != nil {
		return nil, err
	}

	domainLen := len(domain)
	domainProb := make(map[string]float32, domainLen)
	for i := 0; i < domainLen; i++ {
		domainProb[domain[i]] = distribution[i]
	}

	return &RandVar{
		Description:         description,
		isEqualDistribution: false,
		domainProbability:   domainProb,
	}, nil
}

//Probability calculates the probanility of a given event
func (r RandVar) Probability(event string) (float32, error) {
	if prob, contains := r.domainProbability[event]; contains {
		return prob, nil
	}
	return 0, fmt.Errorf("event: %s not in domain", event)
}
