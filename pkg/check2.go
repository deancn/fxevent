package pkg

import "fmt"

type Data2Ensure struct {
}

func NewData2() Checker {
	return &Data2Ensure{}
}

func (o *Data2Ensure) Check(val any) bool {
	o.CheckIndicatorType(val)
	o.CheckIndicatorUdf(val)
	o.CheckIndicatorValue(val)
	return true
}

func (o *Data2Ensure) CheckIndicatorType(val any) {
	fmt.Println("hell run check 2")
}

func (o *Data2Ensure) CheckIndicatorValue(val any) {

}

func (o *Data2Ensure) CheckIndicatorUdf(val any) {

}
