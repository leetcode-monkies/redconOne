package problem0011

func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for left < right {
		currArea := getArea(left, right, height)
		if currArea > max {
			max = currArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func getArea(left, right int, arr []int) int {
	width := right - left
	height := 0
	if arr[left] < arr[right] {
		height = arr[left]
	} else {
		height = arr[right]
	}
	return width * height
}
