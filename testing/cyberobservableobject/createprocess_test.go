package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestProcessCyberObservableObjectSTIX(t *testing.T) {
	np := methodstixobjects.NewProcessCyberObservableObjectSTIX()

	assert.Equal(t, np.GetType(), "process")
	_, err := np.Get()
	assert.NoError(t, err)

	np.SetAnyIsHidden(true)
	assert.True(t, np.GetIsHidden())
	np.SetValueIsHidden(false)
	assert.False(t, np.GetIsHidden())

	np.SetAnyPID(3641)
	assert.Equal(t, np.GetPID(), 3641)
	np.SetValuePID(31)
	assert.Equal(t, np.GetPID(), 31)

	cwde := "/nedv/btgn"
	np.SetAnyCwd(cwde)
	assert.Equal(t, np.GetCwd(), cwde)
	np.SetValueCwd("/vc/as/fd")
	assert.Equal(t, np.GetCwd(), "/vc/as/fd")

	cle := "command line example"
	np.SetAnyCommandLine(cle)
	assert.Equal(t, np.GetCommandLine(), cle)
	np.SetValueCommandLine("n12")
	assert.Equal(t, np.GetCommandLine(), "n12")

	//--- Mtime
	ct := "2024-02-14T12:03:06+00:00"
	err = np.SetAnyCreatedTime(ct)
	assert.NoError(t, err)
	assert.Equal(t, np.GetCreatedTime(), ct)

	err = np.SetAnyCreatedTime(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, np.GetCreatedTime(), "2024-05-28T12:33:44+03:00")

	cure := stixhelpers.IdentifierTypeSTIX("creator user ref example")
	np.SetValueCreatorUserRef(cure)
	assert.Equal(t, np.GetCreatorUserRef(), cure)

	iure := stixhelpers.IdentifierTypeSTIX("image user ref example")
	np.SetValueImageRef(iure)
	assert.Equal(t, np.GetImageRef(), iure)

	pure := stixhelpers.IdentifierTypeSTIX("parent user ref example")
	np.SetValueParentRef(pure)
	assert.Equal(t, np.GetParentRef(), pure)

	np.SetValueChildRefs([]stixhelpers.IdentifierTypeSTIX{
		"child_ref_1",
		"child_ref_2",
		"child_ref_3",
	})
	assert.Equal(t, len(np.GetChildRefs()), 3)

	np.SetValueOpenedConnectionRefs([]stixhelpers.IdentifierTypeSTIX{
		"oc_ref_1",
		"oc_ref_2",
	})
	assert.Equal(t, len(np.GetOpenedConnectionRefs()), 2)

	np.SetAnyEnvironmentVariables("env_key_1", "env_11")
	np.SetAnyEnvironmentVariables("env_key_2", "env_22")
	assert.Equal(t, len(np.GetEnvironmentVariables()), 2)

	np.SetAnyExtensions("ext_1", 123)
	np.SetAnyExtensions("ext_2", 234)
	np.SetAnyExtensions("ext_3", 345)
	assert.Equal(t, len(np.GetExtensions()), 3)
}

/*
// ProcessCyberObservableObjectSTIX объект "Process Object", по терминологии STIX, содержит общие свойства экземпляра компьютерной программы,
// выполняемой в операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно свойство (отличное от типа) этого объекта (или одного из его расширений).
// Extensions - определяет расширения windows-process-exit или windows-service-ext. В дополнение к ним производители МОГУТ создавать свои собственные.
// ключи словаря windows-process-exit, windows-service-ext ДОЛЖНЫ идентифицировать тип расширения по имени. Соответствующие значения словаря ДОЛЖНЫ
// содержать содержимое экземпляра расширения.
// IsHidden - определяет является ли процесс скрытым.
// PID - униальный идентификатор процесса.
// CreatedTime - время, в формате "2016-05-12T08:17:27.000Z", создания процесса.
// Cwd - текущая рабочая директория процесса.
// CommandLine - поределяет полный перечень команд используемых для запуска процесса, включая имя процесса и аргументы.
// EnvironmentVariables - определяет список переменных окружения, в виде словаря, ассоциируемых с приложением.
// OpenedConnectionRefs - определяет список открытых, процессом, сетевых соединений ка одну или более ссылку на объект типа network-traffic.
// CreatorUserRef - определяет что за пользователь создал объект, ссылка ДОЛЖНА ссылатся на объект типа user-account.
// ImageRef - указывает исполняемый двоичный файл, который был выполнен как образ процесса, как ссылка на файловый объект. Объект, на который ссылается
// это свойство, ДОЛЖЕН иметь тип file.
// ParentRef - указывает другой процесс, который породил (т. е. является родителем) этот процесс, как ссылку на объект процесса. Объект, на который
// ссылается это свойство, ДОЛЖЕН иметь тип process.
// ChildRefs - указывает другие процессы, которые были порождены (т. е. дочерние) этим процессом, в качестве ссылки на один или несколько других
// объектов процесса. Объекты, на которые ссылается этот список, ДОЛЖНЫ иметь тип process.
type ProcessCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsHidden             bool                             `json:"is_hidden" bson:"is_hidden"`
	PID                  int                              `json:"pid" bson:"pid"`
	Cwd                  string                           `json:"cwd" bson:"cwd"`
	CommandLine          string                           `json:"command_line" bson:"command_line"`
	CreatedTime          string                           `json:"created_time" bson:"created_time"`
	CreatorUserRef       stixhelpers.IdentifierTypeSTIX   `json:"creator_user_ref" bson:"creator_user_ref"`
	ImageRef             stixhelpers.IdentifierTypeSTIX   `json:"image_ref" bson:"image_ref"`
	ParentRef            stixhelpers.IdentifierTypeSTIX   `json:"parent_ref" bson:"parent_ref"`
	ChildRefs            []stixhelpers.IdentifierTypeSTIX `json:"child_refs" bson:"child_refs"`
	OpenedConnectionRefs []stixhelpers.IdentifierTypeSTIX `json:"opened_connection_refs" bson:"opened_connection_refs"`
	EnvironmentVariables map[string]string                `json:"environment_variables" bson:"environment_variables"`
	Extensions           map[string]interface{}           `json:"extensions" bson:"extensions"`
}
*/
