package cyberobservableobject

import (
	"testing"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/stretchr/testify/assert"
)

func TestUserAccountCyberObservableObjectSTIX(t *testing.T) {
	nua := methodstixobjects.NewUserAccountCyberObservableObjectSTIX()

	assert.Equal(t, nua.GetType(), "user-account")
	_, err := nua.Get()
	assert.NoError(t, err)

	nua.SetAnyIsServiceAccount(true)
	assert.True(t, nua.GetIsServiceAccount())
	nua.SetValueIsServiceAccount(false)
	assert.False(t, nua.GetIsServiceAccount())

	nua.SetAnyIsPrivileged(true)
	assert.True(t, nua.GetIsPrivileged())
	nua.SetValueIsPrivileged(false)
	assert.False(t, nua.GetIsPrivileged())

	nua.SetAnyIsDisabled(true)
	assert.True(t, nua.GetIsDisabled())
	nua.SetValueIsDisabled(false)
	assert.False(t, nua.GetIsDisabled())

	nua.SetAnyCanEscalatePrivs(true)
	assert.True(t, nua.GetCanEscalatePrivs())
	nua.SetValueCanEscalatePrivs(false)
	assert.False(t, nua.GetCanEscalatePrivs())

	nua.SetAnyUserID("uid_yd7w73r5")
	assert.Equal(t, nua.GetUserID(), "uid_yd7w73r5")
	nua.SetValueUserID("uid_0990")
	assert.Equal(t, nua.GetUserID(), "uid_0990")

	nua.SetAnyCredential("credential_example")
	assert.Equal(t, nua.GetCredential(), "credential_example")
	nua.SetValueCredential("cre example")
	assert.Equal(t, nua.GetCredential(), "cre example")

	nua.SetAnyAccountLogin("user-login")
	assert.Equal(t, nua.GetAccountLogin(), "user-login")
	nua.SetValueAccountLogin("login_user")
	assert.Equal(t, nua.GetAccountLogin(), "login_user")

	nua.SetAnyDisplayName("user_name")
	assert.Equal(t, nua.GetDisplayName(), "user_name")
	nua.SetValueDisplayName("u_name")
	assert.Equal(t, nua.GetDisplayName(), "u_name")

	// --- AccountCreated
	act := "2024-02-14T12:03:06+00:00"
	err = nua.SetAnyAccountCreated(act)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountCreated(), act)

	err = nua.SetAnyAccountCreated(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountCreated(), "2024-05-28T12:33:44+03:00")

	// --- AccountExpires
	aet := "2024-02-15T20:10:16+00:00"
	err = nua.SetAnyAccountExpires(aet)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountExpires(), aet)

	err = nua.SetAnyAccountExpires(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountExpires(), "2024-05-28T12:33:44+03:00")

	// --- CredentialLastChanged
	clt := "2024-02-14T15:53:56+00:00"
	err = nua.SetAnyCredentialLastChanged(clt)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetCredentialLastChanged(), clt)

	err = nua.SetAnyCredentialLastChanged(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetCredentialLastChanged(), "2024-05-28T12:33:44+03:00")

	// --- AccountFirstLogin
	aflt := "2024-02-14T16:13:46+00:00"
	err = nua.SetAnyAccountFirstLogin(aflt)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountFirstLogin(), aflt)

	err = nua.SetAnyAccountFirstLogin(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountFirstLogin(), "2024-05-28T12:33:44+03:00")

	// --- AccountLastLogin
	allt := "2024-02-14T16:21:51+00:00"
	err = nua.SetAnyAccountLastLogin(allt)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetAccountLastLogin(), allt)

	err = nua.SetAnyCredentialLastChanged(1716888824134)
	assert.NoError(t, err)
	assert.Equal(t, nua.GetCredentialLastChanged(), "2024-05-28T12:33:44+03:00")

	nua.SetValueAccountType(stixhelpers.OpenVocabTypeSTIX("just user"))
	assert.Equal(t, nua.GetAccountType(), stixhelpers.OpenVocabTypeSTIX("just user"))

	nua.SetFullValueExtensions(map[string]someextensionsstixco.UNIXAccountExtensionSTIX{
		"key_0": {GID: 734},
	})
	assert.Equal(t, len(nua.GetExtensions()), 1)
	nua.SetValueExtensions("key_1", someextensionsstixco.UNIXAccountExtensionSTIX{GID: 3215, HomeDir: "/home/user_1"})
	nua.SetValueExtensions("key_2", someextensionsstixco.UNIXAccountExtensionSTIX{GID: 2241, HomeDir: "/home/user_2"})
	assert.Equal(t, len(nua.GetExtensions()), 3)
}

