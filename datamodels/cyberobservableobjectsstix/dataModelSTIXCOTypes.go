package cyberobservableobjectsstix

import (
	"encoding/json"
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/commonproperties"
	"github.com/av-belyakov/methodstixobjects/datamodels/commonpropertiesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/********** 			Cyber-observable Objects STIX (ТИПЫ)			**********/

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

// AutonomousSystemCyberObservableObjectSTIX объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Number - содержит номер присвоенный Автономной системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Name - название Автономной системы
// RIR - содержит название регионального Интернет-реестра (Regional Internet Registry) которым было дано имя Автономной системы
type AutonomousSystemCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Number int    `json:"number" bson:"number" required:"true"`
	Name   string `json:"name" bson:"name"`
	RIR    string `json:"rir" bson:"rir"`
}

// DirectoryCyberObservableObjectSTIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
// Path - указывает путь, как было первоначально замечено, к каталогу в файловой системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PathEnc - указывает наблюдаемую кодировку для пути. Значение ДОЛЖНО быть указано, если путь хранится в кодировке, отличной от Unicode.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания директории
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации или записи в директорию
// Atime - время, в формате "2016-05-12T08:17:27.000Z", последнего обращения к директории
// ContainsRefs - содержит список файловых объектов или директорий содержащихся внутри директории
type DirectoryCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Path         string                           `json:"path" bson:"path" required:"true"`
	PathEnc      string                           `json:"path_enc" bson:"path_enc"`
	Ctime        string                           `json:"ctime" bson:"ctime"`
	Mtime        string                           `json:"mtime" bson:"mtime"`
	Atime        string                           `json:"atime" bson:"atime"`
	ContainsRefs []stixhelpers.IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
}

// DomainNameCyberObservableObjectSTIX объект "Domain Name", по терминалогии STIX, содержит сетевое доменное имя
// Value - сетевое доменное имя (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - список ссылок на один или несколько IP-адресов или доменных имен, на которые разрешается доменное имя
type DomainNameCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value" required:"true"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
}

// EmailAddressCyberObservableObjectSTIX объект "Email Address", по терминалогии STIX, содержит представление единственного email адреса
// Value - email адрес (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// DisplayName - содержит единственное почтовое имя которое видит человек при просмотре письма
// BelongsToRef - учетная запись пользователя, которой принадлежит адрес электронной почты, в качестве ссылки на объект учетной записи пользователя
type EmailAddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value        string                         `json:"value" bson:"value"`
	DisplayName  string                         `json:"display_name" bson:"display_name"`
	BelongsToRef stixhelpers.IdentifierTypeSTIX `json:"belongs_to_ref" bson:"belongs_to_ref"`
}

