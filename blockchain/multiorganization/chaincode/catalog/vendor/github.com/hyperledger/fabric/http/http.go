/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021-01-07 15:19 
# @File : http.go
# @Description : 
# @Attention : 
*/
package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/config"
	"github.com/pkg/errors"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	logger      = flogging.MustGetLogger("http")
	modeEncrypt bool
)

type BidsunJsonKeyInfo struct {
	ServerId         string `json:"serverId"`
	BidsunTicketUrl  string `json:"bidsunTicketUrl"`
	DoeDecUrl        string `json:"doeDecUrl"`
	CryptKeyLabel    string `json:"cryptKeyLabel"`
	CryptLicToken    string `json:"cryptLicToken"`
	CryptPubKey      string `json:"cryptPubKey"`
	ServerUserKey    string `json:"serverUserKey"`
	SignKeyLabel     string `json:"signKeyLabel"`
	SignPubKey       string `json:"signPubKey"`
	TicketPubKey     string `json:"ticketPublicKey"`
	TicketPrivateKey string `json:"ticketPrivateKey"`
}

type BidsunKeyInfoConfiguration struct {
	ServerId         string `json:"serverId"`
	BidsunTicketUrl  string `json:"bidsunTicketUrl"`
	DoeDecUrl        string `json:"doeDecUrl"`
	CryptKeyLabel    string `json:"cryptKeyLabel"`
	CryptLicToken    string `json:"cryptLicToken"`
	CryptPubKey      *sm2.PublicKey
	ServerUserPrvKey *sm2.PrivateKey
	SignKeyLabel     string `json:"signKeyLabel"`
	SignPubKey       *sm2.PublicKey
	TicketPubKey     string
	TicketPrvKey     *sm2.PrivateKey
}

type DoeDecRespDTO struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
	Data string `json:"data"`
}

type DoeDecRequest struct {
	Ticket        string   `json:"ticket"`
	CryptKeyLabel string   `json:"cryptKeyLabel"`
	LicToken      string   `json:"licToken"`
	Data          []string `json:"data"`
}

type HttpResult struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

var (
	DefaultBidsunHttpHandler func(response *http.Response) (result *HttpResult, b bool, e error) = func(response *http.Response) (result *HttpResult, b bool, e error) {
		dataBytes, e := ioutil.ReadAll(response.Body)
		if nil != e {
			return nil, true, errors.New("ioutil读取失败:" + e.Error())
		}
		if response.StatusCode != 200 {
			msg := "未知"
			if len(dataBytes) != 0 {
				msg = string(dataBytes)
			}
			return nil, true, errors.New("响应码状态不为200,状态信息为:" + response.Status + ",msg:" + msg)
		}
		errCode := response.Header.Get("errorCode")
		if len(errCode) != 0 && errCode != "0" {
			errorDesc := response.Header.Get("errorDesc")
			if len(errorDesc) == 0 {
				errorDesc = "未知错误"
			}
			var msg string
			s, e := url.QueryUnescape(errorDesc)
			// errorStrBytes, e := base64.URLEncoding.DecodeString(errorDesc)
			if nil != e {
				errorDesc = "url decode 未知错误:" + e.Error() + ",原先:" + errorDesc
			}
			msg = s
			return nil, false, errors.New("逻辑错误了,错误码为:" + errCode + ",错误信息为:" + msg)
		}
		var res TicketResp
		if e = json.Unmarshal(dataBytes, &res); nil != e {
			return nil, false, errors.New("json反序列化失败:" + e.Error())
		}
		return &HttpResult{
			Data: res,
			Code: 1,
			Msg:  "success",
		}, true, nil
	}
	DefaultDotHttpHandler = func(response *http.Response) (result *HttpResult, b bool, e error) {
		dataBytes, e := ioutil.ReadAll(response.Body)
		if nil != e {
			return nil, true, errors.New("ioutil读取失败:" + e.Error())
		}
		if response.StatusCode != 200 {
			msg := "未知"
			if len(dataBytes) != 0 {
				msg = string(dataBytes)
			}
			return nil, true, errors.New("响应码状态不为200,状态信息为:" + response.Status + ",msg:" + msg)
		}
		logger.Info("doe返回值:" + string(dataBytes))
		var resp DoeDecRespDTO
		if e = json.Unmarshal(dataBytes, &resp); nil != e {
			return nil, false, errors.New("反序列化失败:" + e.Error())
		}
		logger.Info(resp)
		if resp.Code != 0 {
			return nil, false, errors.New("调用深思失败:" + resp.Desc)
		}
		maps := make([]map[string]interface{}, 0)
		if e = json.Unmarshal([]byte(resp.Data), &maps); nil != e {
			return nil, false, errors.New("反序列化为map失败:" + e.Error())
		}
		return &HttpResult{
			Data: maps,
			Code: 1,
			Msg:  "success",
		}, true, nil
	}
)

