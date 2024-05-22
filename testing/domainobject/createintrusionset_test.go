package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIntrusionSetDomainObjectsSTIX(t *testing.T) {
	ni := methodstixobjects.NewIntrusionSetDomainObjectsSTIX()

	assert.Equal(t, ni.GetType(), "intrusion-set")
	_, err := ni.Get()
	assert.Error(t, err)

	ni.SetAnyName("intrusion-set name")
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

	ni.SetAnyGoals("11")
	ni.SetAnyGoals("22")
	assert.Equal(t, len(ni.GetGoals()), 2)

	rl := stixhelpers.OpenVocabTypeSTIX("resource level")
	ni.SetValueResourceLevel(rl)
	assert.Equal(t, ni.GetResourceLevel(), rl)

	pm := stixhelpers.OpenVocabTypeSTIX("primary motivation")
	ni.SetValuePrimaryMotivation(pm)
	assert.Equal(t, ni.GetPrimaryMotivation(), pm)

	ni.SetValueSecondaryMotivations([]stixhelpers.OpenVocabTypeSTIX{"a", "b", "c", "d", "e", "f"})
	assert.Equal(t, len(ni.GetSecondaryMotivations()), 6)
}

/*
// IntrusionSetDomainObjectsSTIX объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный набор враждебного поведения и ресурсов
// с общими свойствами, который, как считается, управляется одной организацией
// Name - имя используемое для идентификации "Intrusion Set" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Aliases - альтернативные имена используемые для идентификации набора вторжения
// FirstSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный набор вторжения впервые был зафиксирован
// LastSeen - время, в формате "2016-05-12T08:17:27.000Z", когда данный набор вторжения был зафиксирован в последний раз
// Goals - высокоуровневые цели этого набора вторжения
// ResourceLevel - заранее определенный (предложенный) перечень уровней, на которых обычно работает данный набор вторжений,
// который, в свою очередь, определяет ресурсы, доступные этому набору вторжений для использования в атаке
// PrimaryMotivation - одно, из заранее определенных (предложенных) перечней причин, мотиваций или целей определяющий данный набор вторжения
// SecondaryMotivations - заранее определенный (предложенный) вторичный перечень причин, мотиваций или целей определяющий данный набор вторжений
type IntrusionSetDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name                 string                          `json:"name" bson:"name" required:"true"`
	Description          string                          `json:"description" bson:"description"`
	FirstSeen            string                          `json:"first_seen" bson:"first_seen"`
	LastSeen             string                          `json:"last_seen" bson:"last_seen"`
	Aliases              []string                        `json:"aliases" bson:"aliases"`
	Goals                []string                        `json:"goals" bson:"goals"`
	ResourceLevel        stixhelpers.OpenVocabTypeSTIX   `json:"resource_level" bson:"resource_level"`
	PrimaryMotivation    stixhelpers.OpenVocabTypeSTIX   `json:"primary_motivation" bson:"primary_motivation"`
	SecondaryMotivations []stixhelpers.OpenVocabTypeSTIX `json:"secondary_motivations" bson:"secondary_motivations"`
}
*/
