package main

import "fmt"

func main() {
	park := Constructor(1, 1, 0)
	res := park.AddCar(1)
	fmt.Println(res)
	res = park.AddCar(2)
	fmt.Println(res)
	res = park.AddCar(3)
	fmt.Println(res)
	res = park.AddCar(1)
	fmt.Println(res)
}

type ParkingSystem struct {
	park [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{park: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.park[carType-1] >= 1 {
		this.park[carType-1] = this.park[carType-1] - 1
		return true
	} else {
		return false
	}
}
