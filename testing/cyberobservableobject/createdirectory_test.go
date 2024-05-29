package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestDirectoryCyberObservableObjectSTIX(t *testing.T) {
	nd := methodstixobjects.NewDirectoryCyberObservableObjectSTIX()

	assert.Equal(t, nd.GetType(), "directory")
	_, err := nd.Get()
	assert.Error(t, err)

	p := "/opt/path/example"
	nd.SetAnyPath(p)
	_, err = nd.Get()
	assert.NoError(t, err)
	assert.Equal(t, nd.GetPath(), p)
	nd.SetValuePath("/pa")
	assert.Equal(t, nd.GetPath(), "/pa")

	nd.SetAnyPathEnc("utf8")
	assert.Equal(t, nd.GetPathEnc(), "utf8")
	nd.SetValuePathEnc("win-1258")
	assert.Equal(t, nd.GetPathEnc(), "win-1258")

	ct := "2024-02-11T07:01:01+00:00"
	err = nd.SetAnyCtime(ct)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetCtime(), ct)

	err = nd.SetAnyCtime(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetCtime(), "2024-05-22T11:43:27+03:00")

	mt := "2024-02-10T11:21:01+00:00"
	err = nd.SetAnyMtime(mt)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetMtime(), mt)

	err = nd.SetAnyMtime(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetMtime(), "2024-05-22T11:43:27+03:00")

	at := "2024-02-10T11:21:01+00:00"
	err = nd.SetAnyAtime(at)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetAtime(), at)

	err = nd.SetAnyAtime(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nd.GetAtime(), "2024-05-22T11:43:27+03:00")

	nd.SetValueContainsRefs(stixhelpers.IdentifierTypeSTIX("000q"))
	assert.Equal(t, len(nd.GetContainsRefs()), 1)
	nd.SetFullValueContainsRefs([]stixhelpers.IdentifierTypeSTIX{"111a", "222b", "333c"})
	assert.Equal(t, len(nd.GetContainsRefs()), 3)
}

/*
// DirectoryCyberObservableObjectSTIX объект "Directory", по терминалогии STIX, содержит свойства, общие для каталога файловой системы
// Path - указывает путь, как было первоначально замечено, к каталогу в файловой системе (ОБЯЗАТЕЛЬНОЕ ЗНАЧЕНИЕ)
// PathEnc - указывает наблюдаемую кодировку для пути. Значение ДОЛЖНО быть указано, если путь хранится в кодировке, отличной от Unicode.
// Ctime - время, в формате "2016-05-12T08:17:27.000Z", создания директории
// Mtime - время, в формате "2016-05-12T08:17:27.000Z", модификации или записи в директорию
// Atime - время, в формате "2016-05-12T08:17:27.000Z", последнего обращения к директории
// ContainsRefs - содержит список файловых объектов или директорий содержащихся внутри директории
type DirectoryCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Path         string                           `json:"path" bson:"path" required:"true"`
	PathEnc      string                           `json:"path_enc" bson:"path_enc"`
	Ctime        string                           `json:"ctime" bson:"ctime"`
	Mtime        string                           `json:"mtime" bson:"mtime"`
	Atime        string                           `json:"atime" bson:"atime"`
	ContainsRefs []stixhelpers.IdentifierTypeSTIX `json:"contains_refs" bson:"contains_refs"`
}
*/
