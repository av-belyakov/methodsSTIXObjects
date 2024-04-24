package datamodels

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// decodingExtensionsSTIX декодирует следующие типы STIX расширений:
// - "archive-ext"
// - "ntfs-ext"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "windows-process-ext"
// - "windows-service-ext"
// - "unix-account-ext"
func decodingExtensionsSTIX(extType string, rawMsg *json.RawMessage) (interface{}, error) {
	var err error

	//fmt.Printf("func 'decodingExtensionsSTIX', extType = %s, \n", extType)

	switch extType {
	case "archive-ext":
		var archiveExt someextensionsstixco.ArchiveFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &archiveExt)

		return archiveExt, err
	case "ntfs-ext":
		var ntfsExt someextensionsstixco.NTFSFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &ntfsExt)

		return ntfsExt, err
	case "pdf-ext":
		var pdfExt someextensionsstixco.PDFFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &pdfExt)

		return pdfExt, err
	case "raster-image-ext":
		var rasterImageExt someextensionsstixco.RasterImageFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &rasterImageExt)

		//fmt.Printf("func 'decodingExtensionsSTIX', rasterImageExt = %v, \n", rasterImageExt)

		return rasterImageExt, err
	case "windows-pebinary-ext":
		var windowsPebinaryExt someextensionsstixco.WindowsPEBinaryFileExtensionSTIX
		err = json.Unmarshal(*rawMsg, &windowsPebinaryExt)

		return windowsPebinaryExt, err
	case "http-request-ext":
		var httpRequestExt someextensionsstixco.HTTPRequestExtensionSTIX
		err = json.Unmarshal(*rawMsg, &httpRequestExt)

		return httpRequestExt, err
	case "icmp-ext":
		var icmpExt someextensionsstixco.ICMPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &icmpExt)

		return icmpExt, err
	case "socket-ext":
		var socketExt someextensionsstixco.NetworkSocketExtensionSTIX
		err := json.Unmarshal(*rawMsg, &socketExt)

		return socketExt, err
	case "tcp-ext":
		var tcpExt someextensionsstixco.TCPExtensionSTIX
		err := json.Unmarshal(*rawMsg, &tcpExt)

		return tcpExt, err
	case "windows-process-ext":
		var windowsProcessExt someextensionsstixco.WindowsProcessExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsProcessExt)

		return windowsProcessExt, err
	case "windows-service-ext":
		var windowsServiceExt someextensionsstixco.WindowsServiceExtensionSTIX
		err := json.Unmarshal(*rawMsg, &windowsServiceExt)

		return windowsServiceExt, err
	case "unix-account-ext":
		var unixAccountExt someextensionsstixco.UNIXAccountExtensionSTIX
		err := json.Unmarshal(*rawMsg, &unixAccountExt)

		return unixAccountExt, err
	default:
		return struct{}{}, nil
	}
}

