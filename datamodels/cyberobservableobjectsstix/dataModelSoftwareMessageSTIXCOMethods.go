package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- SoftwareCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e SoftwareCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e SoftwareCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Software Object", по терминалогии STIX, содержит свойства, связанные с
// программным обеспечением, включая программные продукты.
func (e *SoftwareCyberObservableObjectSTIX) Get() (*SoftwareCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- Name property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *SoftwareCyberObservableObjectSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *SoftwareCyberObservableObjectSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- CPE property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetCPE() string {
	return e.CPE
}

// SetValueCPE устанавливает значение для поля CPE
func (e *SoftwareCyberObservableObjectSTIX) SetValueCPE(v string) {
	e.CPE = v
}

// SetAnyCPE устанавливает ЛЮБОЕ значение для поля CPE
func (e *SoftwareCyberObservableObjectSTIX) SetAnyCPE(i interface{}) {
	e.CPE = fmt.Sprint(i)
}

// -------- SwID property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetSwID() string {
	return e.SwID
}

// SetValueSwID устанавливает значение для поля SwID
func (e *SoftwareCyberObservableObjectSTIX) SetValueSwID(v string) {
	e.SwID = v
}

// SetAnySwID устанавливает ЛЮБОЕ значение для поля SwID
func (e *SoftwareCyberObservableObjectSTIX) SetAnySwID(i interface{}) {
	e.SwID = fmt.Sprint(i)
}

// -------- Vendor property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetVendor() string {
	return e.Vendor
}

// SetValueVendor устанавливает значение для поля Vendor
func (e *SoftwareCyberObservableObjectSTIX) SetValueVendor(v string) {
	e.Vendor = v
}

// SetAnyVendor устанавливает ЛЮБОЕ значение для поля Vendor
func (e *SoftwareCyberObservableObjectSTIX) SetAnyVendor(i interface{}) {
	e.Vendor = fmt.Sprint(i)
}

// -------- Version property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetVersion() string {
	return e.Version
}

// SetValueVersion устанавливает значение для поля Version
func (e *SoftwareCyberObservableObjectSTIX) SetValueVersion(v string) {
	e.Version = v
}

// SetAnyVersion устанавливает ЛЮБОЕ значение для поля Version
func (e *SoftwareCyberObservableObjectSTIX) SetAnyVersion(i interface{}) {
	e.Version = fmt.Sprint(i)
}

// -------- Languages property ---------
func (e *SoftwareCyberObservableObjectSTIX) GetLanguages() []string {
	return e.Languages
}

// SetValueLanguages устанавливает значение для поля Languages
func (e *SoftwareCyberObservableObjectSTIX) SetValueLanguages(v string) {
	e.Languages = append(e.Languages, v)
}

// SetAnyLanguages устанавливает ЛЮБОЕ значение для поля Languages
func (e *SoftwareCyberObservableObjectSTIX) SetAnyProtocols(i interface{}) {
	e.Languages = append(e.Languages, fmt.Sprint(i))
}

// ValidateStruct является валидатором параметров содержащихся в типе SoftwareCyberObservableObjectSTIX
func (e SoftwareCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(software--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if e.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e SoftwareCyberObservableObjectSTIX) SanitizeStruct() SoftwareCyberObservableObjectSTIX {
	e.Name = commonlibs.StringSanitize(e.Name)
	e.CPE = commonlibs.StringSanitize(e.CPE)
	e.SwID = commonlibs.StringSanitize(e.SwID)
	sizel := len(e.Languages)

	if sizel > 0 {
		tmp := make([]string, 0, sizel)
		for _, v := range e.Languages {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}

		e.Languages = tmp
	}

	e.Vendor = commonlibs.StringSanitize(e.Vendor)
	e.Version = commonlibs.StringSanitize(e.Version)

	return e
}

// GetID возвращает ID STIX объекта
func (e SoftwareCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e SoftwareCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e SoftwareCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'cpe': '%s'\n", e.CPE))
	str.WriteString(fmt.Sprintf("'swid': '%s'\n", e.SwID))
	str.WriteString(fmt.Sprintf("'languages': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'language '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(e.Languages)))
	str.WriteString(fmt.Sprintf("'vendor': '%s'\n", e.Vendor))
	str.WriteString(fmt.Sprintf("'version': '%s'\n", e.Version))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e SoftwareCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Name != "" {
		dataForIndex["name"] = e.Name
	}

	return dataForIndex
}
