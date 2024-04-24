package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/*****************************************************************************/
/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/
/*****************************************************************************/

func (oc *OptionalCommonPropertiesRelationshipObjectSTIX) validateStructCommonFields() bool {
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(oc.SpecVersion)) {
		return false
	}

	if oc.Created.Unix() < 0 {
		return false
	}

	if oc.Modified.Unix() < 0 {
		return false
	}

	return true
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (oc *OptionalCommonPropertiesRelationshipObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("'spec_version': '%s'\n", oc.SpecVersion))
	str.WriteString(fmt.Sprintf("'created': '%v'\n", oc.Created))
	str.WriteString(fmt.Sprintf("'modified': '%v'\n", oc.Modified))

	return str.String()
}

/* --- RelationshipObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (rstix RelationshipObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &rstix); err != nil {
		return nil, err
	}

	return rstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (rstix RelationshipObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(rstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе RelationshipObjectSTIX
func (rstix RelationshipObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(relationship--)[0-9a-f|-]+$`).MatchString(rstix.ID)) {
		return false
	}

	if !rstix.validateStructCommonFields() {
		return false
	}

	if !(regexp.MustCompile(`^[0-9a-z|-]+$`).MatchString(rstix.RelationshipType)) {
		return false
	}

	if !rstix.SourceRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !rstix.TargetRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return rstix.validateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (rstix RelationshipObjectSTIX) SanitizeStruct() RelationshipObjectSTIX {
	rstix.Description = commonlibs.StringSanitize(rstix.Description)

	return rstix
}

// GetID возвращает ID STIX объекта
func (rstix RelationshipObjectSTIX) GetID() string {
	return rstix.ID
}

// GetType возвращает Type STIX объекта
func (rstix RelationshipObjectSTIX) GetType() string {
	return rstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (rstix RelationshipObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(rstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(rstix.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'relationship_type': '%s'\n", rstix.RelationshipType))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", rstix.Description))
	str.WriteString(fmt.Sprintf("'source_ref': '%v'\n", rstix.SourceRef))
	str.WriteString(fmt.Sprintf("'target_ref': '%v'\n", rstix.TargetRef))
	str.WriteString(fmt.Sprintf("'start_time': '%v'\n", rstix.StartTime))
	str.WriteString(fmt.Sprintf("'stop_time': '%v'\n", rstix.StopTime))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (rstix RelationshipObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   rstix.ID,
		"type": rstix.Type,
	}
}

/* --- SightingObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (sstix SightingObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &sstix); err != nil {
		return nil, err
	}

	return sstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (sstix SightingObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(sstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе SightingObjectSTIX
func (sstix SightingObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(sighting--)[0-9a-f|-]+$`).MatchString(sstix.ID)) {
		return false
	}

	if !sstix.validateStructCommonFields() {
		return false
	}

	if !sstix.SightingOfRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(sstix.ObservedDataRefs) > 0 {
		for _, v := range sstix.ObservedDataRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(sstix.WhereSightedRefs) > 0 {
		for _, v := range sstix.WhereSightedRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return sstix.validateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (sstix SightingObjectSTIX) SanitizeStruct() SightingObjectSTIX {
	sstix.Description = commonlibs.StringSanitize(sstix.Description)

	return sstix
}

// GetID возвращает ID STIX объекта
func (sstix SightingObjectSTIX) GetID() string {
	return sstix.ID
}

// GetType возвращает Type STIX объекта
func (sstix SightingObjectSTIX) GetType() string {
	return sstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (sstix SightingObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(sstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(sstix.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'description': '%s'\n", sstix.Description))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", sstix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", sstix.LastSeen))
	str.WriteString(fmt.Sprintf("'count': '%d'\n", sstix.Count))
	str.WriteString(fmt.Sprintf("'sighting_of_ref': '%v'\n", sstix.SightingOfRef))
	str.WriteString(fmt.Sprintf("'observed_data_refs': \n%v", func(l []*IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'observed_data_ref '%d'': '%v'\n", k, *v))
		}

		return str.String()
	}(sstix.ObservedDataRefs)))
	str.WriteString(fmt.Sprintf("'where_sighted_refs': \n%v", func(l []*IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'where_sighted_ref '%d'': '%v'\n", k, *v))
		}

		return str.String()
	}(sstix.WhereSightedRefs)))
	str.WriteString(fmt.Sprintf("'summary': '%v'\n", sstix.Summary))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (sstix SightingObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   sstix.ID,
		"type": sstix.Type,
	}
}
