package commonpropertiesstixdo

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func (e *CommonPropertiesDomainObjectSTIX) Get() *CommonPropertiesDomainObjectSTIX {
	return e
}

// -------- Revoked property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetRevoked() bool {
	return e.Revoked
}

// SetValueRevoked устанавливает значение для поля Revoked
func (e *CommonPropertiesDomainObjectSTIX) SetValueRevoked(v bool) {
	e.Revoked = v
}

// SetAnyRevoked устанавливает ЛЮБОЕ значение для поля Revoked
func (e *CommonPropertiesDomainObjectSTIX) SetAnyRevoked(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Revoked = v
	}
}

// -------- Defanged property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetDefanged() bool {
	return e.Defanged
}

// SetValueDefanged устанавливает значение для поля Defanged
func (e *CommonPropertiesDomainObjectSTIX) SetValueDefanged(v bool) {
	e.Defanged = v
}

// SetAnyDefanged устанавливает ЛЮБОЕ значение для поля Defanged
func (e *CommonPropertiesDomainObjectSTIX) SetAnyDefanged(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Defanged = v
	}
}

// -------- Сonfidence property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetСonfidence() int {
	return e.Сonfidence
}

// SetValueСonfidence устанавливает значение для поля Сonfidence
func (e *CommonPropertiesDomainObjectSTIX) SetValueConfidence(v int) {
	e.Сonfidence = v
}

// SetAnyСonfidence устанавливает ЛЮБОЕ значение для поля Сonfidence
func (e *CommonPropertiesDomainObjectSTIX) SetAnyСonfidence(i interface{}) {
	e.Сonfidence = commonlibs.ConversionAnyToInt(i)
}

// -------- Lang property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetLang() string {
	return e.Lang
}

// SetValueLang устанавливает значение для поля Lang
func (e *CommonPropertiesDomainObjectSTIX) SetValueLang(v string) {
	e.Lang = v
}

// SetAnyLang устанавливает ЛЮБОЕ значение для поля Lang
func (e *CommonPropertiesDomainObjectSTIX) SetAnyLang(i interface{}) {
	e.Lang = fmt.Sprint(i)
}

// -------- SpecVersion property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetSpecVersion() string {
	return e.SpecVersion
}

// SetValueSpecVersion устанавливает значение для поля SpecVersion
func (e *CommonPropertiesDomainObjectSTIX) SetValueSpecVersion(v string) {
	e.SpecVersion = v
}

// SetAnySpecVersion устанавливает ЛЮБОЕ значение для поля SpecVersion
func (e *CommonPropertiesDomainObjectSTIX) SetAnySpecVersion(i interface{}) {
	e.SpecVersion = fmt.Sprint(i)
}

// -------- Created property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetCreated() string {
	return e.Created
}

// SetValueCreated устанавливает значение в формате RFC3339 для поля Created
func (e *CommonPropertiesDomainObjectSTIX) SetValueCreated(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Created = v

	return nil
}

// SetAnyCreated устанавливает ЛЮБОЕ значение для поля Created
func (e *CommonPropertiesDomainObjectSTIX) SetAnyCreated(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Created = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Modified property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetModified() string {
	return e.Modified
}

// SetValueModified устанавливает значение в формате RFC3339 для поля Modified
func (e *CommonPropertiesDomainObjectSTIX) SetValueModified(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Modified = v

	return nil
}

// SetAnyModified устанавливает ЛЮБОЕ значение для поля Modified
func (e *CommonPropertiesDomainObjectSTIX) SetAnyModified(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Modified = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Labels property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetLabels() []string {
	return e.Labels
}

// SetValueLabels устанавливает значение для поля Labels
func (e *CommonPropertiesDomainObjectSTIX) SetValueLabels(v string) {
	e.Labels = append(e.Labels, v)
}

// SetAnyLabels устанавливает ЛЮБОЕ значение для поля Labels
func (e *CommonPropertiesDomainObjectSTIX) SetAnyLabels(i interface{}) {
	e.Labels = append(e.Labels, fmt.Sprint(i))
}

// -------- Extensions property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetExtensions() map[string]string {
	return e.Extensions
}

// SetValueExtensions устанавливает значение для поля Extensions
func (e *CommonPropertiesDomainObjectSTIX) SetValueExtensions(k, v string) {
	e.Extensions[k] = v
}

// SetAnyExtensions устанавливает ЛЮБОЕ значение для поля Extensions
func (e *CommonPropertiesDomainObjectSTIX) SetAnyExtensions(k string, i interface{}) {
	e.Extensions[k] = fmt.Sprint(i)
}

// -------- CreatedByRef property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetCreatedByRef() stixhelpers.IdentifierTypeSTIX {
	return e.CreatedByRef
}

