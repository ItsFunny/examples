/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-19 09:25 
# @File : blockchain.go
# @Description :   blockchain的配置
1个channel可以加入多个组织,一个组织下有多个peer节点,每个peer节点可以注册多个cc
# @Attention : 
*/
package blockchain

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"strconv"
	"strings"
	base2 "vlink.com/v2/vlink-common/base/fabric"
	"vlink.com/v2/vlink-common/base/service"
	error2 "vlink.com/v2/vlink-common/error"
	"vlink.com/v2/vlink-rpc/base"
)

// FIXME 命名规范
type TClientMap map[interface{}]*channel.Client
type TAdminMap map[base.OrganizationID]*resmgmt.Client
type TEventMap map[interface{}]*event.Client

type VlinkBlockChainConfiguration struct {
	Version base2.Version
	sdk     *fabsdk.FabricSDK
	clients map[base.ChannelID]*channel.Client
	admins  map[base.ChannelID]TAdminMap
	events  map[base.ChannelID]*event.Client
	ledgers map[base.ChannelID]*ledger.Client
}

func (setUp *VlinkBlockChainConfiguration) Config(p VlinkBlockChainProperties) error2.IVlinkError {
	if e := setUp.initialize(p); nil != e {
		return e
	}

	if e := setUp.InstallAndInstantiateCC(p); nil != e {
		return e
	}

	return nil
}

func (setUp *VlinkBlockChainConfiguration) initialize(p VlinkBlockChainProperties) error2.IVlinkError {
	fmt.Println("begin 初始化SDK")
	c := config.FromFile(p.ConfigPath)
	sdk, e := fabsdk.New(c)
	if nil != e {
		panic(e)
	}
	// defer sdk.Close()
	setUp.sdk = sdk

	fmt.Println("end 初始化SDK")

	fmt.Println("begin 初始化资源管理器")
	// 多个组织有多个不同的organization ,所以需要for 遍历多次初始化
	for _, channel := range p.Channels {
		for _, organization := range channel.Organizations {
			fmt.Println(fmt.Sprintf("begin 初始化组织为[%s]的资源管理器", organization.OrganizationID))
			resourceManagerClientContext := sdk.Context(fabsdk.WithOrg(string(organization.OrganizationID)), fabsdk.WithUser(organization.OrganizationAdmin))
			admin, e := resmgmt.New(resourceManagerClientContext)
			if nil != e {
				s := fmt.Sprintf("初始化组织为[%s]的资源管理器失败:%s", organization.OrganizationID, e.Error())
				return error2.NewConfigError(e, s)
			}
			if setUp.admins == nil {
				setUp.admins = make(map[base.ChannelID]TAdminMap)
			}
			m, _ := setUp.admins[channel.ChannelID]
			if m == nil {
				setUp.admins[channel.ChannelID] = make(map[base.OrganizationID]*resmgmt.Client)
			}
			// if !exist {
			// }
			setUp.admins[channel.ChannelID][organization.OrganizationID] = admin
			fmt.Println(fmt.Sprintf("end 初始化组织为[%s]的资源管理器", organization.OrganizationID))

			fmt.Println("begin 开始初始化admin-mspclient")
			mspClient, e := mspclient.New(sdk.Context(), mspclient.WithOrg(string(organization.OrganizationID)))
			if nil != e {
				panic(e)
			}
			fmt.Println("begin 组装identites")
			identites := make([]msp.SigningIdentity, 0)
			identity, e := mspClient.GetSigningIdentity(organization.OrganizationAdmin)
			if nil != e {
				panic(e)
			}
			identites = append(identites, identity)
			fmt.Println("end 组装identites")
			fmt.Println("end 初始化msp-client")

			for _, peer := range organization.Peers {
				fmt.Println("begin 查询已经存在的channel")
				channelResp, e := admin.QueryChannels(resmgmt.WithTargetEndpoints(peer.AnchorPeers[0].Address))
				if nil != e {
					return error2.NewConfigError(e, fmt.Sprintf("查询anchorpeer=[%s]上的channel失败:%s", peer.AnchorPeers[0].Address, e.Error()))
				}
				fmt.Println("end 查询已经存在的channel")

				fmt.Println("begin 判断channel是否已经存在")
				isChannelExist := false
				if nil != channelResp {
					for _, c := range channelResp.Channels {
						if strings.EqualFold(c.ChannelId, string(channel.ChannelID)) {
							isChannelExist = true
							break
						}
					}
				}

				if isChannelExist {
					fmt.Println("[channel已经存在]")
				} else {
					fmt.Println("begin 创建channel")
					saveChanReq := resmgmt.SaveChannelRequest{ChannelID: string(channel.ChannelID), ChannelConfigPath: channel.ChannelConfigPath, SigningIdentities: identites}
					// 获取所有的order
					keys := channel.getAllOrderKeys()
					saveChanResp, e := admin.SaveChannel(saveChanReq, keys...)
					if nil != e || saveChanResp.TransactionID == "" {
						panic(e)
					}
					fmt.Println("end 创建channel,channel创建成功")

					fmt.Println("begin 将节点加入channel")
					if e = admin.JoinChannel(string(channel.ChannelID), keys...); nil != e {
						s := fmt.Sprintf("channelId=[%s]加入通道失败:%s", channel.ChannelID, e.Error())
						return error2.NewConfigError(e, s)
					}
					fmt.Println("end 将节点加入channel")
				}
			}

			fmt.Println("begin 创建区块链账本相关")

			cCtx := make([]fabsdk.ContextOption, 0)
			cCtx = append(cCtx, fabsdk.WithOrg(string(organization.OrganizationID)))
			cCtx = append(cCtx, fabsdk.WithUser(organization.OrganizationAdmin))
			ledgerContext := sdk.ChannelContext(string(channel.ChannelID), cCtx...)
			ledgerClient, e := ledger.New(ledgerContext)
			if nil != e {
				panic(e)
			}
			if setUp.ledgers == nil {
				setUp.ledgers = make(map[base.ChannelID]*ledger.Client)
			}
			setUp.ledgers[channel.ChannelID] = ledgerClient
			fmt.Println("end 创建区块链账本相关")
		}
	}

	return nil
}

