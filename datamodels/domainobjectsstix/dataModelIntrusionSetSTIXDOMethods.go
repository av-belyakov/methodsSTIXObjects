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

/* --- IntrusionSetDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e IntrusionSetDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e IntrusionSetDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *IntrusionSetDomainObjectsSTIX) Get() (*IntrusionSetDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &IntrusionSetDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *IntrusionSetDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *IntrusionSetDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *IntrusionSetDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *IntrusionSetDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Aliases property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *IntrusionSetDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *IntrusionSetDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- Goals property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetGoals() []string {
	return e.Goals
}

// SetValueGoals устанавливает значение для поля Goals
func (e *IntrusionSetDomainObjectsSTIX) SetValueGoals(v string) {
	e.Goals = append(e.Goals, v)
}

// SetAnyGoals устанавливает ЛЮБОЕ значение для поля Aliases
func (e *IntrusionSetDomainObjectsSTIX) SetAnyGoals(i interface{}) {
	e.Goals = append(e.Goals, fmt.Sprint(i))
}

// -------- FirstSeen property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetFirstSeen() string {
	return e.FirstSeen
}

// SetValueFirstSeen устанавливает значение в формате RFC3339 для поля FirstSeen
func (e *IntrusionSetDomainObjectsSTIX) SetValueFirstSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.FirstSeen = v

	return nil
}

// SetAnyFirstSeen устанавливает ЛЮБОЕ значение для поля FirstSeen
func (e *IntrusionSetDomainObjectsSTIX) SetAnyFirstSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- LastSeen property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetLastSeen() string {
	return e.LastSeen
}

// SetValueLastSeen устанавливает значение в формате RFC3339 для поля LastSeen
func (e *IntrusionSetDomainObjectsSTIX) SetValueLastSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.LastSeen = v

	return nil
}

// SetAnyLastSeen устанавливает ЛЮБОЕ значение для поля LastSeen
func (e *IntrusionSetDomainObjectsSTIX) SetAnyLastSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- ResourceLevel property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetResourceLevel() stixhelpers.OpenVocabTypeSTIX {
	return e.ResourceLevel
}

func (e *IntrusionSetDomainObjectsSTIX) SetValueResourceLevel(v stixhelpers.OpenVocabTypeSTIX) {
	e.ResourceLevel = v
}

// -------- PrimaryMotivation property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetPrimaryMotivation() stixhelpers.OpenVocabTypeSTIX {
	return e.PrimaryMotivation
}

func (e *IntrusionSetDomainObjectsSTIX) SetValuePrimaryMotivation(v stixhelpers.OpenVocabTypeSTIX) {
	e.PrimaryMotivation = v
}

// -------- SecondaryMotivations property ---------
func (e *IntrusionSetDomainObjectsSTIX) GetSecondaryMotivations() []stixhelpers.OpenVocabTypeSTIX {
	return e.SecondaryMotivations
}

func (e *IntrusionSetDomainObjectsSTIX) SetValueSecondaryMotivations(v []stixhelpers.OpenVocabTypeSTIX) {
	e.SecondaryMotivations = v
}

// ValidateStruct является валидатором параметров содержащихся в типе IntrusionSetDomainObjectsSTIX
func (e IntrusionSetDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(intrusion-set--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if e.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e IntrusionSetDomainObjectsSTIX) SanitizeStruct() IntrusionSetDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(e.Aliases))
		for _, v := range e.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}

		e.Aliases = aliasesTmp
	}

	if len(e.Goals) > 0 {
		goalsTmp := make([]string, 0, len(e.Goals))
		for _, v := range e.Goals {
			goalsTmp = append(goalsTmp, commonlibs.StringSanitize(v))
		}

		e.Goals = goalsTmp
	}

	e.ResourceLevel = e.ResourceLevel.SanitizeStructOpenVocabTypeSTIX()
	e.PrimaryMotivation = e.PrimaryMotivation.SanitizeStructOpenVocabTypeSTIX()

	if len(e.SecondaryMotivations) > 0 {
		sm := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.SecondaryMotivations))
		for _, v := range e.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			sm = append(sm, tmp)
		}

		e.SecondaryMotivations = sm
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e IntrusionSetDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e IntrusionSetDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e IntrusionSetDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Aliases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'goals': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'goal '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Goals)))
	str.WriteString(fmt.Sprintf("'resource_level': '%s'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'primary_motivation': '%s'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'secondary_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'secondary_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.SecondaryMotivations)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e IntrusionSetDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
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

	if len(e.Aliases) > 0 {
		var strTmp string

		for _, v := range e.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}
