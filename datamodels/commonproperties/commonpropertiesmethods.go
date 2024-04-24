package commonproperties

import "fmt"

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (cp CommonPropertiesObjectSTIX) ToStringBeautiful() string {
	return fmt.Sprintf("'type': '%s'\n'id': '%s'\n", cp.Type, cp.ID)
}
