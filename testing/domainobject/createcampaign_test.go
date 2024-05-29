package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCampaignDomainObjectsSTIX(t *testing.T) {
	cdo := methodstixobjects.NewCampaignDomainObjectsSTIX()

	assert.Equal(t, cdo.GetType(), "campaign")
	_, err := cdo.Get()
	assert.Error(t, err)

	cdo.SetAnyName("campaign name")
	_, err = cdo.Get()
	assert.NoError(t, err)
	cdo.SetValueName("cap_name")
	assert.Equal(t, cdo.GetName(), "cap_name")

	cdo.SetAnyObjective("example_objective")
	assert.Equal(t, cdo.GetObjective(), "example_objective")
	cdo.SetValueObjective("e_o")
	assert.Equal(t, cdo.GetObjective(), "e_o")

	cdo.SetAnyDescription("example_description")
	assert.Equal(t, cdo.GetDescription(), "example_description")
	cdo.SetValueDescription("exm_description")
	assert.Equal(t, cdo.GetDescription(), "exm_description")

	// --- FirstSeen
	fst := "2024-02-14T12:03:06+00:00"
	err = cdo.SetAnyFirstSeen(fst)
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetFirstSeen(), fst)

	err = cdo.SetAnyFirstSeen(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetFirstSeen(), "2024-05-28T12:33:44+03:00")

	// --- LastSeen
	lst := "2024-02-15T20:10:16+00:00"
	err = cdo.SetAnyLastSeen(lst)
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetLastSeen(), lst)

	err = cdo.SetAnyLastSeen(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetLastSeen(), "2024-05-28T12:33:44+03:00")

	cdo.SetAnyAliases("vfff")
	cdo.SetAnyAliases("bfff")
	cdo.SetAnyAliases("xfff")
	cdo.SetValueAliases("vujdv")
	assert.Equal(t, len(cdo.GetAliases()), 4)
}

/*
// CampaignDomainObjectsSTIX объект "Campaign", по терминалогии STIX, это набор действий определяющих злонамеренную деятельность или атаки
// Name - имя используемое для идентификации "Campaign" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации "Campaign"
// FirstSeen - время когда компания была впервые обнаружена, в формате "2016-05-12T08:17:27.000Z"
// LastSeen - время когда компания была зафиксирована в последний раз, в формате "2016-05-12T08:17:27.000Z"
// Objective - основная цель, желаемый результат или предполагаемый эффект
type CampaignDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name        string   `json:"name" bson:"name" required:"true"`
	Objective   string   `json:"objective" bson:"objective"`
	Description string   `json:"description" bson:"description"`
	FirstSeen   string   `json:"first_seen" bson:"first_seen"`
	LastSeen    string   `json:"last_seen" bson:"last_seen"`
	Aliases     []string `json:"aliases" bson:"aliases"`
}
*/
