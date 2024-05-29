package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestOpinionDomainObjectsSTIX(t *testing.T) {
	no := methodstixobjects.NewOpinionDomainObjectsSTIX()

	assert.Equal(t, no.GetType(), "opinion")
	_, err := no.Get()
	assert.Error(t, err)

	oe := "opinion example"
	no.SetAnyOpinion(oe)
	assert.Equal(t, no.GetOpinion(), stixhelpers.EnumTypeSTIX(oe))
	no.SetValueOpinion("opinion")
	assert.Equal(t, no.GetOpinion(), stixhelpers.EnumTypeSTIX("opinion"))

	no.SetFullValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"r11", "r22", "r33"})
	no.SetValueObjectRefs(stixhelpers.IdentifierTypeSTIX("r44"))
	assert.Equal(t, len(no.GetObjectRefs()), 4)
	_, err = no.Get()
	assert.NoError(t, err)

	e := "explanation example message"
	no.SetAnyExplanation(e)
	assert.Equal(t, no.GetExplanation(), e)
	no.SetValueExplanation("example message")
	assert.Equal(t, no.GetExplanation(), "example message")

	no.SetAnyAuthors("author_1")
	no.SetAnyAuthors("author_2")
	no.SetValueAuthors("author_3")
	assert.Equal(t, len(no.GetAuthors()), 3)
}

/*
// OpinionDomainObjectsSTIX объект "Opinion", по терминалогии STIX, содержит оценку информации в приведенной в каком либо другом объекте STIX,
// которую произвел другой участник анализа.
// Explanation - объясняет почему обработчик оставил это мнение
// Authors - список авторов этого мнения
// Opinion - мнение обо всех STIX объектах перечисленных в ObjectRefs (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов на другие STIX объекты (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type OpinionDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Explanation string                           `json:"explanation" bson:"explanation"`
	Authors     []string                         `json:"authors" bson:"authors"`
	Opinion     stixhelpers.EnumTypeSTIX         `json:"opinion" bson:"opinion" required:"true"`
	ObjectRefs  []stixhelpers.IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}
*/
