package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestIndicatorDomainObjectsSTIX(t *testing.T) {
	ni := methodstixobjects.NewIndicatorDomainObjectsSTIX()

	assert.Equal(t, ni.GetType(), "indicator")
	_, err := ni.Get()
	assert.Error(t, err)

	ni.SetAnyName("indicator name")
	_, err = ni.Get()
	assert.NoError(t, err)

	ni.SetAnyDescription("example_description")
	assert.Equal(t, ni.GetDescription(), "example_description")

	pattern := "pattern example"
	ni.SetAnyPattern(pattern)
	assert.Equal(t, ni.GetPattern(), pattern)

	patternVersion := "pattern version example"
	ni.SetAnyPatternVersion(patternVersion)
	assert.Equal(t, ni.GetPatternVersion(), patternVersion)

	vf := "2024-02-11T07:01:01+00:00"
	err = ni.SetAnyValidFrom(vf)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetValidFrom(), vf)

	err = ni.SetAnyValidFrom(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetValidFrom(), "2024-05-22T11:43:27+03:00")

	vu := "2024-02-10T11:21:01+00:00"
	err = ni.SetAnyValidUntil(vu)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetValidUntil(), vu)

	err = ni.SetAnyValidUntil(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, ni.GetValidUntil(), "2024-05-22T11:43:27+03:00")

	pt := stixhelpers.OpenVocabTypeSTIX("pattern type example")
	ni.SetValuePatternType(pt)
	assert.Equal(t, ni.GetPatternType(), pt)

	ni.SetValueKillChainPhases([]stixhelpers.KillChainPhasesTypeElementSTIX{
		{KillChainName: "green chan", PhaseName: "first phase"},
		{KillChainName: "red chan", PhaseName: "second phase"},
		{KillChainName: "blue chan", PhaseName: "third phase"},
	})
	assert.Equal(t, len(ni.GetKillChainPhases()), 3)

	ni.SetValueIndicatorTypes([]stixhelpers.OpenVocabTypeSTIX{"1q", "2x", "3a", "4d"})
	assert.Equal(t, len(ni.GetIndicatorTypes()), 4)
}

/*
// IndicatorDomainObjectsSTIX объект "Indicator", по терминалогии STIX, содержит шаблон который может быть использован для обнаружения
// подозрительной или вредоносной киберактивности
// Name - имя используемое для идентификации "Indicator" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// IndicatorTypes - заранее определенный (предложенный) перечень категорий индикаторов
// Pattern - шаблон для обнаружения индикаторов (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternType - одно, из заранее определенных (предложенных) значений языкового шаблона, используемого в этом индикаторе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PatternVersion - версия языка шаблонов
// ValidFrom - время с которого этот индикатор считается валидным, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// ValidUntil - время начиная с которого этот индикатор не может считаться валидным, в формате "2016-05-12T08:17:27.000Z"
// KillChainPhases - список цепочки фактов, к которым можно отнести соответствующте индикаторы
type IndicatorDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name            string                              `json:"name" bson:"name" required:"true"`
	Pattern         string                              `json:"pattern" bson:"pattern" required:"true"`
	PatternVersion  string                              `json:"pattern_version" bson:"pattern_version"`
	Description     string                              `json:"description" bson:"description"`
	ValidFrom       string                              `json:"valid_from" bson:"valid_from" required:"true"`
	ValidUntil      string                              `json:"valid_until" bson:"valid_until"`
	PatternType     stixhelpers.OpenVocabTypeSTIX       `json:"pattern_type" bson:"pattern_type" required:"true"`
	KillChainPhases []stixhelpers.KillChainPhasesTypeElementSTIX `json:"kill_chain_phases" bson:"kill_chain_phases"`
	IndicatorTypes  []stixhelpers.OpenVocabTypeSTIX     `json:"indicator_types" bson:"indicator_types"`
}
*/
