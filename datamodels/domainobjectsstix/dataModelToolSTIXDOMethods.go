package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- ToolDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (tstix ToolDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &tstix); err != nil {
		return nil, err
	}

	return tstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (tstix ToolDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(tstix)

	return &result, err
}

// -------- Name property ---------
func (e *ToolDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *ToolDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *ToolDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *ToolDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *ToolDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *ToolDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- ToolVersion property ---------
func (e *ToolDomainObjectsSTIX) GetToolVersion() string {
	return e.ToolVersion
}

// SetValueToolVersion устанавливает значение для поля ToolVersion
func (e *ToolDomainObjectsSTIX) SetValueToolVersion(v string) {
	e.ToolVersion = v
}

// SetAnyToolVersion устанавливает ЛЮБОЕ значение для поля ToolVersion
func (e *ToolDomainObjectsSTIX) SetAnyToolVersion(i interface{}) {
	e.ToolVersion = fmt.Sprint(i)
}

// -------- Aliases property ---------
func (e *ToolDomainObjectsSTIX) GetAliases() []string {
	return e.Aliases
}

// SetValueAliases устанавливает значение для поля Aliases
func (e *ToolDomainObjectsSTIX) SetValueAliases(v string) {
	e.Aliases = append(e.Aliases, v)
}

// SetAnyAliases устанавливает ЛЮБОЕ значение для поля Aliases
func (e *ToolDomainObjectsSTIX) SetAnyAliases(i interface{}) {
	e.Aliases = append(e.Aliases, fmt.Sprint(i))
}

// -------- KillChainPhases property ---------
func (e *ToolDomainObjectsSTIX) GetKillChainPhases() stixhelpers.KillChainPhasesTypeSTIX {
	return e.KillChainPhases
}

func (e *ToolDomainObjectsSTIX) SetValueKillChainPhases(v stixhelpers.KillChainPhasesTypeSTIX) {
	e.KillChainPhases = v
}

// -------- ToolTypes property ---------
func (e *ToolDomainObjectsSTIX) GetToolTypes() []stixhelpers.OpenVocabTypeSTIX {
	return e.ToolTypes
}

func (e *ToolDomainObjectsSTIX) SetValueToolTypes(v []stixhelpers.OpenVocabTypeSTIX) {
	e.ToolTypes = v
}

// ValidateStruct является валидатором параметров содержащихся в типе ToolDomainObjectsSTIX
func (e ToolDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(tool--)[0-9a-f|-]+$`).MatchString(e.ID)) {
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
func (e ToolDomainObjectsSTIX) SanitizeStruct() ToolDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.ToolTypes) > 0 {
		t := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.ToolTypes))
		for _, v := range e.ToolTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			t = append(t, tmp)
		}

		e.ToolTypes = t
	}

	if len(e.Aliases) > 0 {
		mTmp := make([]string, 0, len(e.Aliases))
		for _, v := range e.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}

		e.Aliases = mTmp
	}

	e.KillChainPhases = e.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()
	e.ToolVersion = commonlibs.StringSanitize(e.ToolVersion)

	return e
}

// GetID возвращает ID STIX объекта
func (e ToolDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e ToolDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ToolDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'tool_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'tool_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ToolTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Aliases)))
	str.WriteString(fmt.Sprintf("'kill_chain_phases': \n%v", func(l []stixhelpers.KillChainPhasesTypeElementSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'kill_chain_phase '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.KillChainPhases)))
	str.WriteString(fmt.Sprintf("'tool_version': '%s'\n", e.ToolVersion))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e ToolDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
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
