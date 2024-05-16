package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- DirectoryCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e DirectoryCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e DirectoryCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// -------- Path property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetPath() string {
	return e.Path
}

// SetValuePath устанавливает значение для поля Path
func (e *DirectoryCyberObservableObjectSTIX) SetValuePath(v string) {
	e.Path = v
}

// SetAnyPath устанавливает ЛЮБОЕ значение для поля Path
func (e *DirectoryCyberObservableObjectSTIX) SetAnyPath(i interface{}) {
	e.Path = fmt.Sprint(i)
}

// -------- PathEnc property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetPathEnc() string {
	return e.PathEnc
}

// SetValuePathEnc устанавливает значение для поля PathEnc
func (e *DirectoryCyberObservableObjectSTIX) SetValuePathEnc(v string) {
	e.PathEnc = v
}

// SetAnyPathEnc устанавливает ЛЮБОЕ значение для поля PathEnc
func (e *DirectoryCyberObservableObjectSTIX) SetAnyPathEnc(i interface{}) {
	e.PathEnc = fmt.Sprint(i)
}

// -------- Ctime property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetCtime() string {
	return e.Ctime
}

// SetValueCtime устанавливает значение в формате RFC3339 для поля Ctime
func (e *DirectoryCyberObservableObjectSTIX) SetValueCtime(v string) {
	e.Ctime = v
}

// SetAnyCtime устанавливает ЛЮБОЕ значение для поля Ctime
func (e *DirectoryCyberObservableObjectSTIX) SetAnyCtime(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Ctime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Mtime property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetMtime() string {
	return e.Mtime
}

// SetValueMtime устанавливает значение в формате RFC3339 для поля Mtime
func (e *DirectoryCyberObservableObjectSTIX) SetValueMtime(v string) {
	e.Mtime = v
}

// SetAnyMtime устанавливает ЛЮБОЕ значение для поля Mtime
func (e *DirectoryCyberObservableObjectSTIX) SetAnyMtime(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Mtime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Atime property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetAtime() string {
	return e.Atime
}

// SetValueAtime устанавливает значение в формате RFC3339 для поля Atime
func (e *DirectoryCyberObservableObjectSTIX) SetValueAtime(v string) {
	e.Atime = v
}

// SetAnyAtime устанавливает ЛЮБОЕ значение для поля Atime
func (e *DirectoryCyberObservableObjectSTIX) SetAnyAtime(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.Atime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- ContainsRefs property ---------
func (e *DirectoryCyberObservableObjectSTIX) GetContainsRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ContainsRefs
}

func (e *DirectoryCyberObservableObjectSTIX) SetValueContainsRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ContainsRefs = v
}

// ValidateStruct является валидатором параметров содержащихся в типе DirectoryCyberObservableObjectSTIX
func (e DirectoryCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(directory--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if e.Path == "" {
		return false
	}

	isUnixPath := govalidator.IsUnixFilePath(e.Path)
	isWinPath := govalidator.IsWinFilePath(e.Path)
	if !isUnixPath && !isWinPath {
		return false
	}

	if len(e.ContainsRefs) > 0 {
		for _, v := range e.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e DirectoryCyberObservableObjectSTIX) SanitizeStruct() DirectoryCyberObservableObjectSTIX {
	e.PathEnc = commonlibs.StringSanitize(e.PathEnc)

	return e
}

// GetID возвращает ID STIX объекта
func (e DirectoryCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e DirectoryCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e DirectoryCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'path': '%s'\n", e.Path))
	str.WriteString(fmt.Sprintf("'path_enc': '%s'\n", e.PathEnc))
	str.WriteString(fmt.Sprintf("'ctime': '%v'\n", e.Ctime))
	str.WriteString(fmt.Sprintf("'mtime': '%v'\n", e.Mtime))
	str.WriteString(fmt.Sprintf("'atime': '%s'\n", e.Atime))
	str.WriteString(fmt.Sprintf("'contains_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'contains_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.ContainsRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e DirectoryCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