func (e *CommonPropertiesDomainObjectSTIX) SetValueCreatedByRef(v stixhelpers.IdentifierTypeSTIX) {
	e.CreatedByRef = v
}

// -------- ExternalReferences property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetExternalReferences() []stixhelpers.ExternalReferenceTypeElementSTIX {
	return e.ExternalReferences
}

func (e *CommonPropertiesDomainObjectSTIX) SetValueExternalReferences(v []stixhelpers.ExternalReferenceTypeElementSTIX) {
	e.ExternalReferences = v
}

// -------- ObjectMarkingRefs property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetObjectMarkingRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectMarkingRefs
}

func (e *CommonPropertiesDomainObjectSTIX) SetValueObjectMarkingRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectMarkingRefs = v
}

// -------- GranularMarkings property ---------
func (e *CommonPropertiesDomainObjectSTIX) GetGranularMarkings() []stixhelpers.GranularMarkingsTypeSTIX {
	return e.GranularMarkings
}

func (e *CommonPropertiesDomainObjectSTIX) SetValueGranularMarkings(v []stixhelpers.GranularMarkingsTypeSTIX) {
	e.GranularMarkings = v
}

// ValidateStructCommonFields выполняет проверку полей типа на соответствие корректным значениям
func (e *CommonPropertiesDomainObjectSTIX) ValidateStructCommonFields() bool {
	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(e.SpecVersion)) {
		return false
	}

	//валидация содержимого поля CreatedByRef
	if len(fmt.Sprint(e.CreatedByRef)) > 0 {
		if !(regexp.MustCompile(`^[0-9a-zA-Z-_]+(--)[0-9a-f|-]+$`).MatchString(fmt.Sprint(e.CreatedByRef))) {
			return false
		}
	}

	//для поля Lang
	if len(e.Lang) > 0 {
		if !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(e.Lang) {
			return false
		}
	}

	//вызываем метод проверки полей типа ExternalReferences
	if ok := checkExternalReferencesTypeSTIX(e.ExternalReferences); !ok {
		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(e.ObjectMarkingRefs) > 0 {
		for _, value := range e.ObjectMarkingRefs {
			if !value.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.GranularMarkings) > 0 {
		for _, value := range e.GranularMarkings {
			//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
			if !value.CheckGranularMarkingsTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct выполняет очистку объекта от 'нежелательных' символов
func (e CommonPropertiesDomainObjectSTIX) SanitizeStruct() CommonPropertiesDomainObjectSTIX {
	//обработка содержимого списка поля Labels
	if len(e.Labels) > 0 {
		nl := make([]string, 0, len(e.Labels))

		for _, l := range e.Labels {
			nl = append(nl, commonlibs.StringSanitize(l))
		}

		e.Labels = nl
	}

	//обработка содержимого списка поля ExternalReferences
	e.ExternalReferences = sanitizeStructExternalReferencesTypeSTIX(e.ExternalReferences)

	//обработка содержимого списка поля Extension
	if len(e.Extensions) > 0 {
		newExtension := make(map[string]string, len(e.Extensions))
		for extKey, extValue := range e.Extensions {
			newExtension[extKey] = commonlibs.StringSanitize(extValue)
		}

		e.Extensions = newExtension
	}

	//время модификации объекта
	e.Modified = commonlibs.GetDateTimeFormatRFC3339(time.Now().UnixMilli())

	return e
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e CommonPropertiesDomainObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}

	ws := commonlibs.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'spec_version': '%s'\n", ws, e.SpecVersion))
	str.WriteString(fmt.Sprintf("%s'created': '%v'\n", ws, e.Created))
	str.WriteString(fmt.Sprintf("%s'modified': '%v'\n", ws, e.Modified))
	str.WriteString(fmt.Sprintf("%s'created_by_ref': '%s'\n", ws, e.CreatedByRef))
	str.WriteString(fmt.Sprintf("%s'revoked': '%v'\n", ws, e.Revoked))
	str.WriteString(fmt.Sprintf("%s'labels': \n%v", ws, func(l []string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'label '%d'': '%s'\n", ws, k, v))
		}

		return str.String()
	}(e.Labels, num+1)))
	str.WriteString(fmt.Sprintf("%s'external_references': \n%v", ws, func(l []stixhelpers.ExternalReferenceTypeElementSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)
		dubleWs := commonlibs.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'external_references element '%d'':\n", ws, k))
			str.WriteString(fmt.Sprintf("%s'source_name': '%s'\n", dubleWs, v.SourceName))
			str.WriteString(fmt.Sprintf("%s'description': '%s'\n", dubleWs, v.Description))
			str.WriteString(fmt.Sprintf("%s'url': '%s'\n", dubleWs, v.URL))
			str.WriteString(fmt.Sprintf("%s'hashes': '%s'\n", dubleWs, v.Hashes))
			str.WriteString(fmt.Sprintf("%s'external_id': '%s'\n", dubleWs, v.ExternalID))
		}

		return str.String()
	}(e.ExternalReferences, num+1)))
	str.WriteString(fmt.Sprintf("%s'object_marking_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.ObjectMarkingRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'granular_markings': \n%v", ws, func(l []stixhelpers.GranularMarkingsTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'granular_markings number %d.'", ws, k))
			str.WriteString(fmt.Sprintf("%s'lang': '%s'\n", ws, v.Lang))
			str.WriteString(fmt.Sprintf("%s'marking_ref': '%v'\n", ws, v.MarkingRef))
			str.WriteString(fmt.Sprintf("%s'selectors': \n%v", ws, func(l []string, num int) string {
				str := strings.Builder{}
				ws := commonlibs.GetWhitespace(num)

				for k, v := range l {
					str.WriteString(fmt.Sprintf("%s'selector '%d'': '%s'\n", ws, k, v))
				}

				return str.String()
			}(v.Selectors, num+1)))
		}

		return str.String()
	}(e.GranularMarkings, num+1)))
	str.WriteString(fmt.Sprintf("%s'defanged': '%v'\n", ws, e.Defanged))
	str.WriteString(fmt.Sprintf("%s'extensions': \n%v", ws, func(l map[string]string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s': '%s'\n", ws, k, v))
		}

		return str.String()
	}(e.Extensions, num+1)))

	return str.String()
}

func checkExternalReferencesTypeSTIX(list []stixhelpers.ExternalReferenceTypeElementSTIX) bool {
	if len(list) == 0 {
		return true
	}

	for _, v := range list {
		if !v.CheckExternalReferenceTypeElementSTIX() {
			return false
		}
	}

	return true
}

func sanitizeStructExternalReferencesTypeSTIX(list []stixhelpers.ExternalReferenceTypeElementSTIX) []stixhelpers.ExternalReferenceTypeElementSTIX {
	size := len(list)
	if size == 0 {
		return list
	}

	ert := make([]stixhelpers.ExternalReferenceTypeElementSTIX, 0, size)
	for _, v := range list {
		tmp := v.SanitizeStructExternalReferenceTypeElementSTIX()
		ert = append(ert, tmp)
	}

	return ert
}
