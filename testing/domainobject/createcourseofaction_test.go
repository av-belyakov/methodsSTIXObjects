package domainobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCourseOfActionDomainObjectsSTIX(t *testing.T) {
	nca := methodstixobjects.NewCourseOfActionDomainObjectsSTIX()

	assert.Equal(t, nca.GetType(), "course-of-action")
	_, err := nca.Get()
	assert.Error(t, err)

	nca.SetAnyName("course-of-action name")
	_, err = nca.Get()
	assert.NoError(t, err)
	nca.SetValueName("cof name")
	assert.Equal(t, nca.GetName(), "cof name")

	nca.SetAnyDescription("example_description")
	assert.Equal(t, nca.GetDescription(), "example_description")
	nca.SetValueDescription("exm_description")
	assert.Equal(t, nca.GetDescription(), "exm_description")
}

/*
// CourseOfActionDomainObjectsSTIX объект "Course of Action", по терминалогии STIX, описывающий совокупность действий
// направленных на предотвращение (защиту) либо реагирование на текущую атаку
// Name - имя используемое для идентификации "Course of Action" (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// Description - более подробное описание
// Action - ЗАРЕЗЕРВИРОВАНО
type CourseOfActionDomainObjectsSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixdo.CommonPropertiesDomainObjectSTIX
	Name        string      `json:"name" bson:"name" required:"true"`
	Description string      `json:"description" bson:"description"`
	Action      interface{} `json:"action" bson:"action"`
}
*/
