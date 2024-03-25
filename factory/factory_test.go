package factory

import (
	"testing"
)

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("qing feng")
	NewRestaurant("qingfeng")
	NewRestaurant(("donglaishun"))
}