package someextensionsstixco

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/********** 			Некоторые расширения, относящиеся к Cyber-observable Objects STIX 			**********/

// ArchiveFileExtensionSTIX тип "archive-ext", по терминалогии STIX, содержит рассширение архивного файла. В ней задается расширение по умолчанию для захвата свойств,
//
//	специфичных для архивных файлов
//
// ContainsRefs - данное свойство определяет файлы содержащиеся в архиве. ДОЛЖНО содержать список типа file или directory (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Comment - определяет комментарий включенный как часть архивного файла
type ArchiveFileExtensionSTIX struct {
	ContainsRefs []stixhelpers.IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
	Comment      string                           `json:"comment" bson:"comment"`
}

// NTFSFileExtensionSTIX  тип "ntfs-ext", по терминалогии STIX, содержит расширение определяющее значения свойств, характерные для файловой системы NTFS
// SID - опреляет безопасный идентификатор, связанный с файлом
// AlternateDataStreams - определяет список альтернативных NTFS потоков данных связанных с файлом
type NTFSFileExtensionSTIX struct {
	SID                  string                                               `json:"sid" bson:"sid"`
	AlternateDataStreams []somecomplextypesstixco.AlternateDataStreamTypeSTIX `json:"alternate_data_streams" bson:"alternate_data_streams"`
}

// PDFFileExtensionSTIX тип "pdf-ext", по терминалогии STIX, содержит свойства специфичные для файлов в формате PDF
// Version - номер версии, берется из заголовка PDF документа
// IsOptimized - определяетс был ли PDF файл оптимизирован
// DocumentInfoDict - содержит информацию полученную из document information dictionary (DID) PDF документа
// Pdfid0 - определяет первый файловый идентификатор найденный в PDF файле
// Pdfid1 - определяет второй файловый идентификатор найденный в PDF файле
type PDFFileExtensionSTIX struct {
	Version          string            `json:"version" bson:"version"`
	IsOptimized      bool              `json:"is_optimized" bson:"is_optimized"`
	DocumentInfoDict map[string]string `json:"document_info_dict" bson:"document_info_dict"`
	Pdfid0           string            `json:"pdfid0" bson:"pdfid0"`
	Pdfid1           string            `json:"pdfid1" bson:"pdfid1"`
}

// RasterImageFileExtensionSTIX тип "raster-image-ext", по терминалогии STIX, определяет специфичные расширения для растровых, графических файлов
// ImageHeight - определяет высоту изображения в файле, в пикселях
// ImageWidth - определяет ширину изображения в файле, в пикселях
// BitsPerPixel - определяет сумму бит для каждого цветового туннеля в файлес изображением, а также общее количество пикселей, используемых для выражения глубины цвета изображения.
// ExifTags - задает набор тегов EXIF, найденных в файле изображения, в виде словаря. Каждая пара ключ/значение в словаре представляет имя/значение одного тега EXIF.
//
//	Соответственно, каждый ключ словаря ДОЛЖЕН быть сохраненной в регистре версией имени тега EXIF, например . Каждое значение словаря ДОЛЖНО быть либо (для типов данных int* EXIF),
//	либо строкой (для всех других типов данных EXIF).
type RasterImageFileExtensionSTIX struct {
	ImageHeight  int                             `json:"image_height" bson:"image_height"`
	ImageWidth   int                             `json:"image_width" bson:"image_width"`
	BitsPerPixel int                             `json:"bits_per_pixel" bson:"bits_per_pixel"`
	ExifTags     somecomplextypesstixco.ExifTags `json:"exif_tags" bson:"exif_tags"`
}