var (
	keyConfiguration *BidsunKeyInfoConfiguration
)

func init() {
	cfgBytes := config.GetOrDefaultFileBytes(config.ENV_DOE_CONFIG_FILEPATH, config.DEFAULT_DOE_CONFIG_FILEPATH, logger, func(err error) {
		modeEncrypt = false
		logger.Info("加密机相关配置失败,因此采取的是非加密机的形式")
	})
	var keyInfo BidsunJsonKeyInfo
	if e := json.Unmarshal(cfgBytes, &keyInfo); nil != e {
		modeEncrypt = false
		logger.Warn("反序列化成keyInfo失败,因此采用非加密形式")
		return
	}
	modeEncrypt = true
	keyConfiguration = &BidsunKeyInfoConfiguration{
		ServerId:         keyInfo.ServerId,
		BidsunTicketUrl:  keyInfo.BidsunTicketUrl,
		CryptKeyLabel:    keyInfo.CryptKeyLabel,
		CryptLicToken:    keyInfo.CryptLicToken,
		CryptPubKey:      parsePubKeyFromString(keyInfo.CryptPubKey),
		ServerUserPrvKey: nil,
		SignKeyLabel:     keyInfo.SignKeyLabel,
		SignPubKey:       parsePubKeyFromString(keyInfo.SignPubKey),
		TicketPubKey:     keyInfo.TicketPubKey,
		TicketPrvKey:     parseSm2PrvKeyFromString(keyInfo.TicketPrivateKey),
		DoeDecUrl:        keyInfo.DoeDecUrl,
	}
}

type TicketReq struct {
	ServerId string `json:"serverId"`
}

type TicketResp struct {
	AppId     string `json:"appId"`
	Ticket    string `json:"ticket"`
	AppSecret string `json:"appSecret"`
}

func getTicket() (TicketResp, error) {
	defaultHttpReq := HttpReq{
		TimeOutSeconds:  10,
		MaxFailCount:    5,
		HttpUrl:         keyConfiguration.BidsunTicketUrl,
		BytesData:       nil,
		Headers:         nil,
		HandlerResponse: DefaultBidsunHttpHandler,
	}
	headers := make(map[string]string, 0)
	privateKey := keyConfiguration.TicketPrvKey
	originBytes := []byte(keyConfiguration.ServerId)
	userId := []byte("1234567812345678")
	r, s, err := sm2.Sm2Sign(privateKey, originBytes, userId)
	if nil != err {
		return TicketResp{}, err
	}
	signature := encodeSignatureRS(r, s)
	headers["signStr"] = string(signature)
	headers["pubKey"] = keyConfiguration.TicketPubKey
	headers["Content-Type"] = "application/json"
	defaultHttpReq.Headers = headers
	req := TicketReq{
		ServerId: keyConfiguration.ServerId,
	}
	reqBytes, _ := json.Marshal(req)
	defaultHttpReq.BytesData = reqBytes
	resString, err := PostPerClient(defaultHttpReq)
	if nil != err {
		panic("获取ticket失败:" + err.Error())
	}
	logger.Info(resString)
	return resString.(TicketResp), err
}

