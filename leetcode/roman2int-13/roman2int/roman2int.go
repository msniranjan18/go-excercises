package roman2int

import (
	"fmt"
)

var (
	Roman2IntMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

func Roman2Int(romanNo string) int {
	fmt.Println("roman:", romanNo)
	result := 0
	for i := len(romanNo) - 1; i >= 0; i-- {
		currVal := Roman2IntMap[romanNo[i]]
		if i != len(romanNo)-1 {
			oldVal := Roman2IntMap[romanNo[i+1]]
			if currVal < oldVal {
				currVal = currVal * -1
			}
		}
		result += currVal
	}
	return result
}

func Roman2Intv2(romanNo string) int {
	fmt.Println("roman:", romanNo)
	result := 0
	prevVal := 0
	for i := len(romanNo) - 1; i >= 0; i-- {
		val := Roman2IntMap[romanNo[i]]

		if val < prevVal {
			result += -val
		} else {
			result += val
		}
		prevVal = val
	}
	return result
}
