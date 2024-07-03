package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- IPv4AddressCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e IPv4AddressCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e IPv4AddressCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "IPv4 Address Object", по терминалогии STIX, содержит один или
// более IPv4 адресов, выраженных с помощью нотации CIDR.
// Обязательные значения в полях Value
func (e *IPv4AddressCyberObservableObjectSTIX) Get() (*IPv4AddressCyberObservableObjectSTIX, error) {
	if e.GetValue() == "" {
		err := fmt.Errorf("the required value 'Value' must not be empty")

		return &IPv4AddressCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- Value property ---------
func (e *IPv4AddressCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *IPv4AddressCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *IPv4AddressCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// -------- ResolvesToRefs property ---------
func (e *IPv4AddressCyberObservableObjectSTIX) GetResolvesToRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ResolvesToRefs
}

func (e *IPv4AddressCyberObservableObjectSTIX) SetValueResolvesToRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ResolvesToRefs = append(e.ResolvesToRefs, v)
}

func (e *IPv4AddressCyberObservableObjectSTIX) SetFullValueResolvesToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ResolvesToRefs = v
}

// -------- BelongsToRefs property ---------
func (e *IPv4AddressCyberObservableObjectSTIX) GetBelongsToRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.BelongsToRefs
}

func (e *IPv4AddressCyberObservableObjectSTIX) SetValueBelongsToRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.BelongsToRefs = append(e.BelongsToRefs, v)
}

func (e *IPv4AddressCyberObservableObjectSTIX) SetFullValueBelongsToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.BelongsToRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе IPv4AddressCyberObservableObjectSTIX
func (e IPv4AddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(ipv4-addr--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Value == "" {
		return false
	}

	isIPv4 := commonlibs.IsIPv4Address(e.Value)
	isNetworkIPv4 := commonlibs.IsComputerNetAddrIPv4Range(e.Value)
	if !isIPv4 && !isNetworkIPv4 {
		return false
	}

	if len(e.ResolvesToRefs) > 0 {
		for _, v := range e.ResolvesToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.BelongsToRefs) > 0 {
		for _, v := range e.BelongsToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e IPv4AddressCyberObservableObjectSTIX) SanitizeStruct() IPv4AddressCyberObservableObjectSTIX {
	return e
}

// GetID возвращает ID STIX объекта
func (e IPv4AddressCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e IPv4AddressCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e IPv4AddressCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'value': '%s'\n", ws, e.Value))
	str.WriteString(fmt.Sprintf("%s'resolves_to_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'resolves_to_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.ResolvesToRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'belongs_to_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'belongs_to_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.BelongsToRefs, num+1)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e IPv4AddressCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Value != "" {
		dataForIndex["value"] = e.Value
	}

	return dataForIndex
}
