package relationshipobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/*****************************************************************************/
/********** 			Relationship Objects STIX (МЕТОДЫ)			**********/
/*****************************************************************************/

func (e *OptionalCommonPropertiesRelationshipObjectSTIX) Get() (*OptionalCommonPropertiesRelationshipObjectSTIX, error) {
	if e.GetCreated() == "1970-01-01T00:00:00+00:00" {
		err := fmt.Errorf("the required value 'Created' must not be empty")

		return &OptionalCommonPropertiesRelationshipObjectSTIX{}, err
	}

	if e.GetModified() == "1970-01-01T00:00:00+00:00" {
		err := fmt.Errorf("the required value 'Modified' must not be empty")

		return &OptionalCommonPropertiesRelationshipObjectSTIX{}, err
	}

	return e, nil
}

// -------- SpecVersion property ---------
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) GetSpecVersion() string {
	return e.SpecVersion
}

// SetValueSpecVersion устанавливает значение для поля SpecVersion
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetValueSpecVersion(v string) {
	e.SpecVersion = v
}

// SetAnySpecVersion устанавливает ЛЮБОЕ значение для поля SpecVersion
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetAnySpecVersion(i interface{}) {
	e.SpecVersion = fmt.Sprint(i)
}

// -------- Created property ---------
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) GetCreated() string {
	return e.Created
}

// SetValueCreated устанавливает значение в формате RFC3339 для поля Created
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetValueCreated(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Created = v

	return nil
}

// SetAnyCreated устанавливает ЛЮБОЕ значение для поля Created
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetAnyCreated(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Created = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Modified property ---------
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) GetModified() string {
	return e.Modified
}

// SetValueModified устанавливает значение в формате RFC3339 для поля Modified
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetValueModified(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Modified = v

	return nil
}

// SetAnyModified устанавливает ЛЮБОЕ значение для поля Modified
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) SetAnyModified(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Modified = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

func (e *OptionalCommonPropertiesRelationshipObjectSTIX) ValidateStructCommonFields() bool {
	return !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(e.SpecVersion))
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e *OptionalCommonPropertiesRelationshipObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("'spec_version': '%s'\n", e.SpecVersion))
	str.WriteString(fmt.Sprintf("'created': '%v'\n", e.Created))
	str.WriteString(fmt.Sprintf("'modified': '%v'\n", e.Modified))

	return str.String()
}

