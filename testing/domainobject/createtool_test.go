package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestToolDomainObjectsSTIX(t *testing.T) {
	nt := methodstixobjects.NewToolDomainObjectsSTIX()

	assert.Equal(t, nt.GetType(), "tool")
	_, err := nt.Get()
	assert.Error(t, err)

	nt.SetAnyName("tool name")
	_, err = nt.Get()
	assert.NoError(t, err)
	nt.SetValueName("t_name")
	assert.Equal(t, nt.GetName(), "t_name")

	nt.SetAnyDescription("example_description")
	assert.Equal(t, nt.GetDescription(), "example_description")
	nt.SetValueDescription("exm_description")
	assert.Equal(t, nt.GetDescription(), "exm_description")

	nt.SetAnyToolVersion("tool v1.2")
	assert.Equal(t, nt.GetToolVersion(), "tool v1.2")
	nt.SetValueToolVersion("tool v2.3")
	assert.Equal(t, nt.GetToolVersion(), "tool v2.3")

	nt.SetAnyAliases("zxc")
	nt.SetAnyAliases("asd")
	nt.SetAnyAliases("qwe")
	nt.SetValueAliases("qwe")
	assert.Equal(t, len(nt.GetAliases()), 4)

	nt.SetFullValueKillChainPhases([]stixhelpers.KillChainPhasesTypeElementSTIX{
		{KillChainName: "green chan", PhaseName: "first phase"},
		{KillChainName: "red chan", PhaseName: "second phase"},
		{KillChainName: "blue chan", PhaseName: "third phase"},
	})
	nt.SetValueKillChainPhases(stixhelpers.KillChainPhasesTypeElementSTIX{
		KillChainName: "grey chan", PhaseName: "first first phase",
	})
	assert.Equal(t, len(nt.GetKillChainPhases()), 4)

	nt.SetFullValueToolTypes([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	nt.SetValueToolTypes(stixhelpers.OpenVocabTypeSTIX("g"))
	assert.Equal(t, len(nt.GetToolTypes()), 7)
}

/*
// ToolDomainObjectsSTIX объект "Tool", по терминалогии STIX, содержит информацию о легитимном ПО которое может быть
// использованно для реализации компьютерных угроз
// Name - имя используемое для идентификации "Tool" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ToolTypes - заранее определенный (предложенный) перечень типов инструментов
// Aliases - альтернативные имена используемые для идентификации инструментов
// KillChainPhases - список цепочки фактов, к которым может быть отнесен этот инструмент
// ToolVersion - версия инструмента
type ToolDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name            string                                       `json:"name" bson:"name" required:"true"`
	Description     string                                       `json:"description" bson:"description"`
	ToolVersion     string                                       `json:"tool_version" bson:"tool_version"`
	Aliases         []string                                     `json:"aliases" bson:"aliases"`
	KillChainPhases []stixhelpers.KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	ToolTypes       []stixhelpers.OpenVocabTypeSTIX              `json:"tool_types" bson:"tool_types"`
}
*/
