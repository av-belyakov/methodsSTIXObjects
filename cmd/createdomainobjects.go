package cmd

import (
	"github.com/av-belyakov/methodstixobjects/datamodels/domainobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// NewAttackPatternDomainObjectsSTIX создает STIX объект "Attack Pattern", по терминалогии STIX, описывающий способы компрометации цели
func NewAttackPatternDomainObjectsSTIX() *domainobjectsstix.AttackPatternDomainObjectsSTIX {
	return &domainobjectsstix.AttackPatternDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Aliases:                          make([]string, 0),
		KillChainPhases:                  make([]stixhelpers.KillChainPhasesTypeElementSTIX, 0),
	}
}

// NewCampaignDomainObjectsSTIX создает STIX объект "Campaign", по терминалогии STIX, это набор действий определяющих злонамеренную деятельность или атаки
func NewCampaignDomainObjectsSTIX() *domainobjectsstix.CampaignDomainObjectsSTIX {
	return &domainobjectsstix.CampaignDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
	}
}

// NewCourseOfActionDomainObjectsSTIX создает STIX объект "Course of Action", по терминалогии STIX, описывающий совокупность действий
// направленных на предотвращение (защиту) либо реагирование на текущую атаку
func NewCourseOfActionDomainObjectsSTIX() *domainobjectsstix.CourseOfActionDomainObjectsSTIX {
	return &domainobjectsstix.CourseOfActionDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}

