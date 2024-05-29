package relationshipobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestCreateOptionalCommonPropertiesRelationshipObjectSTIX(t *testing.T) {
	nocp := methodstixobjects.NewOptionalCommonPropertiesRelationshipObjectSTIX()

	_, err := nocp.Get()
	assert.Error(t, err)

	nocp.SetValueSpecVersion("1.1")
	err = nocp.SetValueCreated("2024-03-02T10:45:01+00:00")
	assert.NoError(t, err)
	err = nocp.SetValueModified("2024-03-12T03:12:51+00:00")
	assert.NoError(t, err)

	assert.NoError(t, err)
}

func TestRelationshipObjectSTIX(t *testing.T) {
	nro := methodstixobjects.NewRelationshipObjectSTIX()

	_, err := nro.Get()
	assert.Error(t, err)

	nro.SetValueRelationshipType("rt")
	nro.SetValueSourceRef("reference_source")
	nro.SetValueTargetRef("reference_target")

	_, err = nro.Get()
	assert.NoError(t, err)

	nro.SetValueDescription("any message")
	assert.Equal(t, nro.GetDescription(), "any message")

	ct := "2024-01-10T12:25:01+00:00"
	err = nro.SetValueCreated(ct)
	assert.NoError(t, err)
	assert.Equal(t, nro.GetCreated(), ct)

	mt := "2024-02-01T13:35:21+00:00"
	err = nro.SetValueModified(mt)
	assert.NoError(t, err)
	assert.Equal(t, nro.GetModified(), mt)

	startt := "2024-02-04T14:31:21+00:00"
	err = nro.SetValueStartTime(startt)
	assert.NoError(t, err)
	assert.Equal(t, nro.GetStartTime(), startt)

	stopt := "2024-02-05T21:11:01+00:00"
	err = nro.SetValueStartTime(stopt)
	assert.NoError(t, err)
	assert.Equal(t, nro.GetStartTime(), stopt)

	sr := stixhelpers.IdentifierTypeSTIX("source_ref_example")
	nro.SetValueSourceRef(sr)
	assert.Equal(t, nro.GetSourceRef(), sr)

	tr := stixhelpers.IdentifierTypeSTIX("target_ref_example")
	nro.SetValueSourceRef(tr)
	assert.Equal(t, nro.GetSourceRef(), tr)

	assert.Equal(t, nro.GetType(), "relationship")
}

func TestSightingObjectSTIX(t *testing.T) {
	nso := methodstixobjects.NewSightingObjectSTIX()

	_, err := nso.Get()
	assert.Error(t, err)

	nso.SetValueSightingOfRef("reference_sigting")
	_, err = nso.Get()
	assert.NoError(t, err)
	assert.Equal(t, nso.GetType(), "sighting")

	nso.SetValueSummary(true)
	assert.True(t, nso.GetSummary())

	nso.SetValueCount(3)
	assert.Equal(t, nso.GetCount(), 3)

	vd := "full description"
	nso.SetValueDescription(vd)
	assert.Equal(t, nso.GetDescription(), vd)

	startt := "2024-02-04T14:31:21+00:00"
	err = nso.SetValueFirstSeen(startt)
	assert.NoError(t, err)
	assert.Equal(t, nso.GetFirstSeen(), startt)

	stopt := "2024-02-05T21:11:01+00:00"
	err = nso.SetValueLastSeen(stopt)
	assert.NoError(t, err)
	assert.Equal(t, nso.GetLastSeen(), stopt)

	nso.SetValueObservedDataRefs([]stixhelpers.IdentifierTypeSTIX{"wwwe", "fffd"})
	assert.Equal(t, len(nso.GetObservedDataRefs()), 2)

	nso.SetValueWhereSightedRefs([]stixhelpers.IdentifierTypeSTIX{"1dwwexz", "2ddert", "3vdrgrgh"})
	assert.Equal(t, len(nso.GetWhereSightedRefs()), 3)
}
