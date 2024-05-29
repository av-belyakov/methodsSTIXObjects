package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- ObservedDataDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e ObservedDataDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e ObservedDataDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Observed Data", по терминалогии STIX, содержит информацию о сущностях связанных с
// кибер безопасностью, таких как файлы, системы или сети. Наблюдаемые данные это не результат анализа или заключение искусственного
// интеллекта, это просто сырая информация без какого-либо контекста.
// Обязательные значения в полях FirstObserved, LastObserved, NumberObserved
func (e *ObservedDataDomainObjectsSTIX) Get() (*ObservedDataDomainObjectsSTIX, error) {
	if e.GetNumberObserved() == 0 {
		err := fmt.Errorf("the required value 'NumberObserved' must not be equal 0")

		return &ObservedDataDomainObjectsSTIX{}, err
	}

	if e.GetFirstObserved() == "1970-01-01T00:00:00+00:00" {
		err := fmt.Errorf("the required value 'FirstObserved' must not be empty")

		return &ObservedDataDomainObjectsSTIX{}, err
	}

	if e.GetLastObserved() == "1970-01-01T00:00:00+00:00" {
		err := fmt.Errorf("the required value 'LastObserved' must not be empty")

		return &ObservedDataDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- NumberObserved property ---------
func (e *ObservedDataDomainObjectsSTIX) GetNumberObserved() int {
	return e.NumberObserved
}

// SetValueNumberObserved устанавливает значение для поля NumberObserved
func (e *ObservedDataDomainObjectsSTIX) SetValueNumberObserved(v int) {
	e.NumberObserved = v
}

// SetAnyNumberObserved устанавливает ЛЮБОЕ значение для поля NumberObserved
func (e *ObservedDataDomainObjectsSTIX) SetAnyNumberObserved(i interface{}) {
	e.NumberObserved = commonlibs.ConversionAnyToInt(i)
}

// -------- FirstObserved property ---------
func (e *ObservedDataDomainObjectsSTIX) GetFirstObserved() string {
	return e.FirstObserved
}

// SetValueFirstObserved устанавливает значение в формате RFC3339 для поля FirstObserved
func (e *ObservedDataDomainObjectsSTIX) SetValueFirstObserved(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.FirstObserved = v

	return nil
}

// SetAnyFirstObserved устанавливает значение для поля FirstObserved
// принимает число (timestamp 13 символов) или строку в формате RFC3339
func (e *ObservedDataDomainObjectsSTIX) SetAnyFirstObserved(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueFirstObserved(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstObserved = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- LastObserved property ---------
func (e *ObservedDataDomainObjectsSTIX) GetLastObserved() string {
	return e.LastObserved
}

// SetValueLastObserved устанавливает значение в формате RFC3339 для поля LastObserved
func (e *ObservedDataDomainObjectsSTIX) SetValueLastObserved(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.LastObserved = v

	return nil
}

// SetAnyLastObserved устанавливает значение для поля LastObserved
// принимает число (timestamp 13 символов) или строку в формате RFC3339
func (e *ObservedDataDomainObjectsSTIX) SetAnyLastObserved(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueLastObserved(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastObserved = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- ObjectRefs property ---------
func (e *ObservedDataDomainObjectsSTIX) GetObjectRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectRefs
}

func (e *ObservedDataDomainObjectsSTIX) SetValueObjectRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = append(e.ObjectRefs, v)
}

func (e *ObservedDataDomainObjectsSTIX) SetFullValueObjectRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе ObservedDataDomainObjectsSTIX
func (e ObservedDataDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(observed-data--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if e.NumberObserved <= 0 {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	for _, v := range e.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e ObservedDataDomainObjectsSTIX) SanitizeStruct() ObservedDataDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	return e
}

// GetID возвращает ID STIX объекта
func (e ObservedDataDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e ObservedDataDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ObservedDataDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'first_observed': '%v'\n", e.FirstObserved))
	str.WriteString(fmt.Sprintf("'last_observed': '%v'\n", e.LastObserved))
	str.WriteString(fmt.Sprintf("'number_observed': '%d'\n", e.NumberObserved))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e ObservedDataDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
