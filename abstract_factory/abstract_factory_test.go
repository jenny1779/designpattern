package abstract_factory

import "testing"

func TestNewWaiterFactory(t *testing.T) {
	factory := NewWaiterFactory()
	customer := factory.CookFood()
	customer.ChooseWaiter()
	leader := factory.CleanRoom()
	leader.ChooseWaiter()
}

