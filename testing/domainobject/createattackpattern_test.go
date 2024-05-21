package domainobject

import (
	"testing"

	"github.com/stretchr/testify/assert"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func TestAttackPatternDomainObjectsSTIX(t *testing.T) {
	ap := methodstixobjects.NewAttackPatternDomainObjectsSTIX()

	assert.Equal(t, ap.GetType(), "attack-pattern")

	_, err := ap.Get()
	assert.Error(t, err)

	ap.SetAnyName("attack-pattern name")
	_, err = ap.Get()
	assert.NoError(t, err)

	ap.SetAnyDescription("any description")
	assert.Equal(t, ap.GetDescription(), "any description")

	ap.SetAnyAliases("one")
	ap.SetAnyAliases("two")
	ap.SetAnyAliases("three")
	assert.Equal(t, len(ap.GetAliases()), 3)

	ap.SetValueKillChainPhases([]stixhelpers.KillChainPhasesTypeElementSTIX{
		{KillChainName: "ddd", PhaseName: "name_ee"},
		{KillChainName: "dfg", PhaseName: "name_zx"},
	})
	assert.Equal(t, len(ap.GetKillChainPhases()), 2)
}

/*
// AttackPatternDomainObjectsSTIX объект "Attack Pattern", по терминалогии STIX, описывающий способы компрометации цели
// Name - имя используемое для идентификации "Attack Pattern" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание атаки
// Aliases - альтернативные имена
// KillChainPhases - список цепочки фактов, в которых используется этот шаблон атак
type AttackPatternDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name            string                                       `json:"name" bson:"name" required:"true"`
	Description     string                                       `json:"description" bson:"description"`
	Aliases         []string                                     `json:"aliases" bson:"aliases"`
	KillChainPhases []stixhelpers.KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
}
*/
