package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- X509CertificateCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &x509sstix); err != nil {
		return nil, err
	}

	return x509sstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (x509sstix X509CertificateCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(x509sstix)

	return &result, err
}

// Get возвращает объект "X.509 Certificate Object", по терминалогии STIX, представлет свойства
// сертификата X.509, определенные в рекомендациях ITU X.509 [X.509]. X.509  Certificate объект должен содержать по крайней
// мере одно cвойство специфичное для этого объекта (помимо type).
func (e *X509CertificateCyberObservableObjectSTIX) Get() (*X509CertificateCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- IsSelfSigned property ---------
func (a *X509CertificateCyberObservableObjectSTIX) GetIsSelfSigned() bool {
	return a.IsSelfSigned
}

// SetValueIsSelfSigned устанавливает BOOL значение для поля IsSelfSigned
func (a *X509CertificateCyberObservableObjectSTIX) SetValueIsSelfSigned(v bool) {
	a.IsSelfSigned = v
}

// SetAnyIsSelfSigned устанавливает ЛЮБОЕ значение для поля IsSelfSigned
func (a *X509CertificateCyberObservableObjectSTIX) SetAnyIsSelfSigned(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsSelfSigned = v
	}
}

// -------- SubjectPublicKeyExponent property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSubjectPublicKeyExponent() int {
	return e.SubjectPublicKeyExponent
}

// SetValueSubjectPublicKeyExponent устанавливает значение для поля SubjectPublicKeyExponent
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSubjectPublicKeyExponent(v int) {
	e.SubjectPublicKeyExponent = v
}

// SetAnySubjectPublicKeyExponent устанавливает ЛЮБОЕ значение для поля SubjectPublicKeyExponent
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySubjectPublicKeyExponent(i interface{}) {
	e.SubjectPublicKeyExponent = commonlibs.ConversionAnyToInt(i)
}

// -------- Version property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetVersion() string {
	return e.Version
}

// SetValueVersion устанавливает значение для поля Version
func (e *X509CertificateCyberObservableObjectSTIX) SetValueVersion(v string) {
	e.Version = v
}

// SetAnyVersion устанавливает ЛЮБОЕ значение для поля Version
func (e *X509CertificateCyberObservableObjectSTIX) SetAnyVersion(i interface{}) {
	e.Version = fmt.Sprint(i)
}

// -------- SerialNumber property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSerialNumber() string {
	return e.SerialNumber
}

// SetValueSerialNumber устанавливает значение для поля SerialNumber
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSerialNumber(v string) {
	e.SerialNumber = v
}

// SetAnySerialNumber устанавливает ЛЮБОЕ значение для поля SerialNumber
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySerialNumber(i interface{}) {
	e.SerialNumber = fmt.Sprint(i)
}

// -------- SignatureAlgorithm property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSignatureAlgorithm() string {
	return e.SignatureAlgorithm
}

// SetValueSignatureAlgorithm устанавливает значение для поля SignatureAlgorithm
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSignatureAlgorithm(v string) {
	e.SignatureAlgorithm = v
}

// SetAnySignatureAlgorithm устанавливает ЛЮБОЕ значение для поля SignatureAlgorithm
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySignatureAlgorithm(i interface{}) {
	e.SignatureAlgorithm = fmt.Sprint(i)
}

// -------- Issuer property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetIssuer() string {
	return e.Issuer
}

// SetValueIssuer устанавливает значение для поля Issuer
func (e *X509CertificateCyberObservableObjectSTIX) SetValueIssuer(v string) {
	e.Issuer = v
}

// SetAnyIssuer устанавливает ЛЮБОЕ значение для поля Issuer
func (e *X509CertificateCyberObservableObjectSTIX) SetAnyIssuer(i interface{}) {
	e.Issuer = fmt.Sprint(i)
}

// -------- Subject property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSubject() string {
	return e.Subject
}

// SetValueSubject устанавливает значение для поля Subject
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSubject(v string) {
	e.Subject = v
}

// SetAnySubject устанавливает ЛЮБОЕ значение для поля Subject
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySubject(i interface{}) {
	e.Subject = fmt.Sprint(i)
}

// -------- SubjectPublicKeyAlgorithm property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSubjectPublicKeyAlgorithm() string {
	return e.SubjectPublicKeyAlgorithm
}

