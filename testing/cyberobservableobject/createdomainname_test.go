package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestDomainNameCyberObservableObjectSTIX(t *testing.T) {
	ndn := methodstixobjects.NewDomainNameCyberObservableObjectSTIX()

	assert.Equal(t, ndn.GetType(), "domain-name")
	_, err := ndn.Get()
	assert.Error(t, err)

	v := "any value"
	ndn.SetAnyValue(v)
	_, err = ndn.Get()
	assert.NoError(t, err)
	assert.Equal(t, ndn.GetValue(), v)
	ndn.SetValueValue("av")
	assert.Equal(t, ndn.GetValue(), "av")

	ndn.SetValueResolvesToRefs(stixhelpers.IdentifierTypeSTIX("ab"))
	assert.Equal(t, len(ndn.GetResolvesToRefs()), 1)
	ndn.SetFullValueResolvesToRefs([]stixhelpers.IdentifierTypeSTIX{"a", "b"})
	assert.Equal(t, len(ndn.GetResolvesToRefs()), 2)
}

/*
// DomainNameCyberObservableObjectSTIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Value - сетевое доменное имя (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - список ссылок на один или несколько IP-адресов или доменных имен, на которые разрешается доменное имя
type DomainNameCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value" required:"true"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
}
*/
