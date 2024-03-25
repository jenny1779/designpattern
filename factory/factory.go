package factory

import "fmt"

type restaurant interface {
	GetFood()
}

type dongLaiShun struct {}

func (d *dongLaiShun) GetFood()  {
	fmt.Println("processing by dong lai shun")
}

type qingFeng struct {}

func (q *qingFeng) GetFood()  {
	fmt.Println("processing by qing feng")
}

func NewRestaurant(name string) restaurant {
	switch name {
	case "qingfeng":
		return &qingFeng{}
	case "donglaishun":
		return &dongLaiShun{}
	}
	return nil
}

