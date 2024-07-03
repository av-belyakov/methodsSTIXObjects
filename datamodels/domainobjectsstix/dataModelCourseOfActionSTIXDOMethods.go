package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- CourseOfActionDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e CourseOfActionDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e CourseOfActionDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Course of Action", по терминалогии STIX, описывающий совокупность действий
// Обязательное значение в поле Name
func (e *CourseOfActionDomainObjectsSTIX) Get() (*CourseOfActionDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &CourseOfActionDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *CourseOfActionDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *CourseOfActionDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *CourseOfActionDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *CourseOfActionDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *CourseOfActionDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *CourseOfActionDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// ValidateStruct является валидатором параметров содержащихся в типе CourseOfActionDomainObjectsSTIX
func (e CourseOfActionDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(course-of-action--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	//обязательное поле
	if e.Name == "" {
		return false
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e CourseOfActionDomainObjectsSTIX) SanitizeStruct() CourseOfActionDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)
	//cstix.Action - ЗАРЕЗЕРВИРОВАНО

	return e
}

// GetID возвращает ID STIX объекта
func (e CourseOfActionDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e CourseOfActionDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e CourseOfActionDomainObjectsSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, e.Name))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'action': '%v'\n", ws, e.Action))

	return str.String()
}