/*
// UserAccountCyberObservableObjectSTIX объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя, включая,
// учетные записи операционной системы, устройства, службы обмена сообщениями и платформы социальных сетей и других прочих учетных записей
// Поскольку все свойства этого объекта являются необязательными, по крайней мере одно из свойств, определенных ниже, ДОЛЖНО быть инициализировано
// при использовании этого объекта
// Extensions - содержит словарь расширяющий тип "User Account Object" одно из расширений "unix-account-ext", реализуется описанным ниже
// типом, UNIXAccountExtensionSTIX кроме этого производитель может созавать свои собственные типы расширений. Ключи данного словаря
// идентифицируют тип расширения по имени, значения являются содержимым экземпляра расширения
// UserID - содержит идентификатор учетной записи. Формат идентификатора зависит от системы в которой находится данная учетная запись пользователя,
// и может быть числовым идентификатором, идентификатором GUID, именем учетной записи, адресом электронной почты и т.д. Свойство  UserId должно
// быть заполнено любым значанием, являющимся уникальным идентификатором системы, членом которой является учетная запись. Например, в системах UNIX он
// будет заполнено значением UID
// Credential - содержит учетные данные пользователя в открытом виде. Предназначено только для закрытого применения при изучении метаданных вредоносных
// программ при их исследовании (например, жестко закодированный пароль администратора домена, который вредоносная программа пытается использовать
// реализации тактики для бокового (латерального) перемещения) и не должно применяться для совместного пользования PII
// AccountLogin - содержит логин пользователя. Используется в тех случаях,когда свойство UserId указывает другие данные, чем то, что пользователь вводит
// при входе в систему
// AccountType - содержит одно, из заранее определенных (предложенных) значений. Является типом аккаунта. Значения этого свойства берутся из множества
// закрепленного в открытом словаре, account-type-ov
// DisplayName - содержит отображаемое имя учетной записи, которое будет отображаться в пользовательских интерфейсах. В Unix, это равносильно полю
// gecos (gecos это поле учётной записи пользователя в файле /etc/passwd )
// IsServiceAccount - содержит индикатор, сигнализирующий что, учетная запись связана с сетевой службой или системным процессом (демоном), а не с
// конкретным человеком. (системный пользователь)
// IsPrivileged - содержит индикатор, сигнализирующий что, учетная запись имеет повышенные привилегии (например, в случае root в Unix или учетной
// записи администратора Windows)
// CanEscalatePrivs  - содержит индикатор, сигнализирующий что, учетная запись имеет возможность повышать привилегии (например, в случае sudo в
// Unix или учетной записи администратора домена Windows)
// IsDisabled  - содержит индикатор, сигнализирующий что, учетная запись отключена
// AccountCreated - время, в формате "2016-05-12T08:17:27.000Z", создания аккаунта
// AccountExpires - время, в формате "2016-05-12T08:17:27.000Z", истечения срока действия учетной записи.
// CredentialLastChanged - время, в формате "2016-05-12T08:17:27.000Z", когда учетные данные учетной записи были изменены в последний раз.
// AccountFirstLogin - время, в формате "2016-05-12T08:17:27.000Z", первого доступа к учетной записи
// AccountLastLogin - время, в формате "2016-05-12T08:17:27.000Z", когда к учетной записи был последний доступ.
type UserAccountCyberObservableObjectSTIX struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	IsServiceAccount      bool                                                     `json:"is_service_account" bson:"is_service_account"`
	IsPrivileged          bool                                                     `json:"is_privileged" bson:"is_privileged"`
	CanEscalatePrivs      bool                                                     `json:"can_escalate_privs" bson:"can_escalate_privs"`
	IsDisabled            bool                                                     `json:"is_disabled" bson:"is_disabled"`
	UserID                string                                                   `json:"user_id" bson:"user_id"`
	Credential            string                                                   `json:"credential" bson:"credential"`
	AccountLogin          string                                                   `json:"account_login" bson:"account_login"`
	DisplayName           string                                                   `json:"display_name" bson:"display_name"`
	AccountCreated        string                                                   `json:"account_created" bson:"account_created"`
	AccountExpires        string                                                   `json:"account_expires" bson:"account_expires"`
	CredentialLastChanged string                                                   `json:"credential_last_changed" bson:"credential_last_changed"`
	AccountFirstLogin     string                                                   `json:"account_first_login" bson:"account_first_login"`
	AccountLastLogin      string                                                   `json:"account_last_login" bson:"account_last_login"`
	AccountType           stixhelpers.OpenVocabTypeSTIX                            `json:"account_type" bson:"account_type"`
	Extensions            map[string]someextensionsstixco.UNIXAccountExtensionSTIX `json:"extensions" bson:"extensions"`
}
*/
