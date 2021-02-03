/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"encoding/pem"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// OrganizationalUnitIdentifiersConfiguration is used to represent an OU
// and an associated trusted certificate
type OrganizationalUnitIdentifiersConfiguration struct {
	// Certificate is the path to a root or intermediate certificate
	Certificate string `yaml:"Certificate,omitempty"`
	// OrganizationalUnitIdentifier is the name of the OU
	OrganizationalUnitIdentifier string `yaml:"OrganizationalUnitIdentifier,omitempty"`
}

// NodeOUs contains information on how to tell apart clients, peers and orderers
// based on OUs. If the check is enforced, by setting Enabled to true,
// the MSP will consider an identity valid if it is an identity of a client, a peer or
// an orderer. An identity should have only one of these special OUs.
type NodeOUs struct {
	// Enable activates the OU enforcement
	Enable bool `yaml:"Enable,omitempty"`
	// ClientOUIdentifier specifies how to recognize clients by OU
	ClientOUIdentifier *OrganizationalUnitIdentifiersConfiguration `yaml:"ClientOUIdentifier,omitempty"`
	// PeerOUIdentifier specifies how to recognize peers by OU
	PeerOUIdentifier *OrganizationalUnitIdentifiersConfiguration `yaml:"PeerOUIdentifier,omitempty"`
}

// Configuration represents the accessory configuration an MSP can be equipped with.
// By default, this configuration is stored in a yaml file
type Configuration struct {
	// OrganizationalUnitIdentifiers is a list of OUs. If this is set, the MSP
	// will consider an identity valid only it contains at least one of these OUs
	OrganizationalUnitIdentifiers []*OrganizationalUnitIdentifiersConfiguration `yaml:"OrganizationalUnitIdentifiers,omitempty"`
	// NodeOUs enables the MSP to tell apart clients, peers and orderers based
	// on the identity's OU.
	NodeOUs *NodeOUs `yaml:"NodeOUs,omitempty"`
}

func readFile(file string) ([]byte, error) {
	fileCont, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read file %s", file)
	}

	return fileCont, nil
}

func readPemFile(file string) ([]byte, error) {
	bytes, err := readFile(file)
	if err != nil {
		return nil, errors.Wrapf(err, "reading from file %s failed", file)
	}

	b, _ := pem.Decode(bytes)
	if b == nil { // TODO: also check that the type is what we expect (cert vs key..)
		return nil, errors.Errorf("no pem content for file %s", file)
	}

	return bytes, nil
}

func getPemMaterialFromDir(dir string) ([][]byte, error) {
	mspLogger.Debugf("Reading directory %s", dir)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return nil, err
	}

	content := make([][]byte, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "could not read directory %s", dir)
	}

	for _, f := range files {
		fullName := filepath.Join(dir, f.Name())

		f, err := os.Stat(fullName)
		if err != nil {
			mspLogger.Warningf("Failed to stat %s: %s", fullName, err)
			continue
		}
		if f.IsDir() {
			continue
		}

		mspLogger.Debugf("Inspecting file %s", fullName)

		item, err := readPemFile(fullName)
		if err != nil {
			mspLogger.Warningf("Failed reading file %s: %s", fullName, err)
			continue
		}

		content = append(content, item)
	}

	return content, nil
}

const (
	cacerts              = "cacerts"
	admincerts           = "admincerts"
	signcerts            = "signcerts"
	keystore             = "keystore"
	intermediatecerts    = "intermediatecerts"
	crlsfolder           = "crls"
	configfilename       = "config.yaml"
	tlscacerts           = "tlscacerts"
	tlsintermediatecerts = "tlsintermediatecerts"
)

func SetupBCCSPKeystoreConfig(bccspConfig *factory.FactoryOpts, keystoreDir string) *factory.FactoryOpts {
	if bccspConfig == nil {
		bccspConfig = factory.GetDefaultOpts()
	}

	if bccspConfig.ProviderName == "SW" {
		if bccspConfig.SwOpts == nil {
			bccspConfig.SwOpts = factory.GetDefaultOpts().SwOpts
		}

		// Only override the KeyStorePath if it was left empty
		if bccspConfig.SwOpts.FileKeystore == nil ||
			bccspConfig.SwOpts.FileKeystore.KeyStorePath == "" {
			bccspConfig.SwOpts.Ephemeral = false
			bccspConfig.SwOpts.FileKeystore = &factory.FileKeystoreOpts{KeyStorePath: keystoreDir}
		}
	}

	return bccspConfig
}

