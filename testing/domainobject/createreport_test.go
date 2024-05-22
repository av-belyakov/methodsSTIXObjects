package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestReportDomainObjectsSTIX(t *testing.T) {
	ni := methodstixobjects.NewReportDomainObjectsSTIX()

	assert.Equal(t, ni.GetType(), "report")
	_, err := ni.Get()
	assert.Error(t, err)

	ni.SetAnyName("report name")
	_, err = ni.Get()
	assert.NoError(t, err)

	ni.SetAnyDescription("example_description")
	assert.Equal(t, ni.GetDescription(), "example_description")

	vf := "2024-02-11T07:01:01+00:00"
	errt := ni.SetAnyPublished(vf)
	assert.NoError(t, errt)
	assert.Equal(t, ni.GetPublished(), vf)

	errt = ni.SetAnyPublished(1716367407123)
	assert.NoError(t, errt)
	assert.Equal(t, ni.GetPublished(), "2024-05-22T11:43:27+03:00")

	ni.SetValueReportTypes([]stixhelpers.OpenVocabTypeSTIX{"1q", "2x", "3a", "4d"})
	assert.Equal(t, len(ni.GetReportTypes()), 4)

	ni.SetValueObjectRefs([]stixhelpers.IdentifierTypeSTIX{"r11", "r22", "r33"})
	assert.Equal(t, len(ni.GetObjectRefs()), 3)
}

/*
// ReportDomainObjectsSTIX объект "Report", по терминалогии STIX, содержит совокупность данных об угрозах, сосредоточенных на одной
// или нескольких темах, таких как описание исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним детали.
// Применяется для группировки информации связанной с кибер угрозой. Может быть использован для дальнейшей публикации данной
// информации как истории расследования.
// Name - имя используемое для идентификации "Report" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// ReportTypes - заранее определенный (предложенный) перечень возможных типов контента содержащихся в этом отчете
// Published - время, в формате "2016-05-12T08:17:27.000Z", когда данный отчет был официально опубликован (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ObjectRefs - список идентификаторов STIX объектов, которые ссылаются на этот отчет (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type ReportDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name        string                           `json:"name" bson:"name" required:"true"`
	Description string                           `json:"description" bson:"description"`
	Published   string                           `json:"published" bson:"published" required:"true"`
	ReportTypes []stixhelpers.OpenVocabTypeSTIX  `json:"report_types" bson:"report_types"`
	ObjectRefs  []stixhelpers.IdentifierTypeSTIX `json:"object_refs" bson:"object_refs" required:"true"`
}
*/
