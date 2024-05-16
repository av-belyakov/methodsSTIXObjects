package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- IPv6AddressCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e IPv6AddressCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e IPv6AddressCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *IPv6AddressCyberObservableObjectSTIX) Get() *IPv6AddressCyberObservableObjectSTIX {
	return e
}

// -------- Value property ---------
func (e *IPv6AddressCyberObservableObjectSTIX) GetValue() string {
	return e.Value
}

// SetValueValue устанавливает значение для поля Value
func (e *IPv6AddressCyberObservableObjectSTIX) SetValueValue(v string) {
	e.Value = v
}

// SetAnyValue устанавливает ЛЮБОЕ значение для поля Value
func (e *IPv6AddressCyberObservableObjectSTIX) SetAnyValue(i interface{}) {
	e.Value = fmt.Sprint(i)
}

// -------- ResolvesToRefs property ---------
func (e *IPv6AddressCyberObservableObjectSTIX) GetContainsRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ResolvesToRefs
}

func (e *IPv6AddressCyberObservableObjectSTIX) SetValueResolvesToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ResolvesToRefs = v
}

// -------- BelongsToRefs property ---------
func (e *IPv6AddressCyberObservableObjectSTIX) GetBelongsToRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.BelongsToRefs
}

func (e *IPv6AddressCyberObservableObjectSTIX) SetValueBelongsToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.BelongsToRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе IPv6AddressCyberObservableObjectSTIX
func (e IPv6AddressCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(ipv6-addr--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Value == "" {
		return false
	}

	if ipv6Addr, _, err := net.ParseCIDR(e.Value); err == nil {
		if !govalidator.IsIPv6(ipv6Addr.String()) {
			return false
		}
	} else {
		if !govalidator.IsIPv6(e.Value) {
			return false
		}
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
func (e IPv6AddressCyberObservableObjectSTIX) SanitizeStruct() IPv6AddressCyberObservableObjectSTIX {
	return e
}

// GetID возвращает ID STIX объекта
func (e IPv6AddressCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e IPv6AddressCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e IPv6AddressCyberObservableObjectSTIX) ToStringBeautiful() string {
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
	str.WriteString(fmt.Sprintf("'belongs_to_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'belongs_to_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.BelongsToRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e IPv6AddressCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Value != "" {
		dataForIndex["value"] = e.Value
	}

	return dataForIndex
}
