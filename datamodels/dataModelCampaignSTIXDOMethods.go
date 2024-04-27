package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- CampaignDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e CampaignDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return e, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e CampaignDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *CampaignDomainObjectsSTIX) Get() *CampaignDomainObjectsSTIX {
	return e
}

// -------- Name property ---------
func (e *CampaignDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *CampaignDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *CampaignDomainObjectsSTIX) GetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Objective property ---------
func (e *CampaignDomainObjectsSTIX) GetObjective() string {
	return e.Objective
}

// SetValueObjective устанавливает значение для поля Objective
func (e *CampaignDomainObjectsSTIX) SetValueObjective(v string) {
	e.Objective = v
}

// SetAnyObjective устанавливает ЛЮБОЕ значение для поля Objective
func (e *CampaignDomainObjectsSTIX) GetAnyObjective(i interface{}) {
	e.Objective = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *CampaignDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *CampaignDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *CampaignDomainObjectsSTIX) GetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Aliases property ---------
func (e *CampaignDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *CampaignDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *CampaignDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- FirstSeen property ---------
func (e *CampaignDomainObjectsSTIX) GetFirstSeen() string {
	return e.FirstSeen
}

// SetValueFirstSeen устанавливает значение в формате RFC3339 для поля FirstSeen
func (e *CampaignDomainObjectsSTIX) SetValueFirstSeen(v string) {
	e.FirstSeen = v
}

// SetAnyFirstSeen устанавливает ЛЮБОЕ значение для поля FirstSeen
func (e *CampaignDomainObjectsSTIX) SetAnyFirstSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- LastSeen property ---------
func (e *CampaignDomainObjectsSTIX) GetLastSeen() string {
	return e.LastSeen
}

// SetValueLastSeen устанавливает значение в формате RFC3339 для поля LastSeen
func (e *CampaignDomainObjectsSTIX) SetValueLastSeen(v string) {
	e.LastSeen = v
}

// SetAnyLastSeen устанавливает ЛЮБОЕ значение для поля LastSeen
func (e *CampaignDomainObjectsSTIX) SetAnyLastSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// ValidateStruct является валидатором параметров содержащихся в типе CampaignDomainObjectsSTIX
func (cstix CampaignDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(campaign--)[0-9a-f|-]+$`).MatchString(cstix.ID)) {
		return false
	}

	return cstix.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (cstix CampaignDomainObjectsSTIX) SanitizeStruct() CampaignDomainObjectsSTIX {
	cstix.CommonPropertiesDomainObjectSTIX = cstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	cstix.Name = commonlibs.StringSanitize(cstix.Name)
	cstix.Description = commonlibs.StringSanitize(cstix.Description)

	if len(cstix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(cstix.Aliases))
		for _, v := range cstix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		cstix.Aliases = aliasesTmp
	}

	cstix.Objective = commonlibs.StringSanitize(cstix.Objective)

	return cstix
}

// GetID возвращает ID STIX объекта
func (cstix CampaignDomainObjectsSTIX) GetID() string {
	return cstix.ID
}

// GetType возвращает Type STIX объекта
func (cstix CampaignDomainObjectsSTIX) GetType() string {
	return cstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (cstix CampaignDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(cstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(cstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", cstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", cstix.Description))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(cstix.Aliases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%s'\n", cstix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%s'\n", cstix.LastSeen))
	str.WriteString(fmt.Sprintf("'objective': '%s'\n", cstix.Objective))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (cstix CampaignDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   cstix.ID,
		"type": cstix.Type,
	}

	if cstix.Name != "" {
		dataForIndex["name"] = cstix.Name
	}

	if cstix.Description != "" {
		dataForIndex["description"] = cstix.Description
	}

	if len(cstix.Aliases) > 0 {
		var strTmp string

		for _, v := range cstix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}
