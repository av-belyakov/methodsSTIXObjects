package stixhelpers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/**********			 Некоторые примитивные типы STIX			 **********/

// CheckExternalReferencesTypeSTIX выполняет проверку значений типа ExternalReferencesTypeSTIX
func (ertstix *ExternalReferencesTypeSTIX) CheckExternalReferencesTypeSTIX() bool {
	if len(*ertstix) == 0 {
		return true
	}

	for _, v := range *ertstix {
		if !v.CheckExternalReferenceTypeElementSTIX() {
			return false
		}
	}

	return true
}

// SanitizeStructExternalReferencesTypeSTIX для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (ertstix ExternalReferencesTypeSTIX) SanitizeStructExternalReferencesTypeSTIX() ExternalReferencesTypeSTIX {
	size := len(ertstix)
	if size == 0 {
		return ertstix
	}

	ert := make([]ExternalReferenceTypeElementSTIX, 0, size)
	for _, v := range ertstix {
		tmp := v.SanitizeStructExternalReferenceTypeElementSTIX()
		ert = append(ert, tmp)
	}

	return ert
}

// CheckExternalReferenceTypeElementSTIX выполняет проверку значений типа ExternalReferenceTypeElementSTIX
func (ertestix *ExternalReferenceTypeElementSTIX) CheckExternalReferenceTypeElementSTIX() bool {
	if ertestix.URL != "" && !govalidator.IsURL(ertestix.URL) {
		return false
	}

	return true
}

// SanitizeStructExternalReferenceTypeElementSTIX выполняет проверку значений типа ExternalReferenceTypeElementSTIX
func (ertestix ExternalReferenceTypeElementSTIX) SanitizeStructExternalReferenceTypeElementSTIX() ExternalReferenceTypeElementSTIX {
	return ExternalReferenceTypeElementSTIX{
		SourceName:  commonlibs.StringSanitize(ertestix.SourceName),
		Description: commonlibs.StringSanitize(ertestix.Description),
		URL:         ertestix.URL,
		Hashes:      ertestix.Hashes,
		ExternalID:  commonlibs.StringSanitize(ertestix.ExternalID),
	}
}

// CheckHashesTypeSTIX выполняет проверку значений типа HashesTypeSTIX
func (htstix *HashesTypeSTIX) CheckHashesTypeSTIX() bool {
	if len(*htstix) == 0 {
		return true
	}

	pattern := regexp.MustCompile(`^[0-9a-zA-Z-_=]+$`)
	for k, v := range *htstix {
		if !pattern.MatchString(k) {
			return false
		}

		if !pattern.MatchString(v) {
			return false
		}
	}

	return true
}

// CheckIdentifierTypeSTIX выполняет проверку значения типа IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) CheckIdentifierTypeSTIX() bool {
	if len(*itstix) == 0 {
		return true
	}

	return regexp.MustCompile(`^[0-9a-zA-Z-_]+(--)[0-9a-f-]+$`).MatchString(fmt.Sprint(*itstix))
}

// AddValue добавляет значение в тип IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) AddValue(str string) {
	var i IdentifierTypeSTIX = IdentifierTypeSTIX(str)
	itstix = &i
}

// SanitizeStructKillChainPhasesTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (kcptstix KillChainPhasesTypeSTIX) SanitizeStructKillChainPhasesTypeSTIX() KillChainPhasesTypeSTIX {
	for k := range kcptstix {
		kcptstix[k].SanitizeStructKillChainPhasesTypeElementSTIX()
	}

	return kcptstix
}

// SanitizeStructKillChainPhasesTypeElementSTIX выполлняет проверку значения типа KillChainPhasesTypeElementSTIX
func (kcptestix KillChainPhasesTypeElementSTIX) SanitizeStructKillChainPhasesTypeElementSTIX() KillChainPhasesTypeElementSTIX {
	if kcptestix.KillChainName == "" {
		return kcptestix
	}

	if kcptestix.PhaseName == "" {
		return kcptestix
	}

	kcptestix.KillChainName = commonlibs.StringSanitize(kcptestix.KillChainName)
	kcptestix.PhaseName = commonlibs.StringSanitize(kcptestix.PhaseName)

	return kcptestix
}

// SanitizeStructOpenVocabTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (ovtstix OpenVocabTypeSTIX) SanitizeStructOpenVocabTypeSTIX() OpenVocabTypeSTIX {
	return OpenVocabTypeSTIX(commonlibs.StringSanitize(fmt.Sprint(ovtstix)))
}

// UnmarshalJSON дополнительный метод декодирования
func (dtstix *DictionaryTypeSTIX) UnmarshalJSON(data []byte) error {
	var (
		err       error
		str       string
		num       int
		listStr   []string
		listInt   []int
		mapStrStr map[string]string
		mapStrInt map[string]int
		mapIntInt map[int]int
		mapIntStr map[int]string
	)

	if err = json.Unmarshal(data, &str); err == nil {
		dtstix.dictionary = str

		return nil
	}
	if err = json.Unmarshal(data, &num); err == nil {
		dtstix.dictionary = num
		return nil
	}
	if err = json.Unmarshal(data, &listStr); err == nil {
		dtstix.dictionary = listStr
		return nil
	}
	if err = json.Unmarshal(data, &listInt); err == nil {
		dtstix.dictionary = listInt
		return nil
	}
	if err = json.Unmarshal(data, &mapStrStr); err == nil {
		dtstix.dictionary = mapStrStr
		return nil
	}
	if err = json.Unmarshal(data, &mapStrInt); err == nil {
		dtstix.dictionary = mapStrInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntInt); err == nil {
		dtstix.dictionary = mapIntInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntStr); err == nil {
		dtstix.dictionary = mapIntStr
		return nil
	}
	//		dtstix.dictionary = fmt.Sprintln(mapIntStr)
	dtstix.dictionary = data

	return fmt.Errorf("JSON message parsing error, undefined value found in the DictionaryTypeSTIX type")
}

func (gmtstix *GranularMarkingsTypeSTIX) CheckGranularMarkingsTypeSTIX() bool {
	//для поля Lang
	if gmtstix.Lang != "" && !(regexp.MustCompile(`^[a-zA-Z]+$`)).MatchString(gmtstix.Lang) {
		return false
	}

	if !gmtstix.MarkingRef.CheckIdentifierTypeSTIX() {
		return false
	}

	selectorTmp := make([]string, 0, len(gmtstix.Selectors))
	for _, v := range gmtstix.Selectors {
		selectorTmp = append(selectorTmp, commonlibs.StringSanitize(v))
	}
	gmtstix.Selectors = selectorTmp

	return true
}
