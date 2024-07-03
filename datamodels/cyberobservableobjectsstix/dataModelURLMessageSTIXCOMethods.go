package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- URLCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e URLCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e URLCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "URL Object", по терминалогии STIX, содержит унифицированный указатель информационного ресурса (URL).
func (e *URLCyberObservableObjectSTIX) Get() (*URLCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- Value property ---------
func (e *URLCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *URLCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *URLCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// ValidateStruct является валидатором параметров содержащихся в типе URLCyberObservableObjectSTIX
func (e URLCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(url--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !govalidator.IsURL(e.Value) {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e URLCyberObservableObjectSTIX) SanitizeStruct() URLCyberObservableObjectSTIX {
	return e
}

// GetID возвращает ID STIX объекта
func (e URLCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e URLCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e URLCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'value': '%s'\n", ws, e.Value))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e URLCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Value != "" {
		dataForIndex["value"] = e.Value
	}

	return dataForIndex
}
