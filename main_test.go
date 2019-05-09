package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m-mizutani/deepalert"
	main "github.com/m-mizutani/deepalert-inspector-go-skelton"
)

func TestHandler(t *testing.T) {
	args := main.Arguments{
		Attr: deepalert.Attribute{},
	}
	entity, err := main.Handler(args)
	assert.NoError(t, err)
	assert.NotNil(t, entity)
}
