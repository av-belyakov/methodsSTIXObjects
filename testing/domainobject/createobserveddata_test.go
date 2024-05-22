package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestObservedDataDomainObjectsSTIX(t *testing.T) {
	nod := methodstixobjects.NewObservedDataDomainObjectsSTIX()

	assert.Equal(t, nod.GetType(), "observed-data")
	_, err := nod.Get()
	assert.Error(t, err)

	nod.SetAnyNumberObserved(10)
	assert.Equal(t, nod.GetNumberObserved(), 10)

	vf := "2024-02-11T07:01:01+00:00"
	errt := nod.SetAnyFirstObserved(vf)
	assert.NoError(t, errt)
	assert.Equal(t, nod.GetFirstObserved(), vf)

	errt = nod.SetAnyFirstObserved(1716367407123)
	assert.NoError(t, errt)
	assert.Equal(t, nod.GetFirstObserved(), "2024-05-22T11:43:27+03:00")

	vu := "2024-02-10T11:21:01+00:00"
	errt = nod.SetAnyLastObserved(vu)
	assert.NoError(t, errt)
	assert.Equal(t, nod.GetLastObserved(), vu)

	errt = nod.SetAnyLastObserved(1716367407123)
	assert.NoError(t, errt)
	assert.Equal(t, nod.GetLastObserved(), "2024-05-22T11:43:27+03:00")

	_, err = nod.Get()
	assert.NoError(t, err)

	nod.SetValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"eee"})
	assert.Equal(t, len(nod.GetObjectRefs()), 1)
}

/*
// ObservedDataDomainObjectsSTIX объект "Observed Data", по терминалогии STIX, содержит информацию о сущностях связанных с
// кибер безопасностью, таких как файлы, системы или сети. Наблюдаемые данные это не результат анализа или заключение искусственного
// интеллекта, это просто сырая информация без какого-либо контекста.
// FirstObserved - время, в формате "2016-05-12T08:17:27.000Z", начала временного окна, в течение которого были замечены данные (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// LastObserved - время, в формате "2016-05-12T08:17:27.000Z", окончание временного окна, в течение которого были замечены данные (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// NumberObserved - количество раз, когда фиксировался каждый наблюдаемый кибер объект SCO, представленный в свойстве ObjectRef (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов на другие наблюдаемые кибер объекты SCO
type ObservedDataDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	NumberObserved int                              `json:"number_observed" bson:"number_observed" required:"true"`
	FirstObserved  string                           `json:"first_observed" bson:"first_observed" required:"true"`
	LastObserved   string                           `json:"last_observed" bson:"last_observed" required:"true"`
	ObjectRefs     []stixhelpers.IdentifierTypeSTIX `json:"object_refs" bson:"object_refs"`
}
*/
