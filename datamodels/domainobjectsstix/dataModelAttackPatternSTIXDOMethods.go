package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- AttackPatternDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e AttackPatternDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return e, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e AttackPatternDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *AttackPatternDomainObjectsSTIX) Get() (*AttackPatternDomainObjectsSTIX, error) {
	if e.GetName() == "" {
		err := fmt.Errorf("the required value 'Name' must not be empty")

		return &AttackPatternDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Name property ---------
func (e *AttackPatternDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *AttackPatternDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *AttackPatternDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *AttackPatternDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *AttackPatternDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *AttackPatternDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Aliases property ---------
func (e *AttackPatternDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *AttackPatternDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *AttackPatternDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- KillChainPhases property ---------
func (e *AttackPatternDomainObjectsSTIX) GetKillChainPhases() []stixhelpers.KillChainPhasesTypeElementSTIX {
	return e.KillChainPhases
}

func (e *AttackPatternDomainObjectsSTIX) SetValueKillChainPhases(v []stixhelpers.KillChainPhasesTypeElementSTIX) {
	e.KillChainPhases = v
}

// ValidateStruct является валидатором параметров содержащихся в типе AttackPatternDomainObjectsSTIX
// возвращает ВАЛИДНЫЙ объект AttackPatternDomainObjectsSTIX (к сожалению нельзя править существующий объект
// из-за ошибки 'cannot use e (variable of type datamodels.AttackPatternDomainObjectsSTIX) as datamodels.HandlerSTIXObject
// value in struct literal: missing method ValidateStruct (ValidateStruct has pointer receiver)' возникающей в
// функции GetListSTIXObjectFromJSON если приемник ValidateStruct работает по ссылке)
func (e AttackPatternDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(attack-pattern--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	//обязательное поле
	if e.Name == "" {
		return false
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e AttackPatternDomainObjectsSTIX) SanitizeStruct() AttackPatternDomainObjectsSTIX {
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

	newKillChainPhases := make([]stixhelpers.KillChainPhasesTypeElementSTIX, 0, len(e.KillChainPhases))
	for _, v := range e.KillChainPhases {
		newKillChainPhases = append(newKillChainPhases, v.SanitizeStructKillChainPhasesTypeElementSTIX())
	}

	e.KillChainPhases = newKillChainPhases

	return e
}

// GetID возвращает ID STIX объекта
func (e AttackPatternDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e AttackPatternDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e AttackPatternDomainObjectsSTIX) ToStringBeautiful() string {
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
	str.WriteString(fmt.Sprintf("'kill_chain_phases': \n%v", func(l stixhelpers.KillChainPhasesTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'key': '%v' 'kill_chain_name': '%s'\n", k, v.KillChainName))
			str.WriteString(fmt.Sprintf("\t'key': '%v' 'phase_name': '%s'\n", k, v.PhaseName))
		}

		return str.String()
	}(e.KillChainPhases)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e AttackPatternDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
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
