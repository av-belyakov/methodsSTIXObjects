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

/* --- ReportDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e ReportDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e ReportDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Report", по терминалогии STIX, содержит совокупность данных об угрозах, сосредоточенных на одной
// или нескольких темах, таких как описание исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним детали.
// Обязательные значения в полях Name, Published, ObjectRefs
func (e *ReportDomainObjectsSTIX) Get() (*ReportDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &ReportDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *ReportDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *ReportDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *ReportDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *ReportDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *ReportDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *ReportDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Published property ---------
func (e *ReportDomainObjectsSTIX) GetPublished() string {
	return e.Published
}

// SetValuePublished устанавливает значение в формате RFC3339 для поля Published
func (e *ReportDomainObjectsSTIX) SetValuePublished(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Published = v

	return nil
}

// SetAnyPublished устанавливает значение для поля Published
// принимает число (timestamp 13 символов) или строку в формате RFC3339
func (e *ReportDomainObjectsSTIX) SetAnyPublished(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValuePublished(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Published = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- ReportTypes property ---------
func (e *ReportDomainObjectsSTIX) GetReportTypes() []stixhelpers.OpenVocabTypeSTIX {
	return e.ReportTypes
}

func (e *ReportDomainObjectsSTIX) SetValueReportTypes(v []stixhelpers.OpenVocabTypeSTIX) {
	e.ReportTypes = v
}

// -------- ObjectRefs property ---------
func (e *ReportDomainObjectsSTIX) GetObjectRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectRefs
}

func (e *ReportDomainObjectsSTIX) SetValueObjectRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе ReportDomainObjectsSTIX
func (e ReportDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(report--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if e.Name == "" {
		return false
	}

	//обязательное поле
	if len(e.ObjectRefs) == 0 {
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
func (e ReportDomainObjectsSTIX) SanitizeStruct() ReportDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)
	if len(e.ReportTypes) > 0 {
		r := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.ReportTypes))
		for _, v := range e.ReportTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			r = append(r, tmp)
		}

		e.ReportTypes = r
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e ReportDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e ReportDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ReportDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'report_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'report_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ReportTypes)))
	str.WriteString(fmt.Sprintf("'published': '%v'\n", e.Published))
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
func (e ReportDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Name != "" {
		dataForIndex["name"] = e.Name
	}

	if e.Description != "" {
		dataForIndex["description"] = e.Description
	}

	return dataForIndex
}
