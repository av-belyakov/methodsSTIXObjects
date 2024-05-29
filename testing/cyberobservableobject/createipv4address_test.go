package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIPv4AddressCyberObservableObjectSTIX(t *testing.T) {
	nipv4 := methodstixobjects.NewIPv4AddressCyberObservableObjectSTIX()

	assert.Equal(t, nipv4.GetType(), "ipv4-addr")
	_, err := nipv4.Get()
	assert.Error(t, err)

	nipv4.SetAnyValue("10.14.24.21")
	assert.Equal(t, nipv4.GetValue(), "10.14.24.21")
	nipv4.SetValueValue("59.64.174.23")
	assert.Equal(t, nipv4.GetValue(), "59.64.174.23")

	nipv4.SetFullValueResolvesToRefs([]stixhelpers.IdentifierTypeSTIX{"res_1", "res_2"})
	assert.Equal(t, len(nipv4.GetResolvesToRefs()), 2)
	nipv4.SetValueResolvesToRefs(stixhelpers.IdentifierTypeSTIX("res_3"))
	assert.Equal(t, len(nipv4.GetResolvesToRefs()), 3)

	nipv4.SetFullValueBelongsToRefs([]stixhelpers.IdentifierTypeSTIX{"res_1", "res_2", "res_3"})
	assert.Equal(t, len(nipv4.GetBelongsToRefs()), 3)
	nipv4.SetValueBelongsToRefs(stixhelpers.IdentifierTypeSTIX("res_n"))
	assert.Equal(t, len(nipv4.GetBelongsToRefs()), 4)
}

/*
// IPv4AddressCyberObservableObjectSTIX объект "IPv4 Address Object", по терминалогии STIX, содержит один или
// более IPv4 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv4-адресов, выраженные с помощью нотации CIDR. Если данный
// объект IPv4-адреса представляет собой один IPv4-адрес, суффикс CIDR /32 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2,
// на которые разрешается IPv6-адрес. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv6-адрес.
// Объекты, на которые ссылается этот список, ДОЛЖНЫ быть типа autonomous-system.
type IPv4AddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []stixhelpers.IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}
*/
