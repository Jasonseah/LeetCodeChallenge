package challenge_1431

/**
a solution attempt by me
*/
func KidsWithCandies(candies []int, extraCandies int) []bool {
	var highestValue = 0
	var candiesResult = make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		if candies[i] > highestValue {
			highestValue = candies[i]
		}
	}

	for key, value := range candies {
		var candiesWithExtra = value + extraCandies
		candiesResult[key] = false
		if candiesWithExtra >= highestValue {
			candiesResult[key] = true
		}
	}

	return candiesResult
}