// GetLocalMspConfigWithType returns a local MSP
// configuration for the MSP in the specified
// directory, with the specified ID and type
func GetLocalMspConfigWithType(dirList [] string, bccspConfig *factory.FactoryOpts, ID, mspType string) (*msp.MSPConfig, error) {
	switch mspType {
	case ProviderTypeToString(FABRIC):
		return GetLocalMspConfig(dirList, bccspConfig, ID)
	case ProviderTypeToString(IDEMIX):
		// TODO 国密多证书 idemix暂时用不到 写死
		return GetIdemixMspConfig(dirList[0], ID)
	default:
		return nil, errors.Errorf("unknown MSP type '%s'", mspType)
	}
}

func GetLocalMspConfig(dirList []string, bccspConfig *factory.FactoryOpts, ID string) (*msp.MSPConfig, error) {
	sigidList := make([]*msp.SigningIdentityInfo, len(dirList))
	for index, dir := range dirList {
		signcertDir := filepath.Join(dir, signcerts)
		keystoreDir := filepath.Join(dir, keystore)
		bccspConfig = SetupBCCSPKeystoreConfig(bccspConfig, keystoreDir)

		err := factory.InitFactories(bccspConfig)
		if err != nil {
			return nil, errors.WithMessage(err, "could not initialize BCCSP Factories")
		}

		signcert, err := getPemMaterialFromDir(signcertDir)
		if err != nil || len(signcert) == 0 {
			return nil, errors.Wrapf(err, "could not load a valid signer certificate from directory %s", signcertDir)
		}

		/* FIXME: for now we're making the following assumptions
		1) there is exactly one signing cert
		2) BCCSP's KeyStore has the private key that matches SKI of
		   signing cert
		*/

		sigid := &msp.SigningIdentityInfo{PublicSigner: signcert[0], PrivateSigner: nil}
		sigidList[index] = sigid
	}

	return getMspConfig(dirList, ID, sigidList)
}

// GetVerifyingMspConfig returns an MSP config given directory, ID and type
func GetVerifyingMspConfig(dirList []string, ID, mspType string) (*msp.MSPConfig, error) {
	switch mspType {
	case ProviderTypeToString(FABRIC):
		return getMspConfig(dirList, ID, nil)
	case ProviderTypeToString(IDEMIX):
		// TODO 国密多证书 idemix暂时不用
		return GetIdemixMspConfig(dirList[0], ID)
	default:
		return nil, errors.Errorf("unknown MSP type '%s'", mspType)
	}
}

