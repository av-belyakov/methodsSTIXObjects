package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- EmailAddressCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e EmailAddressCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e EmailAddressCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Обязательные значения в полях Value
func (e *EmailAddressCyberObservableObjectSTIX) Get() (*EmailAddressCyberObservableObjectSTIX, error) {
	if e.GetValue() == "" {
		err := fmt.Errorf("the required value 'Value' must not be empty")

		return &EmailAddressCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- Value property ---------
func (e *EmailAddressCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *EmailAddressCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *EmailAddressCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// -------- Value property ---------
func (e *EmailAddressCyberObservableObjectSTIX) GetDisplayName() string {
	return e.DisplayName
}

// SetValueDisplayName устанавливает значение для поля DisplayName
func (e *EmailAddressCyberObservableObjectSTIX) SetValueDisplayName(v string) {
	e.DisplayName = v
}

// SetAnyDisplayName устанавливает ЛЮБОЕ значение для поля DisplayName
func (e *EmailAddressCyberObservableObjectSTIX) SetAnyDisplayName(i interface{}) {
	e.DisplayName = fmt.Sprint(i)
}

// -------- BelongsToRef property ---------
func (e *EmailAddressCyberObservableObjectSTIX) GetBelongsToRef() stixhelpers.IdentifierTypeSTIX {
	return e.BelongsToRef
}

func (e *EmailAddressCyberObservableObjectSTIX) SetValueBelongsToRef(v stixhelpers.IdentifierTypeSTIX) {
	e.BelongsToRef = v
}

// ValidateStruct является валидатором параметров содержащихся в типе EmailAddressCyberObservableObjectSTIX
func (e EmailAddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(email-addr--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if !govalidator.IsEmail(e.Value) {
		return false
	}

	if !e.BelongsToRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e EmailAddressCyberObservableObjectSTIX) SanitizeStruct() EmailAddressCyberObservableObjectSTIX {
	e.DisplayName = commonlibs.StringSanitize(e.DisplayName)

	return e
}

// GetID возвращает ID STIX объекта
func (e EmailAddressCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e EmailAddressCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e EmailAddressCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'value': '%s'\n", ws, e.Value))
	str.WriteString(fmt.Sprintf("%s'display_name': '%s'\n", ws, e.DisplayName))
	str.WriteString(fmt.Sprintf("%s'belongs_to_ref': '%v'\n", ws, e.BelongsToRef))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e EmailAddressCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Value != "" {
		dataForIndex["value"] = e.Value
	}

	return dataForIndex
}
