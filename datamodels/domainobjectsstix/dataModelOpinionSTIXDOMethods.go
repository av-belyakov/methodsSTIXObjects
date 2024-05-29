package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- OpinionDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e OpinionDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e OpinionDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Opinion", по терминалогии STIX, содержит оценку информации в приведенной в каком либо другом объекте STIX,
// которую произвел другой участник анализа.
// Обязательные значения в полях Opinion, ObjectRefs
func (e *OpinionDomainObjectsSTIX) Get() (*OpinionDomainObjectsSTIX, error) {
	if e.GetOpinion() == "" {
		err := fmt.Errorf("the required value 'Opinion' must not be empty")

		return &OpinionDomainObjectsSTIX{}, err
	}

	if len(e.GetObjectRefs()) == 0 {
		err := fmt.Errorf("the required value 'ObjectRefs' must not be empty")

		return &OpinionDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Explanation property ---------
func (e *OpinionDomainObjectsSTIX) GetExplanation() string {
	return e.Explanation
}

// SetValueExplanation устанавливает значение для поля Explanation
func (e *OpinionDomainObjectsSTIX) SetValueExplanation(v string) {
	e.Explanation = v
}

// SetAnyExplanation устанавливает ЛЮБОЕ значение для поля Explanation
func (e *OpinionDomainObjectsSTIX) SetAnyExplanation(i interface{}) {
	e.Explanation = fmt.Sprint(i)
}

// -------- Authors property ---------
func (e *OpinionDomainObjectsSTIX) GetAuthors() []string {
	return e.Authors
}

// SetValueAuthors устанавливает значение для поля Authors
func (e *OpinionDomainObjectsSTIX) SetValueAuthors(v string) {
	e.Authors = append(e.Authors, v)
}

// SetAnyAuthors устанавливает ЛЮБОЕ значение для поля Authors
func (e *OpinionDomainObjectsSTIX) SetAnyAuthors(i interface{}) {
	e.Authors = append(e.Authors, fmt.Sprint(i))
}

// -------- Opinion property ---------
func (e *OpinionDomainObjectsSTIX) GetOpinion() stixhelpers.EnumTypeSTIX {
	return e.Opinion
}

// SetValueOpinion устанавливает значение для поля Opinion
func (e *OpinionDomainObjectsSTIX) SetValueOpinion(v stixhelpers.EnumTypeSTIX) {
	e.Opinion = v
}

// SetAnyOpinion устанавливает ЛЮБОЕ значение для поля Opinion
func (e *OpinionDomainObjectsSTIX) SetAnyOpinion(i interface{}) {
	e.Opinion = stixhelpers.EnumTypeSTIX(fmt.Sprint(i))
}

// -------- ObjectRefs property ---------
func (e *OpinionDomainObjectsSTIX) GetObjectRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectRefs
}

func (e *OpinionDomainObjectsSTIX) SetValueObjectRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = append(e.ObjectRefs, v)
}

func (e *OpinionDomainObjectsSTIX) SetFullValueObjectRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе OpinionDomainObjectsSTIX
func (e OpinionDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(opinion--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Opinion == "" {
		return false
	}

	for _, v := range e.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e OpinionDomainObjectsSTIX) SanitizeStruct() OpinionDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Explanation = commonlibs.StringSanitize(e.Explanation)

	if len(e.Authors) > 0 {
		mTmp := make([]string, 0, len(e.Authors))
		for _, v := range e.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}

		e.Authors = mTmp
	}

	e.Opinion = stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(e.Opinion)))

	return e
}

// GetID возвращает ID STIX объекта
func (e OpinionDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e OpinionDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e OpinionDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'explanation': '%v'\n", e.Explanation))
	str.WriteString(fmt.Sprintf("'authors': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'author '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Authors)))
	str.WriteString(fmt.Sprintf("'opinion': '%v'\n", e.Opinion))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e OpinionDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if len(e.Authors) > 0 {
		var strTmp string

		for _, v := range e.Authors {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}