// SetValueSubjectPublicKeyAlgorithm устанавливает значение для поля SubjectPublicKeyAlgorithm
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSubjectPublicKeyAlgorithm(v string) {
	e.SubjectPublicKeyAlgorithm = v
}

// SetAnySubjectPublicKeyAlgorithm устанавливает ЛЮБОЕ значение для поля SubjectPublicKeyAlgorithm
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySubjectPublicKeyAlgorithm(i interface{}) {
	e.SubjectPublicKeyAlgorithm = fmt.Sprint(i)
}

// -------- SubjectPublicKeyModulus property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetSubjectPublicKeyModulus() string {
	return e.SubjectPublicKeyModulus
}

// SetValueSubjectPublicKeyModulus устанавливает значение для поля SubjectPublicKeyModulus
func (e *X509CertificateCyberObservableObjectSTIX) SetValueSubjectPublicKeyModulus(v string) {
	e.SubjectPublicKeyModulus = v
}

// SetAnySubjectPublicKeyModulus устанавливает ЛЮБОЕ значение для поля SubjectPublicKeyModulus
func (e *X509CertificateCyberObservableObjectSTIX) SetAnySubjectPublicKeyModulus(i interface{}) {
	e.SubjectPublicKeyModulus = fmt.Sprint(i)
}

// -------- ValidityNotBefore property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetValidityNotBefore() string {
	return e.ValidityNotBefore
}

// SetValueValidityNotBefore устанавливает значение в формате RFC3339 для поля ValidityNotBefore
func (e *X509CertificateCyberObservableObjectSTIX) SetValueValidityNotBefore(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.ValidityNotBefore = v

	return nil
}

// SetAnyValidityNotBefore устанавливает ЛЮБОЕ значение для поля ValidityNotBefore
func (e *X509CertificateCyberObservableObjectSTIX) SetAnyValidityNotBefore(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.ValidityNotBefore = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- ValidityNotAfter property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetValidityNotAfter() string {
	return e.ValidityNotAfter
}

// SetValueValidityNotAfter устанавливает значение в формате RFC3339 для поля ValidityNotAfter
func (e *X509CertificateCyberObservableObjectSTIX) SetValueValidityNotAfter(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.ValidityNotAfter = v

	return nil
}

// SetAnyValidityNotAfter устанавливает ЛЮБОЕ значение для поля ValidityNotAfter
func (e *X509CertificateCyberObservableObjectSTIX) SetAnyValidityNotAfter(i interface{}) {
	tmp := commonlibs.ConversionAnyToInt(i)
	e.ValidityNotAfter = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))
}

// -------- Hashes property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetHashes() stixhelpers.HashesTypeSTIX {
	return e.Hashes
}

func (e *X509CertificateCyberObservableObjectSTIX) SetValueHashes(v stixhelpers.HashesTypeSTIX) {
	e.Hashes = v
}

// -------- X509V3Extensions property ---------
func (e *X509CertificateCyberObservableObjectSTIX) GetX509V3Extensions() somecomplextypesstixco.X509V3ExtensionsTypeSTIX {
	return e.X509V3Extensions
}

func (e *X509CertificateCyberObservableObjectSTIX) SetValueX509V3Extensions(v somecomplextypesstixco.X509V3ExtensionsTypeSTIX) {
	e.X509V3Extensions = v
}

