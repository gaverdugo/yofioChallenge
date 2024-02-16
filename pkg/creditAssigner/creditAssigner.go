package creditassigner

import (
	"errors"
	"math"
)

type CreditAssigner struct{}

func (CreditAssigner) Assign(investment int32) (int32, int32, int32, error) {
	var assignedCredit [3]int32
	amount := investment / 100
	availableCredit := [3]int32{3, 5, 7}

	dp := make([]int32, amount+1)
	lastUsedCredit := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 1; i <= int(amount); i++ {
		for j, credit := range availableCredit {
			if int(credit) <= i {
				res := dp[i-int(credit)]
				if res != math.MaxInt32 && res+1 < dp[i] {
					dp[i] = res + 1
					lastUsedCredit[i] = j
				}
			}
		}
	}

	cur := amount

	for cur > 0 {
		assignedCredit[lastUsedCredit[cur]]++
		cur -= availableCredit[lastUsedCredit[cur]]
	}

	if dp[amount] < math.MaxInt16 {
		return assignedCredit[0], assignedCredit[1], assignedCredit[2], nil
	} else {
		return 0, 0, 0, errors.New("can't find a valid credit combination")
	}
}
