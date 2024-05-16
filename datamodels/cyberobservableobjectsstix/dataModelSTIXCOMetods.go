package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
)

/*********************************************************************************/
/********** 			Cyber-observable Objects STIX (МЕТОДЫ)			**********/
/*********************************************************************************/

/* --- URLCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (urlstix URLCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &urlstix); err != nil {
		return nil, err
	}

	return urlstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (urlstix URLCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(urlstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе URLCyberObservableObjectSTIX
func (urlstix URLCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(url--)[0-9a-f|-]+$`).MatchString(urlstix.ID)) {
		return false
	}

	if !urlstix.validateStructCommonFields() {
		return false
	}

	if !govalidator.IsURL(urlstix.Value) {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (urlstix URLCyberObservableObjectSTIX) SanitizeStruct() URLCyberObservableObjectSTIX {
	return urlstix
}

// GetID возвращает ID STIX объекта
func (urlstix URLCyberObservableObjectSTIX) GetID() string {
	return urlstix.ID
}

// GetType возвращает Type STIX объекта
func (urlstix URLCyberObservableObjectSTIX) GetType() string {
	return urlstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (urlstix URLCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(urlstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(urlstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'value': '%s'\n", urlstix.Value))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (urlstix URLCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   urlstix.ID,
		"type": urlstix.Type,
	}

	if urlstix.Value != "" {
		dataForIndex["value"] = urlstix.Value
	}

	return dataForIndex
}

/* --- UserAccountCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (uastix UserAccountCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &uastix); err != nil {
		return nil, err
	}

	return uastix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (uastix UserAccountCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(uastix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе UserAccountCyberObservableObjectSTIX
func (uastix UserAccountCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(user-account--)[0-9a-f|-]+$`).MatchString(uastix.ID)) {
		return false
	}

	if !uastix.validateStructCommonFields() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (uastix UserAccountCyberObservableObjectSTIX) SanitizeStruct() UserAccountCyberObservableObjectSTIX {
	uastix.UserID = commonlibs.StringSanitize(uastix.UserID)
	uastix.Credential = commonlibs.StringSanitize(uastix.Credential)
	uastix.AccountLogin = commonlibs.StringSanitize(uastix.AccountLogin)
	uastix.AccountType = uastix.AccountType.SanitizeStructOpenVocabTypeSTIX()
	uastix.DisplayName = commonlibs.StringSanitize(uastix.DisplayName)

	esize := len(uastix.Extensions)
	tmp := make(map[string]someextensionsstixco.UNIXAccountExtensionSTIX, esize)
	for k, v := range uastix.Extensions {
		result := sanitizeExtensionsSTIX(v)
		if ct, ok := result.(someextensionsstixco.UNIXAccountExtensionSTIX); ok {
			tmp[k] = ct
		}
	}
	uastix.Extensions = tmp

	return uastix
}

// GetID возвращает ID STIX объекта
func (uastix UserAccountCyberObservableObjectSTIX) GetID() string {
	return uastix.ID
}

// GetType возвращает Type STIX объекта
func (uastix UserAccountCyberObservableObjectSTIX) GetType() string {
	return uastix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (uastix UserAccountCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(uastix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(uastix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'user_id': '%s'\n", uastix.UserID))
	str.WriteString(fmt.Sprintf("'credential': '%s'\n", uastix.Credential))
	str.WriteString(fmt.Sprintf("'account_login': '%s'\n", uastix.AccountLogin))
	str.WriteString(fmt.Sprintf("'account_type': '%v'\n", uastix.AccountType))
	str.WriteString(fmt.Sprintf("'display_name': '%s'\n", uastix.DisplayName))
	str.WriteString(fmt.Sprintf("'is_service_account': '%v'\n", uastix.IsServiceAccount))
	str.WriteString(fmt.Sprintf("'is_privileged': '%v'\n", uastix.IsPrivileged))
	str.WriteString(fmt.Sprintf("'can_escalate_privs': '%v'\n", uastix.CanEscalatePrivs))
	str.WriteString(fmt.Sprintf("'is_disabled': '%v'\n", uastix.IsDisabled))
	str.WriteString(fmt.Sprintf("'account_created': '%v'\n", uastix.AccountCreated))
	str.WriteString(fmt.Sprintf("'account_expires': '%v'\n", uastix.AccountExpires))
	str.WriteString(fmt.Sprintf("'credential_last_changed': '%v'\n", uastix.CredentialLastChanged))
	str.WriteString(fmt.Sprintf("'account_first_login': '%v'\n", uastix.AccountFirstLogin))
	str.WriteString(fmt.Sprintf("'account_last_login': '%v'\n", uastix.AccountLastLogin))
	str.WriteString(fmt.Sprintln("'extensions':"))

	for k, v := range uastix.Extensions {
		str.WriteString(fmt.Sprintf("\t'%s':\n%v\n", k, toStringBeautiful(v)))
	}

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (uastix UserAccountCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   uastix.ID,
		"type": uastix.Type,
	}
}

/* --- WindowsRegistryKeyCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &wrkstix); err != nil {
		return nil, err
	}

	return wrkstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(wrkstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе WindowsRegistryKeyCyberObservableObjectSTIX
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(windows-registry-key--)[0-9a-f|-]+$`).MatchString(wrkstix.ID)) {
		return false
	}

	if !wrkstix.validateStructCommonFields() {
		return false
	}

	if !wrkstix.CreatorUserRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) SanitizeStruct() WindowsRegistryKeyCyberObservableObjectSTIX {
	wrkstix.Key = commonlibs.StringSanitize(wrkstix.Key)

	sizev := len(wrkstix.Values)
	if sizev > 0 {
		tmp := make([]somecomplextypesstixco.WindowsRegistryValueTypeSTIX, 0, sizev)

		for _, v := range wrkstix.Values {
			tmp = append(tmp, v.SanitizeStructWindowsRegistryValueTypeSTIX())
		}

		wrkstix.Values = tmp
	}

	return wrkstix
}

// GetID возвращает ID STIX объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) GetID() string {
	return wrkstix.ID
}

// GetType возвращает Type STIX объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) GetType() string {
	return wrkstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(wrkstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(wrkstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'key': '%s'\n", wrkstix.Key))
	str.WriteString(fmt.Sprintf("'values': \n%v", func(l []somecomplextypesstixco.WindowsRegistryValueTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'value '%d'':\n", k))
			str.WriteString(fmt.Sprintf("\t\t'name': '%s'\n", v.Name))
			str.WriteString(fmt.Sprintf("\t\t'data': '%s'\n", v.Data))
			str.WriteString(fmt.Sprintf("\t\t'data_type': '%s'\n", v.DataType))
		}

		return str.String()
	}(wrkstix.Values)))
	str.WriteString(fmt.Sprintf("'modified_time': '%v'\n", wrkstix.ModifiedTime))
	str.WriteString(fmt.Sprintf("'creator_user_ref': '%v'\n", wrkstix.CreatorUserRef))
	str.WriteString(fmt.Sprintf("'number_of_subkeys': '%d'\n", wrkstix.NumberOfSubkeys))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   wrkstix.ID,
		"type": wrkstix.Type,
	}
}

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

// ValidateStruct является валидатором параметров содержащихся в типе X509CertificateCyberObservableObjectSTIX
func (x509sstix X509CertificateCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(x509-certificate--)[0-9a-f|-]+$`).MatchString(x509sstix.ID)) {
		return false
	}

	if !x509sstix.validateStructCommonFields() {
		return false
	}

	if !x509sstix.Hashes.CheckHashesTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (x509sstix X509CertificateCyberObservableObjectSTIX) SanitizeStruct() X509CertificateCyberObservableObjectSTIX {
	x509sstix.Version = commonlibs.StringSanitize(x509sstix.Version)
	x509sstix.SerialNumber = commonlibs.StringSanitize(x509sstix.SerialNumber)
	x509sstix.SignatureAlgorithm = commonlibs.StringSanitize(x509sstix.SignatureAlgorithm)
	x509sstix.Issuer = commonlibs.StringSanitize(x509sstix.Issuer)
	x509sstix.Subject = commonlibs.StringSanitize(x509sstix.Subject)
	x509sstix.SubjectPublicKeyAlgorithm = commonlibs.StringSanitize(x509sstix.SubjectPublicKeyAlgorithm)
	x509sstix.SubjectPublicKeyModulus = commonlibs.StringSanitize(x509sstix.SubjectPublicKeyModulus)

	x509sstix.X509V3Extensions = x509sstix.X509V3Extensions.SanitizeStructX509V3ExtensionsTypeSTIX()

	return x509sstix
}

// GetID возвращает ID STIX объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) GetID() string {
	return x509sstix.ID
}

// GetType возвращает Type STIX объекта
func (x509sstix X509CertificateCyberObservableObjectSTIX) GetType() string {
	return x509sstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (x509sstix X509CertificateCyberObservableObjectSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(x509sstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(x509sstix.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'is_self_signed': '%v'\n", x509sstix.IsSelfSigned))
	str.WriteString(fmt.Sprintf("'hashes': \n%v", func(l map[string]string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'%s': '%v'\n", k, v))
		}

		return str.String()
	}(x509sstix.Hashes)))
	str.WriteString(fmt.Sprintf("'version': '%s'\n", x509sstix.Version))
	str.WriteString(fmt.Sprintf("'serial_number': '%s'\n", x509sstix.SerialNumber))
	str.WriteString(fmt.Sprintf("'signature_algorithm': '%s'\n", x509sstix.SignatureAlgorithm))
	str.WriteString(fmt.Sprintf("'issuer': '%s'\n", x509sstix.Issuer))
	str.WriteString(fmt.Sprintf("'validity_not_before': '%v'\n", x509sstix.ValidityNotBefore))
	str.WriteString(fmt.Sprintf("'validity_not_after': '%v'\n", x509sstix.ValidityNotAfter))
	str.WriteString(fmt.Sprintf("'subject': '%s'\n", x509sstix.Subject))
	str.WriteString(fmt.Sprintf("'subject_public_key_algorithm': '%s'\n", x509sstix.SubjectPublicKeyAlgorithm))
	str.WriteString(fmt.Sprintf("'subject_public_key_modulus': '%s'\n", x509sstix.SubjectPublicKeyModulus))
	str.WriteString(fmt.Sprintf("'subject_public_key_exponent': '%v'\n", x509sstix.SubjectPublicKeyExponent))
	str.WriteString(fmt.Sprintln("'x509_v3_extensions':"))
	str.WriteString(fmt.Sprintf("\t'basic_constraints': '%s'\n", x509sstix.X509V3Extensions.BasicConstraints))
	str.WriteString(fmt.Sprintf("\t'name_constraints': '%s'\n", x509sstix.X509V3Extensions.NameConstraints))
	str.WriteString(fmt.Sprintf("\t'policy_contraints': '%s'\n", x509sstix.X509V3Extensions.PolicyContraints))
	str.WriteString(fmt.Sprintf("\t'key_usage': '%s'\n", x509sstix.X509V3Extensions.KeyUsage))
	str.WriteString(fmt.Sprintf("\t'extended_key_usage': '%s'\n", x509sstix.X509V3Extensions.ExtendedKeyUsage))
	str.WriteString(fmt.Sprintf("\t'subject_key_identifier': '%s'\n", x509sstix.X509V3Extensions.SubjectKeyIdentifier))
	str.WriteString(fmt.Sprintf("\t'authority_key_identifier': '%s'\n", x509sstix.X509V3Extensions.AuthorityKeyIdentifier))
	str.WriteString(fmt.Sprintf("\t'subject_alternative_name': '%s'\n", x509sstix.X509V3Extensions.SubjectAlternativeName))
	str.WriteString(fmt.Sprintf("\t'issuer_alternative_name': '%s'\n", x509sstix.X509V3Extensions.IssuerAlternativeName))
	str.WriteString(fmt.Sprintf("\t'subject_directory_attributes': '%s'\n", x509sstix.X509V3Extensions.SubjectDirectoryAttributes))
	str.WriteString(fmt.Sprintf("\t'crl_distribution_points': '%s'\n", x509sstix.X509V3Extensions.CrlDistributionPoints))
	str.WriteString(fmt.Sprintf("\t'inhibit_any_policy': '%s'\n", x509sstix.X509V3Extensions.InhibitAnyPolicy))
	str.WriteString(fmt.Sprintf("\t'private_key_usage_period_not_before': '%v'\n", x509sstix.X509V3Extensions.PrivateKeyUsagePeriodNotBefore))
	str.WriteString(fmt.Sprintf("\t'private_key_usage_period_not_after': '%v'\n", x509sstix.X509V3Extensions.PrivateKeyUsagePeriodNotAfter))
	str.WriteString(fmt.Sprintf("\t'certificate_policies': '%s'\n", x509sstix.X509V3Extensions.CertificatePolicies))
	str.WriteString(fmt.Sprintf("\t'policy_mappings': '%s'\n", x509sstix.X509V3Extensions.PolicyMappings))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (x509sstix X509CertificateCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   x509sstix.ID,
		"type": x509sstix.Type,
	}
}