// ValidateStruct является валидатором параметров содержащихся в типе X509CertificateCyberObservableObjectSTIX
func (e X509CertificateCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(x509-certificate--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !e.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e X509CertificateCyberObservableObjectSTIX) SanitizeStruct() X509CertificateCyberObservableObjectSTIX {
	e.Version = commonlibs.StringSanitize(e.Version)
	e.SerialNumber = commonlibs.StringSanitize(e.SerialNumber)
	e.SignatureAlgorithm = commonlibs.StringSanitize(e.SignatureAlgorithm)
	e.Issuer = commonlibs.StringSanitize(e.Issuer)
	e.Subject = commonlibs.StringSanitize(e.Subject)
	e.SubjectPublicKeyAlgorithm = commonlibs.StringSanitize(e.SubjectPublicKeyAlgorithm)
	e.SubjectPublicKeyModulus = commonlibs.StringSanitize(e.SubjectPublicKeyModulus)

	e.X509V3Extensions = e.X509V3Extensions.SanitizeStructX509V3ExtensionsTypeSTIX()

	return e
}

// GetID возвращает ID STIX объекта
func (e X509CertificateCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e X509CertificateCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e X509CertificateCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)
	dubleWs := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'is_self_signed': '%v'\n", ws, e.IsSelfSigned))
	str.WriteString(fmt.Sprintf("%s'hashes': \n%v", ws, func(l map[string]string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.Hashes, num+1)))
	str.WriteString(fmt.Sprintf("%s'version': '%s'\n", ws, e.Version))
	str.WriteString(fmt.Sprintf("%s'serial_number': '%s'\n", ws, e.SerialNumber))
	str.WriteString(fmt.Sprintf("%s'signature_algorithm': '%s'\n", ws, e.SignatureAlgorithm))
	str.WriteString(fmt.Sprintf("%s'issuer': '%s'\n", ws, e.Issuer))
	str.WriteString(fmt.Sprintf("%s'validity_not_before': '%v'\n", ws, e.ValidityNotBefore))
	str.WriteString(fmt.Sprintf("%s'validity_not_after': '%v'\n", ws, e.ValidityNotAfter))
	str.WriteString(fmt.Sprintf("%s'subject': '%s'\n", ws, e.Subject))
	str.WriteString(fmt.Sprintf("%s'subject_public_key_algorithm': '%s'\n", ws, e.SubjectPublicKeyAlgorithm))
	str.WriteString(fmt.Sprintf("%s'subject_public_key_modulus': '%s'\n", ws, e.SubjectPublicKeyModulus))
	str.WriteString(fmt.Sprintf("%s'subject_public_key_exponent': '%v'\n", ws, e.SubjectPublicKeyExponent))
	str.WriteString(fmt.Sprintf("%s'x509_v3_extensions':", ws))
	str.WriteString(fmt.Sprintf("%s'basic_constraints': '%s'\n", dubleWs, e.X509V3Extensions.BasicConstraints))
	str.WriteString(fmt.Sprintf("%s'name_constraints': '%s'\n", dubleWs, e.X509V3Extensions.NameConstraints))
	str.WriteString(fmt.Sprintf("%s'policy_contraints': '%s'\n", dubleWs, e.X509V3Extensions.PolicyContraints))
	str.WriteString(fmt.Sprintf("%s'key_usage': '%s'\n", dubleWs, e.X509V3Extensions.KeyUsage))
	str.WriteString(fmt.Sprintf("%s'extended_key_usage': '%s'\n", dubleWs, e.X509V3Extensions.ExtendedKeyUsage))
	str.WriteString(fmt.Sprintf("%s'subject_key_identifier': '%s'\n", dubleWs, e.X509V3Extensions.SubjectKeyIdentifier))
	str.WriteString(fmt.Sprintf("%s'authority_key_identifier': '%s'\n", dubleWs, e.X509V3Extensions.AuthorityKeyIdentifier))
	str.WriteString(fmt.Sprintf("%s'subject_alternative_name': '%s'\n", dubleWs, e.X509V3Extensions.SubjectAlternativeName))
	str.WriteString(fmt.Sprintf("%s'issuer_alternative_name': '%s'\n", dubleWs, e.X509V3Extensions.IssuerAlternativeName))
	str.WriteString(fmt.Sprintf("%s'subject_directory_attributes': '%s'\n", dubleWs, e.X509V3Extensions.SubjectDirectoryAttributes))
	str.WriteString(fmt.Sprintf("%s'crl_distribution_points': '%s'\n", dubleWs, e.X509V3Extensions.CrlDistributionPoints))
	str.WriteString(fmt.Sprintf("%s'inhibit_any_policy': '%s'\n", dubleWs, e.X509V3Extensions.InhibitAnyPolicy))
	str.WriteString(fmt.Sprintf("%s'private_key_usage_period_not_before': '%v'\n", dubleWs, e.X509V3Extensions.PrivateKeyUsagePeriodNotBefore))
	str.WriteString(fmt.Sprintf("%s'private_key_usage_period_not_after': '%v'\n", dubleWs, e.X509V3Extensions.PrivateKeyUsagePeriodNotAfter))
	str.WriteString(fmt.Sprintf("%s'certificate_policies': '%s'\n", dubleWs, e.X509V3Extensions.CertificatePolicies))
	str.WriteString(fmt.Sprintf("%s'policy_mappings': '%s'\n", dubleWs, e.X509V3Extensions.PolicyMappings))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e X509CertificateCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
