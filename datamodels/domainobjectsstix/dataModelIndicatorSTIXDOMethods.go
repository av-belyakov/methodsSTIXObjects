package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- IndicatorDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e IndicatorDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e IndicatorDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *IndicatorDomainObjectsSTIX) Get() *IndicatorDomainObjectsSTIX {
	return e
}

// -------- Name property ---------
func (e *IndicatorDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *IndicatorDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *IndicatorDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Pattern property ---------
func (e *IndicatorDomainObjectsSTIX) GetPattern() string {
	return e.Pattern
}

// SetValuePattern устанавливает значение для поля Pattern
func (e *IndicatorDomainObjectsSTIX) SetValuePattern(v string) {
	e.Pattern = v
}

// SetAnyPattern устанавливает ЛЮБОЕ значение для поля Pattern
func (e *IndicatorDomainObjectsSTIX) SetAnyPattern(i interface{}) {
	e.Pattern = fmt.Sprint(i)
}

// -------- PatternVersion property ---------
func (e *IndicatorDomainObjectsSTIX) GetPatternVersion() string {
	return e.PatternVersion
}

// SetValuePatternVersion устанавливает значение для поля PatternVersion
func (e *IndicatorDomainObjectsSTIX) SetValuePatternVersion(v string) {
	e.PatternVersion = v
}

// SetAnyPatternVersion устанавливает ЛЮБОЕ значение для поля PatternVersion
func (e *IndicatorDomainObjectsSTIX) SetAnyPatternVersion(i interface{}) {
	e.PatternVersion = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *IndicatorDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *IndicatorDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *IndicatorDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- ValidFrom property ---------
func (e *IndicatorDomainObjectsSTIX) GetValidFrom() string {
	return e.ValidFrom
}

// SetValueValidFrom устанавливает значение в формате RFC3339 для поля ValidFrom
func (e *IndicatorDomainObjectsSTIX) SetValueValidFrom(v string) {
	e.ValidFrom = v
}

// SetAnyValidFrom устанавливает ЛЮБОЕ значение для поля ValidFrom
func (e *IndicatorDomainObjectsSTIX) SetAnyValidFrom(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.ValidFrom = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- ValidUntil property ---------
func (e *IndicatorDomainObjectsSTIX) GetValidUntil() string {
	return e.ValidUntil
}

// SetValueValidUntil устанавливает значение в формате RFC3339 для поля ValidUntil
func (e *IndicatorDomainObjectsSTIX) SetValueValidUntil(v string) {
	e.ValidUntil = v
}

// SetAnyValidUntil устанавливает ЛЮБОЕ значение для поля ValidUntil
func (e *IndicatorDomainObjectsSTIX) SetAnyValidUntil(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.ValidUntil = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- PatternType property ---------
func (e *IndicatorDomainObjectsSTIX) GetPatternType() stixhelpers.OpenVocabTypeSTIX {
	return e.PatternType
}

func (e *IndicatorDomainObjectsSTIX) SetValuePatternType(v stixhelpers.OpenVocabTypeSTIX) {
	e.PatternType = v
}

// -------- KillChainPhases property ---------
func (e *IndicatorDomainObjectsSTIX) GetKillChainPhases() stixhelpers.KillChainPhasesTypeSTIX {
	return e.KillChainPhases
}

func (e *IndicatorDomainObjectsSTIX) SetValueKillChainPhases(v stixhelpers.KillChainPhasesTypeSTIX) {
	e.KillChainPhases = v
}

// -------- IndicatorTypes property ---------
func (e *IndicatorDomainObjectsSTIX) GetIndicatorTypes() []stixhelpers.OpenVocabTypeSTIX {
	return e.IndicatorTypes
}

func (e *IndicatorDomainObjectsSTIX) SetValueIndicatorTypes(v []stixhelpers.OpenVocabTypeSTIX) {
	e.IndicatorTypes = v
}

// ValidateStruct является валидатором параметров содержащихся в типе IndicatorDomainObjectsSTIX
func (e IndicatorDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(indicator--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if e.Pattern == "" {
		return false
	}

	//обязательное поле
	if e.PatternType == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e IndicatorDomainObjectsSTIX) SanitizeStruct() IndicatorDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.IndicatorTypes) > 0 {
		it := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.IndicatorTypes))
		for _, v := range e.IndicatorTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			it = append(it, tmp)
		}

		e.IndicatorTypes = it
	}

	e.Pattern = commonlibs.StringSanitize(e.Pattern)
	e.PatternType = e.PatternType.SanitizeStructOpenVocabTypeSTIX()
	e.PatternVersion = commonlibs.StringSanitize(e.PatternVersion)
	e.KillChainPhases = e.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	return e
}

// GetID возвращает ID STIX объекта
func (e IndicatorDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e IndicatorDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e IndicatorDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'indicator_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'indicator_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.IndicatorTypes)))
	str.WriteString(fmt.Sprintf("'pattern': '%s'\n", e.Pattern))
	str.WriteString(fmt.Sprintf("'pattern_type': '%s'\n", e.PatternType))
	str.WriteString(fmt.Sprintf("'pattern_version': '%s'\n", e.PatternVersion))
	str.WriteString(fmt.Sprintf("'valid_from': '%v'\n", e.ValidFrom))
	str.WriteString(fmt.Sprintf("'valid_until': '%v'\n", e.ValidUntil))
	str.WriteString(fmt.Sprintf("'sectors': \n%v", func(l []stixhelpers.KillChainPhasesTypeElementSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'sector '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.KillChainPhases)))

	return str.String()
}
