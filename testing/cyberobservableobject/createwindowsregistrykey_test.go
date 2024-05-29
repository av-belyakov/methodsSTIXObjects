package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestWindowsRegistryKeyCyberObservableObjectSTIX(t *testing.T) {
	nwrk := methodstixobjects.NewWindowsRegistryKeyCyberObservableObjectSTIX()

	assert.Equal(t, nwrk.GetType(), "windows-registry-key")
	_, err := nwrk.Get()
	assert.NoError(t, err)

	nwrk.SetAnyNumberOfSubkeys(123)
	assert.Equal(t, nwrk.GetNumberOfSubkeys(), 123)
	nwrk.SetValueNumberOfSubkeys(234)
	assert.Equal(t, nwrk.GetNumberOfSubkeys(), 234)

	nwrk.SetAnyKey("HKEY_LOCAL_MACHINE_1111")
	assert.Equal(t, nwrk.GetKey(), "HKEY_LOCAL_MACHINE_1111")
	nwrk.SetValueKey("HKEY_LOCAL_MACHINE_2211")
	assert.Equal(t, nwrk.GetKey(), "HKEY_LOCAL_MACHINE_2211")

	mt := "2024-02-14T12:03:06+00:00"
	err = nwrk.SetAnyModifiedTime(mt)
	assert.NoError(t, err)
	assert.Equal(t, nwrk.GetModifiedTime(), mt)

	err = nwrk.SetAnyModifiedTime(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nwrk.GetModifiedTime(), "2024-05-28T12:33:44+03:00")

	nwrk.SetValueCreatorUserRef("creater user ref")
	assert.Equal(t, nwrk.GetCreatorUserRef(), stixhelpers.IdentifierTypeSTIX("creater user ref"))

	nwrk.SetValueValues(somecomplextypesstixco.WindowsRegistryValueTypeSTIX{Name: "name_one"})
	nwrk.SetValueValues(somecomplextypesstixco.WindowsRegistryValueTypeSTIX{Name: "name_two"})
	assert.Equal(t, len(nwrk.GetValues()), 2)
	nwrk.SetFullValueValues([]somecomplextypesstixco.WindowsRegistryValueTypeSTIX{
		{Name: "oleg", DataType: "user elem"},
	})
	assert.Equal(t, len(nwrk.GetValues()), 1)
}

/*
// WindowsRegistryKeyCyberObservableObjectSTIX объект "Windows Registry Key Object", по терминалогии STIX. Содержит описание значений полей
// раздела реестра Windows. Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств,определенных ниже,
// должно быть инициализировано при использовании этого объекта.
// Key - содержит полный путь к разделу реестра. Значение ключа,должно быть сохранено в регистре. В название ключа все сокращения должны быть
// раскрыты. Например, вместо HKLM следует использовать HKEY_LOCAL_MACHINE.
// Values - содержит значения, найденные в разделе реестра.
// ModifiedTime - время, в формате "2016-05-12T08:17:27.000Z", последнего изменения раздела реестра.
// CreatorUserRef - содержит ссылку на учетную запись пользователя, из под которой создан раздел реестра. Объект, на который ссылается это свойство, должен иметь тип user-account.
// NumberOfSubkeys - Указывает количество подразделов, содержащихся в разделе реестра.
type WindowsRegistryKeyCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	NumberOfSubkeys int                                                   `json:"number_of_subkeys" bson:"number_of_subkeys"`
	Key             string                                                `json:"key" bson:"key"`
	ModifiedTime    string                                                `json:"modified_time" bson:"modified_time"`
	CreatorUserRef  stixhelpers.IdentifierTypeSTIX                        `json:"creator_user_ref" bson:"creator_user_ref"`
	Values          []somecomplextypesstixco.WindowsRegistryValueTypeSTIX `json:"values" bson:"values"`
}
*/
