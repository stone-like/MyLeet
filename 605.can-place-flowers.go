/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

// @lc code=start

//こっちはO(N)
func canPlace(flowerbed []int, expectCount int) bool {
	count := 0

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		count = 1
		return expectCount <= count
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		leftOK := false
		rightOK := false

		if i-1 >= 0 && flowerbed[i-1] == 0 {
			leftOK = true
		}

		if i+1 <= len(flowerbed)-1 && flowerbed[i+1] == 0 {
			rightOK = true
		}

		if i == 0 && rightOK {
			//rightのみ
			flowerbed[i] = 1
			count++
			continue
		}

		if i == len(flowerbed)-1 && leftOK {
			//leftのみ
			flowerbed[i] = 1
			count++
			continue
		}

		if leftOK && rightOK {
			flowerbed[i] = 1
			count++
			continue
		}
	}

	return expectCount <= count
}
func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	return canPlace(flowerbed, n)
}

//下記は愚直解、O(N^2)
// func canPlace(flowerbed []int) bool {

// 	if len(flowerbed) == 1 && flowerbed[0] == 0 {
// 		flowerbed[0] = 1
// 		return true
// 	}

// 	for i := 0; i < len(flowerbed); i++ {
// 		if flowerbed[i] == 1 {
// 			continue
// 		}

// 		leftOK := false
// 		rightOK := false

// 		if i-1 >= 0 && flowerbed[i-1] == 0 {
// 			leftOK = true
// 		}

// 		if i+1 <= len(flowerbed)-1 && flowerbed[i+1] == 0 {
// 			rightOK = true
// 		}

// 		if i == 0 && rightOK {
// 			//rightのみ
// 			flowerbed[i] = 1
// 			return true
// 		}

// 		if i == len(flowerbed)-1 && leftOK {
// 			//leftのみ
// 			flowerbed[i] = 1
// 			return true
// 		}

// 		if leftOK && rightOK {
// 			flowerbed[i] = 1

// 			return true
// 		}
// 	}

// 	return false
// }
// func canPlaceFlowers(flowerbed []int, n int) bool {

// 	if n == 0 {
// 		return true
// 	}

// 	ok := false

// 	for i := 0; i < n; i++ {
// 		if canPlace(flowerbed) {
// 			ok = true
// 		} else {
// 			ok = false
// 		}
// 	}

// 	return ok
// }

// @lc code=end

