package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- FileCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e FileCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonFileCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	e = FileCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Hashes:             commonObject.Hashes,
		Size:               commonObject.Size,
		Name:               commonObject.Name,
		NameEnc:            commonObject.NameEnc,
		MagicNumberHex:     commonObject.MagicNumberHex,
		MimeType:           commonObject.MimeType,
		Ctime:              commonObject.Ctime,
		Mtime:              commonObject.Mtime,
		Atime:              commonObject.Atime,
		ParentDirectoryRef: commonObject.ParentDirectoryRef,
		ContainsRefs:       commonObject.ContainsRefs,
		ContentRef:         commonObject.ContentRef,
	}

	if len(commonObject.Extensions) == 0 {
		return e, nil
	}

	ext := map[string]interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := datamodels.DecodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = e
	}

	e.Extensions = ext

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e FileCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
func (e *FileCyberObservableObjectSTIX) Get() (*FileCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- Size property ---------
func (e *FileCyberObservableObjectSTIX) GetSize() uint64 {
	return e.Size
}

// SetValueSize устанавливает значение для поля Size
func (e *FileCyberObservableObjectSTIX) SetValueSize(v uint64) {
	e.Size = v
}

// SetAnySize устанавливает ЛЮБОЕ значение для поля Size
func (e *FileCyberObservableObjectSTIX) SetAnySize(i interface{}) {
	e.Size = uint64(commonlibs.ConversionAnyToInt(i))
}

// -------- Name property ---------
func (e *FileCyberObservableObjectSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *FileCyberObservableObjectSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *FileCyberObservableObjectSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- NameEnc property ---------
func (e *FileCyberObservableObjectSTIX) GetNameEnc() string {
	return e.NameEnc
}

// SetValueNameEnc устанавливает значение для поля NameEnc
func (e *FileCyberObservableObjectSTIX) SetValueNameEnc(v string) {
	e.NameEnc = v
}

// SetAnyNameEnc устанавливает ЛЮБОЕ значение для поля NameEnc
func (e *FileCyberObservableObjectSTIX) SetAnyNameEnc(i interface{}) {
	e.NameEnc = fmt.Sprint(i)
}

// -------- MagicNumberHex property ---------
func (e *FileCyberObservableObjectSTIX) GetMagicNumberHex() string {
	return e.MagicNumberHex
}

// SetValueMagicNumberHex устанавливает значение для поля MagicNumberHex
func (e *FileCyberObservableObjectSTIX) SetValueMagicNumberHex(v string) {
	e.MagicNumberHex = v
}

// SetAnyMagicNumberHex устанавливает ЛЮБОЕ значение для поля MagicNumberHex
func (e *FileCyberObservableObjectSTIX) SetAnyMagicNumberHex(i interface{}) {
	e.MagicNumberHex = fmt.Sprint(i)
}

// -------- MimeType property ---------
func (e *FileCyberObservableObjectSTIX) GetMimeType() string {
	return e.MimeType
}

// SetValueMimeType устанавливает значение для поля MimeType
func (e *FileCyberObservableObjectSTIX) SetValueMimeType(v string) {
	e.MimeType = v
}

// SetAnyMimeType устанавливает ЛЮБОЕ значение для поля MimeType
func (e *FileCyberObservableObjectSTIX) SetAnyMimeType(i interface{}) {
	e.MimeType = fmt.Sprint(i)
}

// -------- Ctime property ---------
func (e *FileCyberObservableObjectSTIX) GetCtime() string {
	return e.Ctime
}

// SetValueCtime устанавливает значение в формате RFC3339 для поля Ctime
func (e *FileCyberObservableObjectSTIX) SetValueCtime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Ctime = v

	return nil
}

// SetAnyCtime устанавливает ЛЮБОЕ значение для поля Ctime
func (e *FileCyberObservableObjectSTIX) SetAnyCtime(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueCtime(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Ctime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- Mtime property ---------
func (e *FileCyberObservableObjectSTIX) GetMtime() string {
	return e.Mtime
}

// SetValueMtime устанавливает значение в формате RFC3339 для поля Mtime
func (e *FileCyberObservableObjectSTIX) SetValueMtime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Mtime = v

	return nil
}

// SetAnyMtime устанавливает ЛЮБОЕ значение для поля Mtime
func (e *FileCyberObservableObjectSTIX) SetAnyMtime(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueMtime(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Mtime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- Atime property ---------
func (e *FileCyberObservableObjectSTIX) GetAtime() string {
	return e.Atime
}

// SetValueAtime устанавливает значение в формате RFC3339 для поля Atime
func (e *FileCyberObservableObjectSTIX) SetValueAtime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Atime = v

	return nil
}

// SetAnyAtime устанавливает ЛЮБОЕ значение для поля Atime
func (e *FileCyberObservableObjectSTIX) SetAnyAtime(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueAtime(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Atime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- Hashes property ---------
func (e *FileCyberObservableObjectSTIX) GetHashes() stixhelpers.HashesTypeSTIX {
	return e.Hashes
}

func (e *FileCyberObservableObjectSTIX) SetValueHashes(v stixhelpers.HashesTypeSTIX) {
	e.Hashes = v
}

// -------- ParentDirectoryRef property ---------
func (e *FileCyberObservableObjectSTIX) GetParentDirectoryRef() stixhelpers.IdentifierTypeSTIX {
	return e.ParentDirectoryRef
}

func (e *FileCyberObservableObjectSTIX) SetValueParentDirectoryRef(v stixhelpers.IdentifierTypeSTIX) {
	e.ParentDirectoryRef = v
}

// -------- ContentRef property ---------
func (e *FileCyberObservableObjectSTIX) GetContentRef() stixhelpers.IdentifierTypeSTIX {
	return e.ContentRef
}

func (e *FileCyberObservableObjectSTIX) SetValueContentRef(v stixhelpers.IdentifierTypeSTIX) {
	e.ContentRef = v
}

// -------- ContainsRefs property ---------
func (e *FileCyberObservableObjectSTIX) GetContainsRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ContainsRefs
}

func (e *FileCyberObservableObjectSTIX) SetValueContainsRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ContainsRefs = v
}

// -------- Extensions property ---------
func (e *FileCyberObservableObjectSTIX) GetExtensions() map[string]interface{} {
	return e.Extensions
}

func (e *FileCyberObservableObjectSTIX) SetValueExtensions(v map[string]interface{}) {
	e.Extensions = v
}

// ValidateStruct является валидатором параметров содержащихся в типе FileCyberObservableObjectSTIX
func (fstix FileCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(file--)[0-9a-f|-]+$`).MatchString(fstix.ID)) {
		return false
	}

	if !fstix.ValidateStructCommonFields() {
		return false
	}

	if !fstix.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	if !fstix.ParentDirectoryRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(fstix.ContainsRefs) > 0 {
		for _, v := range fstix.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !fstix.ContentRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(fstix.Extensions) > 0 {
		for _, v := range fstix.Extensions {

			if !datamodels.CheckingExtensionsSTIX(v) {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (fstix FileCyberObservableObjectSTIX) SanitizeStruct() FileCyberObservableObjectSTIX {
	fstix.Name = commonlibs.StringSanitize(fstix.Name)
	fstix.NameEnc = commonlibs.StringSanitize(fstix.NameEnc)
	fstix.MagicNumberHex = commonlibs.StringSanitize(fstix.MagicNumberHex)
	fstix.MimeType = commonlibs.StringSanitize(fstix.MimeType)

	esize := len(fstix.Extensions)
	if esize == 0 {
		return fstix
	}

	tmp := make(map[string]interface{}, esize)
	for k, v := range fstix.Extensions {
		result := datamodels.SanitizeExtensionsSTIX(v)
		tmp[k] = result
	}
	fstix.Extensions = tmp

	return fstix
}

// GetID возвращает ID STIX объекта
func (fstix FileCyberObservableObjectSTIX) GetID() string {
	return fstix.ID
}

// GetType возвращает Type STIX объекта
func (fstix FileCyberObservableObjectSTIX) GetType() string {
	return fstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (fstix FileCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(fstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(fstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'hashes': '%v'\n", fstix.Hashes))
	str.WriteString(fmt.Sprintf("'size': '%d'\n", fstix.Size))
	str.WriteString(fmt.Sprintf("'name': '%s'\n", fstix.Name))
	str.WriteString(fmt.Sprintf("'name_enc': '%s'\n", fstix.NameEnc))
	str.WriteString(fmt.Sprintf("'magic_number_hex': '%s'\n", fstix.MagicNumberHex))
	str.WriteString(fmt.Sprintf("'mime_type': '%s'\n", fstix.MimeType))
	str.WriteString(fmt.Sprintf("'ctime': '%v'\n", fstix.Ctime))
	str.WriteString(fmt.Sprintf("'mtime': '%v'\n", fstix.Mtime))
	str.WriteString(fmt.Sprintf("'atime': '%v'\n", fstix.Atime))
	str.WriteString(fmt.Sprintf("'parent_directory_ref': '%v'\n", fstix.ParentDirectoryRef))
	str.WriteString(fmt.Sprintf("'contains_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'contains_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(fstix.ContainsRefs)))
	str.WriteString(fmt.Sprintf("'content_ref': '%v'\n", fstix.ContentRef))
	str.WriteString(fmt.Sprintln("'extensions':"))

	for k, v := range fstix.Extensions {
		str.WriteString(fmt.Sprintf("\t'%s':\n%v\n", k, datamodels.ToStringBeautiful(v)))
	}

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (fstix FileCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   fstix.ID,
		"type": fstix.Type,
	}

	if fstix.Name != "" {
		dataForIndex["name"] = fstix.Name
	}

	return dataForIndex
}
