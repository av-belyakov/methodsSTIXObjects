package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestSoftwareCyberObservableObjectSTIX(t *testing.T) {
	ns := methodstixobjects.NewSoftwareCyberObservableObjectSTIX()

	assert.Equal(t, ns.GetType(), "software")
	_, err := ns.Get()
	assert.NoError(t, err)

	ns.SetAnyName("any name example")
	assert.Equal(t, ns.GetName(), "any name example")
	ns.SetValueName("name example")
	assert.Equal(t, ns.GetName(), "name example")

	ns.SetAnyCPE("any cpe example")
	assert.Equal(t, ns.GetCPE(), "any cpe example")
	ns.SetValueCPE("cpe example")
	assert.Equal(t, ns.GetCPE(), "cpe example")

	ns.SetAnySwID("any swid example")
	assert.Equal(t, ns.GetSwID(), "any swid example")
	ns.SetValueSwID("swid example")
	assert.Equal(t, ns.GetSwID(), "swid example")

	ns.SetAnyVendor("any vendor example")
	assert.Equal(t, ns.GetVendor(), "any vendor example")
	ns.SetValueVendor("vendor example")
	assert.Equal(t, ns.GetVendor(), "vendor example")

	ns.SetAnyVersion("any version example")
	assert.Equal(t, ns.GetVersion(), "any version example")
	ns.SetValueVersion("version example")
	assert.Equal(t, ns.GetVersion(), "version example")

	ns.SetValueLanguages("RU")
	ns.SetValueLanguages("EN")
	assert.Equal(t, len(ns.GetLanguages()), 2)
	ns.SetAnyLanguages("FR")
	assert.Equal(t, len(ns.GetLanguages()), 3)
}

/*
// SoftwareCyberObservableObjectSTIX объект "Software Object", по терминологии STIX, содержит свойства, связанные с
// программным обеспечением, включая программные продукты.
// Name - назвыание программного обеспечения
// CPE - содержит запись Common Platform Enumeration (CPE) для программного обеспечения, если она доступна. Значение этого свойства должно быть значением
// CPE v2.3 из официального словаря NVD CPE [NVD]
// SwID - содержит запись Тегов Software Identification ID (SWID) [SWID] для программного обеспечения, если таковая имеется. SwID помеченный tagId,
// является глобально уникальным идентификатором и ДОЛЖЕН использоваться как полномочие для идентификации помеченного продукта
// Languages -содержит языки, поддерживаемые программным обеспечением. Значение каждого елемента списка ДОЛЖНО быть кодом языка ISO 639-2
// Vendor - содержит название производителя программного обеспечения
// Version - содержит версию ПО
type SoftwareCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Name      string   `json:"name" bson:"name"`
	CPE       string   `json:"cpe" bson:"cpe"`
	SwID      string   `json:"swid" bson:"swid"`
	Vendor    string   `json:"vendor" bson:"vendor"`
	Version   string   `json:"version" bson:"version"`
	Languages []string `json:"languages" bson:"languages"`
}
*/
