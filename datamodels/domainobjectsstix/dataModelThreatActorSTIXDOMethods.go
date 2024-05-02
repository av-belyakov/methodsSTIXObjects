package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- ThreatActorDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e ThreatActorDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e ThreatActorDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// -------- Name property ---------
func (e *ThreatActorDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *ThreatActorDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *ThreatActorDomainObjectsSTIX) GetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *ThreatActorDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *ThreatActorDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *ThreatActorDomainObjectsSTIX) GetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- FirstSeen property ---------
func (e *ThreatActorDomainObjectsSTIX) GetFirstSeen() string {
	return e.FirstSeen
}

// SetValueFirstSeen устанавливает значение в формате RFC3339 для поля FirstSeen
func (e *ThreatActorDomainObjectsSTIX) SetValueFirstSeen(v string) {
	e.FirstSeen = v
}

// SetAnyFirstSeen устанавливает ЛЮБОЕ значение для поля FirstSeen
func (e *ThreatActorDomainObjectsSTIX) SetAnyFirstSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- LastSeen property ---------
func (e *ThreatActorDomainObjectsSTIX) GetLastSeen() string {
	return e.LastSeen
}

// SetValueLastSeen устанавливает значение в формате RFC3339 для поля LastSeen
func (e *ThreatActorDomainObjectsSTIX) SetValueLastSeen(v string) {
	e.LastSeen = v
}

// SetAnyLastSeen устанавливает ЛЮБОЕ значение для поля LastSeen
func (e *ThreatActorDomainObjectsSTIX) SetAnyLastSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Aliases property ---------
func (e *ThreatActorDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *ThreatActorDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *ThreatActorDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- Goals property ---------
func (e *ThreatActorDomainObjectsSTIX) GetGoals() []string {
	return e.Goals
}

// SetValueGoals устанавливает значение для поля Goals
func (e *ThreatActorDomainObjectsSTIX) SetValueGoals(v string) {
	e.Goals = append(e.Goals, v)
}

// SetAnyGoals устанавливает ЛЮБОЕ значение для поля Goals
func (e *ThreatActorDomainObjectsSTIX) SetAnyGoals(i interface{}) {
	e.Goals = append(e.Goals, fmt.Sprint(i))
}

// -------- Sophistication property ---------
func (e *ThreatActorDomainObjectsSTIX) GetSophistication() stixhelpers.OpenVocabTypeSTIX {
	return e.Sophistication
}

func (e *ThreatActorDomainObjectsSTIX) SetValueSophistication(v stixhelpers.OpenVocabTypeSTIX) {
	e.Sophistication = v
}

// -------- ResourceLevel property ---------
func (e *ThreatActorDomainObjectsSTIX) GetResourceLevel() stixhelpers.OpenVocabTypeSTIX {
	return e.ResourceLevel
}

func (e *ThreatActorDomainObjectsSTIX) SetValueResourceLevel(v stixhelpers.OpenVocabTypeSTIX) {
	e.ResourceLevel = v
}

// -------- PrimaryMotivation property ---------
func (e *ThreatActorDomainObjectsSTIX) GetPrimaryMotivation() stixhelpers.OpenVocabTypeSTIX {
	return e.PrimaryMotivation
}

func (e *ThreatActorDomainObjectsSTIX) SetValuePrimaryMotivation(v stixhelpers.OpenVocabTypeSTIX) {
	e.PrimaryMotivation = v
}

// -------- SecondaryMotivations property ---------
func (e *ThreatActorDomainObjectsSTIX) GetSecondaryMotivations() []stixhelpers.OpenVocabTypeSTIX {
	return e.SecondaryMotivations
}

func (e *ThreatActorDomainObjectsSTIX) SetValueSecondaryMotivations(v []stixhelpers.OpenVocabTypeSTIX) {
	e.SecondaryMotivations = v
}

// -------- PersonalMotivations property ---------
func (e *ThreatActorDomainObjectsSTIX) GetPersonalMotivations() []stixhelpers.OpenVocabTypeSTIX {
	return e.PersonalMotivations
}

func (e *ThreatActorDomainObjectsSTIX) SetValuePersonalMotivations(v []stixhelpers.OpenVocabTypeSTIX) {
	e.PersonalMotivations = v
}

// -------- ThreatActorTypes property ---------
func (e *ThreatActorDomainObjectsSTIX) GetThreatActorTypes() []stixhelpers.OpenVocabTypeSTIX {
	return e.ThreatActorTypes
}

func (e *ThreatActorDomainObjectsSTIX) SetValueThreatActorTypes(v []stixhelpers.OpenVocabTypeSTIX) {
	e.ThreatActorTypes = v
}

// -------- Roles property ---------
func (e *ThreatActorDomainObjectsSTIX) GetRoles() []stixhelpers.OpenVocabTypeSTIX {
	return e.Roles
}

func (e *ThreatActorDomainObjectsSTIX) SetValueRoles(v []stixhelpers.OpenVocabTypeSTIX) {
	e.Roles = v
}

// ValidateStruct является валидатором параметров содержащихся в типе ThreatActorDomainObjectsSTIX
func (e ThreatActorDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(threat-actor--)[0-9a-f|-]+$`).MatchString(e.ID)) {
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
func (e ThreatActorDomainObjectsSTIX) SanitizeStruct() ThreatActorDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.ThreatActorTypes) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.ThreatActorTypes))
		for _, v := range e.ThreatActorTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		e.ThreatActorTypes = ta
	}

	if len(e.Aliases) > 0 {
		mTmp := make([]string, 0, len(e.Aliases))
		for _, v := range e.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}

		e.Aliases = mTmp
	}

	if len(e.Roles) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.Roles))
		for _, v := range e.Roles {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		e.Roles = ta
	}

	if len(e.Goals) > 0 {
		mTmp := make([]string, 0, len(e.Goals))
		for _, v := range e.Goals {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}

		e.Goals = mTmp
	}

	e.Sophistication = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(e.Sophistication)))
	e.ResourceLevel = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(e.ResourceLevel)))
	e.PrimaryMotivation = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(e.PrimaryMotivation)))

	if len(e.SecondaryMotivations) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.SecondaryMotivations))
		for _, v := range e.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		e.SecondaryMotivations = ta
	}

	if len(e.PersonalMotivations) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.PersonalMotivations))
		for _, v := range e.PersonalMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		e.PersonalMotivations = ta
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e ThreatActorDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e ThreatActorDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ThreatActorDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'threat_actor_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\tt'hreat_actor_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ThreatActorTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Aliases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'roles': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'role '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.Roles)))
	str.WriteString(fmt.Sprintf("'goals': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'goal '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Goals)))
	str.WriteString(fmt.Sprintf("'sophistication': '%v'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'resource_level': '%v'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'primary_motivation': '%v'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'secondary_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'secondary_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.SecondaryMotivations)))
	str.WriteString(fmt.Sprintf("'personal_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'personal_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.PersonalMotivations)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e ThreatActorDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
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

	if len(e.Goals) > 0 {
		var strTmp string

		for _, v := range e.Goals {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] += " " + strTmp
	}

	return dataForIndex
}
