package methodstixobjects

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/relationshipobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// NewOptionalCommonPropertiesRelationshipObjectSTIX создает объект содержащий общие, опциональные
// свойства для все объектов STIX типа Relationship Objects
func NewOptionalCommonPropertiesRelationshipObjectSTIX() *relationshipobjectsstix.OptionalCommonPropertiesRelationshipObjectSTIX {
	return &relationshipobjectsstix.OptionalCommonPropertiesRelationshipObjectSTIX{
		Created:  "1970-01-01T00:00:00+00:00",
		Modified: "1970-01-01T00:00:00+00:00",
	}
}

// NewRelationshipObjectSTIX создает объект "Relationship", по терминалогии STIX, используется
// для связывания двух Domain Object STIX (SDO) или Cyber-observable Objects STIX (SCO), чтобы
// описать, как они связаны друг с другом. Если SDO и SCO считаются "узлами" или
// "вершинами" в графе, то Объекты отношений (SRO) представляют собой "ребра".
func NewRelationshipObjectSTIX() *relationshipobjectsstix.RelationshipObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("relationship")

	return &relationshipobjectsstix.RelationshipObjectSTIX{
		CommonPropertiesObjectSTIX:                     *cpo.Get(),
		OptionalCommonPropertiesRelationshipObjectSTIX: *NewOptionalCommonPropertiesRelationshipObjectSTIX(),
		StartTime: "1970-01-01T00:00:00+00:00",
		StopTime:  "1970-01-01T00:00:00+00:00",
	}
}

// SightingObjectSTIX создает объект "Sighting", по терминалогии STIX, это особый тип SRO.
// Отношение, которое содержит дополнительные свойства, отсутствующие в объекте Relationship.
func NewSightingObjectSTIX() *relationshipobjectsstix.SightingObjectSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("sighting")

	return &relationshipobjectsstix.SightingObjectSTIX{
		CommonPropertiesObjectSTIX:                     *cpo.Get(),
		OptionalCommonPropertiesRelationshipObjectSTIX: *NewOptionalCommonPropertiesRelationshipObjectSTIX(),
		FirstSeen:        "1970-01-01T00:00:00+00:00",
		LastSeen:         "1970-01-01T00:00:00+00:00",
		ObservedDataRefs: []stixhelpers.IdentifierTypeSTIX(nil),
		WhereSightedRefs: []stixhelpers.IdentifierTypeSTIX(nil),
	}
}
