package testing

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCommonPropertiesDomainObjectSTIX(t *testing.T) {
	ncpdo := methodstixobjects.NewCommonPropertiesDomainObjectSTIX()

	ncpdo.SetValueRevoked(true)
	assert.True(t, ncpdo.GetRevoked())

	ncpdo.SetValueDefanged(true)
	assert.True(t, ncpdo.Defanged)

	ncpdo.SetValueConfidence(111)
	assert.Equal(t, ncpdo.GetСonfidence(), 111)

	le := "lang_example"
	ncpdo.SetValueLang(le)
	assert.Equal(t, ncpdo.GetLang(), le)

	ncpdo.SetValueSpecVersion("1.23")
	assert.Equal(t, ncpdo.GetSpecVersion(), "1.23")

	ct := "2024-02-05T16:11:21+00:00"
	ncpdo.SetValueCreated(ct)
	assert.Equal(t, ncpdo.GetCreated(), ct)

	mt := "2024-02-05T17:31:01+00:00"
	ncpdo.SetValueModified(mt)
	assert.Equal(t, ncpdo.GetModified(), mt)

	ncpdo.SetValueLabels("labels_1")
	ncpdo.SetValueLabels("labels_2")
	assert.Equal(t, len(ncpdo.GetLabels()), 2)

	ncpdo.SetAnyExtensions("key_1", 123)
	ncpdo.SetAnyExtensions("key_2", "234")
	ext := ncpdo.GetExtensions()
	assert.Equal(t, len(ext), 2)
	key1, ok := ext["key_1"]
	assert.True(t, ok)
	assert.Equal(t, key1, "123")

	ncpdo.SetValueCreatedByRef("full_extension")
	assert.Equal(t, ncpdo.GetCreatedByRef(), stixhelpers.IdentifierTypeSTIX("full_extension"))

	ncpdo.SetValueExternalReferences([]stixhelpers.ExternalReferenceTypeElementSTIX{
		{SourceName: "source_1", Description: "source 1 description"},
		{SourceName: "source_2", Description: "source 2 description"},
	})
	assert.Equal(t, len(ncpdo.GetExternalReferences()), 2)

	ncpdo.SetValueObjectMarkingRefs([]stixhelpers.IdentifierTypeSTIX{"1ddc", "2fdf"})
	assert.Equal(t, len(ncpdo.GetObjectMarkingRefs()), 2)

	ncpdo.SetValueGranularMarkings([]stixhelpers.GranularMarkingsTypeSTIX{
		{MarkingRef: stixhelpers.IdentifierTypeSTIX("bg43w1sd")},
		{MarkingRef: stixhelpers.IdentifierTypeSTIX("zz2")},
		{MarkingRef: stixhelpers.IdentifierTypeSTIX("hth43")},
	})
	assert.Equal(t, len(ncpdo.GetGranularMarkings()), 3)
}
