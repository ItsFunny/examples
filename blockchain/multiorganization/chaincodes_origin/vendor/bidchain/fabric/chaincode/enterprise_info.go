package chaincode

import (
	"bidchain/fabric/chaincode/orgcode_util"
	"bidchain/protocol/store/store_base"
	"bidchain/protocol/transport/base"
	"errors"
)

func Transport2StoreEnterpriseInfo(src *base.EnterpriseInfo) *store_base.StoreEnterpriseInfo {
	var dest store_base.StoreEnterpriseInfo
	dest.CompanyName = src.CompanyName
	dest.CompanyLicenseNumber = src.CompanyLicenseNumber
	dest.CompanyOrgCode = src.CompanyOrgCode
	dest.Username = src.Username
	dest.PhoneNumber = src.PhoneNumber
	dest.OpenId = src.OpenId
	return &dest
}

func Store2TransportEnterpriseInfo(src *store_base.StoreEnterpriseInfo) *base.EnterpriseInfo {
	var dest base.EnterpriseInfo
	dest.CompanyName = src.CompanyName
	dest.CompanyLicenseNumber = src.CompanyLicenseNumber
	dest.CompanyOrgCode = src.CompanyOrgCode
	dest.Username = src.Username
	dest.PhoneNumber = src.PhoneNumber
	dest.OpenId = src.OpenId
	return &dest
}

func IsValidEnterpriseInfo(src *base.EnterpriseInfo) error {
	if src == nil {
		return errors.New("EnterpriseInfo is null")
	}
	if src.CompanyName == "" {
		return errors.New("CompanyName is empty")
	}

	if err := orgcode_util.IsValidOrgCodeAndLicenseNumber(src.CompanyOrgCode, src.CompanyLicenseNumber); err != nil {
		return err
	}

	if src.Username == "" {
		return errors.New("Username is empty")
	}
	if src.PhoneNumber == "" {
		return errors.New("PhoneNumber is empty")
	}
	if src.OpenId == "" {
		return errors.New("OpenId is empty")
	}
	return nil
}
