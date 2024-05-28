package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestFileCyberObservableObjectSTIX(t *testing.T) {
	nf := methodstixobjects.NewFileCyberObservableObjectSTIX()

	assert.Equal(t, nf.GetType(), "file")
	_, err := nf.Get()
	assert.NoError(t, err)

	se := 162734
	nf.SetAnySize(se)
	assert.Equal(t, nf.GetSize(), uint64(se))
	nf.SetValueSize(13)
	assert.Equal(t, nf.GetSize(), uint64(13))

	ne := "any name example"
	nf.SetAnyName(ne)
	assert.Equal(t, nf.GetName(), ne)
	nf.SetValueName("name")
	assert.Equal(t, nf.GetName(), "name")

	nee := "any name enc example"
	nf.SetAnyNameEnc(nee)
	assert.Equal(t, nf.GetNameEnc(), nee)
	nf.SetValueNameEnc("ne")
	assert.Equal(t, nf.GetNameEnc(), "ne")

	mnhe := "magic numberhex example"
	nf.SetAnyMagicNumberHex(mnhe)
	assert.Equal(t, nf.GetMagicNumberHex(), mnhe)
	nf.SetValueMagicNumberHex("magic")
	assert.Equal(t, nf.GetMagicNumberHex(), "magic")

	mte := "mime type example"
	nf.SetAnyMimeType(mte)
	assert.Equal(t, nf.GetMimeType(), mte)
	nf.SetValueMimeType("mime type")
	assert.Equal(t, nf.GetMimeType(), "mime type")

	//--- Ctime
	ct := "2024-02-11T07:01:01+00:00"
	err = nf.SetAnyCtime(ct)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetCtime(), ct)

	err = nf.SetAnyCtime(1716367407123)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetCtime(), "2024-05-22T11:43:27+03:00")

	//--- Mtime
	mt := "2024-02-14T12:03:06+00:00"
	err = nf.SetAnyMtime(mt)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetMtime(), mt)

	err = nf.SetAnyCtime(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetCtime(), "2024-05-28T12:33:44+03:00")

	//--- Atime
	at := "2024-02-18T21:11:11+00:00"
	err = nf.SetAnyCtime(at)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetCtime(), at)

	err = nf.SetAnyCtime(1716888850190)
	assert.NoError(t, err)
	assert.Equal(t, nf.GetCtime(), "2024-05-28T12:34:10+03:00")

	pdre := stixhelpers.IdentifierTypeSTIX("parent dirextory ref example")
	nf.SetValueParentDirectoryRef(pdre)
	assert.Equal(t, nf.GetParentDirectoryRef(), pdre)

	cre := stixhelpers.IdentifierTypeSTIX("content ref example")
	nf.SetValueContentRef(cre)
	assert.Equal(t, nf.GetContentRef(), cre)

	nf.SetValueHashes(stixhelpers.HashesTypeSTIX{
		"hash_1": "77qer726r7387rt8w72e7r37321783t8r3",
		"hash_2": "eab43438883bvc840487384y3f3f73rvv3",
	})
	assert.Equal(t, len(nf.GetHashes()), 2)

	nf.SetValueContainsRefs([]stixhelpers.IdentifierTypeSTIX{
		"contains_ref_1",
		"contains_ref_2",
		"contains_ref_3",
	})
	assert.Equal(t, len(nf.GetContainsRefs()), 3)

	nf.SetValueExtensions(map[string]interface{}{
		"ext_1": "zxcvbbb",
		"ext_2": 123,
	})
	assert.Equal(t, len(nf.GetExtensions()), 2)
}
