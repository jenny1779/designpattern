package facade

import "testing"

func TestCreateCompleteCar(t *testing.T) {
	car := NewCarFacade()
	car.CreateCompleteCar()
}
