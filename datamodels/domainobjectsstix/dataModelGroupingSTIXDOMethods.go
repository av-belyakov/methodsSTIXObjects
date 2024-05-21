package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- GroupingDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e GroupingDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e GroupingDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
// Обязательные значения в полях Name, Context, ObjectRefs
func (e *GroupingDomainObjectsSTIX) Get() (*GroupingDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &GroupingDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *GroupingDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *GroupingDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *GroupingDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *GroupingDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *GroupingDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *GroupingDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Context property ---------
func (e *GroupingDomainObjectsSTIX) GetContext() stixhelpers.OpenVocabTypeSTIX {
	return e.Context
}

func (e *GroupingDomainObjectsSTIX) SetValueContext(v stixhelpers.OpenVocabTypeSTIX) {
	e.Context = v
}

// -------- ObjectRefs property ---------
func (e *GroupingDomainObjectsSTIX) GetObjectRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectRefs
}

func (e *GroupingDomainObjectsSTIX) SetValueObjectRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе GroupingDomainObjectsSTIX
func (e GroupingDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(grouping--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if e.Context == "" {
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
func (e GroupingDomainObjectsSTIX) SanitizeStruct() GroupingDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)
	e.Context = e.Context.SanitizeStructOpenVocabTypeSTIX()

	return e
}

// GetID возвращает ID STIX объекта
func (e GroupingDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e GroupingDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e GroupingDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'context': '%s'\n", e.Context))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ObjectRefs)))

	return str.String()
}
