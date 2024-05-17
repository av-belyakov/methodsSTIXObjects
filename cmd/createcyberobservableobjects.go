package cmd

import "github.com/av-belyakov/methodstixobjects/datamodels/cyberobservableobjectsstix"

// NewArtifactCyberObservableObjectSTIX создает STIX объект "Artifact", по терминалогии STIX, позволяет захватывать массив байтов (8 бит) в виде строки
// в кодировке base64 или связывать его с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей PayloadBin или URL
func NewArtifactCyberObservableObjectSTIX() *cyberobservableobjectsstix.ArtifactCyberObservableObjectSTIX {
	return &cyberobservableobjectsstix.ArtifactCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        *NewCommonPropertiesObjectSTIX(),
		OptionalCommonPropertiesCyberObservableObjectSTIX: *NewOptionalCommonPropertiesCyberObservableObjectSTIX(),
	}
}
