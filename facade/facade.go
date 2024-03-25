package facade

import "fmt"

type CarModel struct {}
func NewCarModel() *CarModel{
	return &CarModel{}
}

func (c *CarModel) SetModel()  {
	fmt.Println("car model ----- set model ")
}

type CarEngine struct {}
func NewCarEngine() *CarEngine{
	return &CarEngine{}
}

func (c *CarEngine) SetEngine()  {
	fmt.Println("car engine ----- set engine ")
}

type CarBody struct {}

func NewCarBody() *CarBody {
	return &CarBody{}
}

func (c *CarBody) SetCarBody()  {
	fmt.Println("car body ---- set body.")
}

type CarFacade struct {
	CarModel
	CarEngine
	CarBody
}
func NewCarFacade() *CarFacade {
	return &CarFacade{}
}

/*包装*/
func (c *CarFacade) CreateCompleteCar()  {
	/* one way*/
	//c.CarModel.SetModel()
	//c.CarEngine.SetEngine()
	//c.CarBody.SetCarBody()

	/* another way*/
	c.SetModel()
	c.SetEngine()
	c.SetCarBody()
}