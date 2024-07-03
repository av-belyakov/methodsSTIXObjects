package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- NoteDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e NoteDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e NoteDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую текущий контекст анализа
// либо содержащей результаты дополнительного анализа которые не может быть описан в терминах объектов STIX
// Обязательные значения в полях Content, ObjectRefs
func (e *NoteDomainObjectsSTIX) Get() (*NoteDomainObjectsSTIX, error) {
	if e.GetContent() == "" {
		err := fmt.Errorf("the required value 'Content' must not be empty")

		return &NoteDomainObjectsSTIX{}, err
	}

	return e, nil
}

// -------- Abstract property ---------
func (e *NoteDomainObjectsSTIX) GetAbstract() string {
	return e.Abstract
}

// SetValueAbstract устанавливает значение для поля Abstract
func (e *NoteDomainObjectsSTIX) SetValueAbstract(v string) {
	e.Abstract = v
}

// SetAnyAbstract устанавливает ЛЮБОЕ значение для поля Abstract
func (e *NoteDomainObjectsSTIX) SetAnyAbstract(i interface{}) {
	e.Abstract = fmt.Sprint(i)
}

// -------- Content property ---------
func (e *NoteDomainObjectsSTIX) GetContent() string {
	return e.Content
}

// SetValueContent устанавливает значение для поля Content
func (e *NoteDomainObjectsSTIX) SetValueContent(v string) {
	e.Content = v
}

// SetAnyContent устанавливает ЛЮБОЕ значение для поля Content
func (e *NoteDomainObjectsSTIX) SetAnyContent(i interface{}) {
	e.Content = fmt.Sprint(i)
}

// -------- Authors property ---------
func (e *NoteDomainObjectsSTIX) GetAuthors() []string {
	return e.Authors
}

// SetValueAuthors устанавливает значение для поля Authors
func (e *NoteDomainObjectsSTIX) SetValueAuthors(v string) {
	e.Authors = append(e.Authors, v)
}

// SetAnyAuthors устанавливает ЛЮБОЕ значение для поля Authors
func (e *NoteDomainObjectsSTIX) SetAnyAuthors(i interface{}) {
	e.Authors = append(e.Authors, fmt.Sprint(i))
}

// -------- ObjectRefs property ---------
func (e *NoteDomainObjectsSTIX) GetObjectRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ObjectRefs
}

func (e *NoteDomainObjectsSTIX) SetValueObjectRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = append(e.ObjectRefs, v)
}

func (e *NoteDomainObjectsSTIX) SetFullValueObjectRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ObjectRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе NoteDomainObjectsSTIX
func (e NoteDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(note--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Content == "" {
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
func (e NoteDomainObjectsSTIX) SanitizeStruct() NoteDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Abstract = commonlibs.StringSanitize(e.Abstract)
	e.Content = commonlibs.StringSanitize(e.Content)

	if len(e.Authors) > 0 {
		mTmp := make([]string, 0, len(e.Authors))
		for _, v := range e.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}

		e.Authors = mTmp
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e NoteDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e NoteDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e NoteDomainObjectsSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'abstract': '%s'\n", ws, e.Abstract))
	str.WriteString(fmt.Sprintf("%s'content': '%s'\n", ws, e.Content))
	str.WriteString(fmt.Sprintf("%s'authors': \n%v", ws, func(l []string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'author '%d'': '%s'\n", ws, k, v))
		}

		return str.String()
	}(e.Authors, num+1)))
	str.WriteString(fmt.Sprintf("%s'object_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'object_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.ObjectRefs, num+1)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e NoteDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Abstract != "" {
		dataForIndex["abstract"] = e.Abstract
	}

	if e.Content != "" {
		dataForIndex["content"] = e.Content
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
