package somecomplextypesstixco

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/********** 			Некоторые 'сложные' типы относящиеся к Cyber-observable Objects STIX 			**********/

// AlternateDataStreamTypeSTIX тип "alternate-data-stream-type", по терминалогии STIX, содержит альтернативные потоки данных для NTFS
// Name - определяет имя альтернативного потока данных (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Hashes - задает словарь хэшей для данных, содержащихся в альтернативном потоке данных. Ключи для данного свойства ДОЛЖНЫ братся из открытого словоря hash-algorithm-ov.
// Size - определяет размер альтернативного потока данных, в байтах
type AlternateDataStreamTypeSTIX struct {
	Name   string                     `json:"name" bson:"name"`
	Hashes stixhelpers.HashesTypeSTIX `json:"hashes" bson:"hashes"`
	Size   uint64                     `json:"size" bson:"size"`
}

// ExifTags тип "exif_tags" набор тегов EXIF, найденных в файле изображения в виде словаря
type ExifTags struct {
	Make        string `json:"make" bson:"make"`
	Model       string `json:"model" bson:"model"`
	XResolution int    `json:"xResolution" bson:"xResolution"`
	YResolution int    `json:"yResolution" bson:"yResolution"`
}

// WindowsPEOptionalHeaderTypeSTIX тип "windows-pe-optional-header-type", по терминалогии STIX, содержит свойства опциональные для PE файла
// MagicHex - хеш значение информирующее о типе PE файла
// MajorLinkerVersion - указывает основной номер версии компоновщика.
// MinorLinkerVersion - указывает номер младшей версии компоновщика.
// SizeOfCode - указывает размер кода (текста) секции.
// SizeOfInitializedData - задает размер инициализированного раздела данных.
// SizeOfUninitializedData - задает размер раздела неинициализированных данных.
// AddressOfEntryPoint - указывает адрес точки входа относительно базы изображения при загрузке исполняемого файла в память.
// BaseOfCode - указывает адрес, который находится относительно базы изображений раздела начала кода при его загрузке в память.
// BaseOfData - указывает адрес, который находится относительно базы изображений раздела beginning-of-data при его загрузке в память.
// ImageBase - задает предпочтительный адрес первого байта изображения при загрузке в память.
// SectionAlignment - задает выравнивание (в байтах) PE-секций при их загрузке в память.
// FileAlignment - задает коэффициент (в байтах), который используется для выравнивания необработанных данных разделов в файле изображения.
// MajorOSVersion - указывает основной номер версии требуемой операционной системы.
// MinorOSVersion - указывает младший номер версии требуемой операционной системы.
// MajorImageVersion - указывает основной номер версии образа.
// MinorImageVersion - Указывает младший номер версии образа.
// MajorSubsystemVersion - указывает основной номер версии подсистемы.
// MinorSubsystemVersion - указывает младший номер версии подсистемы.
// Win32VersionValueHex - указывает зарезервированное значение версии win32.
// SizeOfImage - задает размер изображения в байтах, включая все заголовки, по мере загрузки изображения в память.
// SizeOfHeaders - задает комбинированный размер заголовков MS-DOS, PE и разделов, округленный до кратного значения, указанного в заголовке file_alignment.
// ChecksumHex - задает контрольную сумму двоичного файла PE.
// SubsystemHex - указывает подсистему (например, графический интерфейс, драйвер устройства и т.д.), необходимую для запуска этого образа.
// DllCharacteristicsHex - задает флаги, характеризующие двоичный файл PE.
// SizeOfStackReserve -указывает размер стека для резервирования в байтах.
// SizeOfStackCommit - указывает размер стека для фиксации в байтах.
// SizeOfHeapReserve - задает размер локального пространства кучи для резервирования в байтах.
// SizeOfHeapCommit - указывает размер локального пространства кучи для фиксации в байтах.
// LoaderFlagsHex - указывает зарезервированные флаги загрузчика.
// NumberOfRvaAndSizes - указывает количество записей каталога данных в оставшейся части необязательного заголовка.
// Hashes - указывает все хэши, которые были вычислены для необязательного заголовка.
type WindowsPEOptionalHeaderTypeSTIX struct {
	MagicHex                string                       `json:"magic_hex" bson:"magic_hex"`
	MajorLinkerVersion      int                          `json:"major_linker_version" bson:"major_linker_version"`
	MinorLinkerVersion      int                          `json:"minor_linker_version" bson:"minor_linker_version"`
	SizeOfCode              int                          `json:"size_of_code" bson:"size_of_code"`
	SizeOfInitializedData   int                          `json:"size_of_initialized_data" bson:"size_of_initialized_data"`
	SizeOfUninitializedData int                          `json:"size_of_uninitialized_data" bson:"size_of_uninitialized_data"`
	AddressOfEntryPoint     int                          `json:"address_of_entry_point" bson:"address_of_entry_point"`
	BaseOfCode              int                          `json:"base_of_code" bson:"base_of_code"`
	BaseOfData              int                          `json:"base_of_data" bson:"base_of_data"`
	ImageBase               int                          `json:"image_base" bson:"image_base"`
	SectionAlignment        int                          `json:"section_alignment" bson:"section_alignment"`
	FileAlignment           int                          `json:"file_alignment" bson:"file_alignment"`
	MajorOSVersion          int                          `json:"major_os_version" bson:"major_os_version"`
	MinorOSVersion          int                          `json:"minor_os_version" bson:"minor_os_version"`
	MajorImageVersion       int                          `json:"major_image_version" bson:"major_image_version"`
	MinorImageVersion       int                          `json:"minor_image_version" bson:"minor_image_version"`
	MajorSubsystemVersion   int                          `json:"major_subsystem_version" bson:"major_subsystem_version"`
	MinorSubsystemVersion   int                          `json:"minor_subsystem_version" bson:"minor_subsystem_version"`
	Win32VersionValueHex    string                       `json:"win32_version_value_hex" bson:"win32_version_value_hex"`
	SizeOfImage             int                          `json:"size_of_image" bson:"size_of_image"`
	SizeOfHeaders           int                          `json:"size_of_headers" bson:"size_of_headers"`
	ChecksumHex             string                       `json:"checksum_hex" bson:"checksum_hex"`
	SubsystemHex            string                       `json:"subsystem_hex" bson:"subsystem_hex"`
	DllCharacteristicsHex   string                       `json:"dll_characteristics_hex" bson:"dll_characteristics_hex"`
	SizeOfStackReserve      int                          `json:"size_of_stack_reserve" bson:"size_of_stack_reserve"`
	SizeOfStackCommit       int                          `json:"size_of_stack_commit" bson:"size_of_stack_commit"`
	SizeOfHeapReserve       int                          `json:"size_of_heap_reserve" bson:"size_of_heap_reserve"`
	SizeOfHeapCommit        int                          `json:"size_of_heap_commit" bson:"size_of_heap_commit"`
	LoaderFlagsHex          string                       `json:"loader_flags_hex" bson:"loader_flags_hex"`
	NumberOfRvaAndSizes     int                          `json:"number_of_rva_and_sizes" bson:"number_of_rva_and_sizes"`
	Hashes                  []stixhelpers.HashesTypeSTIX `json:"hashes" bson:"hashes"`
}

