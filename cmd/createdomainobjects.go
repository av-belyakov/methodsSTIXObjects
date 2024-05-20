package cmd

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// NewAttackPatternDomainObjectsSTIX создает STIX объект "Attack Pattern", по терминалогии STIX, описывающий способы компрометации цели
func NewAttackPatternDomainObjectsSTIX() *domainobjectsstix.AttackPatternDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("attack-pattern")

	return &domainobjectsstix.AttackPatternDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Aliases:                          []string(nil),
		KillChainPhases:                  []stixhelpers.KillChainPhasesTypeElementSTIX(nil),
	}
}

// NewCampaignDomainObjectsSTIX создает STIX объект "Campaign", по терминалогии STIX, это набор действий определяющих злонамеренную деятельность или атаки
func NewCampaignDomainObjectsSTIX() *domainobjectsstix.CampaignDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("campaign")

	return &domainobjectsstix.CampaignDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
	}
}

// NewCourseOfActionDomainObjectsSTIX создает STIX объект "Course of Action", по терминалогии STIX, описывающий совокупность действий
// направленных на предотвращение (защиту) либо реагирование на текущую атаку
func NewCourseOfActionDomainObjectsSTIX() *domainobjectsstix.CourseOfActionDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("course-of-action")

	return &domainobjectsstix.CourseOfActionDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}

// NewGroupingDomainObjectsSTIX создает STIX объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
func NewGroupingDomainObjectsSTIX() *domainobjectsstix.GroupingDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("grouping")

	return &domainobjectsstix.GroupingDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		ObjectRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewIdentityDomainObjectsSTIX создает STIX объект "Identity", по терминалогии STIX, содержит основную идентификационную информацию физичиских лиц, организаций и т.д.
func NewIdentityDomainObjectsSTIX() *domainobjectsstix.IdentityDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("identity")

	return &domainobjectsstix.IdentityDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Roles:                            []string(nil),
		Sectors:                          []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewIndicatorDomainObjectsSTIX создает STIX объект "Indicator", по терминалогии STIX, содержит шаблон который может быть использован для обнаружения
// подозрительной или вредоносной киберактивности
func NewIndicatorDomainObjectsSTIX() *domainobjectsstix.IndicatorDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("indicator")

	return &domainobjectsstix.IndicatorDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		ValidFrom:                        "1970-01-01T00:00:00+00:00",
		IndicatorTypes:                   []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewInfrastructureDomainObjectsSTIX создает STIX объект "Infrastructure", по терминалогии STIX, содержит описание любых систем, программных
// служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
func NewInfrastructureDomainObjectsSTIX() *domainobjectsstix.InfrastructureDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("infrastructure")

	return &domainobjectsstix.InfrastructureDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          []string(nil),
		InfrastructureTypes:              []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewIntrusionSetDomainObjectsSTIX создает STIX объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный набор враждебного поведения и ресурсов
// с общими свойствами, который, как считается, управляется одной организацией
func NewIntrusionSetDomainObjectsSTIX() *domainobjectsstix.IntrusionSetDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("intrusion-set")

	return &domainobjectsstix.IntrusionSetDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          []string(nil),
		Goals:                            []string(nil),
		SecondaryMotivations:             []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewLocationDomainObjectsSTIX создает STIX объект "Location", по терминалогии STIX, содержит описание географического местоположения
func NewLocationDomainObjectsSTIX() *domainobjectsstix.LocationDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("location")

	return &domainobjectsstix.LocationDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}

// NewMalwareDomainObjectsSTIX создает STIX объект "Malware", по терминалогии STIX, содержит подробную информацию о функционировании вредоносной программы
func NewMalwareDomainObjectsSTIX() *domainobjectsstix.MalwareDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("malware")

	return &domainobjectsstix.MalwareDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          []string(nil),
		MalwareTypes:                     []stixhelpers.OpenVocabTypeSTIX(nil),
		OperatingSystemRefs:              []stixhelpers.IdentifierTypeSTIX(nil),
		ArchitectureExecutionEnvs:        []stixhelpers.OpenVocabTypeSTIX(nil),
		ImplementationLanguages:          []stixhelpers.OpenVocabTypeSTIX(nil),
		Capabilities:                     []stixhelpers.OpenVocabTypeSTIX(nil),
		SampleRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewMalwareAnalysisDomainObjectsSTIX создает STIX объект "Malware Analysis", по терминалогии STIX, содержит анализ вредоносных программ
