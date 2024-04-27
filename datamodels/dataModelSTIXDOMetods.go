package datamodels

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/*************************************************************************/
/********** 			Domain Objects STIX (МЕТОДЫ)			**********/
/*************************************************************************/

/* --- InfrastructureDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (istix InfrastructureDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (istix InfrastructureDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе InfrastructureDomainObjectsSTIX
func (istix InfrastructureDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(infrastructure--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	if !istix.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if istix.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix InfrastructureDomainObjectsSTIX) SanitizeStruct() InfrastructureDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.InfrastructureTypes) > 0 {
		it := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(istix.InfrastructureTypes))
		for _, v := range istix.InfrastructureTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			it = append(it, tmp)
		}

		istix.InfrastructureTypes = it
	}

	if len(istix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(istix.Aliases))
		for _, v := range istix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		istix.Aliases = aliasesTmp
	}

	istix.KillChainPhases = istix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	return istix
}

// GetID возвращает ID STIX объекта
func (istix InfrastructureDomainObjectsSTIX) GetID() string {
	return istix.ID
}

// GetType возвращает Type STIX объекта
func (istix InfrastructureDomainObjectsSTIX) GetType() string {
	return istix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix InfrastructureDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(istix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", istix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", istix.Description))
	str.WriteString(fmt.Sprintf("'infrastructure_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'infrastructure_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(istix.InfrastructureTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(istix.Aliases)))
	str.WriteString(fmt.Sprintf("'sectors': \n%v", func(l []stixhelpers.KillChainPhasesTypeElementSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'sector '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(istix.KillChainPhases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", istix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", istix.LastSeen))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (istix InfrastructureDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   istix.ID,
		"type": istix.Type,
	}

	if istix.Name != "" {
		dataForIndex["name"] = istix.Name
	}

	if istix.Description != "" {
		dataForIndex["description"] = istix.Description
	}

	if len(istix.Aliases) > 0 {
		var strTmp string

		for _, v := range istix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- IntrusionSetDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (istix IntrusionSetDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &istix); err != nil {
		return nil, err
	}

	return istix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (istix IntrusionSetDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(istix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе IntrusionSetDomainObjectsSTIX
func (istix IntrusionSetDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(intrusion-set--)[0-9a-f|-]+$`).MatchString(istix.ID)) {
		return false
	}

	if !istix.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if istix.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (istix IntrusionSetDomainObjectsSTIX) SanitizeStruct() IntrusionSetDomainObjectsSTIX {
	istix.CommonPropertiesDomainObjectSTIX = istix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	istix.Name = commonlibs.StringSanitize(istix.Name)
	istix.Description = commonlibs.StringSanitize(istix.Description)

	if len(istix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(istix.Aliases))
		for _, v := range istix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		istix.Aliases = aliasesTmp
	}

	if len(istix.Goals) > 0 {
		goalsTmp := make([]string, 0, len(istix.Goals))
		for _, v := range istix.Goals {
			goalsTmp = append(goalsTmp, commonlibs.StringSanitize(v))
		}
		istix.Goals = goalsTmp
	}

	istix.ResourceLevel = istix.ResourceLevel.SanitizeStructOpenVocabTypeSTIX()
	istix.PrimaryMotivation = istix.PrimaryMotivation.SanitizeStructOpenVocabTypeSTIX()

	if len(istix.SecondaryMotivations) > 0 {
		sm := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(istix.SecondaryMotivations))
		for _, v := range istix.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			sm = append(sm, tmp)
		}

		istix.SecondaryMotivations = sm
	}

	return istix
}

// GetID возвращает ID STIX объекта
func (istix IntrusionSetDomainObjectsSTIX) GetID() string {
	return istix.ID
}

// GetType возвращает Type STIX объекта
func (istix IntrusionSetDomainObjectsSTIX) GetType() string {
	return istix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (istix IntrusionSetDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(istix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(istix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", istix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", istix.Description))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(istix.Aliases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", istix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", istix.LastSeen))
	str.WriteString(fmt.Sprintf("'goals': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'goal '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(istix.Goals)))
	str.WriteString(fmt.Sprintf("'resource_level': '%s'\n", istix.FirstSeen))
	str.WriteString(fmt.Sprintf("'primary_motivation': '%s'\n", istix.LastSeen))
	str.WriteString(fmt.Sprintf("'secondary_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'secondary_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(istix.SecondaryMotivations)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (istix IntrusionSetDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   istix.ID,
		"type": istix.Type,
	}

	if istix.Name != "" {
		dataForIndex["name"] = istix.Name
	}

	if istix.Description != "" {
		dataForIndex["description"] = istix.Description
	}

	if len(istix.Aliases) > 0 {
		var strTmp string

		for _, v := range istix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- LocationDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (lstix LocationDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &lstix); err != nil {
		return nil, err
	}

	return lstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (lstix LocationDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(lstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе LocationDomainObjectsSTIX
func (lstix LocationDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(location--)[0-9a-f|-]+$`).MatchString(lstix.ID)) {
		return false
	}

	if !lstix.ValidateStructCommonFields() {
		return false
	}

	if (lstix.Latitude > 90.0) || (lstix.Latitude < -90.0) {
		return false
	}

	if (lstix.Longitude > 180.0) || (lstix.Longitude < -180.0) {
		return false
	}

	if lstix.Country != "" {
		if !(regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(lstix.Country)) {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (lstix LocationDomainObjectsSTIX) SanitizeStruct() LocationDomainObjectsSTIX {
	lstix.CommonPropertiesDomainObjectSTIX = lstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	lstix.Name = commonlibs.StringSanitize(lstix.Name)
	lstix.Description = commonlibs.StringSanitize(lstix.Description)
	lstix.Region = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(lstix.Region)))
	lstix.AdministrativeArea = commonlibs.StringSanitize(lstix.AdministrativeArea)
	lstix.City = commonlibs.StringSanitize(lstix.City)
	lstix.StreetAddress = commonlibs.StringSanitize(lstix.StreetAddress)
	lstix.PostalCode = commonlibs.StringSanitize(lstix.PostalCode)

	return lstix
}

// GetID возвращает ID STIX объекта
func (lstix LocationDomainObjectsSTIX) GetID() string {
	return lstix.ID
}

// GetType возвращает Type STIX объекта
func (lstix LocationDomainObjectsSTIX) GetType() string {
	return lstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (lstix LocationDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(lstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(lstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", lstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", lstix.Description))
	str.WriteString(fmt.Sprintf("'latitude': '%v'\n", lstix.Latitude))
	str.WriteString(fmt.Sprintf("'longitude': '%v'\n", lstix.Longitude))
	str.WriteString(fmt.Sprintf("'precision': '%v'\n", lstix.Precision))
	str.WriteString(fmt.Sprintf("'region': '%s'\n", lstix.Region))
	str.WriteString(fmt.Sprintf("'country': '%s'\n", lstix.Country))
	str.WriteString(fmt.Sprintf("'administrative_area': '%s'\n", lstix.AdministrativeArea))
	str.WriteString(fmt.Sprintf("'city': '%s'\n", lstix.City))
	str.WriteString(fmt.Sprintf("'street_address': '%s'\n", lstix.StreetAddress))
	str.WriteString(fmt.Sprintf("'postal_code': '%s'\n", lstix.PostalCode))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (lstix LocationDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   lstix.ID,
		"type": lstix.Type,
	}

	if lstix.Name != "" {
		dataForIndex["name"] = lstix.Name
	}

	if lstix.Description != "" {
		dataForIndex["description"] = lstix.Description
	}

	if lstix.StreetAddress != "" {
		dataForIndex["street_address"] = lstix.StreetAddress
	}

	return dataForIndex
}

/* --- MalwareDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (mstix MalwareDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &mstix); err != nil {
		return nil, err
	}

	return mstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (mstix MalwareDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(mstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе MalwareDomainObjectsSTIX
func (mstix MalwareDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(malware--)[0-9a-f|-]+$`).MatchString(mstix.ID)) {
		return false
	}

	if !mstix.ValidateStructCommonFields() {
		return false
	}

	if len(mstix.OperatingSystemRefs) > 0 {
		for _, v := range mstix.OperatingSystemRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(mstix.SampleRefs) > 0 {
		for _, v := range mstix.SampleRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mstix MalwareDomainObjectsSTIX) SanitizeStruct() MalwareDomainObjectsSTIX {
	mstix.CommonPropertiesDomainObjectSTIX = mstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	mstix.Name = commonlibs.StringSanitize(mstix.Name)
	mstix.Description = commonlibs.StringSanitize(mstix.Description)

	if len(mstix.MalwareTypes) > 0 {
		mt := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(mstix.MalwareTypes))
		for _, v := range mstix.MalwareTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			mt = append(mt, tmp)
		}

		mstix.MalwareTypes = mt
	}

	if len(mstix.Aliases) > 0 {
		aliasesTmp := make([]string, 0, len(mstix.Aliases))
		for _, v := range mstix.Aliases {
			aliasesTmp = append(aliasesTmp, commonlibs.StringSanitize(v))
		}
		mstix.Aliases = aliasesTmp
	}

	mstix.KillChainPhases = mstix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()

	if len(mstix.ArchitectureExecutionEnvs) > 0 {
		aee := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(mstix.ArchitectureExecutionEnvs))
		for _, v := range mstix.ArchitectureExecutionEnvs {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			aee = append(aee, tmp)
		}

		mstix.ArchitectureExecutionEnvs = aee
	}

	if len(mstix.ImplementationLanguages) > 0 {
		il := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(mstix.ImplementationLanguages))
		for _, v := range mstix.ImplementationLanguages {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			il = append(il, tmp)
		}

		mstix.ImplementationLanguages = il
	}

	if len(mstix.Capabilities) > 0 {
		c := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(mstix.Capabilities))
		for _, v := range mstix.Capabilities {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			c = append(c, tmp)
		}

		mstix.Capabilities = c
	}

	return mstix
}

// GetID возвращает ID STIX объекта
func (mstix MalwareDomainObjectsSTIX) GetID() string {
	return mstix.ID
}

// GetType возвращает Type STIX объекта
func (mstix MalwareDomainObjectsSTIX) GetType() string {
	return mstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mstix MalwareDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(mstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(mstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", mstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", mstix.Description))
	str.WriteString(fmt.Sprintf("'malware_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'malware_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.MalwareTypes)))
	str.WriteString(fmt.Sprintf("'is_family': '%v'\n", mstix.IsFamily))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(mstix.Aliases)))
	str.WriteString(fmt.Sprintf("'kill_chain_phases': \n%v", func(l stixhelpers.KillChainPhasesTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'key': '%v' 'kill_chain_name': '%s'\n", k, v.KillChainName))
			str.WriteString(fmt.Sprintf("\t'key': '%v' 'phase_name': '%s'\n", k, v.PhaseName))
		}

		return str.String()
	}(mstix.KillChainPhases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", mstix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", mstix.LastSeen))
	str.WriteString(fmt.Sprintf("'operating_system_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'operating_system_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.OperatingSystemRefs)))
	str.WriteString(fmt.Sprintf("'architecture_execution_envs': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'architecture_execution_env '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.ArchitectureExecutionEnvs)))
	str.WriteString(fmt.Sprintf("'implementation_languages': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'implementation_language '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.ImplementationLanguages)))
	str.WriteString(fmt.Sprintf("'capabilities': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'capabilitie '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.Capabilities)))
	str.WriteString(fmt.Sprintf("'sample_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'sample_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mstix.SampleRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (mstix MalwareDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   mstix.ID,
		"type": mstix.Type,
	}

	if mstix.Name != "" {
		dataForIndex["name"] = mstix.Name
	}

	if mstix.Description != "" {
		dataForIndex["description"] = mstix.Description
	}

	if len(mstix.Aliases) > 0 {
		var strTmp string

		for _, v := range mstix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- MalwareAnalysisDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (mastix MalwareAnalysisDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &mastix); err != nil {
		return nil, err
	}

	return mastix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (mastix MalwareAnalysisDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(mastix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе MalwareAnalysisDomainObjectsSTIX
func (mastix MalwareAnalysisDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(malware-analysis--)[0-9a-f|-]+$`).MatchString(mastix.ID)) {
		return false
	}

	if !mastix.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if mastix.Product == "" {
		return false
	}

	if mastix.Version != "" && !(regexp.MustCompile(`^[0-9a-z.]+$`).MatchString(mastix.Version)) {
		return false
	}

	if !mastix.HostVMRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !mastix.OperatingSystemRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(mastix.InstalledSoftwareRefs) > 0 {
		for _, v := range mastix.InstalledSoftwareRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(mastix.AnalysisScoRefs) > 0 {
		for _, v := range mastix.AnalysisScoRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !mastix.SampleRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (mastix MalwareAnalysisDomainObjectsSTIX) SanitizeStruct() MalwareAnalysisDomainObjectsSTIX {
	mastix.CommonPropertiesDomainObjectSTIX = mastix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	mastix.Product = commonlibs.StringSanitize(mastix.Product)
	mastix.ConfigurationVersion = commonlibs.StringSanitize(mastix.ConfigurationVersion)
	if len(mastix.Modules) > 0 {
		mTmp := make([]string, 0, len(mastix.Modules))
		for _, v := range mastix.Modules {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		mastix.Modules = mTmp
	}
	mastix.AnalysisEngineVersion = commonlibs.StringSanitize(mastix.AnalysisEngineVersion)
	mastix.AnalysisDefinitionVersion = commonlibs.StringSanitize(mastix.AnalysisDefinitionVersion)
	mastix.ResultName = commonlibs.StringSanitize(mastix.ResultName)
	mastix.Result = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(mastix.Result)))
	mastix.AvResult = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(mastix.AvResult)))

	return mastix
}

// GetID возвращает ID STIX объекта
func (mastix MalwareAnalysisDomainObjectsSTIX) GetID() string {
	return mastix.ID
}

// GetType возвращает Type STIX объекта
func (mastix MalwareAnalysisDomainObjectsSTIX) GetType() string {
	return mastix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (mastix MalwareAnalysisDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(mastix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(mastix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'product': '%s'\n", mastix.Product))
	str.WriteString(fmt.Sprintf("'version': '%s'\n", mastix.Version))
	str.WriteString(fmt.Sprintf("'host_vm_ref': '%s'\n", mastix.HostVMRef))
	str.WriteString(fmt.Sprintf("'operating_system_ref': '%s'\n", mastix.OperatingSystemRef))
	str.WriteString(fmt.Sprintf("'installed_software_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'installed_software_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mastix.InstalledSoftwareRefs)))
	str.WriteString(fmt.Sprintf("'configuration_version': '%s'\n", mastix.ConfigurationVersion))
	str.WriteString(fmt.Sprintf("'modules': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'module '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(mastix.Modules)))
	str.WriteString(fmt.Sprintf("'analysis_engine_version': '%s'\n", mastix.AnalysisEngineVersion))
	str.WriteString(fmt.Sprintf("'analysis_definition_version': '%s'\n", mastix.AnalysisDefinitionVersion))
	str.WriteString(fmt.Sprintf("'submitted': '%v'\n", mastix.Submitted))
	str.WriteString(fmt.Sprintf("'analysis_started': '%v'\n", mastix.AnalysisStarted))
	str.WriteString(fmt.Sprintf("'analysis_ended': '%v'\n", mastix.AnalysisEnded))
	str.WriteString(fmt.Sprintf("'result_name': '%s'\n", mastix.ResultName))
	str.WriteString(fmt.Sprintf("'result': '%s'\n", mastix.Result))
	str.WriteString(fmt.Sprintf("'analysis_sco_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'analysis_sco_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(mastix.AnalysisScoRefs)))
	str.WriteString(fmt.Sprintf("'sample_ref': '%v'\n", mastix.SampleRef))
	str.WriteString(fmt.Sprintf("'av_result': '%v'\n", mastix.AvResult))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (mastix MalwareAnalysisDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   mastix.ID,
		"type": mastix.Type,
	}

	if mastix.Result != "" {
		dataForIndex["result_name"] = mastix.ResultName
	}

	return dataForIndex
}

/* --- NoteDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (nstix NoteDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &nstix); err != nil {
		return nil, err
	}

	return nstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (nstix NoteDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(nstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе NoteDomainObjectsSTIX
func (nstix NoteDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(note--)[0-9a-f|-]+$`).MatchString(nstix.ID)) {
		return false
	}

	if !nstix.ValidateStructCommonFields() {
		return false
	}

	if nstix.Content == "" {
		return false
	}

	for _, v := range nstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (nstix NoteDomainObjectsSTIX) SanitizeStruct() NoteDomainObjectsSTIX {
	nstix.CommonPropertiesDomainObjectSTIX = nstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	nstix.Abstract = commonlibs.StringSanitize(nstix.Abstract)
	nstix.Content = commonlibs.StringSanitize(nstix.Content)

	if len(nstix.Authors) > 0 {
		mTmp := make([]string, 0, len(nstix.Authors))
		for _, v := range nstix.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		nstix.Authors = mTmp
	}

	return nstix
}

// GetID возвращает ID STIX объекта
func (nstix NoteDomainObjectsSTIX) GetID() string {
	return nstix.ID
}

// GetType возвращает Type STIX объекта
func (nstix NoteDomainObjectsSTIX) GetType() string {
	return nstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (nstix NoteDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(nstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(nstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'abstract': '%s'\n", nstix.Abstract))
	str.WriteString(fmt.Sprintf("'content': '%s'\n", nstix.Content))
	str.WriteString(fmt.Sprintf("'authors': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'author '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(nstix.Authors)))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(nstix.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (nstix NoteDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   nstix.ID,
		"type": nstix.Type,
	}

	if nstix.Abstract != "" {
		dataForIndex["abstract"] = nstix.Abstract
	}

	if nstix.Content != "" {
		dataForIndex["content"] = nstix.Content
	}

	if len(nstix.Authors) > 0 {
		var strTmp string

		for _, v := range nstix.Authors {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- ObservedDataDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (odstix ObservedDataDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &odstix); err != nil {
		return nil, err
	}

	return odstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (odstix ObservedDataDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(odstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе ObservedDataDomainObjectsSTIX
func (odstix ObservedDataDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(observed-data--)[0-9a-f|-]+$`).MatchString(odstix.ID)) {
		return false
	}

	if odstix.NumberObserved <= 0 {
		return false
	}

	if (odstix.FirstObserved.Unix() < 0) || (odstix.LastObserved.Unix() < 0) {
		return false
	}

	if !odstix.ValidateStructCommonFields() {
		return false
	}

	//if len(odstix.ObjectRefs) > 0 {
	for _, v := range odstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}
	//}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (odstix ObservedDataDomainObjectsSTIX) SanitizeStruct() ObservedDataDomainObjectsSTIX {
	odstix.CommonPropertiesDomainObjectSTIX = odstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	return odstix
}

// GetID возвращает ID STIX объекта
func (odstix ObservedDataDomainObjectsSTIX) GetID() string {
	return odstix.ID
}

// GetType возвращает Type STIX объекта
func (odstix ObservedDataDomainObjectsSTIX) GetType() string {
	return odstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (odstix ObservedDataDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(odstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(odstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'first_observed': '%v'\n", odstix.FirstObserved))
	str.WriteString(fmt.Sprintf("'last_observed': '%v'\n", odstix.LastObserved))
	str.WriteString(fmt.Sprintf("'number_observed': '%d'\n", odstix.NumberObserved))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(odstix.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (odstix ObservedDataDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   odstix.ID,
		"type": odstix.Type,
	}
}

/* --- OpinionDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (ostix OpinionDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &ostix); err != nil {
		return nil, err
	}

	return ostix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (ostix OpinionDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(ostix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе OpinionDomainObjectsSTIX
func (ostix OpinionDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(opinion--)[0-9a-f|-]+$`).MatchString(ostix.ID)) {
		return false
	}

	if !ostix.ValidateStructCommonFields() {
		return false
	}

	if ostix.Opinion == "" {
		return false
	}

	for _, v := range ostix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ostix OpinionDomainObjectsSTIX) SanitizeStruct() OpinionDomainObjectsSTIX {
	ostix.CommonPropertiesDomainObjectSTIX = ostix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	ostix.Explanation = commonlibs.StringSanitize(ostix.Explanation)

	if len(ostix.Authors) > 0 {
		mTmp := make([]string, 0, len(ostix.Authors))
		for _, v := range ostix.Authors {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		ostix.Authors = mTmp
	}

	ostix.Opinion = stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(ostix.Opinion)))

	return ostix
}

// GetID возвращает ID STIX объекта
func (ostix OpinionDomainObjectsSTIX) GetID() string {
	return ostix.ID
}

// GetType возвращает Type STIX объекта
func (ostix OpinionDomainObjectsSTIX) GetType() string {
	return ostix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (ostix OpinionDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(ostix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(ostix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'explanation': '%v'\n", ostix.Explanation))
	str.WriteString(fmt.Sprintf("'authors': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'author '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(ostix.Authors)))
	str.WriteString(fmt.Sprintf("'opinion': '%v'\n", ostix.Opinion))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(ostix.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (ostix OpinionDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   ostix.ID,
		"type": ostix.Type,
	}

	if len(ostix.Authors) > 0 {
		var strTmp string

		for _, v := range ostix.Authors {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- ReportDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (rstix ReportDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &rstix); err != nil {
		return nil, err
	}

	return rstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (rstix ReportDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(rstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе ReportDomainObjectsSTIX
func (rstix ReportDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(report--)[0-9a-f|-]+$`).MatchString(rstix.ID)) {
		return false
	}

	if !rstix.ValidateStructCommonFields() {
		return false
	}

	//обязательное поле
	if rstix.Name == "" {
		return false
	}

	//обязательное поле
	if len(rstix.ObjectRefs) == 0 {
		return false
	}

	for _, v := range rstix.ObjectRefs {
		if !v.CheckIdentifierTypeSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (rstix ReportDomainObjectsSTIX) SanitizeStruct() ReportDomainObjectsSTIX {
	rstix.CommonPropertiesDomainObjectSTIX = rstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	rstix.Name = commonlibs.StringSanitize(rstix.Name)
	rstix.Description = commonlibs.StringSanitize(rstix.Description)
	if len(rstix.ReportTypes) > 0 {
		r := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(rstix.ReportTypes))
		for _, v := range rstix.ReportTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			r = append(r, tmp)
		}

		rstix.ReportTypes = r
	}

	return rstix
}

// GetID возвращает ID STIX объекта
func (rstix ReportDomainObjectsSTIX) GetID() string {
	return rstix.ID
}

// GetType возвращает Type STIX объекта
func (rstix ReportDomainObjectsSTIX) GetType() string {
	return rstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (rstix ReportDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(rstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(rstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", rstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", rstix.Description))
	str.WriteString(fmt.Sprintf("'report_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'report_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(rstix.ReportTypes)))
	str.WriteString(fmt.Sprintf("'published': '%v'\n", rstix.Published))
	str.WriteString(fmt.Sprintf("'object_refs': \n%v", func(l []stixhelpers.IdentifierTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'object_ref '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(rstix.ObjectRefs)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (rstix ReportDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   rstix.ID,
		"type": rstix.Type,
	}

	if rstix.Name != "" {
		dataForIndex["name"] = rstix.Name
	}

	if rstix.Description != "" {
		dataForIndex["description"] = rstix.Description
	}

	return dataForIndex
}

/* --- ThreatActorDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (tastix ThreatActorDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &tastix); err != nil {
		return nil, err
	}

	return tastix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (tastix ThreatActorDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(tastix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе ThreatActorDomainObjectsSTIX
func (tastix ThreatActorDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(threat-actor--)[0-9a-f|-]+$`).MatchString(tastix.ID)) {
		return false
	}

	if !tastix.ValidateStructCommonFields() {
		return false
	}

	if tastix.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (tastix ThreatActorDomainObjectsSTIX) SanitizeStruct() ThreatActorDomainObjectsSTIX {
	tastix.CommonPropertiesDomainObjectSTIX = tastix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	tastix.Name = commonlibs.StringSanitize(tastix.Name)
	tastix.Description = commonlibs.StringSanitize(tastix.Description)

	if len(tastix.ThreatActorTypes) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(tastix.ThreatActorTypes))
		for _, v := range tastix.ThreatActorTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		tastix.ThreatActorTypes = ta
	}

	if len(tastix.Aliases) > 0 {
		mTmp := make([]string, 0, len(tastix.Aliases))
		for _, v := range tastix.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tastix.Aliases = mTmp
	}

	if len(tastix.Roles) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(tastix.Roles))
		for _, v := range tastix.Roles {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		tastix.Roles = ta
	}

	if len(tastix.Goals) > 0 {
		mTmp := make([]string, 0, len(tastix.Goals))
		for _, v := range tastix.Goals {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tastix.Goals = mTmp
	}

	tastix.Sophistication = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.Sophistication)))
	tastix.ResourceLevel = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.ResourceLevel)))
	tastix.PrimaryMotivation = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(tastix.PrimaryMotivation)))

	if len(tastix.SecondaryMotivations) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(tastix.SecondaryMotivations))
		for _, v := range tastix.SecondaryMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		tastix.SecondaryMotivations = ta
	}

	if len(tastix.PersonalMotivations) > 0 {
		ta := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(tastix.PersonalMotivations))
		for _, v := range tastix.PersonalMotivations {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			ta = append(ta, tmp)
		}

		tastix.PersonalMotivations = ta
	}

	return tastix
}

// GetID возвращает ID STIX объекта
func (tastix ThreatActorDomainObjectsSTIX) GetID() string {
	return tastix.ID
}

// GetType возвращает Type STIX объекта
func (tastix ThreatActorDomainObjectsSTIX) GetType() string {
	return tastix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (tastix ThreatActorDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(tastix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(tastix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", tastix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", tastix.Description))
	str.WriteString(fmt.Sprintf("'threat_actor_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\tt'hreat_actor_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tastix.ThreatActorTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(tastix.Aliases)))
	str.WriteString(fmt.Sprintf("'first_seen': '%v'\n", tastix.FirstSeen))
	str.WriteString(fmt.Sprintf("'last_seen': '%v'\n", tastix.LastSeen))
	str.WriteString(fmt.Sprintf("'roles': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'role '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tastix.Roles)))
	str.WriteString(fmt.Sprintf("'goals': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'goal '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(tastix.Goals)))
	str.WriteString(fmt.Sprintf("'sophistication': '%v'\n", tastix.FirstSeen))
	str.WriteString(fmt.Sprintf("'resource_level': '%v'\n", tastix.LastSeen))
	str.WriteString(fmt.Sprintf("'primary_motivation': '%v'\n", tastix.LastSeen))
	str.WriteString(fmt.Sprintf("'secondary_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'secondary_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tastix.SecondaryMotivations)))
	str.WriteString(fmt.Sprintf("'personal_motivations': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'personal_motivation '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tastix.PersonalMotivations)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (tastix ThreatActorDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   tastix.ID,
		"type": tastix.Type,
	}

	if tastix.Name != "" {
		dataForIndex["name"] = tastix.Name
	}

	if tastix.Description != "" {
		dataForIndex["description"] = tastix.Description
	}

	if len(tastix.Aliases) > 0 {
		var strTmp string

		for _, v := range tastix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	if len(tastix.Goals) > 0 {
		var strTmp string

		for _, v := range tastix.Goals {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] += " " + strTmp
	}

	return dataForIndex
}

/* --- ToolDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (tstix ToolDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &tstix); err != nil {
		return nil, err
	}

	return tstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (tstix ToolDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(tstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе ToolDomainObjectsSTIX
func (tstix ToolDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(tool--)[0-9a-f|-]+$`).MatchString(tstix.ID)) {
		return false
	}

	if !tstix.ValidateStructCommonFields() {
		return false
	}

	if tstix.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (tstix ToolDomainObjectsSTIX) SanitizeStruct() ToolDomainObjectsSTIX {
	tstix.CommonPropertiesDomainObjectSTIX = tstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	tstix.Name = commonlibs.StringSanitize(tstix.Name)
	tstix.Description = commonlibs.StringSanitize(tstix.Description)

	if len(tstix.ToolTypes) > 0 {
		t := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(tstix.ToolTypes))
		for _, v := range tstix.ToolTypes {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			t = append(t, tmp)
		}

		tstix.ToolTypes = t
	}

	if len(tstix.Aliases) > 0 {
		mTmp := make([]string, 0, len(tstix.Aliases))
		for _, v := range tstix.Aliases {
			mTmp = append(mTmp, commonlibs.StringSanitize(v))
		}
		tstix.Aliases = mTmp
	}

	tstix.KillChainPhases = tstix.KillChainPhases.SanitizeStructKillChainPhasesTypeSTIX()
	tstix.ToolVersion = commonlibs.StringSanitize(tstix.ToolVersion)

	return tstix
}

// GetID возвращает ID STIX объекта
func (tstix ToolDomainObjectsSTIX) GetID() string {
	return tstix.ID
}

// GetType возвращает Type STIX объекта
func (tstix ToolDomainObjectsSTIX) GetType() string {
	return tstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (tstix ToolDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(tstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(tstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", tstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", tstix.Description))
	str.WriteString(fmt.Sprintf("'tool_types': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'tool_type '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tstix.ToolTypes)))
	str.WriteString(fmt.Sprintf("'aliases': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'aliase '%d'': '%s'\n", k, v))
		}

		return str.String()
	}(tstix.Aliases)))
	str.WriteString(fmt.Sprintf("'kill_chain_phases': \n%v", func(l []stixhelpers.KillChainPhasesTypeElementSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'kill_chain_phase '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(tstix.KillChainPhases)))
	str.WriteString(fmt.Sprintf("'tool_version': '%s'\n", tstix.ToolVersion))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (tstix ToolDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   tstix.ID,
		"type": tstix.Type,
	}

	if tstix.Name != "" {
		dataForIndex["name"] = tstix.Name
	}

	if tstix.Description != "" {
		dataForIndex["description"] = tstix.Description
	}

	if len(tstix.Aliases) > 0 {
		var strTmp string

		for _, v := range tstix.Aliases {
			strTmp += fmt.Sprintf(" %s", v)
		}

		dataForIndex["aliases"] = strTmp
	}

	return dataForIndex
}

/* --- VulnerabilityDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (vstix VulnerabilityDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &vstix); err != nil {
		return nil, err
	}

	return vstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (vstix VulnerabilityDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(vstix)

	return &result, err
}

// ValidateStruct является валидатором параметров содержащихся в типе VulnerabilityDomainObjectsSTIX
func (vstix VulnerabilityDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(vulnerability--)[0-9a-f|-]+$`).MatchString(vstix.ID)) {
		return false
	}

	if !vstix.ValidateStructCommonFields() {
		return false
	}

	if vstix.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (vstix VulnerabilityDomainObjectsSTIX) SanitizeStruct() VulnerabilityDomainObjectsSTIX {
	vstix.CommonPropertiesDomainObjectSTIX = vstix.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	vstix.Name = commonlibs.StringSanitize(vstix.Name)
	vstix.Description = commonlibs.StringSanitize(vstix.Description)

	return vstix
}

// GetID возвращает ID STIX объекта
func (vstix VulnerabilityDomainObjectsSTIX) GetID() string {
	return vstix.ID
}

// GetType возвращает Type STIX объекта
func (vstix VulnerabilityDomainObjectsSTIX) GetType() string {
	return vstix.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (vstix VulnerabilityDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(vstix.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(vstix.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", vstix.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", vstix.Description))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (vstix VulnerabilityDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   vstix.ID,
		"type": vstix.Type,
	}

	if vstix.Name != "" {
		dataForIndex["name"] = vstix.Name
	}

	if vstix.Description != "" {
		dataForIndex["description"] = vstix.Description
	}

	return dataForIndex
}
