package leetcode

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor2(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 3:
		if ps.small != 0 {
			ps.small--
			return true
		}
	case 2:
		if ps.medium != 0 {
			ps.medium--
			return true
		}
	case 1:
		if ps.big != 0 {
			ps.big--
			return true
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
