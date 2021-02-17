package randvar

import "testing"

func TestRoleDiceProb(t *testing.T) {
	diceRole := New("rolling a dice", []string{"one", "two", "three", "four", "five", "six"})

	prob, err := diceRole.Probability("two")

	if err != nil {
		t.Errorf("Event two is not recornized")
	}

	if prob != 1.0/6 {
		t.Error("Rolling the dice two has probability != from 1/6")
	}
}
