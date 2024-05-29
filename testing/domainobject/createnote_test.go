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
	nn.SetValueContent("content")
	assert.Equal(t, nn.GetContent(), "content")

	a := "abstract example"
	nn.SetAnyAbstract(a)
	assert.Equal(t, nn.GetAbstract(), a)
	nn.SetValueAbstract("abstract")
	assert.Equal(t, nn.GetAbstract(), "abstract")

	nn.SetAnyAuthors("z11")
	nn.SetAnyAuthors("z22")
	nn.SetValueAuthors("z33")
	assert.Equal(t, len(nn.GetAuthors()), 3)

	nn.SetFullValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"assdf", "vcvvv", "werrrt"})
	nn.SetValueObjectRefs(stixhelpers.IdentifierTypeSTIX("cdiid"))
	assert.Equal(t, len(nn.GetObjectRefs()), 4)
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
