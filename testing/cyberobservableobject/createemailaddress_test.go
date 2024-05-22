package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestEmailAddressCyberObservableObjectSTIX(t *testing.T) {
	nea := methodstixobjects.NewEmailAddressCyberObservableObjectSTIX()

	assert.Equal(t, nea.GetType(), "email-addr")
	_, err := nea.Get()
	assert.Error(t, err)

	v := "any value"
	nea.SetAnyValue(v)
	_, err = nea.Get()
	assert.NoError(t, err)
	assert.Equal(t, nea.GetValue(), v)

	dn := "example@mail.ru"
	nea.SetAnyDisplayName(dn)
	assert.Equal(t, nea.GetDisplayName(), dn)

	br := stixhelpers.IdentifierTypeSTIX("Belongs To Ref")
	nea.SetValueBelongsToRef(br)
	assert.Equal(t, nea.GetBelongsToRef(), br)
}

/*
// EmailAddressCyberObservableObjectSTIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Value - email адрес (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// DisplayName - содержит единственное почтовое имя которое видит человек при просмотре письма
// BelongsToRef - учетная запись пользователя, которой принадлежит адрес электронной почты, в качестве ссылки на объект учетной записи пользователя
type EmailAddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value        string                         `json:"value" bson:"value"`
	DisplayName  string                         `json:"display_name" bson:"display_name"`
	BelongsToRef stixhelpers.IdentifierTypeSTIX `json:"belongs_to_ref" bson:"belongs_to_ref"`
}
*/
