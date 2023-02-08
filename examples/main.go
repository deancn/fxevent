package main

import (
	"context"
	"fxevent/pkg"
	"log"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

const A = "a"

type deps struct {
	fx.In

	ACheckers []pkg.Checker `group:"a"`
	BCheckers []pkg.Checker `group:"b"`
}

var opts = make([]fx.Option, 0)

func WithOpts(c pkg.Checker) {
	opts = append(opts, fx.Provide(
		fx.Annotated{
			Group: A,
			Target: func() pkg.Checker {
				return c
			},
		},
	))
}

func main() {
	var mvh pkg.ValidationHelper = pkg.NewValidationHelp()

	WithOpts(pkg.NewData1())
	WithOpts(pkg.NewData2())

	opts = append(opts, fx.WithLogger(func() fxevent.Logger { return fxevent.NopLogger }))
	opts = append(opts,
		fx.WithLogger(func() fxevent.Logger { return fxevent.NopLogger }),
		fx.Provide(
			func(d deps) pkg.ValidationHelper {
				vh := pkg.NewValidationHelp()
				vh.SetCheckers(d.ACheckers)
				return vh
			},
		),
		fx.Populate(&mvh))

	app := fx.New(opts...)

	app.Start(context.TODO())
	log.Println("----")
	mvh.Check("xxx")
}
