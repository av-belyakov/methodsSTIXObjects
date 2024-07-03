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

/* --- ArtifactCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e ArtifactCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e ArtifactCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки
// в кодировке base64 или связывать его с полезной нагрузкой, подобной файлу.
func (e *ArtifactCyberObservableObjectSTIX) Get() (*ArtifactCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- MimeType property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetMimeType() string {
	return e.MimeType
}

// SetValueMimeType устанавливает значение для поля MimeType
func (e *ArtifactCyberObservableObjectSTIX) SetValueMimeType(v string) {
	e.MimeType = v
}

// SetAnyMimeType устанавливает ЛЮБОЕ значение для поля MimeType
func (e *ArtifactCyberObservableObjectSTIX) SetAnyMimeType(i interface{}) {
	e.MimeType = fmt.Sprint(i)
}

// -------- PayloadBin property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetPayloadBin() string {
	return e.PayloadBin
}

// SetValuePayloadBin устанавливает значение для поля PayloadBin
func (e *ArtifactCyberObservableObjectSTIX) SetValuePayloadBin(v string) {
	e.PayloadBin = v
}

// SetAnyPayloadBin устанавливает ЛЮБОЕ значение для поля PayloadBin
func (e *ArtifactCyberObservableObjectSTIX) SetAnyPayloadBin(i interface{}) {
	e.PayloadBin = fmt.Sprint(i)
}

// -------- URL property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetURL() string {
	return e.URL
}

// SetValueURL устанавливает значение для поля URL
func (e *ArtifactCyberObservableObjectSTIX) SetValueURL(v string) {
	e.URL = v
}

// SetAnyURL устанавливает ЛЮБОЕ значение для поля URL
func (e *ArtifactCyberObservableObjectSTIX) SetAnyURL(i interface{}) {
	e.URL = fmt.Sprint(i)
}

// -------- DecryptionKey property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetDecryptionKey() string {
	return e.DecryptionKey
}

// SetValueDecryptionKey устанавливает значение для поля DecryptionKey
func (e *ArtifactCyberObservableObjectSTIX) SetValueDecryptionKey(v string) {
	e.DecryptionKey = v
}

// SetAnyDecryptionKey устанавливает ЛЮБОЕ значение для поля DecryptionKey
func (e *ArtifactCyberObservableObjectSTIX) SetAnyDecryptionKey(i interface{}) {
	e.DecryptionKey = fmt.Sprint(i)
}

// -------- Hashes property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetHashes() stixhelpers.HashesTypeSTIX {
	return e.Hashes
}

// SetValueHashes устанавливает значение для поля Hashes
func (e *ArtifactCyberObservableObjectSTIX) SetValueHashes(v stixhelpers.HashesTypeSTIX) {
	e.Hashes = v
}

// -------- EncryptionAlgorithm property ---------
func (e *ArtifactCyberObservableObjectSTIX) GetEncryptionAlgorithm() stixhelpers.EnumTypeSTIX {
	return e.EncryptionAlgorithm
}

// SetValueEncryptionAlgorithm устанавливает значение для поля EncryptionAlgorithm
func (e *ArtifactCyberObservableObjectSTIX) SetValueEncryptionAlgorithm(v stixhelpers.EnumTypeSTIX) {
	e.EncryptionAlgorithm = v
}

// SetAnyEncryptionAlgorithm устанавливает ЛЮБОЕ значение для поля EncryptionAlgorithm
func (e *ArtifactCyberObservableObjectSTIX) SetAnyEncryptionAlgorithm(i interface{}) {
	e.EncryptionAlgorithm = stixhelpers.EnumTypeSTIX(fmt.Sprint(i))
}

// ValidateStruct является валидатором параметров содержащихся в типе ArtifactCyberObservableObjectSTIX
func (e ArtifactCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(artifact--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if e.PayloadBin != "" {
		if !govalidator.IsBase64(e.PayloadBin) {
			return false
		}
	}

	if e.URL != "" {
		if !govalidator.IsURL(e.URL) {
			return false
		}
	}

	if !e.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e ArtifactCyberObservableObjectSTIX) SanitizeStruct() ArtifactCyberObservableObjectSTIX {
	e.MimeType = commonlibs.StringSanitize(e.MimeType)
	e.EncryptionAlgorithm = stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(e.EncryptionAlgorithm)))
	e.DecryptionKey = commonlibs.StringSanitize(e.DecryptionKey)

	return e
}

// GetID возвращает ID STIX объекта
func (e ArtifactCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e ArtifactCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ArtifactCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'mime_type': '%s'\n", ws, e.MimeType))
	str.WriteString(fmt.Sprintf("%s'payload_bin': '%s'\n", ws, e.PayloadBin))
	str.WriteString(fmt.Sprintf("%s'url': '%s'\n", ws, e.URL))
	str.WriteString(fmt.Sprintf("%s'hashes': '%v'\n", ws, e.Hashes))
	str.WriteString(fmt.Sprintf("%s'encryption_algorithm': '%v'\n", ws, e.EncryptionAlgorithm))
	str.WriteString(fmt.Sprintf("%s'decryption_key': '%s'\n", ws, e.DecryptionKey))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e ArtifactCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.URL != "" {
		dataForIndex["value"] = e.URL
	}

	return dataForIndex
}
