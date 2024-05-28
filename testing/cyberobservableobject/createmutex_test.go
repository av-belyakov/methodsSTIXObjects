package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestMutexCyberObservableObjectSTIX(t *testing.T) {
	nm := methodstixobjects.NewMutexCyberObservableObjectSTIX()

	assert.Equal(t, nm.GetType(), "mutex")
	_, err := nm.Get()
	assert.Error(t, err)

	nm.SetAnyName("mutix name example")
	assert.Equal(t, nm.GetName(), "mutix name example")
	nm.SetValueName("name example")
	assert.Equal(t, nm.GetName(), "name example")
}

/*
// MutexCyberObservableObjectSTIX объект "Mutex Object", по терминалогии STIX, содержит свойства объекта взаимного исключения (mutex).
// Name - указывает имя объекта мьютекса (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type MutexCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Name string `json:"name" bson:"name"`
}
*/
