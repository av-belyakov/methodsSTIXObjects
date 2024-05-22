package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestThreatActorDomainObjectsSTIX(t *testing.T) {
	nta := methodstixobjects.NewThreatActorDomainObjectsSTIX()

	assert.Equal(t, nta.GetType(), "threat-actor")
	_, err := nta.Get()
	assert.Error(t, err)

	nta.SetAnyName("threat-actor name")
	_, err = nta.Get()
	assert.NoError(t, err)

	nta.SetAnyDescription("example_description")
	assert.Equal(t, nta.GetDescription(), "example_description")

	vf := "2024-02-11T07:01:01+00:00"
	err = nta.SetAnyFirstSeen(vf)
	assert.NoError(t, err)
	assert.Equal(t, nta.GetFirstSeen(), vf)

	err = nta.SetAnyFirstSeen(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nta.GetFirstSeen(), "2024-05-22T11:43:27+03:00")

	vu := "2024-02-10T11:21:01+00:00"
	err = nta.SetAnyLastSeen(vu)
	assert.NoError(t, err)
	assert.Equal(t, nta.GetLastSeen(), vu)

	err = nta.SetAnyLastSeen(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nta.GetLastSeen(), "2024-05-22T11:43:27+03:00")

	nta.SetAnyAliases("zxc")
	nta.SetAnyAliases("asd")
	nta.SetAnyAliases("qwe")
	assert.Equal(t, len(nta.GetAliases()), 3)

	nta.SetAnyGoals("11")
	nta.SetAnyGoals("22")
	assert.Equal(t, len(nta.GetGoals()), 2)

	s := stixhelpers.OpenVocabTypeSTIX("sophistication example")
	nta.SetValueSophistication(s)
	assert.Equal(t, nta.GetSophistication(), s)

	rl := stixhelpers.OpenVocabTypeSTIX("resource level example")
	nta.SetValueResourceLevel(rl)
	assert.Equal(t, nta.GetResourceLevel(), rl)

	pm := stixhelpers.OpenVocabTypeSTIX("primary motivation")
	nta.SetValuePrimaryMotivation(pm)
	assert.Equal(t, nta.GetPrimaryMotivation(), pm)

	nta.SetValueSecondaryMotivations([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, len(nta.GetSecondaryMotivations()), 6)

	nta.SetValuePersonalMotivations([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, len(nta.GetPersonalMotivations()), 6)

	nta.SetValueThreatActorTypes([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, len(nta.GetThreatActorTypes()), 6)

	nta.SetValueRoles([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, len(nta.GetRoles()), 6)
}

/*
// ThreatActorDomainObjectsSTIX объект "Threat Actor", по терминалогии STIX, содержит информацию о физических лицах или их
// группах и организациях которые могут действовать со злым умыслом.
// Name - имя используемое для идентификации "Threat Actor" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ThreatActorTypes - заранее определенный (предложенный) перечень типов субъектов угрозы
// Aliases - альтернативные имена используемые для этого субъекта угроз
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный субъект угроз был впервые зафиксирован
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный субъект угроз был зафиксирован в последний раз
// Roles - заранее определенный (предложенный) перечень возможных ролей субъекта угроз
// Goals - высокоуровневые цели субъекта угроз.
// Sophistication - один, из заранее определенного (предложенного) перечня навыков, специальных знания, специальной
// подготовки или опыта, которыми должен обладать субъект угрозы, чтобы осуществить атаку
// ResourceLevel - один, из заранее определенного (предложенного) перечня организационных уровней, на котором обычно работает
// этот субъект угрозы, который, в свою очередь, определяет ресурсы, доступные этому субъекту угрозы для использования в атаке.
// PrimaryMotivation - одна, из заранее определенного (предложенного) перечня причин, мотиваций или целей стоящих за этим субъектом угрозы
// SecondaryMotivations - заранее определенный (предложенный) перечень возможных вторичных причин, мотиваций или целей стоящих за этим субъектом угрозы
// PersonalMotivations - заранее определенный (предложенный) перечень возможных персональных причин, мотиваций или целей стоящих за этим субъектом угрозы
type ThreatActorDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name                 string                          `json:"name" bson:"name" required:"true"`
	Description          string                          `json:"description" bson:"description"`
	FirstSeen            string                          `json:"first_seen" bson:"first_seen"`
	LastSeen             string                          `json:"last_seen" bson:"last_seen"`
	Aliases              []string                        `json:"aliases" bson:"aliases"`
	Goals                []string                        `json:"goals" bson:"goals"`
	Sophistication       stixhelpers.OpenVocabTypeSTIX   `json:"sophistication" bson:"sophistication"`
	ResourceLevel        stixhelpers.OpenVocabTypeSTIX   `json:"resource_level" bson:"resource_level"`
	PrimaryMotivation    stixhelpers.OpenVocabTypeSTIX   `json:"primary_motivation" bson:"primary_motivation"`
	SecondaryMotivations []stixhelpers.OpenVocabTypeSTIX `json:"secondary_motivations" bson:"secondary_motivations"`
	PersonalMotivations  []stixhelpers.OpenVocabTypeSTIX `json:"personal_motivations" bson:"personal_motivations"`
	ThreatActorTypes     []stixhelpers.OpenVocabTypeSTIX `json:"threat_actor_types" bson:"threat_actor_types"`
	Roles                []stixhelpers.OpenVocabTypeSTIX `json:"roles" bson:"roles"`
}
*/
