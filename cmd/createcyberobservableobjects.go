package cmd

import (
	"encoding/json"

	"github.com/av-belyakov/methodstixobjects/datamodels/cyberobservableobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// NewArtifactCyberObservableObjectSTIX создает STIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки
// в кодировке base64 или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
func NewArtifactCyberObservableObjectSTIX() *cyberobservableobjectsstix.ArtifactCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("artifact")

	return &cyberobservableobjectsstix.ArtifactCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewAutonomousSystemCyberObservableObjectSTIX создает STIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
func NewAutonomousSystemCyberObservableObjectSTIX() *cyberobservableobjectsstix.AutonomousSystemCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("autonomous-system")

	return &cyberobservableobjectsstix.AutonomousSystemCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewDirectoryCyberObservableObjectSTIX создает STIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
func NewDirectoryCyberObservableObjectSTIX() *cyberobservableobjectsstix.DirectoryCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("directory")

	return &cyberobservableobjectsstix.DirectoryCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Ctime:        "1970-01-01T00:00:00+00:00",
		Mtime:        "1970-01-01T00:00:00+00:00",
		Atime:        "1970-01-01T00:00:00+00:00",
		ContainsRefs: []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewDomainNameCyberObservableObjectSTIX создает STIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
func NewDomainNameCyberObservableObjectSTIX() *cyberobservableobjectsstix.DomainNameCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("domain-name")

	return &cyberobservableobjectsstix.DomainNameCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		ResolvesToRefs: []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewEmailAddressCyberObservableObjectSTIX создает STIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
func NewEmailAddressCyberObservableObjectSTIX() *cyberobservableobjectsstix.EmailAddressCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("email-addr")

	return &cyberobservableobjectsstix.EmailAddressCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewEmailMessageCyberObservableObjectSTIX создает STIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
func NewEmailMessageCyberObservableObjectSTIX() *cyberobservableobjectsstix.EmailMessageCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("email-message")

	return &cyberobservableobjectsstix.EmailMessageCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Date:                   "1970-01-01T00:00:00+00:00",
		ReceivedLines:          []string(nil),
		ToRefs:                 []stixhelpers.IdentifierTypeSTIX(nil),
		CcRefs:                 []stixhelpers.IdentifierTypeSTIX(nil),
		BccRefs:                []stixhelpers.IdentifierTypeSTIX(nil),
		AdditionalHeaderFields: make(map[string]stixhelpers.DictionaryTypeSTIX),
	}
}

// NewCommonFileCyberObservableObjectSTIX создает общий STIX объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
func NewCommonFileCyberObservableObjectSTIX() *cyberobservableobjectsstix.CommonFileCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("file")

	return &cyberobservableobjectsstix.CommonFileCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Ctime:        "1970-01-01T00:00:00+00:00",
		Mtime:        "1970-01-01T00:00:00+00:00",
		Atime:        "1970-01-01T00:00:00+00:00",
		ContainsRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		Extensions:   make(map[string]*json.RawMessage),
	}
}

// NewFileCyberObservableObjectSTIX создает STIX объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
func NewFileCyberObservableObjectSTIX() *cyberobservableobjectsstix.FileCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("file")

	return &cyberobservableobjectsstix.FileCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Ctime:        "1970-01-01T00:00:00+00:00",
		Mtime:        "1970-01-01T00:00:00+00:00",
		Atime:        "1970-01-01T00:00:00+00:00",
		ContainsRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		Extensions:   make(map[string]interface{}),
	}
}

// NewIPv4AddressCyberObservableObjectSTIX создает STIX объект "IPv4 Address Object", по терминалогии STIX, содержит один или
// более IPv4 адресов, выраженных с помощью нотации CIDR.
func NewIPv4AddressCyberObservableObjectSTIX() *cyberobservableobjectsstix.IPv4AddressCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("ipv4-addr")

	return &cyberobservableobjectsstix.IPv4AddressCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		ResolvesToRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		BelongsToRefs:  []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewIPv6AddressCyberObservableObjectSTIX создает STIX объект "IPv6 Address Object", по терминалогии STIX, содержит один или более IPv6 адресов, выраженных с помощью нотации CIDR.
func NewIPv6AddressCyberObservableObjectSTIX() *cyberobservableobjectsstix.IPv6AddressCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("ipv6-addr")

	return &cyberobservableobjectsstix.IPv6AddressCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		ResolvesToRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		BelongsToRefs:  []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewMACAddressCyberObservableObjectSTIX создает STIX объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой
// один адрес управления доступом к среде (MAC).
func NewMACAddressCyberObservableObjectSTIX() *cyberobservableobjectsstix.MACAddressCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("mac-addr")

	return &cyberobservableobjectsstix.MACAddressCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewMutexCyberObservableObjectSTIX создает STIX объект "Mutex Object", по терминалогии STIX, содержит свойства объекта взаимного исключения (mutex).
func NewMutexCyberObservableObjectSTIX() *cyberobservableobjectsstix.MutexCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("mutex")

	return &cyberobservableobjectsstix.MutexCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewCommonNetworkTrafficCyberObservableObjectSTIX создает общий STIX объект "Network Traffic Object", по терминалогии STIX, содержит объект
