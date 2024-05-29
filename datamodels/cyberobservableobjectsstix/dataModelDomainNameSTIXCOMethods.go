package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- DomainNameCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e DomainNameCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e DomainNameCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Обязательные значения в полях Value
func (e *DomainNameCyberObservableObjectSTIX) Get() (*DomainNameCyberObservableObjectSTIX, error) {
	if e.GetValue() == "" {
		err := fmt.Errorf("the required value 'Value' must not be empty")

		return &DomainNameCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- Value property ---------
func (e *DomainNameCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *DomainNameCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *DomainNameCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// -------- ResolvesToRefs property ---------
func (e *DomainNameCyberObservableObjectSTIX) GetResolvesToRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ResolvesToRefs
}

func (e *DomainNameCyberObservableObjectSTIX) SetValueResolvesToRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ResolvesToRefs = append(e.ResolvesToRefs, v)
}

func (e *DomainNameCyberObservableObjectSTIX) SetFullValueResolvesToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ResolvesToRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе DomainNameCyberObservableObjectSTIX
func (e DomainNameCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(domain-name--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if !govalidator.IsDNSName(e.Value) {
		return false
	}

	if len(e.ResolvesToRefs) > 0 {
		for _, v := range e.ResolvesToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e DomainNameCyberObservableObjectSTIX) SanitizeStruct() DomainNameCyberObservableObjectSTIX {
	return e
}

// GetID возвращает ID STIX объекта
func (e DomainNameCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e DomainNameCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e DomainNameCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'value': '%s'\n", e.Value))
	str.WriteString(fmt.Sprintf("'resolves_to_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'resolves_to_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ResolvesToRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e DomainNameCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Value != "" {
		dataForIndex["value"] = e.Value
	}

	return dataForIndex
}