// EmailMessageCyberObservableObjectSTIX объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
// IsMultipart - информирует содержит ли 'тело' email множественные MIME части (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Date - время, в формате "2016-05-12T08:17:27.000Z", когда email сообщение было отправлено
// ContentType - содержит содержимое 'Content-Type' заголовка email сообщения
// FromRef - содержит содержимое 'From:' заголовка email сообщения
// SenderRef - содержит содержимое поля 'Sender:' email сообщения
// ToRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'To:'
// CcRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'CC:'
// BccRefs - содержит список почтовых ящиков, которые являются получателями сообщения электронной почты, содержимое поля 'BCC:'
// MessageID - содержимое Message-ID email сообщения
// Subject - содержит тему сообщения
// ReceivedLines - содержит одно или несколько полей заголовка 'Received', которые могут быть включены в заголовки email
// AdditionalHeaderFields - содержит любые другие поля заголовка (за исключением date, received_lines, content_type, from_ref,
// sender_ref, to_ref, cc_ref, bcc_refs и subject), найденные в сообщении электронной почты в виде словаря
// Body - содержит тело сообщения
// BodyMultipart - содержит адает список MIME-части, которые составляют тело email. Это свойство НЕ ДОЛЖНО использоваться,
// если	is_multipart имеет значение false
// RawEmailRef - содержит 'сырое' бинарное содержимое email сообщения
type EmailMessageCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsMultipart            bool                                           `json:"is_multipart" bson:"is_multipart" required:"true"`
	ContentType            string                                         `json:"content_type" bson:"content_type"`
	MessageID              string                                         `json:"message_id" bson:"message_id"`
	Subject                string                                         `json:"subject" bson:"subject"`
	Body                   string                                         `json:"body" bson:"body"`
	Date                   string                                         `json:"date" bson:"date"`
	ReceivedLines          []string                                       `json:"received_lines" bson:"received_lines"`
	FromRef                stixhelpers.IdentifierTypeSTIX                 `json:"from_ref" bson:"from_ref"`
	SenderRef              stixhelpers.IdentifierTypeSTIX                 `json:"sender_ref" bson:"sender_ref"`
	RawEmailRef            stixhelpers.IdentifierTypeSTIX                 `json:"raw_email_ref" bson:"raw_email_ref"`
	ToRefs                 []stixhelpers.IdentifierTypeSTIX               `json:"to_refs" bson:"to_refs"`
	CcRefs                 []stixhelpers.IdentifierTypeSTIX               `json:"cc_refs" bson:"cc_refs"`
	BccRefs                []stixhelpers.IdentifierTypeSTIX               `json:"bcc_refs" bson:"bcc_refs"`
	BodyMultipart          []somecomplextypesstixco.EmailMIMEPartTypeSTIX `json:"body_multipart" bson:"body_multipart"`
	AdditionalHeaderFields map[string]stixhelpers.DictionaryTypeSTIX      `json:"additional_header_fields" bson:"additional_header_fields"`
}

// CommonFileCyberObservableObjectSTIX общий объект "File Object", по терминалогии STIX, содержит объект со свойствами файла
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext.
// В дополнении к ним пользователь может создавать свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
// Hashes - определяет словарь хешей для файла. При этом ДОЛЖНЫ использоватся ключи из открытого словаря hash-algorithm- ov.
// Size - содержит размер файла в байтах
// Name - содержит имя файла
// NameEnc - определяет кодировку имени файла. Содержимое должно соответствовать ревизии IANA от 2013-12-20.
// MagicNumberHex - указывает шестнадцатеричную константу (“магическое число”), связанную с определенным форматом файла, который соответствует этому файлу, если это применимо.
// MimeType - определяет MIME имени файла, например, application/msword.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания файла
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации файла
// Atime - время, в формате "2016-05-12T08:17:27.000Z", обращения к файлу
// ParentDirectoryRef - определяет родительскую директорию для файла. Объект ссылающийся на это свойство ДОЛЖЕН быть типом directory
// ContainsRefs - содержит ссылки на другие Cyber-observable Objects STIX, содержащиеся в файле, например другой файл, добавленный в конец файла, или IP-адрес, содержащийся где-то в файле.
// ContentRef - определяет контент файла. Данное значение ДОЛЖНО иметь тип artifact, то есть ссылатся на ArtifactCyberObservableObjectSTIX
type CommonFileCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Size               uint64                           `json:"size" bson:"size"`
	Name               string                           `json:"name" bson:"name"`
	NameEnc            string                           `json:"name_enc" bson:"name_enc"`
	MagicNumberHex     string                           `json:"magic_number_hex" bson:"magic_number_hex"`
	MimeType           string                           `json:"mime_type" bson:"mime_type"`
	Ctime              string                           `json:"ctime" bson:"ctime"`
	Mtime              string                           `json:"mtime" bson:"mtime"`
	Atime              string                           `json:"atime" bson:"atime"`
	Hashes             stixhelpers.HashesTypeSTIX       `json:"hashes" bson:"hashes"`
	ParentDirectoryRef stixhelpers.IdentifierTypeSTIX   `json:"parent_directory_ref" bson:"parent_directory_ref"`
	ContentRef         stixhelpers.IdentifierTypeSTIX   `json:"content_ref" bson:"content_ref"`
	ContainsRefs       []stixhelpers.IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
	Extensions         map[string]*json.RawMessage      `json:"extensions" bson:"extensions"`
}

