package commonproperties

import "fmt"

func (e *CommonPropertiesObjectSTIX) Get() *CommonPropertiesObjectSTIX {
	return e
}

// -------- ID property ---------
func (e *CommonPropertiesObjectSTIX) GetID() string {
	return e.ID
}

// SetValueID устанавливает значение для поля ID
func (e *CommonPropertiesObjectSTIX) SetValueID(v string) {
	e.ID = v
}

// SetAnyID устанавливает ЛЮБОЕ значение для поля ID
func (e *CommonPropertiesObjectSTIX) SetAnyID(i interface{}) {
	e.ID = fmt.Sprint(i)
}

// -------- Type property ---------
func (e *CommonPropertiesObjectSTIX) GetType() string {
	return e.Type
}

// SetValueType устанавливает значение для поля Type
func (e *CommonPropertiesObjectSTIX) SetValueType(v string) {
	e.Type = v
}

// SetAnyType устанавливает ЛЮБОЕ значение для поля Type
func (e *CommonPropertiesObjectSTIX) SetAnyType(i interface{}) {
	e.Type = fmt.Sprint(i)
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e CommonPropertiesObjectSTIX) ToStringBeautiful() string {
	return fmt.Sprintf("'type': '%s'\n'id': '%s'\n", e.Type, e.ID)
}
