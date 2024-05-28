package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestURLCyberObservableObjectSTIX(t *testing.T) {
	nurl := methodstixobjects.NewURLCyberObservableObjectSTIX()

	assert.Equal(t, nurl.GetType(), "url")
	_, err := nurl.Get()
	assert.NoError(t, err)

	nurl.SetAnyValue("any value example")
	assert.Equal(t, nurl.GetValue(), "any value example")
	nurl.SetValueValue("value example")
	assert.Equal(t, nurl.GetValue(), "value example")
}

/*
// URLCyberObservableObjectSTIX объект "URL Object", по терминологии STIX, содержит унифицированный указатель информационного ресурса (URL).
// Value - содержит унифицированный указатель информационного ресурса (URL).
type URLCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}
*/
