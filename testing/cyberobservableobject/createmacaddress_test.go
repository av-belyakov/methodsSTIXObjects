package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestMACAddressCyberObservableObjectSTIX(t *testing.T) {
	nmac := methodstixobjects.NewMACAddressCyberObservableObjectSTIX()

	assert.Equal(t, nmac.GetType(), "mac-addr")
	_, err := nmac.Get()
	assert.Error(t, err)

	nmac.SetAnyValue("02:42:0c:07:a0:4c")
	assert.Equal(t, nmac.GetValue(), "02:42:0c:07:a0:4c")
	nmac.SetValueValue("12:43:0c:07:a6:1c")
	assert.Equal(t, nmac.GetValue(), "12:43:0c:07:a6:1c")
}

/*
// MACAddressCyberObservableObjectSTIX объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой
// один адрес управления доступом к среде (MAC).
// Value - Задает значение одного MAC-адреса. Значение MAC - адреса ДОЛЖНО быть представлено в виде одного строчного MAC-48 address,
// разделенного двоеточием, который ДОЛЖЕН включать начальные нули для каждого октета. Пример: 00:00:ab:cd:ef:01. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type MACAddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}
*/
