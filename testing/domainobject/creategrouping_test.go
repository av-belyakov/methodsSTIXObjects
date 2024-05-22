package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestGroupingDomainObjectsSTIX(t *testing.T) {
	ng := methodstixobjects.NewGroupingDomainObjectsSTIX()

	assert.Equal(t, ng.GetType(), "grouping")
	_, err := ng.Get()
	assert.Error(t, err)

	ng.SetAnyName("grouping name")
	_, err = ng.Get()
	assert.NoError(t, err)

	ng.SetAnyDescription("example_description")
	assert.Equal(t, ng.GetDescription(), "example_description")

	ng.SetValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"111a", "222b", "333c"})
	assert.Equal(t, len(ng.GetObjectRefs()), 3)
}

/*
// GroupingDomainObjectsSTIX объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
// Name - имя используемое для идентификации "Grouping" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Context - краткий дескриптор конкретного контекста, совместно используемого содержимым, на которое ссылается группа. Должно быть одно,
// из заранее определенных (предложенных) значений (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - указывает на список объектов STIX, на которые ссылается эта группировка (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type GroupingDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name        string                           `json:"name" bson:"name" required:"true"`
	Description string                           `json:"description" bson:"description"`
	Context     stixhelpers.OpenVocabTypeSTIX    `json:"context" bson:"context" required:"true"`
	ObjectRefs  []stixhelpers.IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}
*/
