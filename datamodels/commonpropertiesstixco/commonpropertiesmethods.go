package commonpropertiesstixco

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) Get() *OptionalCommonPropertiesCyberObservableObjectSTIX {
	return e
}

// -------- Defanged property ---------
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) GetDefanged() bool {
	return e.Defanged
}

// SetValueDefanged устанавливает значение для поля Defanged
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetValueDefanged(v bool) {
	e.Defanged = v
}

// SetAnyDefanged устанавливает ЛЮБОЕ значение для поля Defanged
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetAnyDefanged(i interface{}) {
	if v, ok := i.(bool); ok {
		e.Defanged = v
	}
}

// -------- SpecVersion property ---------
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) GetSpecVersion() string {
	return e.SpecVersion
}

// SetValueSpecVersion устанавливает значение для поля SpecVersion
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetValueSpecVersion(v string) {
	e.SpecVersion = v
}

// SetAnySpecVersion устанавливает ЛЮБОЕ значение для поля SpecVersion
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetAnySpecVersion(i interface{}) {
	e.SpecVersion = fmt.Sprint(i)
}

// -------- ObjectMarkingRefs property ---------
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) GetObjectMarkingRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectMarkingRefs
}

func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetValueObjectMarkingRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectMarkingRefs = v
}

// -------- GranularMarkings property ---------
func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) GetGranularMarkings() []stixhelpers.GranularMarkingsTypeSTIX {
	return e.GranularMarkings
}

func (e *OptionalCommonPropertiesCyberObservableObjectSTIX) SetValueGranularMarkings(v []stixhelpers.GranularMarkingsTypeSTIX) {
	e.GranularMarkings = v
}

func (ocpcstix *OptionalCommonPropertiesCyberObservableObjectSTIX) ValidateStructCommonFields() bool {
	//валидация содержимого поля SpecVersion
	if !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(ocpcstix.SpecVersion)) {
		return false
	}

	//проверяем поле ObjectMarkingRefs
	if len(ocpcstix.ObjectMarkingRefs) > 0 {
		for _, value := range ocpcstix.ObjectMarkingRefs {
			if !value.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(ocpcstix.GranularMarkings) > 0 {
		for _, value := range ocpcstix.GranularMarkings {
			//вызываем метод проверки полей типа GranularMarkingsTypeSTIX
			if !value.CheckGranularMarkingsTypeSTIX() {
				return false
			}
		}
	}

	return true
}

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("'spec_version': '%s'\n", ocpcstix.SpecVersion))
	str.WriteString(fmt.Sprintf("'object_marking_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_marking_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(ocpcstix.ObjectMarkingRefs)))
	str.WriteString(fmt.Sprintf("'granular_markings': \n%v", func(l []stixhelpers.GranularMarkingsTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'granular_markings number %d.'", k))
			str.WriteString(fmt.Sprintf("\t'lang': '%s'\n", v.Lang))
			str.WriteString(fmt.Sprintf("\t'marking_ref': '%v'\n", v.MarkingRef))
			str.WriteString(fmt.Sprintf("\t'selectors': \n%v", func(l []string) string {
				str := strings.Builder{}

				for k, v := range l {
					str.WriteString(fmt.Sprintf("\t\t'selector '%d'': '%s'\n", k, v))
				}

				return str.String()
			}(v.Selectors)))
		}

		return str.String()
	}(ocpcstix.GranularMarkings)))
	str.WriteString(fmt.Sprintf("'defanged': '%v'\n", ocpcstix.Defanged))

	return str.String()
}
