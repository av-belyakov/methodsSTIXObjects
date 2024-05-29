package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNetworkTrafficCyberObservableObjectSTIX(t *testing.T) {
	nnt := methodstixobjects.NewNetworkTrafficCyberObservableObjectSTIX()

	assert.Equal(t, nnt.GetType(), "network-traffic")
	_, err := nnt.Get()
	assert.NoError(t, err)

	nnt.SetAnyIsActive(true)
	assert.True(t, nnt.GetIsActive())
	nnt.SetValueIsActive(false)
	assert.False(t, nnt.GetIsActive())

	nnt.SetAnySrcPort(23564)
	assert.Equal(t, nnt.GetSrcPort(), 23564)
	nnt.SetValueSrcPort(125)
	assert.Equal(t, nnt.GetSrcPort(), 125)

	nnt.SetAnyDstPort(8080)
	assert.Equal(t, nnt.GetDstPort(), 8080)
	nnt.SetValueDstPort(8181)
	assert.Equal(t, nnt.GetDstPort(), 8181)

	nnt.SetAnySrcPackets(8913)
	assert.Equal(t, nnt.GetSrcPackets(), 8913)
	nnt.SetValueSrcPackets(1813)
	assert.Equal(t, nnt.GetSrcPackets(), 1813)

	nnt.SetAnyDstPackets(1123)
	assert.Equal(t, nnt.GetDstPackets(), 1123)
	nnt.SetValueDstPackets(1224)
	assert.Equal(t, nnt.GetDstPackets(), 1224)

	nnt.SetAnySrcByteCount(78123331)
	assert.Equal(t, nnt.GetSrcByteCount(), uint64(78123331))
	nnt.SetValueSrcByteCount(14533331)
	assert.Equal(t, nnt.GetSrcByteCount(), uint64(14533331))

	nnt.SetAnyDstByteCount(16954324447)
	assert.Equal(t, nnt.GetDstByteCount(), uint64(16954324447))
	nnt.SetValueDstByteCount(623447)
	assert.Equal(t, nnt.GetDstByteCount(), uint64(623447))

	// --- Start
	mt := "2024-02-14T12:03:06+00:00"
	err = nnt.SetAnyStart(mt)
	assert.NoError(t, err)
	assert.Equal(t, nnt.GetStart(), mt)

	err = nnt.SetAnyStart(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nnt.GetStart(), "2024-05-28T12:33:44+03:00")

	//--- End
	at := "2024-02-18T21:11:11+00:00"
	err = nnt.SetAnyEnd(at)
	assert.NoError(t, err)
	assert.Equal(t, nnt.GetEnd(), at)

	err = nnt.SetAnyEnd(1716888850190)
	assert.NoError(t, err)
	assert.Equal(t, nnt.GetEnd(), "2024-05-28T12:34:10+03:00")

	nnt.SetAnyProtocols("tcp")
	nnt.SetAnyProtocols("udp")
	nnt.SetValueProtocols("http")
	assert.Equal(t, len(nnt.GetProtocols()), 3)

	sre := stixhelpers.IdentifierTypeSTIX("srcref example")
	nnt.SetValueSrcRef(sre)
	assert.Equal(t, nnt.GetSrcRef(), sre)

	dre := stixhelpers.IdentifierTypeSTIX("dstref example")
	nnt.SetValueSrcRef(dre)
	assert.Equal(t, nnt.GetSrcRef(), dre)

	spre := stixhelpers.IdentifierTypeSTIX("src payload ref example")
	nnt.SetValueSrcPayloadRef(spre)
	assert.Equal(t, nnt.GetSrcPayloadRef(), spre)

	dpre := stixhelpers.IdentifierTypeSTIX("dst payload ref example")
	nnt.SetValueDstPayloadRef(dpre)
	assert.Equal(t, nnt.GetDstPayloadRef(), dpre)

	ere := stixhelpers.IdentifierTypeSTIX("Encapsulated by ref example")
	nnt.SetValueEncapsulatedByRef(ere)
	assert.Equal(t, nnt.GetEncapsulatedByRef(), ere)

	nnt.SetFullValueEncapsulatesRefs([]stixhelpers.IdentifierTypeSTIX{
		"encapsulate_ref_1",
		"encapsulate_ref_2",
		"encapsulate_ref_3",
	})
	assert.Equal(t, len(nnt.GetEncapsulatesRefs()), 3)
	nnt.SetValueEncapsulatesRefs(stixhelpers.IdentifierTypeSTIX("encapsulate_ref_4"))
	assert.Equal(t, len(nnt.GetEncapsulatesRefs()), 4)

	nnt.SetAnyIPFix("ipfix_1", "addv_vdjjb")
	nnt.SetAnyIPFix("ipfix_2", "hfhd_dvnie")
	nnt.SetAnyIPFix("ipfix_3", "nvdv_jief9")
	nnt.SetValueIPFix("ipfix_4", "teyu_duueu")
	assert.Equal(t, len(nnt.GetIPFix()), 4)

	nnt.SetFullValueExtensions(map[string]interface{}{"ext_0": 6293})
	nnt.SetValueExtensions("ext_1", 123)
	nnt.SetValueExtensions("ext_2", 234)
	nnt.SetValueExtensions("ext_3", 345)
	assert.Equal(t, len(nnt.GetExtensions()), 4)
}

/*
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
*/
