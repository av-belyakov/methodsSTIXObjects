package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNoteDomainObjectsSTIX(t *testing.T) {
	nn := methodstixobjects.NewNoteDomainObjectsSTIX()

	assert.Equal(t, nn.GetType(), "note")
	_, err := nn.Get()
	assert.Error(t, err)

	nn.SetAnyContent("new content")
	_, err = nn.Get()
	assert.NoError(t, err)
	assert.Equal(t, nn.GetContent(), "new content")

	a := "abstract example"
	nn.SetAnyAbstract(a)
	assert.Equal(t, nn.GetAbstract(), a)

	nn.SetAnyAuthors("z11")
	nn.SetAnyAuthors("z22")
	assert.Equal(t, len(nn.GetAuthors()), 2)

	nn.SetValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"assdf", "vcvvv", "werrrt"})
	assert.Equal(t, len(nn.GetObjectRefs()), 3)
}

/*
// NoteDomainObjectsSTIX объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую текущий контекст анализа
// либо содержащей результаты дополнительного анализа которые не может быть описан в терминах объектов STIX
// Abstract - краткое изложение содержания записки
// Content - основное содержание (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Authors - список авторов
// ObjectRefs - список идентификаторов на других DO STIX объектов к которым применяется замечание (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type NoteDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Abstract   string                           `json:"abstract" bson:"abstract"`
	Content    string                           `json:"content" bson:"content" required:"true"`
	Authors    []string                         `json:"authors" bson:"authors"`
	ObjectRefs []stixhelpers.IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}
*/
