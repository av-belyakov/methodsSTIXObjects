package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
)

/* --- AutonomousSystemCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e AutonomousSystemCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e AutonomousSystemCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Autonomous System", по терминалогии STIX, содержит параметры Автономной системы
// Обязательные значения в полях Number
func (e *AutonomousSystemCyberObservableObjectSTIX) Get() (*AutonomousSystemCyberObservableObjectSTIX, error) {
	if e.GetNumber() == 0 {
		err := fmt.Errorf("the required value 'Number' must not be empty")

		return &AutonomousSystemCyberObservableObjectSTIX{}, err
	}

	return e, nil
}

// -------- NumberObserved property ---------
func (e *AutonomousSystemCyberObservableObjectSTIX) GetNumber() int {
	return e.Number
}

// SetValueNumberObserved устанавливает значение для поля Number
func (e *AutonomousSystemCyberObservableObjectSTIX) SetValueNumber(v int) {
	e.Number = v
}

// SetAnyNumberObserved устанавливает ЛЮБОЕ значение для поля Number
func (e *AutonomousSystemCyberObservableObjectSTIX) SetAnyNumber(i interface{}) {
	e.Number = commonlibs.ConversionAnyToInt(i)
}

// -------- Name property ---------
func (e *AutonomousSystemCyberObservableObjectSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *AutonomousSystemCyberObservableObjectSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *AutonomousSystemCyberObservableObjectSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- RIR property ---------
func (e *AutonomousSystemCyberObservableObjectSTIX) GetRIR() string {
	return e.RIR
}

// SetValueRIR устанавливает значение для поля RIR
func (e *AutonomousSystemCyberObservableObjectSTIX) SetValueRIR(v string) {
	e.RIR = v
}

// SetAnyRIR устанавливает ЛЮБОЕ значение для поля RIR
func (e *AutonomousSystemCyberObservableObjectSTIX) SetAnyRIR(i interface{}) {
	e.RIR = fmt.Sprint(i)
}

// ValidateStruct является валидатором параметров содержащихся в типе AutonomousSystemCyberObservableObjectSTIX
func (e AutonomousSystemCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(autonomous-system--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if e.Name == "" {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e AutonomousSystemCyberObservableObjectSTIX) SanitizeStruct() AutonomousSystemCyberObservableObjectSTIX {
	e.Name = commonlibs.StringSanitize(e.Name)
	e.RIR = commonlibs.StringSanitize(e.RIR)

	return e
}

// GetID возвращает ID STIX объекта
func (e AutonomousSystemCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e AutonomousSystemCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e AutonomousSystemCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'number': '%d'\n", ws, e.Number))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, e.Name))
	str.WriteString(fmt.Sprintf("%s'rir': '%s'\n", ws, e.RIR))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e AutonomousSystemCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Name != "" {
		dataForIndex["name"] = e.Name
	}

	return dataForIndex
}
