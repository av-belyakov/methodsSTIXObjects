package testing

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestOptionalCommonPropertiesCyberObservableObjectSTIX(t *testing.T) {
	nocpc := methodstixobjects.NewOptionalCommonPropertiesCyberObservableObjectSTIX()
	nocpc.SetValueDefanged(true)
	assert.True(t, nocpc.GetDefanged())

	nocpc.SetValueSpecVersion("223")
	assert.Equal(t, nocpc.GetSpecVersion(), "223")

	omr := []stixhelpers.IdentifierTypeSTIX{"2sdff", "chs732", "ydd383"}
	nocpc.SetValueObjectMarkingRefs(omr)
	assert.Equal(t, len(nocpc.GetObjectMarkingRefs()), 3)

	gm := []stixhelpers.GranularMarkingsTypeSTIX{
		{Lang: "lang_example"},
		{Selectors: []string{"1", "2", "3"}},
	}
	nocpc.SetValueGranularMarkings(gm)
	assert.Equal(t, len(nocpc.GetGranularMarkings()), 2)
}