// WindowsPESectionTypeSTIX тип "windows-pe-section-type", по терминалогии STIX, определяет методанные PE-секции файла
// Name - определяет имя секции (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Size - размер секции в байтах.
// Entropy - задает вычисленную энтропию для сечения, рассчитанную с использованием алгоритма Шеннона.
// Hashes - задает любые хэши, вычисляемые по разделу. Свойство ДОЛЖНО содержать ключи из открытого словаря hash-algorithm-ov
type WindowsPESectionTypeSTIX struct {
	Name    string                     `json:"name" bson:"name"`
	Size    uint64                     `json:"size" bson:"size"`
	Entropy float64                    `json:"entropy" bson:"entropy"`
	Hashes  stixhelpers.HashesTypeSTIX `json:"hashes" bson:"hashes"`
}

// EmailMIMEPartTypeSTIX тип "email-mime-part-type", по терминалогии STIX, содержит один компонент тела email из нескольких частей
// Body - содержит содержимое части MIME, если content_type не указан или начинается с text/ (например, в случае обычного текста или HTML-письма)
// BodyRawRef - содержит содержимое нетекстовых частей MIME, то есть тех, чей content_type не начинается с text, в качестве
//
//	ссылки на объект артефакта или Файловый объект
//
// ContentType - содержимое поля 'Content-Type' заголовка MIME части email
// ContentDisposition - содержимое поля 'Content-Disposition' заголовка MIME части email
type EmailMIMEPartTypeSTIX struct {
	Body               string                         `json:"body" bson:"body"`
	BodyRawRef         stixhelpers.IdentifierTypeSTIX `json:"body_raw_ref" bson:"body_raw_ref"`
	ContentType        string                         `json:"content_type" bson:"content_type"`
	ContentDisposition string                         `json:"content_disposition" bson:"content_disposition"`
}