func getMspConfig(dirList []string, ID string, sigidList []*msp.SigningIdentityInfo) (*msp.MSPConfig, error) {
	configMap := make(map[string][]byte)
	for index, dir := range dirList {
		cacertDir := filepath.Join(dir, cacerts)
		admincertDir := filepath.Join(dir, admincerts)
		intermediatecertsDir := filepath.Join(dir, intermediatecerts)
		crlsDir := filepath.Join(dir, crlsfolder)
		configFile := filepath.Join(dir, configfilename)
		tlscacertDir := filepath.Join(dir, tlscacerts)
		tlsintermediatecertsDir := filepath.Join(dir, tlsintermediatecerts)

		cacerts, err := getPemMaterialFromDir(cacertDir)
		if err != nil || len(cacerts) == 0 {
			return nil, errors.WithMessage(err, fmt.Sprintf("could not load a valid ca certificate from directory %s", cacertDir))
		}

		admincert, err := getPemMaterialFromDir(admincertDir)
		if err != nil || len(admincert) == 0 {
			return nil, errors.WithMessage(err, fmt.Sprintf("could not load a valid admin certificate from directory %s", admincertDir))
		}

		intermediatecerts, err := getPemMaterialFromDir(intermediatecertsDir)
		if os.IsNotExist(err) {
			mspLogger.Debugf("Intermediate certs folder not found at [%s]. Skipping. [%s]", intermediatecertsDir, err)
		} else if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("failed loading intermediate ca certs at [%s]", intermediatecertsDir))
		}

		tlsCACerts, err := getPemMaterialFromDir(tlscacertDir)
		tlsIntermediateCerts := [][]byte{}
		if os.IsNotExist(err) {
			mspLogger.Debugf("TLS CA certs folder not found at [%s]. Skipping and ignoring TLS intermediate CA folder. [%s]", tlsintermediatecertsDir, err)
		} else if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("failed loading TLS ca certs at [%s]", tlsintermediatecertsDir))
		} else if len(tlsCACerts) != 0 {
			tlsIntermediateCerts, err = getPemMaterialFromDir(tlsintermediatecertsDir)
			if os.IsNotExist(err) {
				mspLogger.Debugf("TLS intermediate certs folder not found at [%s]. Skipping. [%s]", tlsintermediatecertsDir, err)
			} else if err != nil {
				return nil, errors.WithMessage(err, fmt.Sprintf("failed loading TLS intermediate ca certs at [%s]", tlsintermediatecertsDir))
			}
		} else {
			mspLogger.Debugf("TLS CA certs folder at [%s] is empty. Skipping.", tlsintermediatecertsDir)
		}

		crls, err := getPemMaterialFromDir(crlsDir)
		if os.IsNotExist(err) {
			mspLogger.Debugf("crls folder not found at [%s]. Skipping. [%s]", crlsDir, err)
		} else if err != nil {
			return nil, errors.WithMessage(err, fmt.Sprintf("failed loading crls at [%s]", crlsDir))
		}

		// Load configuration file
		// if the configuration file is there then load it
		// otherwise skip it
		var ouis []*msp.FabricOUIdentifier
		var nodeOUs *msp.FabricNodeOUs
		_, err = os.Stat(configFile)
		if err == nil {
			// load the file, if there is a failure in loading it then
			// return an error
			raw, err := ioutil.ReadFile(configFile)
			if err != nil {
				return nil, errors.Wrapf(err, "failed loading configuration file at [%s]", configFile)
			}

			configuration := Configuration{}
			err = yaml.Unmarshal(raw, &configuration)
			if err != nil {
				return nil, errors.Wrapf(err, "failed unmarshalling configuration file at [%s]", configFile)
			}

			// Prepare OrganizationalUnitIdentifiers
			if len(configuration.OrganizationalUnitIdentifiers) > 0 {
				for _, ouID := range configuration.OrganizationalUnitIdentifiers {
					f := filepath.Join(dir, ouID.Certificate)
					raw, err = readFile(f)
					if err != nil {
						return nil, errors.Wrapf(err, "failed loading OrganizationalUnit certificate at [%s]", f)
					}

					oui := &msp.FabricOUIdentifier{
						Certificate:                  raw,
						OrganizationalUnitIdentifier: ouID.OrganizationalUnitIdentifier,
					}
					ouis = append(ouis, oui)
				}
			}

			// Prepare NodeOUs
			if configuration.NodeOUs != nil && configuration.NodeOUs.Enable {
				mspLogger.Debug("Loading NodeOUs")
				if configuration.NodeOUs.ClientOUIdentifier == nil || len(configuration.NodeOUs.ClientOUIdentifier.OrganizationalUnitIdentifier) == 0 {
					return nil, errors.New("Failed loading NodeOUs. ClientOU must be different from nil.")
				}
				if configuration.NodeOUs.PeerOUIdentifier == nil || len(configuration.NodeOUs.PeerOUIdentifier.OrganizationalUnitIdentifier) == 0 {
					return nil, errors.New("Failed loading NodeOUs. PeerOU must be different from nil.")
				}

				nodeOUs = &msp.FabricNodeOUs{
					Enable:             configuration.NodeOUs.Enable,
					ClientOuIdentifier: &msp.FabricOUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.ClientOUIdentifier.OrganizationalUnitIdentifier},
					PeerOuIdentifier:   &msp.FabricOUIdentifier{OrganizationalUnitIdentifier: configuration.NodeOUs.PeerOUIdentifier.OrganizationalUnitIdentifier},
				}

				// Read certificates, if defined

				// ClientOU
				f := filepath.Join(dir, configuration.NodeOUs.ClientOUIdentifier.Certificate)
				raw, err = readFile(f)
				if err != nil {
					mspLogger.Infof("Failed loading ClientOU certificate at [%s]: [%s]", f, err)
				} else {
					nodeOUs.ClientOuIdentifier.Certificate = raw
				}

				// PeerOU
				f = filepath.Join(dir, configuration.NodeOUs.PeerOUIdentifier.Certificate)
				raw, err = readFile(f)
				if err != nil {
					mspLogger.Debugf("Failed loading PeerOU certificate at [%s]: [%s]", f, err)
				} else {
					nodeOUs.PeerOuIdentifier.Certificate = raw
				}
			}
		} else {
			mspLogger.Debugf("MSP configuration file not found at [%s]: [%s]", configFile, err)
		}

		var signatureHashFamily string
		var identityIdentifierHashFunction string
		_, ok, err := isSM2Cert(cacerts[0])
		if err == nil && ok {
			signatureHashFamily = bccsp.SM3
			identityIdentifierHashFunction = bccsp.SM3
		} else {
			signatureHashFamily = bccsp.SHA2
			identityIdentifierHashFunction = bccsp.SHA256
		}
		// Set FabricCryptoConfig
		cryptoConfig := &msp.FabricCryptoConfig{
			SignatureHashFamily:            signatureHashFamily,
			IdentityIdentifierHashFunction: identityIdentifierHashFunction,
		}

		var sigId *msp.SigningIdentityInfo
		if sigidList == nil {
			sigId = nil
		} else {
			sigId = sigidList[index]
		}
		// Compose FabricMSPConfig
		fmspconf := &msp.FabricMSPConfig{
			Admins:                        admincert,
			RootCerts:                     cacerts,
			IntermediateCerts:             intermediatecerts,
			SigningIdentity:               sigId,
			Name:                          ID,
			OrganizationalUnitIdentifiers: ouis,
			RevocationList:                crls,
			CryptoConfig:                  cryptoConfig,
			TlsRootCerts:                  tlsCACerts,
			TlsIntermediateCerts:          tlsIntermediateCerts,
			FabricNodeOus:                 nodeOUs,
		}

		fmpsjs, _ := proto.Marshal(fmspconf)
		// 国密多证书支持
		if _, ok, _ = isSM2Cert(cacerts[0]); ok {
			configMap["sm2"] = fmpsjs
		} else {
			configMap["ecdsa"] = fmpsjs
		}
	}

	mspconf := &msp.MSPConfig{ConfigMap: configMap, Type: int32(FABRIC)}

	return mspconf, nil
}

