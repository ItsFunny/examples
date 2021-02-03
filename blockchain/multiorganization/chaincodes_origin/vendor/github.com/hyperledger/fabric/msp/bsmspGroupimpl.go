package msp

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/utils"
	"github.com/hyperledger/fabric/bssignaturecertmanagement"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
)

type bccspmspGroup struct {
	MSP
	mspMap map[string]*bccspmsp
	// the provider identifier for this MSP
	name    string
	version MSPVersion
}

// newBccspMsp returns an MSP instance backed up by a BCCSP
// crypto provider. It handles x.509 certificates and can
// generate identities and signing identities backed by
// certificates and keypairs
func newBccspMspGroup(version MSPVersion) (MSP, error) {
	mspLogger.Debugf("Creating BCCSP-based MSP instance")

	theMsp := &bccspmspGroup{}
	theMsp.version = version

	if version != MSPv1_0 && version != MSPv1_1 && version != MSPv1_3 {
		return nil, fmt.Errorf("invalid MSPVersion[%d]", version)
	}
	return theMsp, nil
}

func (mspGroup *bccspmspGroup) DeserializeIdentity(serializedIdentity []byte) (Identity, error) {
	var err error
	var identity Identity
	if m, ok := mspGroup.mspMap["sm2"]; ok {
		identity, err = m.DeserializeIdentity(serializedIdentity)
		if err == nil {
			if identity.GetIdentifier().SignatureAlgorithm == "sm2" {
				return identity, nil
			}
		}
	}

	if m, ok := mspGroup.mspMap["ecdsa"]; ok {
		identity, err = m.DeserializeIdentity(serializedIdentity)
		if err == nil {
			return identity, nil
		}
	}

	return nil, fmt.Errorf("bccspmspGroup faield to DeserializeIdentity, desc=%v", err)
}

// IsWellFormed checks if the given identity can be deserialized into its provider-specific form
func (mspGroup *bccspmspGroup) IsWellFormed(identity *msp.SerializedIdentity) error {
	bl, _ := pem.Decode(identity.IdBytes)
	if bl == nil {
		return errors.New("PEM decoding resulted in an empty block")
	}
	// Important: This method looks very similar to getCertFromPem(idBytes []byte) (*x509.Certificate, error)
	// But we:
	// 1) Must ensure PEM block is of type CERTIFICATE or is empty
	// 2) Must not replace getCertFromPem with this method otherwise we will introduce
	//    a change in validation logic which will result in a chain fork.
	if bl.Type != "CERTIFICATE" && bl.Type != "" {
		return errors.Errorf("pem type is %s, should be 'CERTIFICATE' or missing", bl.Type)
	}
	_, err := sm2.ParseCertificate(bl.Bytes)
	if err == nil {
		return nil
	}
	_, err = x509.ParseCertificate(bl.Bytes)
	return err
}

// Setup the MSP instance according to configuration information
// Setup sets up the internal data structures
// for this MSP, given an MSPConfig ref; it
// returns nil in case of success or an error otherwise
func (mspGroup *bccspmspGroup) Setup(conf1 *msp.MSPConfig) error {
	if conf1 == nil {
		return errors.New("Setup error: nil conf reference")
	}
	mspGroup.mspMap = make(map[string]*bccspmsp)
	// given that it's an msp of type fabric, extract the MSPConfig instance

	if conf1.ConfigMap != nil && conf1.Config != nil {
		panic("bccspmspGroup Setup configMap and Config cannot both exist")
	}

	if conf1.Config != nil {
		bcmsp, err := mspGroup.buildBccmsp(conf1.Config)
		if err == nil {
			signatureCertType := "ecdsa"
			if bcmsp.useSM2  {
				signatureCertType = "sm2"
			}
			mspGroup.mspMap[signatureCertType] = bcmsp
		} else {
			var mspConfigMap msp.FabricMSPConfigMap
			err := proto.Unmarshal(conf1.Config, &mspConfigMap)
			if err != nil {
				return err
			}
			for signatureCertType, conf := range mspConfigMap.ConfigMap {
				data, err := proto.Marshal(conf)
				if err != nil {
					return err
				}
				bcmsp, err := mspGroup.buildBccmsp(data)
				if err != nil {
					return err
				}
				mspGroup.mspMap[signatureCertType] = bcmsp
			}
		}

	}


	if conf1.ConfigMap != nil {
		for signatureCertType, conf := range conf1.ConfigMap {
			bcmsp, err := mspGroup.buildBccmsp(conf)
			if err != nil {
				return err
			}
			mspGroup.mspMap[signatureCertType] = bcmsp
		}
	}

	return nil
}

func (mspGroup *bccspmspGroup) buildBccmsp(conf []byte) (*bccspmsp, error) {
	var mspConfig msp.FabricMSPConfig
	err := proto.Unmarshal(conf, &mspConfig)
	if err != nil   {
		return nil, errors.Wrap(err, "failed unmarshalling fabric msp config")
	}
	if mspConfig.Name == "" {
		return nil, errors.New("mspConfig.Name is empty")
	}
	bcmsp, err := newBccspMsp(mspGroup.version)
	if err != nil {
		return nil, err
	}
	bcmsp.name = mspConfig.Name
	// set the name for this msp
	if mspGroup.name == "" {
		mspGroup.name = mspConfig.Name
	} else if mspGroup.name != mspConfig.Name {
		msg := fmt.Sprintf("diffrent mspName in bccspmspGroup, [%s] != [%s]", mspGroup.name, mspConfig.Name)
		panic(msg)
	}
	mspLogger.Debugf("Setting up MSP instance %s", mspGroup.name)
	// !!! 这里很重要 通过判断是否是国密 以国密的形式初始化
	if _, ok, err := isSM2Cert(mspConfig.RootCerts[0]); ok && err == nil {
		bcmsp.useSM2 = true
	}

	// setup
	err = bcmsp.internalSetupFunc(&mspConfig)
	return bcmsp, err
}