func BuildSignDataHeader(appSecret, appId, timeStamp, originData string) string {
	sb := strings.Builder{}
	sb.WriteString(appId)
	sb.WriteString(timeStamp)
	sb.WriteString(originData)
	msg := sb.String()
	logger.Info("hmac加密的原数据为:" + msg)
	return base64.StdEncoding.EncodeToString(encodeHmacSHA256([]byte(appSecret), []byte(msg)))
}
func encodeHmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func Decrypt(data []byte) (result []byte) {
	defer func() {
		if e := recover(); nil != e {
			logger.Error(fmt.Sprintf("调用加密机失败,因此返回原来的数据,或许数据不需要加密,原先的数据为:"+string(data)+",错误信息为:%s", e))
			result = data
		}
	}()
	if !modeEncrypt {
		return data
	}
	logger.Info("开始调用加密机,数据为:" + base64.StdEncoding.EncodeToString(data))
	ticketResp, e := getTicket()
	if nil != e {
		panic("获取ticket失败:" + e.Error())
	}
	req := DoeDecRequest{
		Ticket:        ticketResp.Ticket,
		CryptKeyLabel: keyConfiguration.CryptKeyLabel,
		LicToken:      keyConfiguration.CryptLicToken,
	}
	req.Data = []string{base64.StdEncoding.EncodeToString(data)}
	logger.Info("加密数据为:" + req.Data[0])
	marshal, _ := json.Marshal(req)
	logger.Info("请求参数数据为:" + string(marshal))
	originData := string(marshal)
	// originData="123"
	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	headers["appId"] = ticketResp.AppId
	headers["timestamp"] = timeStamp
	signDataHeader := BuildSignDataHeader(ticketResp.AppSecret, ticketResp.AppId, timeStamp, originData)
	logger.Info("timeStamp:" + timeStamp)
	logger.Info("signDataHeader:" + signDataHeader)
	headers["signData"] = signDataHeader
	s, e := PostPerClient(HttpReq{
		TimeOutSeconds:  10,
		MaxFailCount:    5,
		HttpUrl:         keyConfiguration.DoeDecUrl,
		BytesData:       marshal,
		Headers:         headers,
		HandlerResponse: DefaultDotHttpHandler,
	})
	if nil != e {
		panic("调用深思失败:" + e.Error())
	}
	resp := s.([]map[string]interface{})
	if len(req.Data) != len(resp) {
		panic("返回值不一致,请求了:" + strconv.Itoa(len(req.Data)) + ",但是实际对象只返回了:" + strconv.Itoa(len(resp)))
	}
	for i := 0; i < len(req.Data); i++ {
		v := resp[i]
		if code, exist := v["code"]; !exist {
			panic("调用深思失败,没有code码返回")
		} else {
			switch code.(type) {
			case float64:
				codeInt := int(code.(float64))
				if codeInt != 0 {
					msg := ""
					desc := v["desc"]
					if desc != nil {
						msg = desc.(string)
					}
					panic("调用深思失败,对内部的数据解密失败:" + msg)
				}
			}
		}
		pureSec := v[req.Data[i]]
		if pureSec == nil {
			panic("返回的密钥为空")
		}
		bytes, e := base64.StdEncoding.DecodeString(pureSec.(string))
		if nil != e {
			panic("返回的密钥base64 解码失败:" + e.Error())
		}
		logger.Info("密钥为:" + string(bytes))
		result = bytes
		return
	}
	logger.Info(resp)
	panic("数据为空")
}

func parsePubKeyFromString(b64Str string) *sm2.PublicKey {
	pubBytes, e := base64.StdEncoding.DecodeString(b64Str)
	if nil != e {
		log.Fatal(e)
	}
	x := pubBytes[:32]
	y := pubBytes[32:64]
	pubX := new(big.Int).SetBytes(x)
	pubY := new(big.Int).SetBytes(y)
	pub := sm2.PublicKey{
		Curve: sm2.P256Sm2(),
		X:     pubX,
		Y:     pubY,
	}
	return &pub
}
func parseSm2PrvKeyFromString(b64String string) *sm2.PrivateKey {
	decodeString, e := base64.StdEncoding.DecodeString(b64String)
	if nil != e {
		panic("解析私钥失败")
	}
	if len(decodeString) == 32 {
		var (
			raw []byte
		)
		if strings.Contains(string(decodeString), "BEGIN") {
			var block *pem.Block
			block, _ = pem.Decode(decodeString)
			if block == nil {
				panic("failed to decode private key")
			}
			raw = block.Bytes
		}
		raw = decodeString
		if len(raw) != 32 {
			panic("标准私钥为32个字节")
		}
		var privateKey sm2.PrivateKey
		d := raw
		privateKey.D = new(big.Int).SetBytes(d)
		curve := sm2.P256Sm2()
		privateKey.Curve = curve
		x, y := curve.ScalarBaseMult(d)
		privateKey.X, privateKey.Y = x, y
		return &privateKey
	}
	key, e := sm2.ParseSm2PrivateKey(decodeString)
	if nil != e {
		panic("ParseSm2PrivateKey#解析私钥失败")
	}
	return key
}

// 将两个大整数拼接成字符串
func encodeSignatureRS(r, s *big.Int) []byte {
	// 缺位补足
	r1 := r.Text(16)
	if len(r1) < 64 {
		for i := len(r1); i < 64; i++ {
			r1 = "0" + r1
		}
	}

	// 缺位补足
	s1 := s.Text(16)
	if len(s1) < 64 {
		for i := len(s1); i < 64; i++ {
			s1 = "0" + s1
		}
	}

	return []byte(fmt.Sprintf("%s%s", r1, s1))
}
