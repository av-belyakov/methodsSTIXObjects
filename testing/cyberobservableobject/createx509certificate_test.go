package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestX509CertificateCyberObservableObjectSTIX(t *testing.T) {
	nxc := methodstixobjects.NewX509CertificateCyberObservableObjectSTIX()

	assert.Equal(t, nxc.GetType(), "x509-certificate")
	_, err := nxc.Get()
	assert.NoError(t, err)

	nxc.SetAnyIsSelfSigned(true)
	assert.True(t, nxc.GetIsSelfSigned())
	nxc.SetValueIsSelfSigned(false)
	assert.False(t, nxc.GetIsSelfSigned())

	nxc.SetAnySubjectPublicKeyExponent(123)
	assert.Equal(t, nxc.GetSubjectPublicKeyExponent(), 123)
	nxc.SetValueSubjectPublicKeyExponent(234)
	assert.Equal(t, nxc.GetSubjectPublicKeyExponent(), 234)

	nxc.SetAnyVersion("v1.2")
	assert.Equal(t, nxc.GetVersion(), "v1.2")
	nxc.SetValueVersion("v1.44")
	assert.Equal(t, nxc.GetVersion(), "v1.44")

	nxc.SetAnySerialNumber("1.2.43")
	assert.Equal(t, nxc.GetSerialNumber(), "1.2.43")
	nxc.SetValueSerialNumber("41244")
	assert.Equal(t, nxc.GetSerialNumber(), "41244")

	nxc.SetAnySignatureAlgorithm("gg_y63")
	assert.Equal(t, nxc.GetSignatureAlgorithm(), "gg_y63")
	nxc.SetValueSignatureAlgorithm("977s-7s-v")
	assert.Equal(t, nxc.GetSignatureAlgorithm(), "977s-7s-v")

	nxc.SetAnyIssuer("adfvzzz")
	assert.Equal(t, nxc.GetIssuer(), "adfvzzz")
	nxc.SetValueIssuer("w233333")
	assert.Equal(t, nxc.GetIssuer(), "w233333")

	nxc.SetAnySubject("8dddu-cfv")
	assert.Equal(t, nxc.GetSubject(), "8dddu-cfv")
	nxc.SetValueSubject("__dwdw__")
	assert.Equal(t, nxc.GetSubject(), "__dwdw__")

	nxc.SetAnySubjectPublicKeyAlgorithm("Public-Key-Algorithm")
	assert.Equal(t, nxc.GetSubjectPublicKeyAlgorithm(), "Public-Key-Algorithm")
	nxc.SetValueSubjectPublicKeyAlgorithm("public-key-algorithm")
	assert.Equal(t, nxc.GetSubjectPublicKeyAlgorithm(), "public-key-algorithm")

	nxc.SetAnySubjectPublicKeyModulus("Public-Key-Modulus")
	assert.Equal(t, nxc.GetSubjectPublicKeyModulus(), "Public-Key-Modulus")
	nxc.SetValueSubjectPublicKeyModulus("public-key-modulus")
	assert.Equal(t, nxc.GetSubjectPublicKeyModulus(), "public-key-modulus")

	nxc.SetValueHashes(stixhelpers.HashesTypeSTIX{
		"hash_one":   "value_one",
		"hash_two":   "value_two",
		"hash_three": "value_three",
	})

	nxc.SetValueX509V3Extensions(somecomplextypesstixco.X509V3ExtensionsTypeSTIX{
		BasicConstraints: "Basic_Constraints",
		NameConstraints:  "Name_Constraints",
	})
	assert.Equal(t, nxc.GetX509V3Extensions().BasicConstraints, "Basic_Constraints")
	assert.Equal(t, nxc.GetX509V3Extensions().NameConstraints, "Name_Constraints")
}

/*
// X509CertificateCyberObservableObjectSTIX объект "X.509 Certificate Object", по терминологии STIX, представлет свойства
// сертификата X.509, определенные в рекомендациях ITU X.509 [X.509]. X.509  Certificate объект должен содержать по крайней
// мере одно cвойство специфичное для этого объекта (помимо type).
// IsSelfSigned - содержит индикатор, является ли сертификат самоподписным, то есть подписан ли он тем же субъектом, личность которого он удостоверяет.
// Hashes - содержит любые хэши, которые были вычислены для всего содержимого сертификата. Является типом данных словар, значения ключей которого должны
// быть из открытого словаря hash-algorithm-ov.
// Version- содержит версию закодированного сертификата
// SerialNumber - содержит уникальный идентификатор сертификата, выданного конкретным Центром сертификации.
// SignatureAlgorithm - содержит имя алгоритма, используемого для подписи сертификата.
// Issuer - содержит название удостоверяющего центра выдавшего сертификат
// ValidityNotBefore - время, в формате "2016-05-12T08:17:27.000Z", начала действия сертификата.
// ValidityNotAfter - время, в формате "2016-05-12T08:17:27.000Z", окончания действия сертификата.
// Subject - содержит имя сущности, связанной с открытым ключом, хранящимся в поле "subject public key" открого ключа сертификата.
// SubjectPublicKeyAlgorithm - содержит название алгоритма применяемого для шифрования данных, отправляемых субъекту.
// SubjectPublicKeyModulus - указывает модульную часть открытого ключа RSA.
// SubjectPublicKeyExponent - указывает экспоненциальную часть открытого ключа RSA субъекта в виде целого числа.
// X509V3Extension - указывает любые стандартные расширения X.509 v3, которые могут использоваться в сертификате.
type X509CertificateCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsSelfSigned              bool                                            `json:"is_self_signed" bson:"is_self_signed"`
	SubjectPublicKeyExponent  int                                             `json:"subject_public_key_exponent" bson:"subject_public_key_exponent"`
	Version                   string                                          `json:"version" bson:"version"`
	SerialNumber              string                                          `json:"serial_number" bson:"serial_number"`
	SignatureAlgorithm        string                                          `json:"signature_algorithm" bson:"signature_algorithm"`
	Issuer                    string                                          `json:"issuer" bson:"issuer"`
	Subject                   string                                          `json:"subject" bson:"subject"`
	SubjectPublicKeyAlgorithm string                                          `json:"subject_public_key_algorithm" bson:"subject_public_key_algorithm"`
	SubjectPublicKeyModulus   string                                          `json:"subject_public_key_modulus" bson:"subject_public_key_modulus"`
	ValidityNotBefore         string                                          `json:"validity_not_before" bson:"validity_not_before"`
	ValidityNotAfter          string                                          `json:"validity_not_after" bson:"validity_not_after"`
	Hashes                    stixhelpers.HashesTypeSTIX                      `json:"hashes" bson:"hashes"`
	X509V3Extensions          somecomplextypesstixco.X509V3ExtensionsTypeSTIX `json:"x509_v3_extensions" bson:"x509_v3_extensions"`
}
*/
