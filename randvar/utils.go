package randvar

import "errors"

// VerifyDist verifies if distribution sum is 1
func VerifyDist(event string, distribution map[string]float64) error {
	var sum float64 = 0.0
	for _, d := range distribution {
		sum += d
	}
	if sum != 1.0 {
		return errors.New("Distribution sum of %s is not equal to 1")
	}
	return nil
}