// GetVersion returns the version of this MSP
func (mspGroup *bccspmspGroup) GetVersion() MSPVersion {
	return mspGroup.version
}

// GetType returns the provider type
func (mspGroup *bccspmspGroup) GetType() ProviderType {
	return FABRIC
}

// GetIdentifier returns the provider identifier
func (mspGroup *bccspmspGroup) GetIdentifier() (string, error) {
	return mspGroup.name, nil
}

// GetSigningIdentity returns a signing identity corresponding to the provided identifier
func (mspGroup *bccspmspGroup) GetSigningIdentity(identifier *IdentityIdentifier) (SigningIdentity, error) {
	return nil, errors.Errorf("no signing identity for %#v", identifier)
}

func (mspGroup *bccspmspGroup) deserializeIdentityInternal(serializedIdentity []byte) (Identity, error) {
	// This MSP will always deserialize certs this way
	bl, _ := pem.Decode(serializedIdentity)
	if bl == nil {
		return nil, errors.New("could not decode the PEM structure")
	}
	sm2Cert, isSm2Cert, err := utils.IsSM2Cert(bl.Bytes)
	if err != nil {
		return nil, err
	}
	if isSm2Cert {
		sm2Bccmsp := mspGroup.mspMap["sm2"]
		pub, err := sm2Bccmsp.bccsp.KeyImport(sm2Cert, &bccsp.SM2X509PublicKeyImportOpts{Temporary: true})
		if err != nil {
			return nil, errors.WithMessage(err, "failed to import certificate's public key")
		}

		id, err := newSM2Identity(sm2Cert, pub, sm2Bccmsp)
		return id, err
	} else {
		cert, err := x509.ParseCertificate(bl.Bytes)
		if err != nil {
			return nil, errors.Wrap(err, "parseCertificate failed")
		}

		// Now we have the certificate; make sure that its fields
		// (e.g. the Issuer.OU or the Subject.OU) match with the
		// MSP id that this MSP has; otherwise it might be an attack
		// TODO!
		// We can't do it yet because there is no standardized way
		// (yet) to encode the MSP ID into the x.509 body of a cert
		ecdsaBccmsp := mspGroup.mspMap["ecdsa"]
		pub, err := ecdsaBccmsp.bccsp.KeyImport(cert, &bccsp.X509PublicKeyImportOpts{Temporary: true})
		if err != nil {
			return nil, errors.WithMessage(err, "failed to import certificate's public key")
		}

		id, err := newIdentity(cert, pub, ecdsaBccmsp)
		return id, err
	}
}

// GetDefaultSigningIdentity returns the default signing identity
func (mspGroup *bccspmspGroup) GetDefaultSigningIdentity(channelName string) (SigningIdentity, error) {
	signatureCertType := bssignaturecertmanagement.GetChannelSignatureCertType(channelName)
	return mspGroup.mspMap[signatureCertType].signer, nil
}

func (mspGroup *bccspmspGroup) GetSigningIdentityBySignatureCertType(signatureCertType string) (SigningIdentity, error) {
	bccmspVal := mspGroup.mspMap[signatureCertType]
	if bccmspVal == nil {
		return nil, errors.New("signatureCertType not loaded")
	}
	return bccmspVal.signer, nil
}

// GetTLSRootCerts returns the TLS root certificates for this MSP
func (mspGroup *bccspmspGroup) GetTLSRootCerts() [][]byte {
	// 目前多证书共用一套tls证书
	if m, ok := mspGroup.mspMap["ecdsa"]; ok {
		return m.tlsRootCerts
	}
	if m, ok := mspGroup.mspMap["sm2"]; ok {
		return m.tlsRootCerts
	}
	return nil
}

// GetTLSIntermediateCerts returns the TLS intermediate root certificates for this MSP
func (mspGroup *bccspmspGroup) GetTLSIntermediateCerts() [][]byte {
	if m, ok := mspGroup.mspMap["ecdsa"]; ok {
		return m.tlsIntermediateCerts
	}
	if m, ok := mspGroup.mspMap["sm2"]; ok {
		return m.tlsIntermediateCerts
	}
	return nil
}

// Validate checks whether the supplied identity is valid
func (mspGroup *bccspmspGroup) Validate(id Identity) error {
	signatureAlgorithm := id.GetIdentifier().SignatureAlgorithm
	return mspGroup.mspMap[signatureAlgorithm].Validate(id)
}

// SatisfiesPrincipal checks whether the identity matches
// the description supplied in MSPPrincipal. The check may
// involve a byte-by-byte comparison (if the principal is
// a serialized identity) or may require MSP validation
func (mspGroup *bccspmspGroup) SatisfiesPrincipal(id Identity, principal *msp.MSPPrincipal) error {
	signatureAlgorithm := id.GetIdentifier().SignatureAlgorithm
	return mspGroup.mspMap[signatureAlgorithm].SatisfiesPrincipal(id, principal)
}
