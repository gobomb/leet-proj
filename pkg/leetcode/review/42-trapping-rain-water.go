package review

// 42. trapping-rain-water

func trap(height []int) int {
	return trapBetween(height, 0, len(height)-1)
}
func trapBetween(height []int, left, right int) (rs int) {
	if right-left <= 1 {
		return 0
	}
	fi, si := left, left
	h1, h2 := 0, 0
	// 1. find highest and second highest index
	for i := left; i <= right; i++ {
		if height[i] >= h1 {
			si = fi
			h2 = h1
			fi = i
			h1 = height[i]
		}
		if height[i] > h2 && height[i] < h1 {
			h2 = height[i]
			si = i
		}
	}
	if fi == si {
		return 0
	}
	// fmt.Printf("%v %v\n", h1, h2)
	// 2. recursive do 1
	leftTrap := trapBetween(height, left, min(fi, si))
	rightTrap := trapBetween(height, max(fi, si), right)

	// 3. sub the other bars between highest and second highest bars
	subed := subOtherBars(height, fi, si)
	// fmt.Println(h1, h2, left, right, fi, si, subed)

	return leftTrap + rightTrap + subed
}

func subOtherBars(height []int, h, l int) int {
	var sub int
	// if h==l ?
	for i := min(h, l) + 1; i < max(h, l); i++ {
		sub += height[i]
	}

	w := h - l
	if w < 0 {
		w = -w
	}

	return height[l]*(w-1) - sub
}

func trap2(height []int) int {
	maxH := 0
	dp1 := make([]int, len(height))
	dp2 := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		maxH = max(maxH, height[i])
		dp1[i] = maxH
	}

	maxH = 0

	for j := len(height) - 1; j >= 0; j-- {
		maxH = max(maxH, height[j])
		dp2[j] = maxH
	}

	wt := 0
	for i := 1; i < len(height)-1; i++ {
		h := min(dp1[i], dp2[i])
		wt += h - height[i]
	}

	return wt
}