// HTTPRequestExtensionSTIX тип "http-request-ext", по терминалогии STIX, определяет специфичное расширение HTTP-запроса, задает расширение по умолчанию для захвата свойств
//
//	сетевого трафика, специфичных для HTTP-запросов. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть http-request-ext.
//
// RequestMethod - указывает часть HTTP-метода строки HTTP - запроса в виде строчной строки. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// RequestValue - указывает часть значения (обычно путь к ресурсу) строки HTTP-запроса. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// RequestVersion - указывает часть HTTP-версии строки HTTP - запроса в виде строчной строки.
// RequestHeader - задает все поля заголовка HTTP, которые могут быть найдены в запросе клиента HTTP, в виде словаря.
// MessageBodyLength - указывает длину тела HTTP-сообщения, если оно включено, в байтах.
// MessageBodyDataRef - указывает данные, содержащиеся в теле HTTP-сообщения, если они включены. Объект, на который ссылается это свойство, ДОЛЖЕН иметь тип artifact.
type HTTPRequestExtensionSTIX struct {
	RequestMethod      string                         `json:"request_method" bson:"request_method"`
	RequestValue       string                         `json:"request_value" bson:"request_value"`
	RequestVersion     string                         `json:"request_version" bson:"request_version"`
	RequestHeader      map[string]string              `json:"request_header" bson:"request_header"`
	MessageBodyLength  int                            `json:"message_body_length" bson:"message_body_length"`
	MessageBodyDataRef stixhelpers.IdentifierTypeSTIX `json:"message_body_data_ref" bson:"message_body_data_ref"`
}

// ICMPExtensionSTIX тип "icmp-ext", по терминалогии STIX, определяет специфичное расширение по умолчанию для захвата свойств сетевого трафика, специфичных для ICMP. Ключ для этого
//
//	расширения при использовании в словаре расширений ДОЛЖЕН быть icmp-ext.
//
// ICMPTypeHex - указывает тип ICMP байт.
// ICMPCodeHex - задает байт кода ICMP.
type ICMPExtensionSTIX struct {
	ICMPTypeHex string `json:"icmp_type_hex" bson:"icmp_type_hex"`
	ICMPCodeHex string `json:"icmp_code_hex" bson:"icmp_code_hex"`
}

// NetworkSocketExtensionSTIX тип "socket-ext", по терминалогии STIX, определяет специфичное расширение по умолчанию для захвата свойств сетевого трафика, связанных с сетевыми сокетами.
//
//	Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть socket-ext.
//
// AddressFamily -  указывает семейство адресов (AF_*), для которого настроен сокет. Значения этого свойства ДОЛЖНЫ быть получены из network-socket-address-family-enum. (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// IsBlocking - указывает, находится ли сокет в режиме блокировки.
// IsListening - указывает, находится ли сокет в режиме прослушивания.
// Options - указывает любые параметры (например, SO_*), которые могут использоваться сокетом в качестве словаря. Каждый ключ в словаре ДОЛЖЕН быть сохраненной в регистре версией имени
//
//	параметра, например SO_ACCEPTCONN. Каждое значение ключа в словаре ДОЛЖНО быть значением для соответствующего ключа опций. Каждое значение словаря ДОЛЖНО быть целым числом.
//
// SocketType - указывает тип сокета. Значения этого свойства ДОЛЖНЫ быть получены из network-socket-type-enum.
// SocketDescriptor - указывается значение дескриптора файла сокета, связанное с сокетом, как неотрицательное целое число.
// SocketHandle - указывает значение дескриптора или индекса, связанное с сокетом.
type NetworkSocketExtensionSTIX struct {
	AddressFamily    stixhelpers.EnumTypeSTIX `json:"address_family" bson:"address_family"`
	IsBlocking       bool                     `json:"is_blocking" bson:"is_blocking"`
	IsListening      bool                     `json:"is_listening" bson:"is_listening"`
	Options          map[string]int           `json:"options" bson:"options"`
	SocketType       stixhelpers.EnumTypeSTIX `json:"socket_type" bson:"socket_type"`
	SocketDescriptor int                      `json:"socket_descriptor" bson:"socket_descriptor"`
	SocketHandle     int                      `json:"socket_handle" bson:"socket_handle"`
}