// FileCyberObservableObjectSTIX объект "File Object", по терминалогии STIX, последекодирования из JSON (основной, рабочий объект)
// Extensions - определяет следующие расширения pdf-ext, archive-ext, ntfs-ext, raster-image-ext, windows-pebinary-ext. В дополнении
// к ним пользователь может создавать свои расширения. При этом ключ словаря должен однозначно идентифицировать тип расширения.
// Hashes - определяет словарь хешей для файла. При этом ДОЛЖНЫ использоватся ключи из открытого словаря hash-algorithm- ov.
// Size - содержит размер файла в байтах
// Name - содержит имя файла
// NameEnc - определяет кодировку имени файла. Содержимое должно соответствовать ревизии IANA от 2013-12-20.
// MagicNumberHex - указывает шестнадцатеричную константу (“магическое число”), связанную с определенным форматом файла, который соответствует этому файлу, если это применимо.
// MimeType - определяет MIME имени файла, например, application/msword.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания файла
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации файла
// Atime - время, в формате "2016-05-12T08:17:27.000Z", обращения к файлу
// ParentDirectoryRef - определяет родительскую директорию для файла. Объект ссылающийся на это свойство ДОЛЖЕН быть типом directory
// ContainsRefs - содержит ссылки на другие Cyber-observable Objects STIX, содержащиеся в файле, например другой файл, добавленный в конец файла, или IP-адрес, содержащийся где-то в файле.
// ContentRef - определяет контент файла. Данное значение ДОЛЖНО иметь тип artifact, то есть ссылатся на ArtifactCyberObservableObjectSTIX
type FileCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Size               uint64                           `json:"size" bson:"size"`
	Name               string                           `json:"name" bson:"name"`
	NameEnc            string                           `json:"name_enc" bson:"name_enc"`
	MagicNumberHex     string                           `json:"magic_number_hex" bson:"magic_number_hex"`
	MimeType           string                           `json:"mime_type" bson:"mime_type"`
	Ctime              string                           `json:"ctime" bson:"ctime"`
	Mtime              string                           `json:"mtime" bson:"mtime"`
	Atime              string                           `json:"atime" bson:"atime"`
	ParentDirectoryRef stixhelpers.IdentifierTypeSTIX   `json:"parent_directory_ref" bson:"parent_directory_ref"`
	Hashes             stixhelpers.HashesTypeSTIX       `json:"hashes" bson:"hashes"`
	ContentRef         stixhelpers.IdentifierTypeSTIX   `json:"content_ref" bson:"content_ref"`
	ContainsRefs       []stixhelpers.IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
	Extensions         map[string]interface{}           `json:"extensions" bson:"extensions"`
}

// IPv4AddressCyberObservableObjectSTIX объект "IPv4 Address Object", по терминалогии STIX, содержит один или
// более IPv4 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv4-адресов, выраженные с помощью нотации CIDR. Если данный
// объект IPv4-адреса представляет собой один IPv4-адрес, суффикс CIDR /32 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2,
// на которые разрешается IPv6-адрес. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv6-адрес.
// Объекты, на которые ссылается этот список, ДОЛЖНЫ быть типа autonomous-system.
type IPv4AddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []stixhelpers.IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}

