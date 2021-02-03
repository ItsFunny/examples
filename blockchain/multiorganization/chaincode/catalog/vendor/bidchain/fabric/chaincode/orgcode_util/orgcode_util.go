package orgcode_util

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	validOrgCodeLength       = 9
	validLicenseNumberLength = 18
)

var (
	orgCodeStr      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	orgCodeWs       = []int{3, 7, 9, 10, 5, 8, 4, 2}
	orgCodeRegexStr = "([0-9A-Z]){8}-?[0-9|X]$"
	orgCodeRegex    = regexp.MustCompile(orgCodeRegexStr)
)

func IsValidOrgCode(orgCode string) bool {
	ok := orgCodeRegex.MatchString(orgCode)
	if !ok {
		return false
	}
	// 最后一位是校验码
	var checkCode byte = orgCode[len(orgCode)-1]
	sum := 0
	for i := 0; i < 8; i++ {
		ci := strings.IndexByte(orgCodeStr, orgCode[i])
		wi := orgCodeWs[i]
		sum += ci * wi
	}
	c9 := 11 - (sum % 11)
	var c9Ch byte
	if c9 == 10 {
		c9Ch = 'X'
	} else if c9 == 11 {
		c9Ch = '0'
	} else {
		c9Ch = '0' + byte(c9)
	}
	return c9Ch == checkCode
}

func getValidOrgCode(orgCode string) (string, error) {
	if strings.IndexByte(orgCode, '-') != -1 {
		orgCode = strings.Replace(orgCode, "-", "", 1)
	}
	if len(orgCode) != validOrgCodeLength {
		return "", fmt.Errorf("orgCode[%s]'s length: %d != %d", orgCode, len(orgCode), validOrgCodeLength)
	}
	if !IsValidOrgCode(orgCode) {
		return "", fmt.Errorf("invliad orgCode[%s]", orgCode)
	}
	return orgCode, nil
}

// 1.orgCode和licenseNumber不能同时为空
// 2.如果orgCode不为空，需要验证其是否有效
// 3.如果licenseNumber不为空，需要验证其长度是否为18位
// 4. 如果两者都不为空，需要验证 orgCode和licenseNumber是否匹配
func IsValidOrgCodeAndLicenseNumber(orgCode string, licenseNumber string) error {
	var err error
	if orgCode == "" && licenseNumber == "" {
		return  errors.New("orgCode and licenseNumber both empty")
	}

	// 验证orgCode
	if orgCode != "" {
		orgCode, err = getValidOrgCode(orgCode)
		if err != nil {
			return err
		}
	}

	// 验证三证合一
	if licenseNumber != "" {
		if len(licenseNumber) != validLicenseNumberLength {
			return fmt.Errorf("invalid licenseNubmer[%s]'s length[%d] != %d", licenseNumber, len(licenseNumber), validLicenseNumberLength)
		}
	}

	// 如果两者都不为空，验证是否匹配
	if orgCode != "" && licenseNumber != "" {
		cmpOrgCode := getOrgCodeFromLicenseNumber(licenseNumber)
		if orgCode != cmpOrgCode {
			return  fmt.Errorf("orgCode[%s] dismatch with licenseNumber[%s]", orgCode, licenseNumber)
		}
	}
	return nil
}

func getOrgCodeFromLicenseNumber(licenseNumber string) string {
	return licenseNumber[8:17]
}
