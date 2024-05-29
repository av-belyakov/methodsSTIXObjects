package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIdentityDomainObjectsSTIX(t *testing.T) {
	ni := methodstixobjects.NewIdentityDomainObjectsSTIX()

	assert.Equal(t, ni.GetType(), "identity")
	_, err := ni.Get()
	assert.Error(t, err)

	ni.SetAnyName("identity name")
	_, err = ni.Get()
	assert.NoError(t, err)
	ni.SetValueName("i name")
	assert.Equal(t, ni.GetName(), "i name")

	ni.SetAnyDescription("example_description")
	assert.Equal(t, ni.GetDescription(), "example_description")
	ni.SetValueDescription("exm_description")
	assert.Equal(t, ni.GetDescription(), "exm_description")

	cimessage := "example contact information"
	ni.SetAnyContactInformation(cimessage)
	assert.Equal(t, ni.GetContactInformation(), cimessage)
	ni.SetValueContactInformation("e c information")
	assert.Equal(t, ni.GetContactInformation(), "e c information")

	ni.SetAnyRoles("one")
	ni.SetAnyRoles("two")
	ni.SetValueRoles("three")
	assert.Equal(t, len(ni.GetRoles()), 3)

	icmessage := stixhelpers.OpenVocabTypeSTIX("test identity class")
	ni.SetValueIdentityClass(icmessage)
	assert.Equal(t, ni.GetIdentityClass(), icmessage)

	ni.SetValueSectors("section")
	assert.Equal(t, len(ni.GetSectors()), 1)
}

/*
// IdentityDomainObjectsSTIX объект "Identity", по терминалогии STIX, содержит основную идентификационную информацию физичиских лиц, организаций и т.д.
// Name - имя используемое для идентификации "Identity" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Roles - список ролей для идентификации действий
// IdentityClass - одно, из заранее определенных (предложенных) значений физического лица или организации
// Sectors - заранее определенный (предложенный) перечень отраслей промышленности, к которой принадлежит физическое лицо или организация
// ContactInformation - любая контактная информация (email, phone number and etc.)
type IdentityDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name               string                          `json:"name" bson:"name" required:"true"`
	Description        string                          `json:"description" bson:"description"`
	ContactInformation string                          `json:"contact_information" bson:"contact_information"`
	Roles              []string                        `json:"roles" bson:"roles"`
	IdentityClass      stixhelpers.OpenVocabTypeSTIX   `json:"identity_class" bson:"identity_class"`
	Sectors            []stixhelpers.OpenVocabTypeSTIX `json:"sectors" bson:"sectors"`
}
*/
