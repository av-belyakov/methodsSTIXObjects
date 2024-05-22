package stixhelpers

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/**********			 Некоторые примитивные типы STIX			 **********/

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

// SetValue добавляет значение в тип IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) SetValue(v string) {
	var str IdentifierTypeSTIX = IdentifierTypeSTIX(v)
	itstix = &str
}

// SetAny добавляет значение в тип IdentifierTypeSTIX
func (itstix *IdentifierTypeSTIX) SetAny(i interface{}) {
	var e IdentifierTypeSTIX = IdentifierTypeSTIX(fmt.Sprint(i))
	itstix = &e
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
		dtstix.Dictionary = str

		return nil
	}
	if err = json.Unmarshal(data, &num); err == nil {
		dtstix.Dictionary = num
		return nil
	}
	if err = json.Unmarshal(data, &listStr); err == nil {
		dtstix.Dictionary = listStr
		return nil
	}
	if err = json.Unmarshal(data, &listInt); err == nil {
		dtstix.Dictionary = listInt
		return nil
	}
	if err = json.Unmarshal(data, &mapStrStr); err == nil {
		dtstix.Dictionary = mapStrStr
		return nil
	}
	if err = json.Unmarshal(data, &mapStrInt); err == nil {
		dtstix.Dictionary = mapStrInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntInt); err == nil {
		dtstix.Dictionary = mapIntInt
		return nil
	}
	if err = json.Unmarshal(data, &mapIntStr); err == nil {
		dtstix.Dictionary = mapIntStr
		return nil
	}

	dtstix.Dictionary = data

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
