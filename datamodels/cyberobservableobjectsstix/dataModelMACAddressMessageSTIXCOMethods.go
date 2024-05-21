package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
)

/* --- MACAddressCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e MACAddressCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e MACAddressCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой
// один адрес управления доступом к среде (MAC).
// Обязательные значения в полях Value
func (e *MACAddressCyberObservableObjectSTIX) Get() (*MACAddressCyberObservableObjectSTIX, error) {
	if e.GetValue() == "" {
		err := fmt.Errorf("the required value 'Value' must not be empty")

		return &MACAddressCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- Value property ---------
func (e *MACAddressCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *MACAddressCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *MACAddressCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// ValidateStruct является валидатором параметров содержащихся в типе MACAddressCyberObservableObjectSTIX
func (e MACAddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(mac-addr--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !govalidator.IsMAC(e.Value) {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e MACAddressCyberObservableObjectSTIX) SanitizeStruct() MACAddressCyberObservableObjectSTIX {
	return e
}

// GetID возвращает ID STIX объекта
func (e MACAddressCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e MACAddressCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e MACAddressCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'value': '%s'\n", e.Value))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e MACAddressCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