// checkingExtensionsSTIX выполняет проверку полей следующих типов STIX расширений:
// - "archive-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "windows-service-ext"
func checkingExtensionsSTIX(extType interface{}) bool {
	switch et := extType.(type) {
	case someextensionsstixco.ArchiveFileExtensionSTIX:
		for _, v := range et.ContainsRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}

	case someextensionsstixco.NTFSFileExtensionSTIX:
		if len(et.AlternateDataStreams) == 0 {
			return true
		}

		for _, v := range et.AlternateDataStreams {
			if !v.Hashes.CheckHashesTypeSTIX() {
				return false
			}
		}

	case someextensionsstixco.WindowsPEBinaryFileExtensionSTIX:
		if !et.FileHeaderHashes.CheckHashesTypeSTIX() {
			return false
		}

	case someextensionsstixco.HTTPRequestExtensionSTIX:
		if !et.MessageBodyDataRef.CheckIdentifierTypeSTIX() {
			return false
		}

	case someextensionsstixco.WindowsServiceExtensionSTIX:
		for _, v := range et.ServiceDllRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// sanitizeExtensionsSTIX для ряда полей следующих типов STIX расширений:
// - "archive-ext"
// - "ntfs-ext"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "windows-process-ext"
// - "windows-service-ext"
// - "unix-account-ext"
// выполняет замену некоторых специальных символов на их HTML код
func sanitizeExtensionsSTIX(extType interface{}) interface{} {
	sanitizeList := func(rh map[string]string) map[string]string {
		mapTmp := make(map[string]string, len(rh))
		for k, v := range rh {
			mapTmp[k] = commonlibs.StringSanitize(v)
		}

		return mapTmp
	}

	sanitizeExifTags := func(et somecomplextypesstixco.ExifTags) somecomplextypesstixco.ExifTags {
		et.Make = commonlibs.StringSanitize(et.Make)
		et.Model = commonlibs.StringSanitize(et.Model)

		return et
	}

	switch et := extType.(type) {
	case someextensionsstixco.ArchiveFileExtensionSTIX:
		return commonlibs.StringSanitize(et.Comment)

	case someextensionsstixco.NTFSFileExtensionSTIX:
		var tmpType someextensionsstixco.NTFSFileExtensionSTIX
		tmpType.SID = commonlibs.StringSanitize(et.SID)

		size := len(et.AlternateDataStreams)

		if size == 0 {
			return tmpType
		}

		ads := make([]somecomplextypesstixco.AlternateDataStreamTypeSTIX, 0, size)
		for _, v := range et.AlternateDataStreams {
			ads = append(ads, somecomplextypesstixco.AlternateDataStreamTypeSTIX{
				Name:   commonlibs.StringSanitize(v.Name),
				Hashes: v.Hashes,
				Size:   v.Size,
			})
		}
		tmpType.AlternateDataStreams = ads

		return tmpType

	case someextensionsstixco.PDFFileExtensionSTIX:
		return someextensionsstixco.PDFFileExtensionSTIX{
			Version:          commonlibs.StringSanitize(et.Version),
			IsOptimized:      et.IsOptimized,
			DocumentInfoDict: sanitizeList(et.DocumentInfoDict),
			Pdfid0:           commonlibs.StringSanitize(et.Pdfid0),
			Pdfid1:           commonlibs.StringSanitize(et.Pdfid1),
		}

	case someextensionsstixco.RasterImageFileExtensionSTIX:
		return someextensionsstixco.RasterImageFileExtensionSTIX{
			ImageHeight:  et.ImageHeight,
			ImageWidth:   et.ImageWidth,
			BitsPerPixel: et.BitsPerPixel,
			ExifTags:     sanitizeExifTags(et.ExifTags),
		}

	case someextensionsstixco.WindowsPEBinaryFileExtensionSTIX:
		ssize := len(et.Sections)
		ns := make([]somecomplextypesstixco.WindowsPESectionTypeSTIX, 0, ssize)
		if ssize > 0 {
			for _, v := range et.Sections {
				tmp := v.SanitizeStructWindowsPESectionTypeSTIX()
				ns = append(ns, tmp)
			}
			et.Sections = ns
		}

		return someextensionsstixco.WindowsPEBinaryFileExtensionSTIX{
			PeType:                  et.PeType.SanitizeStructOpenVocabTypeSTIX(),
			Imphash:                 commonlibs.StringSanitize(et.Imphash),
			MachineHex:              commonlibs.StringSanitize(et.MachineHex),
			NumberOfSections:        et.NumberOfSections,
			TimeDateStamp:           et.TimeDateStamp,
			PointerToSymbolTableHex: commonlibs.StringSanitize(et.PointerToSymbolTableHex),
			NumberOfSymbols:         et.NumberOfSymbols,
			SizeOfOptionalHeader:    et.SizeOfOptionalHeader,
			CharacteristicsHex:      commonlibs.StringSanitize(et.CharacteristicsHex),
			FileHeaderHashes: func(l stixhelpers.HashesTypeSTIX) stixhelpers.HashesTypeSTIX {
				fhh := make(stixhelpers.HashesTypeSTIX, len(l))

				for k, v := range l {
					fhh[k] = commonlibs.StringSanitize(string(v))
				}

				return fhh
			}(et.FileHeaderHashes),
			OptionalHeader: et.OptionalHeader.SanitizeStructWindowsPEOptionalHeaderTypeSTIX(),
			Sections:       ns,
		}

	case someextensionsstixco.HTTPRequestExtensionSTIX:
		return someextensionsstixco.HTTPRequestExtensionSTIX{
			RequestMethod:      commonlibs.StringSanitize(et.RequestMethod),
			RequestValue:       commonlibs.StringSanitize(et.RequestValue),
			RequestVersion:     commonlibs.StringSanitize(et.RequestVersion),
			RequestHeader:      sanitizeList(et.RequestHeader),
			MessageBodyLength:  et.MessageBodyLength,
			MessageBodyDataRef: et.MessageBodyDataRef,
		}

	case someextensionsstixco.ICMPExtensionSTIX:
		return someextensionsstixco.ICMPExtensionSTIX{
			ICMPTypeHex: commonlibs.StringSanitize(et.ICMPTypeHex),
			ICMPCodeHex: commonlibs.StringSanitize(et.ICMPCodeHex),
		}

	case someextensionsstixco.NetworkSocketExtensionSTIX:
		return someextensionsstixco.NetworkSocketExtensionSTIX{
			AddressFamily:    stixhelpers.EnumTypeSTIX((commonlibs.StringSanitize(string(et.AddressFamily)))),
			IsBlocking:       et.IsBlocking,
			IsListening:      et.IsListening,
			Options:          et.Options,
			SocketType:       stixhelpers.EnumTypeSTIX((commonlibs.StringSanitize(string(et.SocketType)))),
			SocketDescriptor: et.SocketDescriptor,
			SocketHandle:     et.SocketHandle,
		}

	case someextensionsstixco.TCPExtensionSTIX:
		return someextensionsstixco.TCPExtensionSTIX{
			SrcFlagsHex: commonlibs.StringSanitize(et.SrcFlagsHex),
			DstFlagsHex: commonlibs.StringSanitize(et.DstFlagsHex),
		}

	case someextensionsstixco.WindowsProcessExtensionSTIX:
		return someextensionsstixco.WindowsProcessExtensionSTIX{
			ASLREnabled:    et.ASLREnabled,
			DEPEnabled:     et.DEPEnabled,
			Priority:       commonlibs.StringSanitize(et.Priority),
			OwnerSID:       commonlibs.StringSanitize(et.OwnerSID),
			WindowTitle:    commonlibs.StringSanitize(et.WindowTitle),
			StartupInfo:    sanitizeList(et.StartupInfo),
			IntegrityLevel: stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(et.IntegrityLevel))),
		}

	case someextensionsstixco.WindowsServiceExtensionSTIX:
		return someextensionsstixco.WindowsServiceExtensionSTIX{
			ServiceName: commonlibs.StringSanitize(et.ServiceName),
			Descriptions: func(l []string) []string {
				size := len(l)
				nl := make([]string, 0, size)
				for _, v := range l {
					nl = append(nl, commonlibs.StringSanitize(v))
				}

				return nl
			}(et.Descriptions),
			DisplayName:    commonlibs.StringSanitize(et.DisplayName),
			GroupName:      commonlibs.StringSanitize(et.GroupName),
			StartType:      stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(et.StartType))),
			ServiceDllRefs: et.ServiceDllRefs,
			ServiceType:    stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(et.ServiceType))),
			ServiceStatus:  stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(et.ServiceStatus))),
		}

	case someextensionsstixco.UNIXAccountExtensionSTIX:
		return someextensionsstixco.UNIXAccountExtensionSTIX{
			GID: et.GID,
			Groups: func(l []string) []string {
				size := len(l)
				nl := make([]string, 0, size)
				for _, v := range l {
					nl = append(nl, commonlibs.StringSanitize(v))
				}

				return nl
			}(et.Groups),
			HomeDir: commonlibs.StringSanitize(et.HomeDir),
			Shell:   commonlibs.StringSanitize(et.Shell),
		}
	}

	return extType
}

