package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestInfrastructureDomainObjectsSTIX(t *testing.T) {
	ni := methodstixobjects.NewInfrastructureDomainObjectsSTIX()

	assert.Equal(t, ni.GetType(), "infrastructure")
	_, err := ni.Get()
	assert.Error(t, err)

	ni.SetAnyName("infrastructure name")
	_, err = ni.Get()
	assert.NoError(t, err)

	ni.SetAnyDescription("example_description")
	assert.Equal(t, ni.GetDescription(), "example_description")

	vf := "2024-02-11T07:01:01+00:00"
	err = ni.SetAnyFirstSeen(vf)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetFirstSeen(), vf)

	err = ni.SetAnyFirstSeen(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetFirstSeen(), "2024-05-22T11:43:27+03:00")

	vu := "2024-02-10T11:21:01+00:00"
	err = ni.SetAnyLastSeen(vu)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetLastSeen(), vu)

	err = ni.SetAnyLastSeen(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetLastSeen(), "2024-05-22T11:43:27+03:00")

	ni.SetAnyAliases("zxc")
	ni.SetAnyAliases("asd")
	ni.SetAnyAliases("qwe")
	assert.Equal(t, len(ni.GetAliases()), 3)

	ni.SetValueKillChainPhases([]stixhelpers.KillChainPhasesTypeElementSTIX{
		{KillChainName: "green chan", PhaseName: "first phase"},
		{KillChainName: "red chan", PhaseName: "second phase"},
		{KillChainName: "blue chan", PhaseName: "third phase"},
	})
	assert.Equal(t, len(ni.GetKillChainPhases()), 3)

	ni.SetValueIndicatorTypes([]stixhelpers.OpenVocabTypeSTIX{"1q", "2x", "3a", "4d"})
	assert.Equal(t, len(ni.GetIndicatorTypes()), 4)
}

/*
// InfrastructureDomainObjectsSTIX объект "Infrastructure", по терминалогии STIX, содержит описание любых систем, программных
// служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
// Name - имя используемое для идентификации "Infrastructure" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// InfrastructureTypes - заранее определенный (предложенный) перечень описываемых инфраструктур
// Aliases - альтернативные имена используемые для идентификации этой инфраструктуры
// KillChainPhases - список цепочки фактов, к которым может быть отнесена эта инфраструктура
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данная инфраструктура была впервые замечена за осуществлением вредоносной активности
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данная инфраструктура в последний раз была замечена за осуществлением вредоносной активности
type InfrastructureDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name                string                                       `json:"name" bson:"name" required:"true"`
	Description         string                                       `json:"description" bson:"description"`
	FirstSeen           string                                       `json:"first_seen" bson:"first_seen"`
	LastSeen            string                                       `json:"last_seen" bson:"last_seen"`
	Aliases             []string                                     `json:"aliases" bson:"aliases"`
	KillChainPhases     []stixhelpers.KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	InfrastructureTypes []stixhelpers.OpenVocabTypeSTIX              `json:"infrastructure_types" bson:"infrastructure_types"`
}
*/