// TCPExtensionSTIX тип "tcp-ext", по терминалогии STIX, определяет специфичное расширение задает расширение по умолчанию для захвата свойств сетевого трафика, специфичных для TCP.
//
//	Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть tcp-ext.
//
// SrcFlagsHex - указывает исходные флаги TCP как объединение всех флагов TCP, наблюдаемых между началом трафика (как определено свойством start) и концом трафика (как определено свойством end).
// DstFlagsHex - указывает флаги TCP назначения как объединение всех флагов TCP, наблюдаемых между началом трафика (как определено свойством start) и концом трафика (как определено свойством end).
type TCPExtensionSTIX struct {
	SrcFlagsHex string `json:"src_flags_hex" bson:"src_flags_hex"`
	DstFlagsHex string `json:"dst_flags_hex" bson:"dst_flags_hex"`
}

// WindowsPEBinaryFileExtensionSTIX тип "windows-pebinary-ext", по терминалогии STIX, определяет специфичные свойства для Windows portable executable (PE) файлов
// PeType - определяет тип PE binary, в данном поле СЛЕДУЕТ использовать значение из открытого словаря windows-pebinary-type-ov (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ).
// Imphash - определяет специальный 'импортируемый' хеш или imphash подсчитываемый для основанных на PE Binary импортируемых библиотек и функций.
// MachineHex - определяет специфичный, машинный хеш.
// NumberOfSections - определяет количество секций в PE Binary
// TimeDateStamp - время, в формате "2016-05-12T08:17:27.000Z", когда PE Binary был создан.
// PointerToSymbolTableHex - задает смещение файла таблицы символов COFF.
// NumberOfSymbols - задает номер записи в таблице символов PE binary.
// SizeOfOptionalHeader - задает размер заголовка в PE binary.
// CharacteristicsHex - задает флаги, определяющие характеристики файла.
// FileHeaderHashes - определяет любой, подсчитанный для этого файла, хеш.
// OptionalHeader - задает заголовок PE binary файла. Должно использоватся свойство из windows-pe-optional-header-type.
// Sections - задает список методанных касательно секций PE файла.
type WindowsPEBinaryFileExtensionSTIX struct {
	PeType                  stixhelpers.OpenVocabTypeSTIX                          `json:"pe_type" bson:"pe_type"`
	Imphash                 string                                                 `json:"imphash" bson:"imphash"`
	MachineHex              string                                                 `json:"machine_hex" bson:"machine_hex"`
	NumberOfSections        int                                                    `json:"number_of_sections" bson:"number_of_sections"`
	TimeDateStamp           time.Time                                              `json:"time_date_stamp" bson:"time_date_stamp"`
	PointerToSymbolTableHex string                                                 `json:"pointer_to_symbol_table_hex" bson:"pointer_to_symbol_table_hex"`
	NumberOfSymbols         int                                                    `json:"number_of_symbols" bson:"number_of_symbols"`
	SizeOfOptionalHeader    int                                                    `json:"size_of_optional_header" bson:"size_of_optional_header"`
	CharacteristicsHex      string                                                 `json:"characteristics_hex" bson:"characteristics_hex"`
	FileHeaderHashes        stixhelpers.HashesTypeSTIX                             `json:"file_header_hashes" bson:"file_header_hashes"`
	OptionalHeader          somecomplextypesstixco.WindowsPEOptionalHeaderTypeSTIX `json:"optional_header" bson:"optional_header"`
	Sections                []somecomplextypesstixco.WindowsPESectionTypeSTIX      `json:"sections" bson:"sections"`
}

