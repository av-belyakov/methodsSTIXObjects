package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIPv6AddressCyberObservableObjectSTIX(t *testing.T) {
	nipv4 := methodstixobjects.NewIPv6AddressCyberObservableObjectSTIX()

	assert.Equal(t, nipv4.GetType(), "ipv6-addr")
	_, err := nipv4.Get()
	assert.Error(t, err)

	nipv4.SetAnyValue("2001:cb:1a:fc:10:10:abc")
	assert.Equal(t, nipv4.GetValue(), "2001:cb:1a:fc:10:10:abc")
	nipv4.SetValueValue("2001:cb::fc:0:0:abc")
	assert.Equal(t, nipv4.GetValue(), "2001:cb::fc:0:0:abc")

	nipv4.SetValueResolvesToRefs([]stixhelpers.IdentifierTypeSTIX{
		"resolves_ref_1",
		"resolves_ref_2",
		"resolves_ref_3",
	})
	assert.Equal(t, len(nipv4.GetResolvesToRefs()), 3)

	nipv4.SetValueBelongsToRefs([]stixhelpers.IdentifierTypeSTIX{"res_bel_1", "res_bel_2", "res_bel_3"})
	assert.Equal(t, len(nipv4.GetBelongsToRefs()), 3)
}

/*
// IPv6AddressCyberObservableObjectSTIX объект "IPv6 Address Object", по терминалогии STIX, содержит один или более IPv6 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv6-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv6-адреса представляет собой один IPv6-адрес,
// суффикс CIDR /128 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
// на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv4-адрес. Объекты, на которые ссылается
// этот список,ДОЛЖНЫ быть типа autonomous-system.
type IPv6AddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []stixhelpers.IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}
*/
