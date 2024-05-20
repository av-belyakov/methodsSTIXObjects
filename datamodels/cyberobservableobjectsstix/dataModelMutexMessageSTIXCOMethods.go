package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- MutexCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e MutexCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e MutexCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *MutexCyberObservableObjectSTIX) Get() (*MutexCyberObservableObjectSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &MutexCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *MutexCyberObservableObjectSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *MutexCyberObservableObjectSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *MutexCyberObservableObjectSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// ValidateStruct является валидатором параметров содержащихся в типе MutexCyberObservableObjectSTIX
func (e MutexCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(mutex--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e MutexCyberObservableObjectSTIX) SanitizeStruct() MutexCyberObservableObjectSTIX {
	e.Name = commonlibs.StringSanitize(e.Name)

	return e
}

// GetID возвращает ID STIX объекта
func (e MutexCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e MutexCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e MutexCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e MutexCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Name != "" {
		dataForIndex["name"] = e.Name
	}

	return dataForIndex
}
