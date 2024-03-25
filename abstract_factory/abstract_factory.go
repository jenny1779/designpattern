package abstract_factory

import "fmt"

/*代码定义了一个接口 Dinner 和两个结构体 Customer 和 Leader，
它们实现了 Dinner 接口中的 ChooseWaiter() 方法。
然后定义了一个接口 WaiterTask，其中包括两个方法 CreateFood() 和 CleanRoom()，分别用于食物烹饪和打扫卫生。
最后，实现了 waiter 结构体来实现 NewWaiterFactory 接口中的这两个方法，分别返回 Customer 和 Leader 的实例。
*/

type Dinner interface {
	ChooseWaiter()
}

type Customer struct {}

func (c *Customer) ChooseWaiter()  {
	fmt.Println("customer choose a waiter to serve")
}

type Leader struct {}

func (l *Leader) ChooseWaiter()  {
	fmt.Println("leader choose a waiter to work")
}

type WaiterTask interface {
	CookFood() Dinner
	CleanRoom() Dinner
}

type Waiter struct {}

func NewWaiterFactory() WaiterTask{
	return &Waiter{}
}

func (w *Waiter) CookFood() Dinner{
	return &Customer{}
}

func (w *Waiter) CleanRoom() Dinner {
	return &Leader{}
}

