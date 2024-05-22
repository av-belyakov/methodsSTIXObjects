package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestAutonomousSystemCyberObservableObjectSTIX(t *testing.T) {
	nas := methodstixobjects.NewAutonomousSystemCyberObservableObjectSTIX()

	assert.Equal(t, nas.GetType(), "autonomous-system")
	_, err := nas.Get()
	assert.Error(t, err)

	nas.SetAnyNumber(1324010)
	_, err = nas.Get()
	assert.NoError(t, err)

	asn := "autonomous-system name"
	nas.SetAnyName(asn)
	assert.Equal(t, nas.GetName(), asn)

	nas.SetAnyRIR("RE RIR")
	assert.Equal(t, nas.GetRIR(), "RE RIR")
}

/*
// AutonomousSystemCyberObservableObjectSTIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Number - содержит номер присвоенный Автономной системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - название Автономной системы
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системы
type AutonomousSystemCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Number int    `json:"number" bson:"number" required:"true"`
	Name   string `json:"name" bson:"name"`
	RIR    string `json:"rir" bson:"rir"`
}
*/