// toStringBeautiful выполняет красивое представление информации содержащейся в следующих типах
// - "archive-ext"
// - "ntfs-ext"
// - "pdf-ext"
// - "raster-image-ext"
// - "windows-pebinary-ext"
// - "http-request-ext"
// - "icmp-ext"
// - "socket-ext"
// - "tcp-ext"
// - "windows-process-ext"
// - "windows-service-ext"
// - "unix-account-ext"
func toStringBeautiful(extType interface{}) string {
	str := strings.Builder{}
	str.WriteString("\t\t")

	switch et := extType.(type) {
	case someextensionsstixco.ArchiveFileExtensionSTIX:
		str.WriteString(fmt.Sprintln("'contains_refs':"))
		for k, v := range et.ContainsRefs {
			str.WriteString(fmt.Sprintf("\t\t\t'contains_ref '%d'': '%v'", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'comment': '%v'\n", et.Comment))

	case someextensionsstixco.NTFSFileExtensionSTIX:
		str.WriteString(fmt.Sprintf("'sid': '%s'\n", et.SID))
		str.WriteString(fmt.Sprintln("\t\t'alternate_data_streams':"))
		for k, v := range et.AlternateDataStreams {
			str.WriteString(fmt.Sprintf("\t\t\t'alternate_data_stream '%d'':\n", k))
			str.WriteString(fmt.Sprintf("\t\t\t\t'name': '%s'\n", v.Name))
			str.WriteString(fmt.Sprintln("\t\t\t\t'hashes':"))
			for a, b := range v.Hashes {
				str.WriteString(fmt.Sprintf("\t\t\t\t\t'%s': '%v'\n", a, b))
			}
			str.WriteString(fmt.Sprintf("\t\t\t\t'size': '%d'\n", v.Size))
		}

	case someextensionsstixco.PDFFileExtensionSTIX:
		str.WriteString(fmt.Sprintf("'version': '%s'\n", et.Version))
		str.WriteString(fmt.Sprintf("\t\t'is_optimized': '%v'\n", et.IsOptimized))
		str.WriteString(fmt.Sprintln("\t\t'document_info_dict':"))
		for k, v := range et.DocumentInfoDict {
			str.WriteString(fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'pdfid0': '%s'\n", et.Pdfid0))
		str.WriteString(fmt.Sprintf("\t\t'pdfid1': '%s'\n", et.Pdfid1))

	case someextensionsstixco.RasterImageFileExtensionSTIX:
		str.WriteString(fmt.Sprintf("'image_height': '%d'\n", et.ImageHeight))
		str.WriteString(fmt.Sprintf("\t\t'image_width': '%d'\n", et.ImageWidth))
		str.WriteString(fmt.Sprintf("\t\t'bits_per_pixel': '%d'\n", et.BitsPerPixel))
		str.WriteString(fmt.Sprintln("\t\t'exif_tags':"))
		str.WriteString(fmt.Sprintf("\t\t\t'make': '%s'\n", et.ExifTags.Make))
		str.WriteString(fmt.Sprintf("\t\t\t'model': '%s'\n", et.ExifTags.Model))
		str.WriteString(fmt.Sprintf("\t\t\t'xResolution': '%d'\n", et.ExifTags.XResolution))
		str.WriteString(fmt.Sprintf("\t\t\t'yResolution': '%d'\n", et.ExifTags.YResolution))

	case someextensionsstixco.WindowsPEBinaryFileExtensionSTIX:
		str.WriteString(fmt.Sprintf("'pe_type': '%v'\n", et.PeType))
		str.WriteString(fmt.Sprintf("\t\t'imphash': '%s'\n", et.Imphash))
		str.WriteString(fmt.Sprintf("\t\t'machine_hex': '%s'\n", et.MachineHex))
		str.WriteString(fmt.Sprintf("\t\t'number_of_sections': '%d'\n", et.NumberOfSections))
		str.WriteString(fmt.Sprintf("\t\t'time_date_stamp': '%v'\n", et.TimeDateStamp))
		str.WriteString(fmt.Sprintf("\t\t'pointer_to_symbol_table_hex': '%s'\n", et.PointerToSymbolTableHex))
		str.WriteString(fmt.Sprintf("\t\t'number_of_symbols': '%d'\n", et.NumberOfSymbols))
		str.WriteString(fmt.Sprintf("\t\t'size_of_optional_header': '%d'\n", et.SizeOfOptionalHeader))
		str.WriteString(fmt.Sprintf("\t\t'characteristics_hex': '%s'\n", et.CharacteristicsHex))
		str.WriteString(fmt.Sprintln("\t\t'file_header_hashes':"))
		for k, v := range et.FileHeaderHashes {
			str.WriteString(fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintln("\t\t'optional_header':"))
		str.WriteString(fmt.Sprintf("\t\t\t'magic_hex': '%s'\n", et.OptionalHeader.MagicHex))
		str.WriteString(fmt.Sprintf("\t\t\t'major_linker_version': '%d'\n", et.OptionalHeader.MajorLinkerVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'minor_linker_version': '%d'\n", et.OptionalHeader.MinorLinkerVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_code': '%d'\n", et.OptionalHeader.SizeOfCode))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_initialized_data': '%d'\n", et.OptionalHeader.SizeOfInitializedData))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_uninitialized_data': '%d'\n", et.OptionalHeader.SizeOfUninitializedData))
		str.WriteString(fmt.Sprintf("\t\t\t'address_of_entry_point': '%d'\n", et.OptionalHeader.AddressOfEntryPoint))
		str.WriteString(fmt.Sprintf("\t\t\t'base_of_code': '%d'\n", et.OptionalHeader.BaseOfCode))
		str.WriteString(fmt.Sprintf("\t\t\t'image_base': '%d'\n", et.OptionalHeader.ImageBase))
		str.WriteString(fmt.Sprintf("\t\t\t'section_alignment': '%d'\n", et.OptionalHeader.SectionAlignment))
		str.WriteString(fmt.Sprintf("\t\t\t'file_alignment': '%d'\n", et.OptionalHeader.FileAlignment))
		str.WriteString(fmt.Sprintf("\t\t\t'major_os_version': '%d'\n", et.OptionalHeader.MajorOSVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'minor_os_version': '%d'\n", et.OptionalHeader.MinorOSVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'major_image_version': '%d'\n", et.OptionalHeader.MajorImageVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'minor_image_version': '%d'\n", et.OptionalHeader.MinorImageVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'major_subsystem_version': '%d'\n", et.OptionalHeader.MajorSubsystemVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'minor_subsystem_version': '%d'\n", et.OptionalHeader.MinorSubsystemVersion))
		str.WriteString(fmt.Sprintf("\t\t\t'win32_version_value_hex': '%s'\n", et.OptionalHeader.Win32VersionValueHex))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_image': '%d'\n", et.OptionalHeader.SizeOfImage))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_headers': '%d'\n", et.OptionalHeader.SizeOfHeaders))
		str.WriteString(fmt.Sprintf("\t\t\t'checksum_hex': '%s'\n", et.OptionalHeader.ChecksumHex))
		str.WriteString(fmt.Sprintf("\t\t\t'subsystem_hex': '%s'\n", et.OptionalHeader.SubsystemHex))
		str.WriteString(fmt.Sprintf("\t\t\t'dll_characteristics_hex': '%s'\n", et.OptionalHeader.DllCharacteristicsHex))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_stack_reserve': '%d'\n", et.OptionalHeader.SizeOfStackReserve))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_stack_commit': '%d'\n", et.OptionalHeader.SizeOfStackCommit))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_heap_reserve': '%d'\n", et.OptionalHeader.SizeOfHeapReserve))
		str.WriteString(fmt.Sprintf("\t\t\t'size_of_heap_commit': '%d'\n", et.OptionalHeader.SizeOfHeapCommit))
		str.WriteString(fmt.Sprintf("\t\t\t'loader_flags_hex': '%s'\n", et.OptionalHeader.LoaderFlagsHex))
		str.WriteString(fmt.Sprintf("\t\t\t'number_of_rva_and_aizes': '%d'\n", et.OptionalHeader.NumberOfRvaAndSizes))
		str.WriteString(fmt.Sprintln("\t\t\t'hashes':"))
		for k, v := range et.OptionalHeader.Hashes {
			str.WriteString(fmt.Sprintf("\t\t\t\t'hashe '%d'': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintln("\t\t'sections':"))
		for k, v := range et.Sections {
			str.WriteString(fmt.Sprintf("\t\t\t'section '%d'':\n", k))
			str.WriteString(fmt.Sprintf("\t\t\t\t'name': '%s'\n", v.Name))
			str.WriteString(fmt.Sprintf("\t\t\t\t'size': '%v'\n", v.Size))
			str.WriteString(fmt.Sprintf("\t\t\t\t'entropy': '%v'\n", v.Entropy))
			str.WriteString(fmt.Sprintln("\t\t\t\t'hashes':", v.Hashes))
			for k, v := range v.Hashes {
				str.WriteString(fmt.Sprintf("\t\t\t\t\t'%s': '%s'\n", k, v))
			}
		}

	case someextensionsstixco.HTTPRequestExtensionSTIX:
		str.WriteString(fmt.Sprintf("'request_method': '%s'\n", et.RequestMethod))
		str.WriteString(fmt.Sprintf("\t\t'request_value': '%s'\n", et.RequestValue))
		str.WriteString(fmt.Sprintf("\t\t'request_version': '%s'\n", et.RequestVersion))
		str.WriteString(fmt.Sprintln("\t\t'request_header':"))
		for k, v := range et.RequestHeader {
			str.WriteString(fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'message_body_length': '%d'\n", et.MessageBodyLength))
		str.WriteString(fmt.Sprintf("\t\t'message_body_data_ref': '%v'\n", et.MessageBodyDataRef))

	case someextensionsstixco.ICMPExtensionSTIX:
		str.WriteString(fmt.Sprintf("'icmp_type_hex': '%s'\n", et.ICMPTypeHex))
		str.WriteString(fmt.Sprintf("\t\t'icmp_code_hex': '%s'\n", et.ICMPCodeHex))

	case someextensionsstixco.NetworkSocketExtensionSTIX:
		str.WriteString(fmt.Sprintf("'address_family': '%v'\n", et.AddressFamily))
		str.WriteString(fmt.Sprintf("\t\t'is_blocking': '%v'\n", et.IsBlocking))
		str.WriteString(fmt.Sprintf("\t\t'is_listening': '%v'\n", et.IsListening))
		str.WriteString(fmt.Sprintln("\t\t'options':"))
		for k, v := range et.Options {
			str.WriteString(fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'socket_type': '%v'\n", et.SocketType))
		str.WriteString(fmt.Sprintf("\t\t'socket_descriptor': '%d'\n", et.SocketDescriptor))
		str.WriteString(fmt.Sprintf("\t\t'socket_handle': '%d'\n", et.SocketHandle))

	case someextensionsstixco.TCPExtensionSTIX:
		str.WriteString(fmt.Sprintf("'src_flags_hex': '%v'\n", et.SrcFlagsHex))
		str.WriteString(fmt.Sprintf("\t\t'dst_flags_hex': '%v'\n", et.DstFlagsHex))

	case someextensionsstixco.WindowsProcessExtensionSTIX:
		str.WriteString(fmt.Sprintf("'aslr_enabled': '%v'\n", et.ASLREnabled))
		str.WriteString(fmt.Sprintf("\t\t'dep_enabled': '%v'\n", et.DEPEnabled))
		str.WriteString(fmt.Sprintf("\t\t'priority': '%s'\n", et.Priority))
		str.WriteString(fmt.Sprintf("\t\t'owner_sid': '%s'\n", et.OwnerSID))
		str.WriteString(fmt.Sprintf("\t\t'window_title': '%s'\n", et.WindowTitle))
		str.WriteString(fmt.Sprintln("\t\t'startup_info':"))
		for k, v := range et.StartupInfo {
			str.WriteString(fmt.Sprintf("\t\t\t'%s': '%v'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'integrity_level': '%v'\n", et.IntegrityLevel))

	case someextensionsstixco.WindowsServiceExtensionSTIX:
		str.WriteString(fmt.Sprintf("'service_name': '%s'\n", et.ServiceName))
		str.WriteString(fmt.Sprintln("\t\t'descriptions':"))
		for k, v := range et.Descriptions {
			str.WriteString(fmt.Sprintf("\t\t\t'description '%d'': '%s'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'display_name': '%s'\n", et.DisplayName))
		str.WriteString(fmt.Sprintf("\t\t'group_name': '%s'\n", et.GroupName))
		str.WriteString(fmt.Sprintf("\t\t'start_type': '%v'\n", et.StartType))
		str.WriteString(fmt.Sprintln("\t\t'service_dll_refs':"))
		for k, v := range et.ServiceDllRefs {
			str.WriteString(fmt.Sprintf("\t\t\t'service_dll_ref '%d'': '%v'", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'service_type': '%v'\n", et.ServiceType))
		str.WriteString(fmt.Sprintf("\t\t'service_status': '%v'\n", et.ServiceStatus))

	case someextensionsstixco.UNIXAccountExtensionSTIX:
		str.WriteString(fmt.Sprintf("'gid': '%d'\n", et.GID))
		str.WriteString(fmt.Sprintln("\t\t'groups':"))
		for k, v := range et.Groups {
			str.WriteString(fmt.Sprintf("\t\t\t'group '%d'': '%s'\n", k, v))
		}
		str.WriteString(fmt.Sprintf("\t\t'home_dir': '%s'\n", et.HomeDir))
		str.WriteString(fmt.Sprintf("\t\t'shell': '%s'\n", et.Shell))

	default:
		str.WriteString(fmt.Sprintf("'%v'\n", et))

	}

	return str.String()
}
