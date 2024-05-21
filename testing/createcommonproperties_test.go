package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
)

func TestCommonPropertiesObjectSTIX(t *testing.T) {
	cpo := methodstixobjects.NewCommonPropertiesObjectSTIX()

	typeExample := "common_time_example"
	cpo.SetValueType(typeExample)
	assert.Equal(t, cpo.GetType(), typeExample)

	idExample := "ye72-rr49-d244tsq-23"
	cpo.SetValueID(idExample)
	assert.Equal(t, cpo.GetID(), idExample)
}
