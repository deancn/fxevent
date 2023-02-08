package pkg

import "fmt"

type Data1Ensure struct {
}

func NewData1() Checker {
	return &Data1Ensure{}
}

func (o *Data1Ensure) Check(val any) bool {
	o.CheckIndicatorType(val)
	o.CheckIndicatorUdf(val)
	o.CheckIndicatorValue(val)
	return true
}

func (o *Data1Ensure) CheckIndicatorType(val any) {
	fmt.Println("hell run check 1")
}

func (o *Data1Ensure) CheckIndicatorValue(val any) {

}

func (o *Data1Ensure) CheckIndicatorUdf(val any) {

}
