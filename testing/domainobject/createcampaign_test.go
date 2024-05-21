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

	cdo.SetAnyObjective("example_objective")
	assert.Equal(t, cdo.GetObjective(), "example_objective")

	cdo.SetAnyDescription("example_description")
	assert.Equal(t, cdo.GetDescription(), "example_description")

	err = cdo.SetAnyFirstSeen("2024-02-05T17:31:01+00:00")
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetFirstSeen(), "2024-02-05T17:31:01+00:00")

	err = cdo.SetAnyLastSeen(1716302427100)
	assert.NoError(t, err)
	assert.Equal(t, cdo.GetLastSeen(), "2024-05-21T17:40:27+03:00")

	cdo.SetAnyAliases("vfff")
	cdo.SetAnyAliases("bfff")
	cdo.SetAnyAliases("xfff")
	assert.Equal(t, len(cdo.GetAliases()), 3)
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
