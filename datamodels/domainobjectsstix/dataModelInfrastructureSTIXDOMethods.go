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

/* --- InfrastructureDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e InfrastructureDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e InfrastructureDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Infrastructure", по терминалогии STIX, содержит описание любых систем, программных
// служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
// Обязательные значения в полях Name
func (e *InfrastructureDomainObjectsSTIX) Get() (*InfrastructureDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &InfrastructureDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *InfrastructureDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *InfrastructureDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *InfrastructureDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *InfrastructureDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *InfrastructureDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *InfrastructureDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Aliases property ---------
func (e *InfrastructureDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *InfrastructureDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *InfrastructureDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- FirstSeen property ---------
func (e *InfrastructureDomainObjectsSTIX) GetFirstSeen() string {
	return e.FirstSeen
}

// SetValueFirstSeen устанавливает значение в формате RFC3339 для поля FirstSeen
func (e *InfrastructureDomainObjectsSTIX) SetValueFirstSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.FirstSeen = v

	return nil
}

// SetAnyFirstSeen устанавливает значение для поля FirstSeen
// принимает число (timestamp 13 символов) или строку в формате RFC3339
func (e *InfrastructureDomainObjectsSTIX) SetAnyFirstSeen(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueFirstSeen(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- LastSeen property ---------
func (e *InfrastructureDomainObjectsSTIX) GetLastSeen() string {
	return e.LastSeen
}

// SetValueLastSeen устанавливает значение в формате RFC3339 для поля LastSeen
func (e *InfrastructureDomainObjectsSTIX) SetValueLastSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.LastSeen = v

	return nil
}

// SetAnyLastSeen устанавливает значение для поля LastSeen
// принимает число (timestamp 13 символов) или строку в формате RFC3339
func (e *InfrastructureDomainObjectsSTIX) SetAnyLastSeen(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueLastSeen(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- KillChainPhases property ---------
func (e *InfrastructureDomainObjectsSTIX) GetKillChainPhases() []stixhelpers.KillChainPhasesTypeElementSTIX {
	return e.KillChainPhases
}

func (e *InfrastructureDomainObjectsSTIX) SetValueKillChainPhases(v []stixhelpers.KillChainPhasesTypeElementSTIX) {
	e.KillChainPhases = v
}

// -------- InfrastructureTypes property ---------
func (e *InfrastructureDomainObjectsSTIX) GetIndicatorTypes() []stixhelpers.OpenVocabTypeSTIX {
	return e.InfrastructureTypes
}

func (e *InfrastructureDomainObjectsSTIX) SetValueIndicatorTypes(v []stixhelpers.OpenVocabTypeSTIX) {
	e.InfrastructureTypes = v
}

// ValidateStruct является валидатором параметров содержащихся в типе InfrastructureDomainObjectsSTIX
func (e InfrastructureDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(infrastructure--)[0-9a-f|-]+$`).MatchString(e.ID)) {
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
func (e InfrastructureDomainObjectsSTIX) SanitizeStruct() InfrastructureDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.InfrastructureTypes) > 0 {
		it := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.InfrastructureTypes))
		for _, v := range e.InfrastructureTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			it = append(it, tmp)
		}

		e.InfrastructureTypes = it
	}

	if len(e.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(e.Aliases))
		for _, v := range e.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}

		e.Aliases = aliasesTmp
	}

	killChainPhases := []stixhelpers.KillChainPhasesTypeElementSTIX(nil)
	for _, v := range e.KillChainPhases {
		killChainPhases = append(killChainPhases, v.SanitizeStructKillChainPhasesTypeElementSTIX())
	}
	e.KillChainPhases = killChainPhases

	return e
}

// GetID возвращает ID STIX объекта
func (e InfrastructureDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e InfrastructureDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e InfrastructureDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'infrastructure_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'infrastructure_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.InfrastructureTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Aliases)))
	str.WriteString(fmt.Sprintf("'sectors': \n%v", func(l []stixhelpers.KillChainPhasesTypeElementSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'sector '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.KillChainPhases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", e.LastSeen))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (istix InfrastructureDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   istix.ID,
		"type": istix.Type,
	}

	if istix.Name != "" {
		dataForIndex["name"] = istix.Name
	}

	if istix.Description != "" {
		dataForIndex["description"] = istix.Description
	}

	if len(istix.Aliases) > 0 {
		var strTmp string

		for _, v := range istix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}