// WindowsRegistryValueTypeSTIX объект "Windows Registry Value Type", по терминалогии STIX. Данный тип фиксирует
//
//	значения свойств находящихся в разделе реестра Windows. Поскольку все свойства этого типа являются необязательными,
//
// по крайней мере одно из свойств, определенных ниже, должно быть инициализировано при использовании этого типа.
// Name - содержит название параметра реестра. Для указания значения ключа реестра по умолчанию необходимо использовать пустую строку.
// Data - содержит данные, содержащиеся в значении реестра.
// DataType - содержит тип данных реестра (REG_*), используемый в значении реестра. Значения этого свойства должны быть получены из перечисления windows-registry-datatype enum.
type WindowsRegistryValueTypeSTIX struct {
	Name     string                   `json:"name" bson:"name"`
	Data     string                   `json:"data" bson:"data"`
	DataType stixhelpers.EnumTypeSTIX `json:"data_type" bson:"data_type"`
}

// X509V3ExtensionsTypeSTIX тип "X.509 v3 Extensions Type", по терминалогии STIX. Описывает поля расширения X.509 v3, фиксрует дополнительную информацию,
//
//	такую как альтернативные имена субъектов. Объект, использующий тип "x509-v3-extensions-type", должен определить хотя бы одно из этих полей в нем.
//	Данный тип расширяет только объекты "X.509 Certificate Object".
//
// BasicConstraints - задает многозначное расширение, которое указывает, является ли сертификат сертификатом Удостоверяющего центра (CA). Первое (обязательное)
//
//	название CA сопровождается истинным или ложным. Если CA имеет значение TRUE, то может быть включено необязательное имя pathlen, за которым следует
//	неотрицательное значение. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.19.
//
// NameConstraints - указывает пространство имен, в котором должны располагаться все имена  применяемые в сертификатах указанных в пути сертификации. Также
//
//	эквивалентно значению идентификатора объекта (OID) 2.5.29.30.
//
// PolicyConstraints - содержит любые ограничения на проверку сертификатов, выданных Удостоверяющим центром.  Также эквивалентно значению идентификатора
//
//	объекта (OID) 2.5.29.36.
//
// KeyUsage - задает многозначное расширение, состоящее из списка имен разрешенных для использования ключей. Также эквивалентно значению идентификатора
//
//	объекта (OID) 2.5.29.15.
//
// ExtendedKeyUsage - содержит список целей, для которых может использоваться открытый ключ сертификата. Также эквивалентно значению идентификатора объекта
//
//	(OID) 2.5.29.37.
//
// SubjectKeyIdentifier - указывает идентификатор, который обеспечивает средство идентификации сертификатов, содержащих определенный открытый ключ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.14.
// AuthorityKeyIdentifier - содержит идентификатор, который обеспечивает средство идентификации открытого ключа, соответствующего закрытому ключу, используемому
//
//	для подписи сертификата. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.35.
//
// SubjectAlternativeName - указывает дополнительные идентификаторы, которые должны быть привязаны к субъекту сертификата. Также эквивалентно значению
//
//	идентификатора объекта (OID) 2.5.29.17.
//
// IssuerAlternativeName - указывает дополнительные идентификаторы, которые должны быть привязаны к эмитенту сертификата. Также эквивалентно значению
//
//	идентификатора объекта (OID) 2.5.29.18.
//
// SubjectDirectoryAttributes - указывает идентификационные признаки (например, национальность) субъекта. Также эквивалентно значению идентификатора
//
//	объекта (OID) 2.5.29.9.
//
// CrlDistributionPoints - указывает способ получения информации CRL. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.31.
// InhibitAnyPolicy - содержит количество дополнительных сертификатов, которые могут появиться в пути до того, как любая Политика больше не будет разрешена.
//
//	Также эквивалентно значению идентификатора объекта (OID) 2.5.29.54.
//
// PrivateKeyUsagePeriodNotBefore - время, в формате "2016-05-12T08:17:27.000Z", начала срока действия закрытого ключа, если он отличается от срока действия сертификата.
// PrivateKeyUsagePeriodNotAfter -  время, в формате "2016-05-12T08:17:27.000Z", окончания срока действия закрытого ключа, если он отличается от срока действия сертификата.
// CertificatePolicies - содержит последовательность из одного или нескольких терминов информации о политике, каждый из которых состоит из идентификатора
//
//	объекта (OID) и необязательных квалификаторов. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.32.
//
// PolicyMappings - содержит одну или несколько пар OID; каждая пара включает issuerDomainPolicy и subjectDomainPolicy. Пары индикаторов указывают на то,
//
//	считает ли выдающий УЦ свою issuerDomainPolicy эквивалентной subjectDomainPolicy субъекта УЦ. Также эквивалентно значению идентификатора объекта (OID) 2.5.29.33.
type X509V3ExtensionsTypeSTIX struct {
	BasicConstraints               string    `json:"basic_constraints" bson:"basic_constraints"`
	NameConstraints                string    `json:"name_constraints" bson:"name_constraints"`
	PolicyContraints               string    `json:"policy_contraints" bson:"policy_contraints"`
	KeyUsage                       string    `json:"key_usage" bson:"key_usage"`
	ExtendedKeyUsage               string    `json:"extended_key_usage" bson:"extended_key_usage"`
	SubjectKeyIdentifier           string    `json:"subject_key_identifier" bson:"subject_key_identifier"`
	AuthorityKeyIdentifier         string    `json:"authority_key_identifier" bson:"authority_key_identifier"`
	SubjectAlternativeName         string    `json:"subject_alternative_name" bson:"subject_alternative_name"`
	IssuerAlternativeName          string    `json:"issuer_alternative_name" bson:"issuer_alternative_name"`
	SubjectDirectoryAttributes     string    `json:"subject_directory_attributes" bson:"subject_directory_attributes"`
	CrlDistributionPoints          string    `json:"crl_distribution_points" bson:"crl_distribution_points"`
	InhibitAnyPolicy               string    `json:"inhibit_any_policy" bson:"inhibit_any_policy"`
	PrivateKeyUsagePeriodNotBefore time.Time `json:"private_key_usage_period_not_before" bson:"private_key_usage_period_not_before"`
	PrivateKeyUsagePeriodNotAfter  time.Time `json:"private_key_usage_period_not_after" bson:"private_key_usage_period_not_after"`
	CertificatePolicies            string    `json:"certificate_policies" bson:"certificate_policies"`
	PolicyMappings                 string    `json:"policy_mappings" bson:"policy_mappings"`
}
