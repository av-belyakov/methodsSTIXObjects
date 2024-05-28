package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestArtifactCyberObservableObjectSTIX(t *testing.T) {
	na := methodstixobjects.NewArtifactCyberObservableObjectSTIX()

	assert.Equal(t, na.GetType(), "artifact")
	_, err := na.Get()
	assert.NoError(t, err)

	na.SetAnyMimeType("text/json")
	assert.Equal(t, na.GetMimeType(), "text/json")
	na.SetValueMimeType("t/j")
	assert.Equal(t, na.GetMimeType(), "t/j")

	pb := "6wdt7w8ft89r9373f="
	na.SetAnyPayloadBin(pb)
	assert.Equal(t, na.GetPayloadBin(), pb)
	na.SetValuePayloadBin("yyd2")
	assert.Equal(t, na.GetPayloadBin(), "yyd2")

	url := "http://example.net/"
	na.SetAnyURL(url)
	assert.Equal(t, na.GetURL(), url)
	na.SetValueURL("http://")
	assert.Equal(t, na.GetURL(), "http://")

	dk := "Decryption Key Example"
	na.SetAnyDecryptionKey(dk)
	assert.Equal(t, na.GetDecryptionKey(), dk)
	na.SetValueDecryptionKey("dke")
	assert.Equal(t, na.GetDecryptionKey(), "dke")

	na.SetValueHashes(stixhelpers.HashesTypeSTIX{"1": "dd", "2": "ew"})
	assert.Equal(t, len(na.GetHashes()), 2)

	na.SetAnyEncryptionAlgorithm("Encryption Algorithm")
	assert.Equal(t, na.GetEncryptionAlgorithm(), stixhelpers.EnumTypeSTIX("Encryption Algorithm"))
}

/*
// ArtifactCyberObservableObjectSTIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки
// в кодировке base64 или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
// MimeType - по возможности это значение ДОЛЖНО быть одним из значений, определенных в реестре типов носителей IANA. В универсальном
// каталоге всех существующих типов файлов.
// PayloadBin - бинарные данные в base64
// URL - унифицированный указатель ресурса (URL)
// Hashes - словарь хешей для URL или PayloadBin
// EncryptionAlgorithm - тип алгоритма шифрования для бинарных данных
// DecryptionKey - определяет ключ для дешифрования зашифрованных данных
type ArtifactCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	MimeType            string                     `json:"mime_type" bson:"mime_type"`
	PayloadBin          string                     `json:"payload_bin" bson:"payload_bin"`
	URL                 string                     `json:"url" bson:"url"`
	DecryptionKey       string                     `json:"decryption_key" bson:"decryption_key"`
	Hashes              stixhelpers.HashesTypeSTIX `json:"hashes" bson:"hashes"`
	EncryptionAlgorithm stixhelpers.EnumTypeSTIX   `json:"encryption_algorithm" bson:"encryption_algorithm"`
}
*/