// захватывающих метаданные и результаты конкретного статического или динамического анализа, выполненного на экземпляре
// вредоносного ПО или семействе вредоносных программ
func NewMalwareAnalysisDomainObjectsSTIX() *domainobjectsstix.MalwareAnalysisDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("malware-analysis")

	return &domainobjectsstix.MalwareAnalysisDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Submitted:                        "1970-01-01T00:00:00+00:00",
		AnalysisStarted:                  "1970-01-01T00:00:00+00:00",
		AnalysisEnded:                    "1970-01-01T00:00:00+00:00",
		Modules:                          []string(nil),
		InstalledSoftwareRefs:            []stixhelpers.IdentifierTypeSTIX(nil),
		AnalysisScoRefs:                  []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewNoteDomainObjectsSTIX создает STIX объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую текущий контекст анализа
// либо содержащей результаты дополнительного анализа которые не может быть описан в терминах объектов STIX
func NewNoteDomainObjectsSTIX() *domainobjectsstix.NoteDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("note")

	return &domainobjectsstix.NoteDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Authors:                          []string(nil),
		ObjectRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewObservedDataDomainObjectsSTIX создает STIX объект "Observed Data", по терминалогии STIX, содержит информацию о сущностях связанных с
// кибер безопасностью, таких как файлы, системы или сети. Наблюдаемые данные это не результат анализа или заключение искусственного
// интеллекта, это просто сырая информация без какого-либо контекста.
func NewObservedDataDomainObjectsSTIX() *domainobjectsstix.ObservedDataDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("observed-data")

	return &domainobjectsstix.ObservedDataDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstObserved:                    "1970-01-01T00:00:00+00:00",
		LastObserved:                     "1970-01-01T00:00:00+00:00",
		ObjectRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewOpinionDomainObjectsSTIX создает STIX объект "Opinion", по терминалогии STIX, содержит оценку информации в приведенной в каком либо другом объекте STIX
// которую произвел другой участник анализа.
func NewOpinionDomainObjectsSTIX() *domainobjectsstix.OpinionDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("opinion")

	return &domainobjectsstix.OpinionDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Authors:                          []string(nil),
		ObjectRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewReportDomainObjectsSTIX создает STIX объект "Report", по терминалогии STIX, содержит совокупность данных об угрозах, сосредоточенных на одной
// или нескольких темах, таких как описание исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним детали.
// Применяется для группировки информации связанной с кибер угрозой. Может быть использован для дальнейшей публикации данной
// информации как истории расследования.
func NewReportDomainObjectsSTIX() *domainobjectsstix.ReportDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("report")

	return &domainobjectsstix.ReportDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Published:                        "1970-01-01T00:00:00+00:00",
		ReportTypes:                      []stixhelpers.OpenVocabTypeSTIX(nil),
		ObjectRefs:                       []stixhelpers.IdentifierTypeSTIX(nil),
	}
}

// NewThreatActorDomainObjectsSTIX создает STIX объект "Threat Actor", по терминалогии STIX, содержит информацию о физических лицах или их
// группах и организациях которые могут действовать со злым умыслом.
func NewThreatActorDomainObjectsSTIX() *domainobjectsstix.ThreatActorDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("threat-actor")

	return &domainobjectsstix.ThreatActorDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          []string(nil),
		Goals:                            []string(nil),
		SecondaryMotivations:             []stixhelpers.OpenVocabTypeSTIX(nil),
		PersonalMotivations:              []stixhelpers.OpenVocabTypeSTIX(nil),
		ThreatActorTypes:                 []stixhelpers.OpenVocabTypeSTIX(nil),
		Roles:                            []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewToolDomainObjectsSTIX создает STIX объект "Tool", по терминалогии STIX, содержит информацию о легитимном ПО которое может быть
// использованно для реализации компьютерных угроз
func NewToolDomainObjectsSTIX() *domainobjectsstix.ToolDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("tool")

	return &domainobjectsstix.ToolDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Aliases:                          []string(nil),
		ToolTypes:                        []stixhelpers.OpenVocabTypeSTIX(nil),
	}
}

// NewVulnerabilityDomainObjectsSTIX создает STIX объект "Vulnerability", по терминологии STIX, содержит описание уязвимостей полученных в
// результате неверной формализации требований, ошибочном проектировании или некорректной реализации программного кода
// или логики в ПО, а также в компонентах оборудования
func NewVulnerabilityDomainObjectsSTIX() *domainobjectsstix.VulnerabilityDomainObjectsSTIX {
	cpo := NewCommonPropertiesObjectSTIX()
	cpo.SetAnyType("vulnerability")

	return &domainobjectsstix.VulnerabilityDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *cpo.Get(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}
