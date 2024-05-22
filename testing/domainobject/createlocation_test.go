package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestLocationDomainObjectsSTIX(t *testing.T) {
	nl := methodstixobjects.NewLocationDomainObjectsSTIX()

	assert.Equal(t, nl.GetType(), "location")
	_, err := nl.Get()
	assert.NoError(t, err)

	nl.SetAnyName("location name")
	_, err = nl.Get()
	assert.NoError(t, err)

	lat := float32(23.455)
	nl.SetValueLatitude(lat)
	assert.Equal(t, nl.GetLatitude(), lat)

	lon := float32(51.1025)
	nl.SetValueLongitude(lon)
	assert.Equal(t, nl.GetLongitude(), lon)

	pre := float32(13.5232)
	nl.SetValuePrecision(pre)
	assert.Equal(t, nl.GetPrecision(), pre)

	nl.SetAnyDescription("example_description")
	assert.Equal(t, nl.GetDescription(), "example_description")

	nl.SetAnyCountry("RU")
	assert.Equal(t, nl.GetCountry(), "RU")

	nl.SetAnyCity("Moscow")
	assert.Equal(t, nl.GetCity(), "Moscow")

	nl.SetAnyAdministrativeArea("RU_RU")
	assert.Equal(t, nl.GetAdministrativeArea(), "RU_RU")

	sa := "Jucova street, 45c1"
	nl.SetAnyStreetAddress(sa)
	assert.Equal(t, nl.GetStreetAddress(), sa)

	nl.SetAnyPostalCode("105420")
	assert.Equal(t, nl.GetPostalCode(), "105420")

	req := stixhelpers.OpenVocabTypeSTIX("region name")
	nl.SetValueRegion(req)
	assert.Equal(t, nl.GetRegion(), req)
}

/*
// LocationDomainObjectsSTIX объект "Location", по терминалогии STIX, содержит описание географического местоположения
// Name - имя используемое для идентификации "Location"
// Description - более подробное описание
// Latitude - широта
// Longitude - долгота
// Precision - определяет точность координат, заданных свойствами широта и долгота (измеряется в метрах)
// Region - один, из заранее определенного (предложенного) перечня регионов
// Country - наименование страны
// AdministrativeArea - административный округ
// City - наименование города
// StreetAddress - физический адрес
// PostalCode - почтовый адрес
type LocationDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Latitude           float32                       `json:"latitude" bson:"latitude"`
	Longitude          float32                       `json:"longitude" bson:"longitude"`
	Precision          float32                       `json:"precision" bson:"precision"`
	Name               string                        `json:"name" bson:"name"`
	Description        string                        `json:"description" bson:"description"`
	Country            string                        `json:"country" bson:"country"`
	AdministrativeArea string                        `json:"administrative_area" bson:"administrative_area"`
	City               string                        `json:"city" bson:"city"`
	StreetAddress      string                        `json:"street_address" bson:"street_address"`
	PostalCode         string                        `json:"postal_code" bson:"postal_code"`
	Region             stixhelpers.OpenVocabTypeSTIX `json:"region" bson:"region"`
}
*/
