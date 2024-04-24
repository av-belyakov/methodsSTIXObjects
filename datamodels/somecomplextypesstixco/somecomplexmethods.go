package somecomplextypesstixco

import (
	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

//******* WindowsPEOptionalHeaderTypeSTIX ******//
//**********************************************//

// SanitizeStructWindowsPEOptionalHeaderTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wpeoh WindowsPEOptionalHeaderTypeSTIX) SanitizeStructWindowsPEOptionalHeaderTypeSTIX() WindowsPEOptionalHeaderTypeSTIX {
	wpeoh.MagicHex = commonlibs.StringSanitize(wpeoh.MagicHex)
	wpeoh.Win32VersionValueHex = commonlibs.StringSanitize(wpeoh.Win32VersionValueHex)
	wpeoh.ChecksumHex = commonlibs.StringSanitize(wpeoh.ChecksumHex)
	wpeoh.SubsystemHex = commonlibs.StringSanitize(wpeoh.SubsystemHex)
	wpeoh.DllCharacteristicsHex = commonlibs.StringSanitize(wpeoh.DllCharacteristicsHex)

	hsize := len(wpeoh.Hashes)
	if hsize == 0 {
		return wpeoh
	}

	nhashex := make([]stixhelpers.HashesTypeSTIX, 0, hsize)
	for _, value := range wpeoh.Hashes {
		ht := make(stixhelpers.HashesTypeSTIX, len(value))
		for k, v := range value {
			ht[k] = v
		}
		nhashex = append(nhashex, ht)
	}
	wpeoh.Hashes = nhashex

	return wpeoh
}

//******* WindowsPESectionTypeSTIX ******//
//***************************************//

// SanitizeStructWindowsPESectionTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wpes WindowsPESectionTypeSTIX) SanitizeStructWindowsPESectionTypeSTIX() WindowsPESectionTypeSTIX {
	wpes.Name = commonlibs.StringSanitize(wpes.Name)

	hsize := len(wpes.Hashes)
	if hsize == 0 {
		return wpes
	}

	nh := make(stixhelpers.HashesTypeSTIX, hsize)
	for k, v := range wpes.Hashes {
		nh[k] = commonlibs.StringSanitize(v)
	}
	wpes.Hashes = nh

	return wpes
}

//******* WindowsRegistryValueTypeSTIX ******//
//*******************************************//

// SanitizeStructWindowsRegistryValueTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (wrv WindowsRegistryValueTypeSTIX) SanitizeStructWindowsRegistryValueTypeSTIX() WindowsRegistryValueTypeSTIX {
	wrv.Name = commonlibs.StringSanitize(wrv.Name)
	wrv.Data = commonlibs.StringSanitize(wrv.Data)
	wrv.DataType = stixhelpers.EnumTypeSTIX(commonlibs.StringSanitize(string(wrv.DataType)))

	return wrv
}

//******* X509V3ExtensionsTypeSTIX ******//
//***************************************//

// SanitizeStructX509V3ExtensionsTypeSTIX выполняет замену некоторых специальных символов на их HTML код
func (x509v3e X509V3ExtensionsTypeSTIX) SanitizeStructX509V3ExtensionsTypeSTIX() X509V3ExtensionsTypeSTIX {
	x509v3e.BasicConstraints = commonlibs.StringSanitize(x509v3e.BasicConstraints)
	x509v3e.NameConstraints = commonlibs.StringSanitize(x509v3e.NameConstraints)
	x509v3e.PolicyContraints = commonlibs.StringSanitize(x509v3e.PolicyContraints)
	x509v3e.KeyUsage = commonlibs.StringSanitize(x509v3e.KeyUsage)
	x509v3e.ExtendedKeyUsage = commonlibs.StringSanitize(x509v3e.ExtendedKeyUsage)
	x509v3e.SubjectKeyIdentifier = commonlibs.StringSanitize(x509v3e.SubjectKeyIdentifier)
	x509v3e.AuthorityKeyIdentifier = commonlibs.StringSanitize(x509v3e.AuthorityKeyIdentifier)
	x509v3e.SubjectAlternativeName = commonlibs.StringSanitize(x509v3e.SubjectAlternativeName)
	x509v3e.IssuerAlternativeName = commonlibs.StringSanitize(x509v3e.IssuerAlternativeName)
	x509v3e.SubjectDirectoryAttributes = commonlibs.StringSanitize(x509v3e.SubjectDirectoryAttributes)
	x509v3e.CrlDistributionPoints = commonlibs.StringSanitize(x509v3e.CrlDistributionPoints)
	x509v3e.InhibitAnyPolicy = commonlibs.StringSanitize(x509v3e.InhibitAnyPolicy)
	x509v3e.CertificatePolicies = commonlibs.StringSanitize(x509v3e.CertificatePolicies)
	x509v3e.PolicyMappings = commonlibs.StringSanitize(x509v3e.PolicyMappings)

	return x509v3e
}