const (
	IdemixConfigDirMsp                  = "msp"
	IdemixConfigDirUser                 = "user"
	IdemixConfigFileIssuerPublicKey     = "IssuerPublicKey"
	IdemixConfigFileRevocationPublicKey = "RevocationPublicKey"
	IdemixConfigFileSigner              = "SignerConfig"
)

// GetIdemixMspConfig returns the configuration for the Idemix MSP
func GetIdemixMspConfig(dir string, ID string) (*msp.MSPConfig, error) {
	ipkBytes, err := readFile(filepath.Join(dir, IdemixConfigDirMsp, IdemixConfigFileIssuerPublicKey))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read issuer public key file")
	}

	revocationPkBytes, err := readFile(filepath.Join(dir, IdemixConfigDirMsp, IdemixConfigFileRevocationPublicKey))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read revocation public key file")
	}

	idemixConfig := &msp.IdemixMSPConfig{
		Name:         ID,
		Ipk:          ipkBytes,
		RevocationPk: revocationPkBytes,
	}

	signerBytes, err := readFile(filepath.Join(dir, IdemixConfigDirUser, IdemixConfigFileSigner))
	if err == nil {
		signerConfig := &msp.IdemixMSPSignerConfig{}
		err = proto.Unmarshal(signerBytes, signerConfig)
		if err != nil {
			return nil, err
		}
		idemixConfig.Signer = signerConfig
	}

	confBytes, err := proto.Marshal(idemixConfig)
	if err != nil {
		return nil, err
	}

	// TODO 国密多证书 idemix暂时不使用，这里先写死
	configMap := map[string][]byte{
		"ecdsa": confBytes,
	}
	return &msp.MSPConfig{ConfigMap: configMap, Type: int32(IDEMIX)}, nil
}
