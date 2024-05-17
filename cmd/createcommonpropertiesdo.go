package cmd

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/commonpropertiesstixdo"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

func NewCommonPropertiesDomainObjectSTIX() *commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX {
	return &commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX{
		Created:            "1970-01-01T00:00:00+00:00",
		Modified:           "1970-01-01T00:00:00+00:00",
		Labels:             make([]string, 0),
		Extensions:         map[string]string{},
		ExternalReferences: make([]stixhelpers.ExternalReferenceTypeElementSTIX, 0),
		ObjectMarkingRefs:  make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}
