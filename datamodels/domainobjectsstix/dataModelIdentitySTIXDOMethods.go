package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- IdentityDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e IdentityDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e IdentityDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

func (e *IdentityDomainObjectsSTIX) Get() *IdentityDomainObjectsSTIX {
	return e
}

// -------- Name property ---------
func (e *IdentityDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *IdentityDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *IdentityDomainObjectsSTIX) GetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *IdentityDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *IdentityDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *IdentityDomainObjectsSTIX) GetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- ContactInformation property ---------
func (e *IdentityDomainObjectsSTIX) GetContactInformation() string {
	return e.ContactInformation
}

// SetValueContactInformation устанавливает значение для поля ContactInformation
func (e *IdentityDomainObjectsSTIX) SetValueContactInformation(v string) {
	e.ContactInformation = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *IdentityDomainObjectsSTIX) GetAnyContactInformation(i interface{}) {
	e.ContactInformation = fmt.Sprint(i)
}

// -------- Roles property ---------
func (e *IdentityDomainObjectsSTIX) GetRoles() []string {
	return e.Roles
}

// SetValueRoles устанавливает значение для поля Roles
func (e *IdentityDomainObjectsSTIX) SetValueRoles(v string) {
	e.Roles = append(e.Roles, v)
}

// SetAnyRoles устанавливает ЛЮБОЕ значение для поля Roles
func (e *IdentityDomainObjectsSTIX) SetAnyRoles(i interface{}) {
	e.Roles = append(e.Roles, fmt.Sprint(i))
}

// -------- IdentityClass property ---------
func (e *IdentityDomainObjectsSTIX) GetIdentityClass() stixhelpers.OpenVocabTypeSTIX {
	return e.IdentityClass
}

func (e *IdentityDomainObjectsSTIX) SetValueIdentityClass(v stixhelpers.OpenVocabTypeSTIX) {
	e.IdentityClass = v
}

// -------- ObjectRefs property ---------
func (e *IdentityDomainObjectsSTIX) GetSectors() []stixhelpers.OpenVocabTypeSTIX {
	return e.Sectors
}

func (e *IdentityDomainObjectsSTIX) SetValueSectors(v []stixhelpers.OpenVocabTypeSTIX) {
	e.Sectors = v
}

// ValidateStruct является валидатором параметров содержащихся в типе IdentityDomainObjectsSTIX
func (e IdentityDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(identity--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	//обязательное поле
	if e.Name == "" {
		return false
	}

	return e.ValidateStructCommonFields()
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e IdentityDomainObjectsSTIX) SanitizeStruct() IdentityDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.SanitizeStruct().CommonPropertiesDomainObjectSTIX

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)

	if len(e.Roles) > 0 {
		rolesTmp := make([]string, 0, len(e.Roles))
		for _, v := range e.Roles {
			rolesTmp = append(rolesTmp, commonlibs.StringSanitize(v))
		}

		e.Roles = rolesTmp
	}

	e.IdentityClass = e.IdentityClass.SanitizeStructOpenVocabTypeSTIX()

	if len(e.Sectors) > 0 {
		sectorsTmp := make([]stixhelpers.OpenVocabTypeSTIX, 0, len(e.Sectors))
		for _, v := range e.Sectors {
			tmp := v.SanitizeStructOpenVocabTypeSTIX()
			sectorsTmp = append(sectorsTmp, tmp)
		}

		e.Sectors = sectorsTmp
	}

	e.ContactInformation = commonlibs.StringSanitize(e.ContactInformation)

	return e
}

// GetID возвращает ID STIX объекта
func (istix IdentityDomainObjectsSTIX) GetID() string {
	return istix.ID
}

// GetType возвращает Type STIX объекта
func (e IdentityDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e IdentityDomainObjectsSTIX) ToStringBeautiful() string {
	str := strings.Builder{}

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful())
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful())
	str.WriteString(fmt.Sprintf("'name': '%s'\n", e.Name))
	str.WriteString(fmt.Sprintf("'description': '%s'\n", e.Description))
	str.WriteString(fmt.Sprintf("'roles': \n%v", func(l []string) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'role '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.Roles)))
	str.WriteString(fmt.Sprintf("'identity_class': '%s'\n", e.IdentityClass))
	str.WriteString(fmt.Sprintf("'sectors': \n%v", func(l []stixhelpers.OpenVocabTypeSTIX) string {
		str := strings.Builder{}

		for k, v := range l {
			str.WriteString(fmt.Sprintf("\t'sector '%d'': '%v'\n", k, v))
		}

		return str.String()
	}(e.Sectors)))
	str.WriteString(fmt.Sprintf("'contact_information': '%s'\n", e.ContactInformation))

	return str.String()
}
