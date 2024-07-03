package commonpropertiesstixco

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
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

func (ocpcstix OptionalCommonPropertiesCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(fmt.Sprintf("%s'spec_version': '%s'\n", ws, ocpcstix.SpecVersion))
	str.WriteString(fmt.Sprintf("%s'object_marking_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'object_marking_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(ocpcstix.ObjectMarkingRefs, num+1)))
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
	}(ocpcstix.GranularMarkings, num+1)))
	str.WriteString(fmt.Sprintf("%s'defanged': '%v'\n", ws, ocpcstix.Defanged))

	return str.String()
}