func (setUp *VlinkBlockChainConfiguration) InstallAndInstantiateCC(p VlinkBlockChainProperties) error2.IVlinkError {
	fmt.Println("begin InstallAndInstantiateCC")
	for _, ccan := range p.Channels {
		for _, organization := range ccan.Organizations {
			for _, peer := range organization.Peers {
				for _, chaincode := range peer.ChainCodes {
					fmt.Println(fmt.Sprintf("begin 创建chaincode package,chaincodeId=[%s]"), chaincode.ChainCodeID)
					ccPackage, e := packager.NewCCPackage(chaincode.ChainCodePath, p.GoPath)
					if nil != e {
						panic(e)
					}
					fmt.Println("end 创建chaincodepackage")
					ccIsInstall := false

					fmt.Println("begin 查询已经安装的chaincode")
					admin := setUp.admins[ccan.ChannelID][organization.OrganizationID]
					anchorAddress := organization.Peers[0].AnchorPeers[0].Address
					queryInstallCcResp, e := admin.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(anchorAddress))
					if nil != e {
						panic(e)
					}

					fmt.Printf("begin 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)
					for _, c := range queryInstallCcResp.Chaincodes {
						if strings.EqualFold(string(chaincode.ChainCodeID), c.Name) {
							ccIsInstall = true
							break
						}
					}
					fmt.Printf("end 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)

					if ccIsInstall {
						fmt.Printf("该cc[%s]已经安装\n", chaincode.ChainCodeID)
					} else {
						fmt.Printf("begin 安装[%s]链码\n", chaincode.ChainCodeID)
						installCcReq := resmgmt.InstallCCRequest{
							Name:    string(chaincode.ChainCodeID),
							Path:    chaincode.ChainCodePath,
							Version: "0",
							Package: ccPackage,
						}
						responses, e := admin.InstallCC(installCcReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
						if nil != e || responses == nil {
							panic(e)
						}
						fmt.Printf("end 安装[%s]链码\n", chaincode)
					}
					ccIsInstance := false
					fmt.Println("begin 查询已经实例化的链码")
					queryInstanReps, e := admin.QueryInstantiatedChaincodes(string(ccan.ChannelID), resmgmt.WithTargetEndpoints(anchorAddress))
					if nil != e {
						s := fmt.Sprintf("查询通道=[%s]上实例化的链码失败:%s", ccan.ChannelID)
						return error2.NewConfigError(e, s)
					}

					for _, c := range queryInstanReps.Chaincodes {
						if strings.EqualFold(c.Name, string(chaincode.ChainCodeID)) {
							ccIsInstance = true
							break
						}
					}
					fmt.Println("end 查询已经实例化的链码")

					if ccIsInstance {
						fmt.Printf("该[%s]已经实例化\n", chaincode.ChainCodeID)
					} else {
						fmt.Printf("begin cc[%s]实例化\n", chaincode.ChainCodeID)
						ccPolicy := cauthdsl.SignedByMspMember(string(organization.OrganizationID))
						instanReq := resmgmt.InstantiateCCRequest{
							Name:       string(chaincode.ChainCodeID),
							Path:       chaincode.ChainCodePath,
							Version:    "0",
							Args:       [][]byte{[]byte("init"), []byte("init")},
							Policy:     ccPolicy,
							CollConfig: nil,
						}
						instanResp, e := admin.InstantiateCC(string(ccan.ChannelID), instanReq, resmgmt.WithTargetEndpoints(anchorAddress))
						if nil != e || instanResp.TransactionID == "" {
							panic(e)
						}
						fmt.Printf("end cc[%s]实例化\n", chaincode.ChainCodeID)
					}
				}
			}
		}

		fmt.Println("begin 创建用于execute和query的channel client")
		channelContext := setUp.sdk.ChannelContext(string(ccan.ChannelID), ccan.getAllUsers()...)
		channelClient, e := channel.New(channelContext)
		if nil != e {
			return error2.NewConfigError(e, "创建通道:"+string(ccan.ChannelID)+"失败")
		}
		if setUp.clients == nil {
			setUp.clients = make(map[base.ChannelID]*channel.Client)
		}
		setUp.clients[ccan.ChannelID] = channelClient
		fmt.Println("end 创建用于execute和query的channel client")

		fmt.Println("begin 创建event事件客户端")
		eveClient, e := event.New(channelContext, event.WithBlockEvents())
		if nil != e {
			return error2.NewConfigError(e, "创建通道事件:"+string(ccan.ChannelID)+"失败")
		}
		if setUp.events == nil {
			setUp.events = make(map[base.ChannelID]*event.Client)
		}
		setUp.events[ccan.ChannelID] = eveClient
		fmt.Println("begin 创建event事件客户端")
		fmt.Println("end InstallAndInstantiateCC")
	}

	return nil
}

func (this *VlinkBlockChainConfiguration) defaultExecute(b base.ChainBaseReq, req interface{}) (channel.Response, error2.IVlinkError) {
	var d interface{}
	switch req.(type) {
	case service.IVlinkCrypter:
		encrypt, e := req.(service.IVlinkCrypter).Encrypt(this.Version)
		if nil != e {
			return channel.Response{}, error2.NewArguError(e, "参数加密失败")
		}
		d = encrypt
	default:
		d = req
	}
	// response, vlinkError := this.execute(b, d)
	// if nil!=vlinkError{
	// 	return channel.Response{},vlinkError
	// }
	// return response,nil

	return this.execute(b, d)
}

func (this *VlinkBlockChainConfiguration) execute(b base.ChainBaseReq, data interface{}) (channel.Response, error2.IVlinkError) {
	var args []string
	args = append(args, string(b.MethodName))

	bytes, e := json.Marshal(data)
	if e != nil {
		return channel.Response{}, error2.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", data))
	}
	args = append(args, string(bytes))
	response, err := this.clients[b.ChannelID].Execute(channel.Request{
		ChaincodeID: string(b.ChainCodeID),
		Fcn:         args[0],
		Args:        [][]byte{[]byte(args[1]), []byte(strconv.Itoa(int(this.Version)))},
	})

	if nil != err {
		return response, error2.NewFabricError(e, fmt.Sprintf("调用fabric失败,方法名称为:%s", b.MethodName))
	}

	return response, nil
}

func HandleResponse(response channel.Response) (base2.BaseFabricResp, error2.IVlinkError) {
	bytes := response.Payload
	var resp base2.BaseFabricResp

	e := json.Unmarshal(bytes, &resp)
	if nil != e {
		return resp, error2.NewJSONSerializeError(e, "反序列化为 BaseFabricResp 失败")
	}

	return resp, nil
}

type VlinkBlockChainProperties struct {
	// 加密的版本号
	CryptVersion base2.Version       `yaml:"cryptVersion" json:"version"`
	Channels     []ChannelProperties `yaml:"channels" json:"channels"`
	// 2019-12-28 add
	// config的路径,绝对路径
	ConfigPath string `yaml:"configPath" json:"configPath"`
	// GOPATH入境
	GoPath string `yaml:"goPath" json:"goPath"`
}

// channel的配置
type ChannelProperties struct {
	// channel名称
	ChannelID     base.ChannelID           `yaml:"channelId" json:"channelId"`
	Organizations []OrganizationProperties `yaml:"organizations"json:"organizations"`
	// 2019-12-28 add
	// order 信息
	Orders []OrdererProperties `yaml:"orders" json:"orders"`
	// channel的定义路径: 既*.tx 路径
	ChannelConfigPath string `yaml:"channelConfigPath" json:"channelConfigPath"`
}

func (c ChannelProperties) getAllContextOption() []fabsdk.ContextOption {
	return append(c.getAllOrganizaions(), c.getAllUsers()...)
}
func (c ChannelProperties) getAllOrganizaions() []fabsdk.ContextOption {
	res := make([]fabsdk.ContextOption, 0)
	for _, org := range c.Organizations {
		o := fabsdk.WithOrg(string(org.OrganizationID))
		res = append(res, o)
	}
	return res
}
func (c ChannelProperties) getAllUsers() []fabsdk.ContextOption {
	res := make([]fabsdk.ContextOption, 0)
	for _, o := range c.Organizations {
		for _, u := range o.Users {
			res = append(res, fabsdk.WithUser(u))
		}
	}
	return res
}

func (c ChannelProperties) getAllOrderKeys() []resmgmt.RequestOption {
	ends := make([]resmgmt.RequestOption, 0)
	for _, o := range c.Orders {
		endpoint := resmgmt.WithOrdererEndpoint(o.OrdererID)
		ends = append(ends, endpoint)
	}
	return ends
}

type OrdererProperties struct {
	OrdererID      string `yaml:"ordererId" json:"ordererId"`
	OrdererAddress string `yaml:"ordererAddress" json:"ordererAddress"`
	OrdererPort    int    `yaml:"ordererPort" json:"ordererPort"`
}
type OrganizationProperties struct {
	OrganizationID base.OrganizationID `yaml:"organizationId" json:"organizationId"`
	Peers          []PeerProperties    `yaml:"peers" json:"peers"`
	// 2019-12-28 add 组织admin名称
	OrganizationAdmin string `yaml:"organizationAdmin" json:"organizationAdmin"`
	// 用户list,暂且只是一个string切片
	Users []string `yaml:"users" json:"users"`
}

// peer的配置
type PeerProperties struct {
	// anchorpeer的节点地址
	AnchorPeers []AnchorPeer `yaml:"anchorPeers" json:"anchorPeers"`
	// 这个peer上要挂载哪些cc
	ChainCodes []ChainCodeProperties `yaml:"chainCodes" json:"chainCodes"`
}

type AnchorPeer struct {
	Address string `yaml:"address" json:"address"`
	Port    int    `yaml:"port" json:"port"`
}

// cc的配置
type ChainCodeProperties struct {
	ChainCodeID   base.ChainCodeID `yaml:"chainCodeId" json:"chainCodeId"`
	ChainCodePath string           `yaml:"chainCodePath" json:"chainCodePath"`
}
