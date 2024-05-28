package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestEmailMessageCyberObservableObjectSTIX(t *testing.T) {
	nem := methodstixobjects.NewEmailMessageCyberObservableObjectSTIX()

	assert.Equal(t, nem.GetType(), "email-message")
	_, err := nem.Get()
	assert.NoError(t, err)

	nem.SetValueIsMultipart(true)
	assert.True(t, nem.GetIsMultipart())
	nem.SetAnyIsMultipart(false)
	assert.False(t, nem.GetIsMultipart())

	cte := "any content type example"
	nem.SetAnyContentType(cte)
	assert.Equal(t, nem.GetContentType(), cte)
	nem.SetValueContentType("d6w")
	assert.Equal(t, nem.GetContentType(), "d6w")

	mide := "message id example"
	nem.SetAnyMessageID(mide)
	assert.Equal(t, nem.GetMessageID(), mide)
	nem.SetValueMessageID("me")
	assert.Equal(t, nem.GetMessageID(), "me")

	se := "subject example"
	nem.SetAnySubject(se)
	assert.Equal(t, nem.GetSubject(), se)
	nem.SetValueSubject("subject")
	assert.Equal(t, nem.GetSubject(), "subject")

	be := "body example"
	nem.SetAnyBody(be)
	assert.Equal(t, nem.GetBody(), be)
	nem.SetValueBody("body")
	assert.Equal(t, nem.GetBody(), "body")

	dt := "2024-02-11T07:01:01+00:00"
	err = nem.SetAnyDate(dt)
	assert.NoError(t, err)
	assert.Equal(t, nem.GetDate(), dt)

	err = nem.SetAnyDate(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nem.GetDate(), "2024-05-22T11:43:27+03:00")

	nem.SetAnyReceivedLines("R_Lines_1")
	nem.SetAnyReceivedLines("R_Lines_2")
	nem.SetAnyReceivedLines("R_Lines_3")
	nem.SetValueReceivedLines("R_Lines_4")
	assert.Equal(t, len(nem.GetReceivedLines()), 4)

	fre := stixhelpers.IdentifierTypeSTIX("from ref example")
	nem.SetValueFromRef(fre)
	assert.Equal(t, nem.GetFromRef(), fre)

	sre := stixhelpers.IdentifierTypeSTIX("sender ref example")
	nem.SetValueSenderRef(sre)
	assert.Equal(t, nem.GetSenderRef(), sre)

	rere := stixhelpers.IdentifierTypeSTIX("raw email ref example")
	nem.SetValueRawEmailRef(rere)
	assert.Equal(t, nem.GetRawEmailRef(), rere)

	nem.SetValueToRefs([]stixhelpers.IdentifierTypeSTIX{"to-ref_1", "to-ref_2", "to-ref_3"})
	assert.Equal(t, len(nem.GetToRefs()), 3)

	nem.SetValueCcRefs([]stixhelpers.IdentifierTypeSTIX{"cc-ref_1", "cc-ref_2"})
	assert.Equal(t, len(nem.GetCcRefs()), 2)

	nem.SetValueBccRefs([]stixhelpers.IdentifierTypeSTIX{"bbc-ref_1"})
	assert.Equal(t, len(nem.GetBccRefs()), 1)

	nem.SetValueBodyMultipart([]somecomplextypesstixco.EmailMIMEPartTypeSTIX{
		{Body: "any body 1", ContentType: "content type example"},
		{Body: "any body 2", ContentType: "content type example"},
	})
	assert.Equal(t, len(nem.GetBodyMultipart()), 2)

	nem.SetValueAdditionalHeaderFields(map[string]stixhelpers.DictionaryTypeSTIX{
		"book_1": {Dictionary: "dictionary-1"},
		"book_2": {Dictionary: "dictionary-2"},
	})
	assert.Equal(t, nem.GetAdditionalHeaderFields()["book_1"].Dictionary, "dictionary-1")
}

/*
// EmailMessageCyberObservableObjectSTIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
// IsMultipart - информирует содержит ли 'тело' email множественные MIME части (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Date - время, в формате "2016-05-12T08:17:27.000Z", когда email сообщение было отправлено
// ContentType - содержит содержимое 'Content-Type' заголовка email сообщения
// FromRef - содержит содержимое 'From:' заголовка email сообщения
// SenderRef - содержит содержимое поля 'Sender:' email сообщения
// ToRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'To:'
// CcRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'CC:'
// BccRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'BCC:'
// MessageID - содержимое Message-ID email сообщения
// Subject - содержит тему сообщения
// ReceivedLines - содержит одно или несколько полей заголовка 'Received', которые могут быть включены в заголовки email
// AdditionalHeaderFields - содержит любые другие поля заголовка (за исключением date, received_lines, content_type, from_ref,
// sender_ref, to_ref, cc_ref, bcc_refs и subject), найденные в сообщении электронной почты в виде словаря
// Body - содержит тело сообщения
// BodyMultipart - содержит адает список MIME-части, которые составляют тело email. Это свойство НЕ ДОЛЖНО использоваться,
// если	is_multipart имеет значение false
// RawEmailRef - содержит 'сырое' бинарное содержимое email сообщения
type EmailMessageCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsMultipart            bool                                           `json:"is_multipart" bson:"is_multipart" required:"true"`
	ContentType            string                                         `json:"content_type" bson:"content_type"`
	MessageID              string                                         `json:"message_id" bson:"message_id"`
	Subject                string                                         `json:"subject" bson:"subject"`
	Body                   string                                         `json:"body" bson:"body"`
	Date                   string                                         `json:"date" bson:"date"`
	ReceivedLines          []string                                       `json:"received_lines" bson:"received_lines"`
	FromRef                stixhelpers.IdentifierTypeSTIX                 `json:"from_ref" bson:"from_ref"`
	SenderRef              stixhelpers.IdentifierTypeSTIX                 `json:"sender_ref" bson:"sender_ref"`
	RawEmailRef            stixhelpers.IdentifierTypeSTIX                 `json:"raw_email_ref" bson:"raw_email_ref"`
	ToRefs                 []stixhelpers.IdentifierTypeSTIX               `json:"to_refs" bson:"to_refs"`
	CcRefs                 []stixhelpers.IdentifierTypeSTIX               `json:"cc_refs" bson:"cc_refs"`
	BccRefs                []stixhelpers.IdentifierTypeSTIX               `json:"bcc_refs" bson:"bcc_refs"`
	BodyMultipart          []somecomplextypesstixco.EmailMIMEPartTypeSTIX `json:"body_multipart" bson:"body_multipart"`
	AdditionalHeaderFields map[string]stixhelpers.DictionaryTypeSTIX      `json:"additional_header_fields" bson:"additional_header_fields"`
}
*/
