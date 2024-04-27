package commonpropertiesstixdo

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// CommonPropertiesDomainObjectSTIX свойства общие, для всех объектов STIX
// SpecVersion - версия спецификации STIX используемая для представления текущего объекта (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Created - время создания объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Modified - время модификации объекта, в формате "2016-05-12T08:17:27.000Z" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// CreatedByRef - содержит идентификатор источника создавшего данный объект
// Revoked - вернуть к текущему состоянию
// Labels - определяет набор терминов, используемых для описания данного объекта
// Сonfidence - определяет уверенность создателя в правильности своих данных. Доверительное значение ДОЛЖНО быть числом
// в диапазоне 0-100. Если 0 - значение не определено.
// Lang - содержит текстовый код языка, на котором написан контент объекта. Для английского это "en" для русского "ru"
// ExternalReferences - список внешних ссылок не относящихся к STIX информации
// ObjectMarkingRefs - определяет список ID ссылающиеся на объект "marking-definition", по терминалогии STIX, в котором содержатся значения применяющиеся к этому объекту
// GranularMarkings - определяет список "гранулярных меток" (granular_markings) относящихся к этому объекту
// Defanged - определяет были ли определены данные содержащиеся в объекте
// Extensions - может содержать дополнительную информацию, относящуюся к объекту
type CommonPropertiesDomainObjectSTIX struct {
	Revoked            bool                                           `json:"revoked" bson:"revoked"`
	Defanged           bool                                           `json:"defanged" bson:"defanged"`
	Сonfidence         int                                            `json:"confidence" bson:"confidence"`
	Lang               string                                         `json:"lang" bson:"lang"`
	SpecVersion        string                                         `json:"spec_version" bson:"spec_version" required:"true"`
	Created            string                                         `json:"created" bson:"created" required:"true"`
	Modified           string                                         `json:"modified" bson:"modified" required:"true"`
	Labels             []string                                       `json:"labels" bson:"labels"`
	Extensions         map[string]string                              `json:"extensions" bson:"extensions"`
	CreatedByRef       stixhelpers.IdentifierTypeSTIX                 `json:"created_by_ref" bson:"created_by_ref"`
	ExternalReferences []stixhelpers.ExternalReferenceTypeElementSTIX `json:"external_references" bson:"external_references"`
	ObjectMarkingRefs  []stixhelpers.IdentifierTypeSTIX               `json:"object_marking_refs" bson:"object_marking_refs"`
	GranularMarkings   []stixhelpers.GranularMarkingsTypeSTIX         `json:"granular_markings" bson:"granular_markings"`
}