// NewGroupingDomainObjectsSTIX создает STIX объект "Grouping", по терминалогии STIX, объединяет различные объекты STIX в рамках какого то общего контекста
func NewGroupingDomainObjectsSTIX() *domainobjectsstix.GroupingDomainObjectsSTIX {
	return &domainobjectsstix.GroupingDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		ObjectRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewIdentityDomainObjectsSTIX создает STIX объект "Identity", по терминалогии STIX, содержит основную идентификационную информацию физичиских лиц, организаций и т.д.
func NewIdentityDomainObjectsSTIX() *domainobjectsstix.IdentityDomainObjectsSTIX {
	return &domainobjectsstix.IdentityDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Roles:                            make([]string, 0),
		Sectors:                          make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewIndicatorDomainObjectsSTIX создает STIX объект "Indicator", по терминалогии STIX, содержит шаблон который может быть использован для обнаружения
// подозрительной или вредоносной киберактивности
func NewIndicatorDomainObjectsSTIX() *domainobjectsstix.IndicatorDomainObjectsSTIX {
	return &domainobjectsstix.IndicatorDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		ValidFrom:                        "1970-01-01T00:00:00+00:00",
		IndicatorTypes:                   make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewInfrastructureDomainObjectsSTIX создает STIX объект "Infrastructure", по терминалогии STIX, содержит описание любых систем, программных
// служб, а так же любые связанные с ними физические или виртуальные ресурсы, предназначенные для поддержки какой-либо цели
func NewInfrastructureDomainObjectsSTIX() *domainobjectsstix.InfrastructureDomainObjectsSTIX {
	return &domainobjectsstix.InfrastructureDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          make([]string, 0),
		InfrastructureTypes:              make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewIntrusionSetDomainObjectsSTIX создает STIX объект "Intrusion Set", по терминалогии STIX, содержит сгруппированный набор враждебного поведения и ресурсов
// с общими свойствами, который, как считается, управляется одной организацией
func NewIntrusionSetDomainObjectsSTIX() *domainobjectsstix.IntrusionSetDomainObjectsSTIX {
	return &domainobjectsstix.IntrusionSetDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          make([]string, 0),
		Goals:                            make([]string, 0),
		SecondaryMotivations:             make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewLocationDomainObjectsSTIX создает STIX объект "Location", по терминалогии STIX, содержит описание географического местоположения
func NewLocationDomainObjectsSTIX() *domainobjectsstix.LocationDomainObjectsSTIX {
	return &domainobjectsstix.LocationDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}

// NewMalwareDomainObjectsSTIX создает STIX объект "Malware", по терминалогии STIX, содержит подробную информацию о функционировании вредоносной программы
func NewMalwareDomainObjectsSTIX() *domainobjectsstix.MalwareDomainObjectsSTIX {
	return &domainobjectsstix.MalwareDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          make([]string, 0),
		MalwareTypes:                     make([]stixhelpers.OpenVocabTypeSTIX, 0),
		OperatingSystemRefs:              make([]stixhelpers.IdentifierTypeSTIX, 0),
		ArchitectureExecutionEnvs:        make([]stixhelpers.OpenVocabTypeSTIX, 0),
		ImplementationLanguages:          make([]stixhelpers.OpenVocabTypeSTIX, 0),
		Capabilities:                     make([]stixhelpers.OpenVocabTypeSTIX, 0),
		SampleRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewMalwareAnalysisDomainObjectsSTIX создает STIX объект "Malware Analysis", по терминалогии STIX, содержит анализ вредоносных программ
// захватывающих метаданные и результаты конкретного статического или динамического анализа, выполненного на экземпляре
// вредоносного ПО или семействе вредоносных программ
func NewMalwareAnalysisDomainObjectsSTIX() *domainobjectsstix.MalwareAnalysisDomainObjectsSTIX {
	return &domainobjectsstix.MalwareAnalysisDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Submitted:                        "1970-01-01T00:00:00+00:00",
		AnalysisStarted:                  "1970-01-01T00:00:00+00:00",
		AnalysisEnded:                    "1970-01-01T00:00:00+00:00",
		Modules:                          make([]string, 0),
		InstalledSoftwareRefs:            make([]stixhelpers.IdentifierTypeSTIX, 0),
		AnalysisScoRefs:                  make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewNoteDomainObjectsSTIX создает STIX объект "Note", по терминалогии STIX, содержит текстовую информации дополняющую текущий контекст анализа
// либо содержащей результаты дополнительного анализа которые не может быть описан в терминах объектов STIX
func NewNoteDomainObjectsSTIX() *domainobjectsstix.NoteDomainObjectsSTIX {
	return &domainobjectsstix.NoteDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Authors:                          make([]string, 0),
		ObjectRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewObservedDataDomainObjectsSTIX создает STIX объект "Observed Data", по терминалогии STIX, содержит информацию о сущностях связанных с
// кибер безопасностью, таких как файлы, системы или сети. Наблюдаемые данные это не результат анализа или заключение искусственного
// интеллекта, это просто сырая информация без какого-либо контекста.
func NewObservedDataDomainObjectsSTIX() *domainobjectsstix.ObservedDataDomainObjectsSTIX {
	return &domainobjectsstix.ObservedDataDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstObserved:                    "1970-01-01T00:00:00+00:00",
		LastObserved:                     "1970-01-01T00:00:00+00:00",
		ObjectRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewOpinionDomainObjectsSTIX создает STIX объект "Opinion", по терминалогии STIX, содержит оценку информации в приведенной в каком либо другом объекте STIX
// которую произвел другой участник анализа.
func NewOpinionDomainObjectsSTIX() *domainobjectsstix.OpinionDomainObjectsSTIX {
	return &domainobjectsstix.OpinionDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Authors:                          make([]string, 0),
		ObjectRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewReportDomainObjectsSTIX создает STIX объект "Report", по терминалогии STIX, содержит совокупность данных об угрозах, сосредоточенных на одной
// или нескольких темах, таких как описание исполнителя, вредоносного ПО или метода атаки, включая контекст и связанные с ним детали.
// Применяется для группировки информации связанной с кибер угрозой. Может быть использован для дальнейшей публикации данной
// информации как истории расследования.
func NewReportDomainObjectsSTIX() *domainobjectsstix.ReportDomainObjectsSTIX {
	return &domainobjectsstix.ReportDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Published:                        "1970-01-01T00:00:00+00:00",
		ReportTypes:                      make([]stixhelpers.OpenVocabTypeSTIX, 0),
		ObjectRefs:                       make([]stixhelpers.IdentifierTypeSTIX, 0),
	}
}

// NewThreatActorDomainObjectsSTIX создает STIX объект "Threat Actor", по терминалогии STIX, содержит информацию о физических лицах или их
// группах и организациях которые могут действовать со злым умыслом.
func NewThreatActorDomainObjectsSTIX() *domainobjectsstix.ThreatActorDomainObjectsSTIX {
	return &domainobjectsstix.ThreatActorDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		FirstSeen:                        "1970-01-01T00:00:00+00:00",
		LastSeen:                         "1970-01-01T00:00:00+00:00",
		Aliases:                          make([]string, 0),
		Goals:                            make([]string, 0),
		SecondaryMotivations:             make([]stixhelpers.OpenVocabTypeSTIX, 0),
		PersonalMotivations:              make([]stixhelpers.OpenVocabTypeSTIX, 0),
		ThreatActorTypes:                 make([]stixhelpers.OpenVocabTypeSTIX, 0),
		Roles:                            make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewToolDomainObjectsSTIX создает STIX объект "Tool", по терминалогии STIX, содержит информацию о легитимном ПО которое может быть
// использованно для реализации компьютерных угроз
func NewToolDomainObjectsSTIX() *domainobjectsstix.ToolDomainObjectsSTIX {
	return &domainobjectsstix.ToolDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
		Aliases:                          make([]string, 0),
		ToolTypes:                        make([]stixhelpers.OpenVocabTypeSTIX, 0),
	}
}

// NewVulnerabilityDomainObjectsSTIX создает STIX объект "Vulnerability", по терминологии STIX, содержит описание уязвимостей полученных в
// результате неверной формализации требований, ошибочном проектировании или некорректной реализации программного кода
// или логики в ПО, а также в компонентах оборудования
func NewVulnerabilityDomainObjectsSTIX() *domainobjectsstix.VulnerabilityDomainObjectsSTIX {
	return &domainobjectsstix.VulnerabilityDomainObjectsSTIX{
		CommonPropertiesObjectSTIX:       *NewCommonPropertiesObjectSTIX(),
		CommonPropertiesDomainObjectSTIX: *NewCommonPropertiesDomainObjectSTIX(),
	}
}