// WindowsProcessExtensionSTIX тип "windows-process-ext", по терминалогии STIX, содержит рассширения расширение по умолчанию для захвата свойств,
//
//	специфичных для процессов Windows. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-process-ext.
//
// ASLREnabled - указывает, включена ли для процесса рандомизация компоновки адресного пространства (ASLR).
// DEPEnabled - указывает, включено ли для процесса Предотвращение выполнения данных (DEP).
// Priority - задает текущий приоритет процесса в Windows. Это значение ДОЛЖНО быть строкой, которая заканчивается на _CLASS.
// OwnerSID - указывает значение идентификатора безопасности (SID) владельца процесса.
// WindowTitle - задает заголовок главного окна процесса.
// StartupInfo - указывает структуру STARTUP_INFO, используемую процессом в качестве словаря.
// IntegrityLevel - указывает уровень целостности Windows или надежность процесса. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-integrity-level-enum.
type WindowsProcessExtensionSTIX struct {
	ASLREnabled    bool                     `json:"aslr_enabled" bson:"aslr_enabled"`
	DEPEnabled     bool                     `json:"dep_enabled" bson:"dep_enabled"`
	Priority       string                   `json:"priority" bson:"priority"`
	OwnerSID       string                   `json:"owner_sid" bson:"owner_sid"`
	WindowTitle    string                   `json:"window_title" bson:"window_title"`
	StartupInfo    map[string]string        `json:"startup_info" bson:"startup_info"`
	IntegrityLevel stixhelpers.EnumTypeSTIX `json:"integrity_level" bson:"integrity_level"`
}

// WindowsServiceExtensionSTIX тип "windows-service-ext", по терминалогии STIX, содержит рассширения службы Windows задает расширение по умолчанию для
//
//	захвата свойств, специфичных для служб Windows. Ключ для этого расширения при использовании в словаре расширений ДОЛЖЕН быть windows-service-ext.
//
// ServiceName - определяет название сервиса.
// Descriptions - определяет детальное описание сервиса.
// DisplayName - указывает отображаемое имя службы в элементах управления графическим интерфейсом Windows.
// GroupName - указывает имя группы заказа загрузки, членом которой является служба.
// StartType - указывает параметры запуска, определенные для службы. Значения этого свойства ДОЛЖНЫ быть получены из перечисления windows-service-start-type-enum.
// ServiceDllRefs - указывает библиотеки DLL, загруженные службой, в качестве ссылки на один или несколько файловых объектов. Объекты, на которые ссылается
//
//	это свойство, ДОЛЖНЫ иметь тип file.
//
// ServiceType - определяет тип сервиса. Значение данного поля ДОЛЖНО иметь тип windows-service-type-enum.
// ServiceStatus - определяет статус сервиса. Значение данного поля ДОЛЖНО иметь тип windows-service-status-enum.
type WindowsServiceExtensionSTIX struct {
	ServiceName    string                           `json:"service_name" bson:"service_name"`
	Descriptions   []string                         `json:"descriptions" bson:"descriptions"`
	DisplayName    string                           `json:"display_name" bson:"display_name"`
	GroupName      string                           `json:"group_name" bson:"group_name"`
	StartType      stixhelpers.EnumTypeSTIX         `json:"start_type" bson:"start_type"`
	ServiceDllRefs []stixhelpers.IdentifierTypeSTIX `json:"service_dll_refs" bson:"service_dll_refs"`
	ServiceType    stixhelpers.EnumTypeSTIX         `json:"service_type" bson:"service_type"`
	ServiceStatus  stixhelpers.EnumTypeSTIX         `json:"service_status" bson:"service_status"`
}

// UNIXAccountExtensionSTIX тип "unix-account-ext", по терминалогии STIX, содержит рассширения 'по умолчанию' захваченной дополнительной информации
// предназначенной для аккаунтов UNIX систем.
// GID - содержит первичный групповой ID аккаунта
// Groups - содержит список имен групп которые являются членами аккаунта
// HomeDir - содержит домашную директорию аккаунта
// Shell - содержит командную оболочку аккаунта
type UNIXAccountExtensionSTIX struct {
	GID     int      `json:"gid" bson:"gid"`
	Groups  []string `json:"group" bson:"group"`
	HomeDir string   `json:"home_dir" bson:"home_dir"`
	Shell   string   `json:"shell" bson:"shell"`
}
