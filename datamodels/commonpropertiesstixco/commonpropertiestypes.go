package commonpropertiesstixco

import "github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"

// OptionalCommonPropertiesCyberObservableObjectSTIX содержит опциональные общие свойства для Cyber-observable Objects STIX
// SpecVersion - версия STIX спецификации.
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся
// значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// // Extensions - может содержать дополнительную информацию, относящуюся к объекту
type OptionalCommonPropertiesCyberObservableObjectSTIX struct {
	Defanged          bool                                   `json:"defanged" bson:"defanged"`
	SpecVersion       string                                 `json:"spec_version" bson:"spec_version"`
	ObjectMarkingRefs []stixhelpers.IdentifierTypeSTIX       `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings  []stixhelpers.GranularMarkingsTypeSTIX `json:"granular_markings" bson:"granular_markings"`
	//Extensions        map[string]DictionaryTypeSTIX `json:"extensions" bson:"extensions"`
}
