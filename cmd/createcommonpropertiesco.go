package methodstixobjects

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/commonpropertiesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func NewOptionalCommonPropertiesCyberObservableObjectSTIX() *commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX {
	return &commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX{
		ObjectMarkingRefs: make([]stixhelpers.IdentifierTypeSTIX, 0),
		GranularMarkings:  make([]stixhelpers.GranularMarkingsTypeSTIX, 0),
	}
}
