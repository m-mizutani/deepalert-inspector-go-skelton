package main

import (
	"context"

	"github.com/m-mizutani/deepalert"
)

type Arguments struct {
	Attr deepalert.Attribute
}

func Handler(args Arguments) (deepalert.ReportContentEntity, error) {
	hostReport := deepalert.ReportHost{}
	return &hostReport, nil

}

func sampleHandler(ctx context.Context, attr deepalert.Attribute) (deepalert.ReportContentEntity, error) {
	args := Arguments{Attr: attr}
	return Handler(args)
}

func main() {
	deepalert.StartInspector(sampleHandler, "sampleHandler")
}
