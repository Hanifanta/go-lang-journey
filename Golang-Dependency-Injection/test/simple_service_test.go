package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-restful/simple"
	"testing"
)

func TestSimple(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializeService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}