/* --- RelationshipObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e RelationshipObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e RelationshipObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект RelationshipObjectSTIX. Объект должен содержать обязательное значения:
// RelationshipType - содержит наименование, используемое для идентификации объекта.
// SourceRef - устанавливает идентификатор исходного (исходящего) объекта.
// TargetRef - определяет идентификатор целевого (to) объекта.
func (e *RelationshipObjectSTIX) Get() (*RelationshipObjectSTIX, error) {
	if e.GetRelationshipType() == "" {
		err := fmt.Errorf("the required value 'RelationshipType' must not be empty")

		return &RelationshipObjectSTIX{}, err
	}

	if e.GetSourceRef() == "" {
		err := fmt.Errorf("the required value 'SourceRef' must not be empty")

		return &RelationshipObjectSTIX{}, err
	}

	if e.GetTargetRef() == "" {
		err := fmt.Errorf("the required value 'TargetRef' must not be empty")

		return &RelationshipObjectSTIX{}, err
	}

	return e, nil
}

// -------- RelationshipType property ---------
func (e *RelationshipObjectSTIX) GetRelationshipType() string {
	return e.RelationshipType
}

// SetValueRelationshipType устанавливает значение для поля RelationshipType
func (e *RelationshipObjectSTIX) SetValueRelationshipType(v string) {
	e.RelationshipType = v
}

// SetAnyRelationshipType устанавливает ЛЮБОЕ значение для поля RelationshipType
func (e *RelationshipObjectSTIX) SetAnyRelationshipType(i interface{}) {
	e.RelationshipType = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *RelationshipObjectSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *RelationshipObjectSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *RelationshipObjectSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- StartTime property ---------
func (e *RelationshipObjectSTIX) GetStartTime() string {
	return e.StartTime
}

// SetValueStartTime устанавливает значение в формате RFC3339 для поля StartTime
func (e *RelationshipObjectSTIX) SetValueStartTime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.StartTime = v

	return nil
}

// SetAnyStartTime устанавливает ЛЮБОЕ значение для поля StartTime
func (e *RelationshipObjectSTIX) SetAnyStartTime(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.StartTime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- StopTime property ---------
func (e *RelationshipObjectSTIX) GetStopTime() string {
	return e.StopTime
}

// SetValueStopTime устанавливает значение в формате RFC3339 для поля StopTime
func (e *RelationshipObjectSTIX) SetValueStopTime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.StopTime = v

	return nil
}

// SetAnyStopTime устанавливает ЛЮБОЕ значение для поля StopTime
func (e *RelationshipObjectSTIX) SetAnyStopTime(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.StopTime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- SourceRef property ---------
func (e *RelationshipObjectSTIX) GetSourceRef() stixhelpers.IdentifierTypeSTIX {
	return e.SourceRef
}

// SetValueSourceRef устанавливает ссылку для поля SourceRef
func (e *RelationshipObjectSTIX) SetValueSourceRef(v stixhelpers.IdentifierTypeSTIX) {
	e.SourceRef = v
}

// -------- TargetRef property ---------
func (e *RelationshipObjectSTIX) GetTargetRef() stixhelpers.IdentifierTypeSTIX {
	return e.TargetRef
}

// SetValueTargetRef устанавливает ссылку для поля TargetRef
func (e *RelationshipObjectSTIX) SetValueTargetRef(v stixhelpers.IdentifierTypeSTIX) {
	e.TargetRef = v
}

// ValidateStruct является валидатором параметров содержащихся в типе RelationshipObjectSTIX
func (e RelationshipObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(relationship--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !(regexp.MustCompile(`^[0-9a-z|-]+$`).MatchString(e.RelationshipType)) {
		return false
	}

	if !e.SourceRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.TargetRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e RelationshipObjectSTIX) SanitizeStruct() RelationshipObjectSTIX {
	e.Description = commonlibs.StringSanitize(e.Description)

	return e
}

// GetID возвращает ID STIX объекта
func (e RelationshipObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e RelationshipObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e RelationshipObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'relationship_type': '%s'\n", e.RelationshipType))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'source_ref': '%v'\n", e.SourceRef))
	str.WriteString(fmt.Sprintf("'target_ref': '%v'\n", e.TargetRef))
	str.WriteString(fmt.Sprintf("'start_time': '%v'\n", e.StartTime))
	str.WriteString(fmt.Sprintf("'stop_time': '%v'\n", e.StopTime))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e RelationshipObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}

/* --- SightingObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e SightingObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e SightingObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект SightingObjectSTIX. Объект должен содержать обязательное значение
// SightingOfRef (ссылку на объект Domain Object STIX).
func (e *SightingObjectSTIX) Get() (*SightingObjectSTIX, error) {
	if e.GetSightingOfRef() == "" {
		err := fmt.Errorf("the required value 'SightingOfRef' must not be empty")

		return &SightingObjectSTIX{}, err
	}

	return e, nil
}

// -------- Summary property ---------
func (a *SightingObjectSTIX) GetSummary() bool {
	return a.Summary
}

// SetValueSummary устанавливает BOOL значение для поля Summary
func (a *SightingObjectSTIX) SetValueSummary(v bool) {
	a.Summary = v
}

// SetAnySummary устанавливает ЛЮБОЕ значение для поля Summary
func (a *SightingObjectSTIX) SetAnySummary(i interface{}) {
	if v, ok := i.(bool); ok {
		a.Summary = v
	}
}

// -------- Count property ---------
func (e *SightingObjectSTIX) GetCount() int {
	return e.Count
}

// SetValueCount устанавливает значение для поля Count
func (e *SightingObjectSTIX) SetValueCount(v int) {
	e.Count = v
}

// SetAnyCount устанавливает ЛЮБОЕ значение для поля Count
func (e *SightingObjectSTIX) SetAnyCount(i interface{}) {
	e.Count = commonlibs.ConversionAnyToInt(i)
}

// -------- Description property ---------
func (e *SightingObjectSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *SightingObjectSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *SightingObjectSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- FirstSeen property ---------
func (e *SightingObjectSTIX) GetFirstSeen() string {
	return e.FirstSeen
}

// SetValueFirstSeen устанавливает значение в формате RFC3339 для поля FirstSeen
func (e *SightingObjectSTIX) SetValueFirstSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.FirstSeen = v

	return nil
}

// SetAnyFirstSeen устанавливает ЛЮБОЕ значение для поля FirstSeen
func (e *SightingObjectSTIX) SetAnyFirstSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.FirstSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- LastSeen property ---------
func (e *SightingObjectSTIX) GetLastSeen() string {
	return e.LastSeen
}

// SetValueLastSeen устанавливает значение в формате RFC3339 для поля LastSeen
func (e *SightingObjectSTIX) SetValueLastSeen(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.LastSeen = v

	return nil
}

// SetAnyLastSeen устанавливает ЛЮБОЕ значение для поля LastSeen
func (e *SightingObjectSTIX) SetAnyLastSeen(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.LastSeen = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- SightingOfRef property ---------
func (e *SightingObjectSTIX) GetSightingOfRef() stixhelpers.IdentifierTypeSTIX {
	return e.SightingOfRef
}

func (e *SightingObjectSTIX) SetValueSightingOfRef(v stixhelpers.IdentifierTypeSTIX) {
	e.SightingOfRef = v
}

// -------- ObservedDataRefs property ---------
func (e *SightingObjectSTIX) GetObservedDataRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObservedDataRefs
}

func (e *SightingObjectSTIX) SetValueObservedDataRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObservedDataRefs = v
}

// -------- WhereSightedRefs property ---------
func (e *SightingObjectSTIX) GetWhereSightedRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.WhereSightedRefs
}

func (e *SightingObjectSTIX) SetValueWhereSightedRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.WhereSightedRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе SightingObjectSTIX
func (e SightingObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(sighting--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !e.SightingOfRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(e.ObservedDataRefs) > 0 {
		for _, v := range e.ObservedDataRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.WhereSightedRefs) > 0 {
		for _, v := range e.WhereSightedRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e SightingObjectSTIX) SanitizeStruct() SightingObjectSTIX {
	e.Description = commonlibs.StringSanitize(e.Description)

	return e
}

// GetID возвращает ID STIX объекта
func (e SightingObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e SightingObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e SightingObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesRelationshipObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", e.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", e.LastSeen))
	str.WriteString(fmt.Sprintf("'count': '%d'\n", e.Count))
	str.WriteString(fmt.Sprintf("'sighting_of_ref': '%v'\n", e.SightingOfRef))
	str.WriteString(fmt.Sprintf("'observed_data_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'observed_data_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ObservedDataRefs)))
	str.WriteString(fmt.Sprintf("'where_sighted_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'where_sighted_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.WhereSightedRefs)))
	str.WriteString(fmt.Sprintf("'summary': '%v'\n", e.Summary))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e SightingObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
