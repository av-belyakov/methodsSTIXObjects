package domainobjectsstix

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

// ValidateStruct является валидатором параметров содержащихся в типе CampaignDomainObjectsSTIX
func (e CampaignDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(campaign--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e CampaignDomainObjectsSTIX) SanitizeStruct() CampaignDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(e.Aliases))
		for _, v := range e.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}

		e.Aliases = aliasesTmp
	}

	e.Objective = commonlibs.StringSanitize(e.Objective)

	return e
}

// GetID возвращает ID STIX объекта
func (e CampaignDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e CampaignDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e CampaignDomainObjectsSTIX) ToStringBeautiful() string {
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
	str.WriteString(fmt.Sprintf("'first_seen': '%s'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%s'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'objective': '%s'\n", e.Objective))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e CampaignDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
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
