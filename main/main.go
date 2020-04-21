// https://www.hackerrank.com/challenges/icecream-parlor/problem

package main

import (
	"fmt"
	"sort"
)

func contains(items []int, item int) bool {
	for _, anItem := range items {
		if anItem == item {
			return true
		}
	}

	return false
}

func icecreamParlor(m int, arr []int) []int {
	var iceMap = make(map[int][]int)
	var prices = make([]int, 0)
	var result = make([]int, 2)

	// First, let's put all the available ice cream to the map.
	// The map stores the price and it's 1-base index
	for idx, price := range arr {
		idxList, exist := iceMap[price]

		if !exist {
			idxList = make([]int, 0)

			if !contains(prices, price) {
				// Only store the price if it's not yet exist in the price list
				prices = append(prices, price)
			}
		}

		iceMap[price] = append(idxList, idx+1)
	}

	// Sort the prices
	sort.Ints(prices)

	// Now, let's calculate the price combination
	for _, price := range prices {
		var price1 = price
		var price2 = m - price1

		if price1 == price2 {
			// Price1 and Price2 are the same. Let's check the availability of the ice cream
			if len(iceMap[price1]) > 1 {
				var indices = iceMap[price1]
				result[0] = indices[0]
				result[1] = indices[1]

				break
			}
		} else {
			// Now, check if price2 exists
			_, exist := iceMap[price2]

			if exist {
				// Price2 is found.
				var indices1 = iceMap[price1]
				var indices2 = iceMap[price2]

				result[0] = indices1[0]
				result[1] = indices2[0]

				break
			}
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	fmt.Println(icecreamParlor(4, []int{1, 4, 5, 3, 2}))
	fmt.Println(icecreamParlor(4, []int{2, 2, 4, 3}))
}