// Сетевого трафика представляющий собой произвольный сетевой трафик, который исходит из источника и адресуется адресату.
func NewCommonNetworkTrafficCyberObservableObjectSTIX() *cyberobservableobjectsstix.CommonNetworkTrafficCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("network-traffic")

	return &cyberobservableobjectsstix.CommonNetworkTrafficCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Start:            "1970-01-01T00:00:00+00:00",
		End:              "1970-01-01T00:00:00+00:00",
		Protocols:        []string(nil),
		EncapsulatesRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		IPFix:            make(map[string]interface{}),
		Extensions:       make(map[string]*json.RawMessage),
	}
}

// NewNetworkTrafficCyberObservableObjectSTIX создает STIX объект "Network Traffic Object", по терминалогии STIX, содержит объект
// Сетевого трафика представляющий собой произвольный сетевой трафик, который исходит из источника и адресуется адресату.
func NewNetworkTrafficCyberObservableObjectSTIX() *cyberobservableobjectsstix.NetworkTrafficCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("network-traffic")

	return &cyberobservableobjectsstix.NetworkTrafficCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Start:            "1970-01-01T00:00:00+00:00",
		End:              "1970-01-01T00:00:00+00:00",
		Protocols:        []string(nil),
		EncapsulatesRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		IPFix:            make(map[string]string),
		Extensions:       make(map[string]interface{}),
	}
}

// NewCommonProcessCyberObservableObjectSTIX создает общий STIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
// выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
func NewCommonProcessCyberObservableObjectSTIX() *cyberobservableobjectsstix.CommonProcessCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("process")

	return &cyberobservableobjectsstix.CommonProcessCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		CreatedTime:          "1970-01-01T00:00:00+00:00",
		ChildRefs:            []stixhelpers.IdentifierTypeSTIX(nil),
		OpenedConnectionRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		EnvironmentVariables: make(map[string]string),
		Extensions:           make(map[string]*json.RawMessage),
	}
}

// NewProcessCyberObservableObjectSTIX создает STIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
// выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
func NewProcessCyberObservableObjectSTIX() *cyberobservableobjectsstix.ProcessCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("process")

	return &cyberobservableobjectsstix.ProcessCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		CreatedTime:          "1970-01-01T00:00:00+00:00",
		ChildRefs:            []stixhelpers.IdentifierTypeSTIX(nil),
		OpenedConnectionRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		EnvironmentVariables: make(map[string]string),
		Extensions:           make(map[string]interface{}),
	}
}

// NewSoftwareCyberObservableObjectSTIX создает STIX объект "Software Object", по терминологии STIX, содержит свойства, связанные с
// программным обеспечением, включая программные продукты.
func NewSoftwareCyberObservableObjectSTIX() *cyberobservableobjectsstix.SoftwareCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("software")

	return &cyberobservableobjectsstix.SoftwareCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		Languages: []string(nil),
	}
}

// NewURLCyberObservableObjectSTIX создает STIX объект "URL Object", по терминологии STIX, содержит унифицированный указатель информационного ресурса (URL).
func NewURLCyberObservableObjectSTIX() *cyberobservableobjectsstix.URLCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("url")

	return &cyberobservableobjectsstix.URLCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}

// NewUserAccountCyberObservableObjectSTIX создает STIX объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя, включая,
// учетные записи операционной системы, устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей
func NewUserAccountCyberObservableObjectSTIX() *cyberobservableobjectsstix.UserAccountCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("user-account")

	return &cyberobservableobjectsstix.UserAccountCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		AccountCreated:        "1970-01-01T00:00:00+00:00",
		AccountExpires:        "1970-01-01T00:00:00+00:00",
		CredentialLastChanged: "1970-01-01T00:00:00+00:00",
		AccountFirstLogin:     "1970-01-01T00:00:00+00:00",
		AccountLastLogin:      "1970-01-01T00:00:00+00:00",
		Extensions:            make(map[string]someextensionsstixco.UNIXAccountExtensionSTIX),
	}
}

// NewWindowsRegistryKeyCyberObservableObjectSTIX создает STIX объект "Windows Registry Key Object", по терминалогии STIX. Содержит описание значений полей
// раздела реестра Windows.
func NewWindowsRegistryKeyCyberObservableObjectSTIX() *cyberobservableobjectsstix.WindowsRegistryKeyCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("windows-registry-key")

	return &cyberobservableobjectsstix.WindowsRegistryKeyCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		ModifiedTime: "1970-01-01T00:00:00+00:00",
		Values:       []somecomplextypesstixco.WindowsRegistryValueTypeSTIX(nil),
	}
}

// NewX509CertificateCyberObservableObjectSTIX создает STIX объект "X.509 Certificate Object", по терминологии STIX, представлет свойства
// сертификата X.509, определенные в рекомендациях ITU X.509 [X.509]. X.509  Certificate объект должен содержать по крайней
// мере одно cвойство специфичное для этого объекта (помимо type).
func NewX509CertificateCyberObservableObjectSTIX() *cyberobservableobjectsstix.X509CertificateCyberObservableObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("x509-certificate")

	return &cyberobservableobjectsstix.X509CertificateCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *cpo.Get(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
		ValidityNotBefore:                                 "1970-01-01T00:00:00+00:00",
		ValidityNotAfter:                                  "1970-01-01T00:00:00+00:00",
	}
}