// IPv6AddressCyberObservableObjectSTIX объект "IPv6 Address Object", по терминалогии STIX, содержит один или более IPv6 адресов, выраженных с помощью нотации CIDR.
// Value - указывает значения одного или нескольких IPv6-адресов, выраженные с помощью нотации CIDR. Если данный объект IPv6-адреса представляет собой один IPv6-адрес,
// суффикс CIDR /128 МОЖЕТ быть опущен. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ResolvesToRefs - указывает список ссылок на один или несколько MAC-адресов управления доступом к носителям уровня 2, на которые разрешается IPv6-адрес. Объекты,
// на которые ссылается этот список, ДОЛЖНЫ иметь тип macaddr.
// BelongsToRefs - указывает список ссылок на одну или несколько автономных систем (AS), к которым принадлежит IPv4-адрес. Объекты, на которые ссылается
// этот список,ДОЛЖНЫ быть типа autonomous-system.
type IPv6AddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `json:"value" bson:"value"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `json:"resolves_to_refs" bson:"resolves_to_refs"`
	BelongsToRefs  []stixhelpers.IdentifierTypeSTIX `json:"belongs_to_refs" bson:"belongs_to_refs"`
}

// MACAddressCyberObservableObjectSTIX объект "MAC Address Object", по терминалогии STIX, содержит объект MAC-адрес, представляющий собой
// один адрес управления доступом к среде (MAC).
// Value - Задает значение одного MAC-адреса. Значение MAC - адреса ДОЛЖНО быть представлено в виде одного строчного MAC-48 address,
// разделенного двоеточием, который ДОЛЖЕН включать начальные нули для каждого октета. Пример: 00:00:ab:cd:ef:01. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
type MACAddressCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}

// MutexCyberObservableObjectSTIX объект "Mutex Object", по терминалогии STIX, содержит свойства объекта взаимного исключения (mutex).
// Name - указывает имя объекта мьютекса (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
type MutexCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Name string `json:"name" bson:"name"`
}

// CommonNetworkTrafficCyberObservableObjectSTIX общий объект "Network Traffic Object", по терминалогии STIX, содержит объект
// Сетевого трафика представляющий собой произвольный сетевой трафик, который исходит из источника и адресуется адресату.
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать свои собственные.
// ключи словаря http-request-ext, cp-ext, icmp-ext, socket-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения
// словаря ДОЛЖНЫ содержать содержимое экземпляра расширения.
// Start - время, в формате "2016-05-12T08:17:27.000Z", инициирования сетевого трафика, если он известен.
// End - время, в формате "2016-05-12T08:17:27.000Z", окончания сетевого трафика, если он известен.
// IsActive - указывает, продолжается ли сетевой трафик. Если задано свойство end, то это свойство ДОЛЖНО быть false.
// SrcRef - указывает источник сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка, ДОЛЖЕН
// быть типа ipv4-addr, ipv6 - addr, mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// DstRef - указывает место назначения сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается ссылка,
// ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr,	mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// SrcPort - задает исходный порт, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// DstPort - задает порт назначения, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// Protocols - указывает протоколы, наблюдаемые в сетевом трафике, а также их соответствующее состояние.
// SrcByteCount - задает число байтов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstByteCount - задает число байтов в виде положительного целого числа, отправленных из пункта назначения в источник.
// SrcPackets - задает количество пакетов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstPackets - задает количество пакетов в виде положительного целого числа, отправленных от пункта назначения к источнику
// IPFix - указывает любые данные Экспорта информации IP-потока [IPFIX] для трафика в виде словаря. Каждая пара ключ/значение в словаре представляет
// имя/значение одного элемента IPFIX.	Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени элемента IPFIX.
// SrcPayloadRef - указывает байты, отправленные из источника в пункт назначения. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// DstPayloadRef - указывает байты, отправленные из пункта назначения в источник. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// EncapsulatesRefs - ссылки на другие объекты, инкапсулированные этим объектом. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
// EncapsulatedByRef - ссылки на другой объект сетевого трафика, который инкапсулирует этот объект. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
type CommonNetworkTrafficCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsActive          bool                             `json:"is_active" bson:"is_active"`
	SrcPort           int                              `json:"src_port" bson:"src_port"`
	DstPort           int                              `json:"dst_port" bson:"dst_port"`
	SrcPackets        int                              `json:"src_packets" bson:"src_packets"`
	DstPackets        int                              `json:"dst_packets" bson:"dst_packets"`
	SrcByteCount      uint64                           `json:"src_byte_count" bson:"src_byte_count"`
	DstByteCount      uint64                           `json:"dst_byte_count" bson:"dst_byte_count"`
	Start             string                           `json:"start" bson:"start"`
	End               string                           `json:"end" bson:"end"`
	Protocols         []string                         `json:"protocols" bson:"protocols"`
	SrcRef            stixhelpers.IdentifierTypeSTIX   `json:"src_ref" bson:"src_ref"`
	DstRef            stixhelpers.IdentifierTypeSTIX   `json:"dst_ref" bson:"dst_ref"`
	SrcPayloadRef     stixhelpers.IdentifierTypeSTIX   `json:"src_payload_ref" bson:"src_payload_ref"`
	DstPayloadRef     stixhelpers.IdentifierTypeSTIX   `json:"dst_payload_ref" bson:"dst_payload_ref"`
	EncapsulatedByRef stixhelpers.IdentifierTypeSTIX   `json:"encapsulated_by_ref" bson:"encapsulated_by_ref"`
	EncapsulatesRefs  []stixhelpers.IdentifierTypeSTIX `json:"encapsulates_refs" bson:"encapsulates_refs"`
	IPFix             map[string]interface{}           `json:"ipfix" bson:"ipfix"`
	Extensions        map[string]*json.RawMessage      `json:"extensions" bson:"extensions"`
}

// NetworkTrafficCyberObservableObjectSTIX объект "Network Traffic Object", по терминалогии STIX, содержит объект
// Сетевого трафика представляющий собой произвольный сетевой трафик, который исходит из источника и адресуется адресату.
// Extensions - объект Сетевого трафика определяет следующие расширения. В дополнение к ним производители МОГУТ создавать
// свои собственные. ключи словаря http-request-ext, cp-ext,
// Start - время, в формате "2016-05-12T08:17:27.000Z", инициирования сетевого трафика, если он известен.
// End - время, в формате "2016-05-12T08:17:27.000Z", окончания сетевого трафика, если он известен.
// IsActive - указывает, продолжается ли сетевой трафик. Если задано свойство end, то это свойство ДОЛЖНО быть false.
// SrcRef - указывает источник сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается
// ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr, mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// DstRef - указывает место назначения сетевого трафика в качестве ссылки на кибернаблюдаемый объект. Объект, на который ссылается
// ссылка, ДОЛЖЕН быть типа ipv4-addr, ipv6 - addr,	mac-addr или domain-name (для случаев, когда IP-адрес для доменного имени неизвестен).
// SrcPort - задает исходный порт, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// DstPort - задает порт назначения, используемый в сетевом трафике, в виде целого числа. Значение порта ДОЛЖНО находиться в диапазоне от 0 до 65535.
// Protocols - указывает протоколы, наблюдаемые в сетевом трафике, а также их соответствующее состояние.
// SrcByteCount - задает число байтов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstByteCount - задает число байтов в виде положительного целого числа, отправленных из пункта назначения в источник.
// SrcPackets - задает количество пакетов в виде положительного целого числа, отправленных от источника к месту назначения.
// DstPackets - задает количество пакетов в виде положительного целого числа, отправленных от пункта назначения к источнику
// IPFix - указывает любые данные Экспорта информации IP-потока [IPFIX] для трафика в виде словаря. Каждая пара ключ/значение в
// словаре представляет имя/значение одного элемента IPFIX.	Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени элемента IPFIX.
// SrcPayloadRef - указывает байты, отправленные из источника в пункт назначения. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// DstPayloadRef - указывает байты, отправленные из пункта назначения в источник. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
// EncapsulatesRefs - ссылки на другие объекты, инкапсулированные этим объектом. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
// EncapsulatedByRef - ссылки на другой объект сетевого трафика, который инкапсулирует этот объект. Объекты, на которые ссылается это свойство, ДОЛЖНЫ иметь тип network-traffic.
type NetworkTrafficCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsActive          bool                             `json:"is_active" bson:"is_active"`
	SrcPort           int                              `json:"src_port" bson:"src_port"`
	DstPort           int                              `json:"dst_port" bson:"dst_port"`
	SrcPackets        int                              `json:"src_packets" bson:"src_packets"`
	DstPackets        int                              `json:"dst_packets" bson:"dst_packets"`
	SrcByteCount      uint64                           `json:"src_byte_count" bson:"src_byte_count"`
	DstByteCount      uint64                           `json:"dst_byte_count" bson:"dst_byte_count"`
	Start             string                           `json:"start" bson:"start"`
	End               string                           `json:"end" bson:"end"`
	Protocols         []string                         `json:"protocols" bson:"protocols"`
	SrcRef            stixhelpers.IdentifierTypeSTIX   `json:"src_ref" bson:"src_ref"`
	DstRef            stixhelpers.IdentifierTypeSTIX   `json:"dst_ref" bson:"dst_ref"`
	SrcPayloadRef     stixhelpers.IdentifierTypeSTIX   `json:"src_payload_ref" bson:"src_payload_ref"`
	DstPayloadRef     stixhelpers.IdentifierTypeSTIX   `json:"dst_payload_ref" bson:"dst_payload_ref"`
	EncapsulatedByRef stixhelpers.IdentifierTypeSTIX   `json:"encapsulated_by_ref" bson:"encapsulated_by_ref"`
	EncapsulatesRefs  []stixhelpers.IdentifierTypeSTIX `json:"encapsulates_refs" bson:"encapsulates_refs"`
	IPFix             map[string]string                `json:"ipfix" bson:"ipfix"`
	Extensions        map[string]interface{}           `json:"extensions" bson:"extensions"`
}

// CommonProcessCyberObservableObjectSTIX общий объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
// выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные
// ключи словаря windows-process-exit, windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ
// содержать содержимое экземпляра расширения.
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который
// ссылается это свойство, ДОЛЖЕН иметь тип file.
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект,
// на который ссылается это свойство, ДОЛЖЕН иметь тип process.
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
// объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type CommonProcessCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsHidden             bool                             `json:"is_hidden" bson:"is_hidden"`
	PID                  int                              `json:"pid" bson:"pid"`
	Cwd                  string                           `json:"cwd" bson:"cwd"`
	CommandLine          string                           `json:"command_line" bson:"command_line"`
	CreatedTime          string                           `json:"created_time" bson:"created_time"`
	CreatorUserRef       stixhelpers.IdentifierTypeSTIX   `json:"creator_user_ref" bson:"creator_user_ref"`
	ImageRef             stixhelpers.IdentifierTypeSTIX   `json:"image_ref" bson:"image_ref"`
	ParentRef            stixhelpers.IdentifierTypeSTIX   `json:"parent_ref" bson:"parent_ref"`
	ChildRefs            []stixhelpers.IdentifierTypeSTIX `json:"child_refs" bson:"child_refs"`
	OpenedConnectionRefs []stixhelpers.IdentifierTypeSTIX `json:"opened_connection_refs" bson:"opened_connection_refs"`
	EnvironmentVariables map[string]string                `json:"environment_variables" bson:"environment_variables"`
	Extensions           map[string]*json.RawMessage      `json:"extensions" bson:"extensions"`
}

// ProcessCyberObservableObjectSTIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
// выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные.
// ключи словаря windows-process-exit, windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ
// содержать содержимое экземпляра расширения.
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который ссылается
// это свойство, ДОЛЖЕН иметь тип file.
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект, на который
// ссылается это свойство, ДОЛЖЕН иметь тип process.
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
// объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type ProcessCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsHidden             bool                             `json:"is_hidden" bson:"is_hidden"`
	PID                  int                              `json:"pid" bson:"pid"`
	Cwd                  string                           `json:"cwd" bson:"cwd"`
	CommandLine          string                           `json:"command_line" bson:"command_line"`
	CreatedTime          string                           `json:"created_time" bson:"created_time"`
	CreatorUserRef       stixhelpers.IdentifierTypeSTIX   `json:"creator_user_ref" bson:"creator_user_ref"`
	ImageRef             stixhelpers.IdentifierTypeSTIX   `json:"image_ref" bson:"image_ref"`
	ParentRef            stixhelpers.IdentifierTypeSTIX   `json:"parent_ref" bson:"parent_ref"`
	ChildRefs            []stixhelpers.IdentifierTypeSTIX `json:"child_refs" bson:"child_refs"`
	OpenedConnectionRefs []stixhelpers.IdentifierTypeSTIX `json:"opened_connection_refs" bson:"opened_connection_refs"`
	EnvironmentVariables map[string]string                `json:"environment_variables" bson:"environment_variables"`
	Extensions           map[string]interface{}           `json:"extensions" bson:"extensions"`
}

// SoftwareCyberObservableObjectSTIX объект "Software Object", по терминологии STIX, содержит свойства, связанные с
// программным обеспечением, включая программные продукты.
// Name - назвыание программного обеспечения
// CPE - содержит запись Common Platform Enumeration (CPE) для программного обеспечения, если она доступна. Значение этого свойства должно быть значением
// CPE v2.3 из официального словаря NVD CPE [NVD]
// SwID - содержит запись Тегов Software Identification ID (SWID) [SWID] для программного обеспечения, если таковая имеется. SwID помеченный tagId,
// является глобально уникальным идентификатором и ДОЛЖЕН использоваться как полномочие для идентификации помеченного продукта
// Languages -содержит языки, поддерживаемые программным обеспечением. Значение каждого елемента списка ДОЛЖНО быть кодом языка ISO 639-2
// Vendor - содержит название производителя программного обеспечения
// Version - содержит версию ПО
type SoftwareCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Name      string   `json:"name" bson:"name"`
	CPE       string   `json:"cpe" bson:"cpe"`
	SwID      string   `json:"swid" bson:"swid"`
	Vendor    string   `json:"vendor" bson:"vendor"`
	Version   string   `json:"version" bson:"version"`
	Languages []string `json:"languages" bson:"languages"`
}

// URLCyberObservableObjectSTIX объект "URL Object", по терминологии STIX, содержит унифицированный указатель информационного ресурса (URL).
// Value - содержит унифицированный указатель информационного ресурса (URL).
type URLCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `json:"value" bson:"value"`
}

// UserAccountCyberObservableObjectSTIX объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя, включая,
// учетные записи операционной системы, устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей
// Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств, определенных ниже, ДОЛЖНО быть инициализировано
// при использовании этого объекта
// Extensions - содержит словарь расширяющий тип "User Account Object" одно из расширений "unix-account-ext", реализуется описанным ниже
// типом, UNIXAccountExtensionSTIX кроме этого производитель может созавать свои собственные типы расширений. Ключи данного словаря
// идентифицируют тип расширения по имени, значения являются содержимым экземпляра расширения
// UserID - содержит идентификатор учетной записи. Формат идентификатора зависит от системы в которой находится данная учетная запись пользователя,
// и может быть числовым идентификатором, идентификатором GUID, именем учетной записи, адресом электронной почты и т.д. Свойство  UserId должно
// быть заполнено любым значанием, являющимся уникальным идентификатором системы, членом которой является учетная запись. Например, в системах UNIX он
// будет заполнено значением UID
// Credential - содержит учетные данные пользователя в открытом виде. Предназначено только для закрытого применения при изучении метаданных вредоносных
// программ при их исследовании (например, жестко закодированный пароль администратора домена, который вредоносная программа пытается использовать
// реализации тактики для бокового (латерального) перемещения) и не должно применяться для совместного пользования PII
// AccountLogin - содержит логин пользователя. Используется в тех случаях,когда свойство UserId указывает другие данные, чем то, что пользователь вводит
// при входе в систему
// AccountType - содержит одно, из заранее определенных (предложенных) значений. Является типом аккаунта. Значения этого свойства берутся из множества
// закрепленного в открытом словаре, account-type-ov
// DisplayName - содержит отображаемое имя учетной записи, которое будет отображаться в пользовательских интерфейсах. В Unix, это равносильно полю
// gecos (gecos это поле учётной записи пользователя в файле /etc/passwd )
// IsServiceAccount - содержит индикатор, сигнализирующий что, учетная запись связана с сетевой службой или системным процессом (демоном), а не с
// конкретным человеком. (системный пользователь)
// IsPrivileged - содержит индикатор, сигнализирующий что, учетная запись имеет повышенные привилегии (например, в случае root в Unix или учетной
// записи администратора Windows)
// CanEscalatePrivs  - содержит индикатор, сигнализирующий что, учетная запись имеет возможность повышать привилегии (например, в случае sudo в
// Unix или учетной записи администратора домена Windows)
// IsDisabled  - содержит индикатор, сигнализирующий что, учетная запись отключена
// AccountCreated - время, в формате "2016-05-12T08:17:27.000Z", создания аккаунта
// AccountExpires - время, в формате "2016-05-12T08:17:27.000Z", истечения срока действия учетной записи.
// CredentialLastChanged - время, в формате "2016-05-12T08:17:27.000Z", когда учетные данные учетной записи были изменены в последний раз.
// AccountFirstLogin - время, в формате "2016-05-12T08:17:27.000Z", первого доступа к учетной записи
// AccountLastLogin - время, в формате "2016-05-12T08:17:27.000Z", когда к учетной записи был последний доступ.
type UserAccountCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsServiceAccount      bool                                                     `json:"is_service_account" bson:"is_service_account"`
	IsPrivileged          bool                                                     `json:"is_privileged" bson:"is_privileged"`
	CanEscalatePrivs      bool                                                     `json:"can_escalate_privs" bson:"can_escalate_privs"`
	IsDisabled            bool                                                     `json:"is_disabled" bson:"is_disabled"`
	UserID                string                                                   `json:"user_id" bson:"user_id"`
	Credential            string                                                   `json:"credential" bson:"credential"`
	AccountLogin          string                                                   `json:"account_login" bson:"account_login"`
	DisplayName           string                                                   `json:"display_name" bson:"display_name"`
	AccountCreated        time.Time                                                `json:"account_created" bson:"account_created"`
	AccountExpires        time.Time                                                `json:"account_expires" bson:"account_expires"`
	CredentialLastChanged time.Time                                                `json:"credential_last_changed" bson:"credential_last_changed"`
	AccountFirstLogin     time.Time                                                `json:"account_first_login" bson:"account_first_login"`
	AccountLastLogin      time.Time                                                `json:"account_last_login" bson:"account_last_login"`
	AccountType           stixhelpers.OpenVocabTypeSTIX                            `json:"account_type" bson:"account_type"`
	Extensions            map[string]someextensionsstixco.UNIXAccountExtensionSTIX `json:"extensions" bson:"extensions"`
}

// WindowsRegistryKeyCyberObservableObjectSTIX объект "Windows Registry Key Object", по терминалогии STIX. Содержит описание значений полей
// раздела реестра Windows. Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств,определенных ниже,
// должно быть инициализировано при использовании этого объекта.
// Key - содержит полный путь к разделу реестра. Значение ключа,должно быть сохранено в регистре. В название ключа все сокращения должны быть
// раскрыты. Например, вместо HKLM следует использовать HKEY_LOCAL_MACHINE.
// Values - содержит значения, найденные в разделе реестра.
// ModifiedTime - время, в формате "2016-05-12T08:17:27.000Z", последнего изменения раздела реестра.
// CreatorUserRef - содержит ссылку на учетную запись пользователя, из под которой создан раздел реестра. Объект, на который ссылается это свойство, должен иметь тип user-account.
// NumberOfSubkeys - Указывает количество подразделов, содержащихся в разделе реестра.
type WindowsRegistryKeyCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	NumberOfSubkeys int                                                   `json:"number_of_subkeys" bson:"number_of_subkeys"`
	Key             string                                                `json:"key" bson:"key"`
	ModifiedTime    time.Time                                             `json:"modified_time" bson:"modified_time"`
	CreatorUserRef  stixhelpers.IdentifierTypeSTIX                        `json:"creator_user_ref" bson:"creator_user_ref"`
	Values          []somecomplextypesstixco.WindowsRegistryValueTypeSTIX `json:"values" bson:"values"`
}

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
	ValidityNotBefore         time.Time                                       `json:"validity_not_before" bson:"validity_not_before"`
	ValidityNotAfter          time.Time                                       `json:"validity_not_after" bson:"validity_not_after"`
	Hashes                    stixhelpers.HashesTypeSTIX                      `json:"hashes" bson:"hashes"`
	X509V3Extensions          somecomplextypesstixco.X509V3ExtensionsTypeSTIX `json:"x509_v3_extensions" bson:"x509_v3_extensions"`
}
