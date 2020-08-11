/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-24 12:43
# @File : blockchian.go
# @Description :
# @Attention :
*/
package config

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	error2 "myLibrary/go-library/blockchain/error"
	"myLibrary/go-library/blockchain/model"
	utils2 "myLibrary/go-library/blockchain/utils"
	"myLibrary/go-library/blockchain/wrapper"
	"myLibrary/go-library/common"
	"myLibrary/go-library/common/blockchain/base"
	error3 "myLibrary/go-library/common/error"
	"myLibrary/go-library/go/converters"
	"myLibrary/go-library/go/log"
	"myLibrary/go-library/go/wallet"
	"strconv"
	"strings"
	"time"
)

type AfterSetupOption func(BlockChainProperties) error3.IBaseError

type ConfigWrapper struct {
	// BlockHandler     handler.IBlockHandler
	SetupBlockEventExecutor SetupBlockEventExecutor
	AfterSetupOption        []AfterSetupOption
}

type OrganizationChannelClientInfo struct {
	Client *channel.Client
}

type ChannelClientInfo struct {
	Clients map[base.OrganizationID]*OrganizationChannelClientInfo
}

// 选择channel->选择channel下的organization->获取client
type ChannelClientWrapper struct {
	clients map[base.ChannelID]*ChannelClientInfo
}

// 组织的资源admin,用于控制channel中的资源
type OrganizationResourceAdmin struct {
	Admin *resmgmt.Client
}

type ResourceAdminInfo struct {
	admins map[base.OrganizationID]*OrganizationResourceAdmin
}

type ResourceAdminWrapper struct {
	admins map[base.ChannelID]*ResourceAdminInfo
}

// //

type ChannelEventInfo struct {
	blockRegistration      fab.Registration
	blockStopEventFlagChan chan struct{}
	EventClient            *event.Client
}

type ChannelEventWrapper struct {
	Events map[base.ChannelID]*ChannelEventInfo
}

// ///////
type ChannelLedgerInfo struct {
	Ledger *ledger.Client
}
type ChannelLedgerWrapper struct {
	ledgers map[base.ChannelID]*ChannelLedgerInfo
}

//
type BlockChainConfiguration struct {
	Log                  log.Logger
	Version              base.Version
	Properties           *BlockChainProperties
	sdk                  *fabsdk.FabricSDK
	clientWrapper        *ChannelClientWrapper
	adminResourceWrapper *ResourceAdminWrapper
	events               *ChannelEventWrapper
	ledgerWrapper        *ChannelLedgerWrapper
	msps                 *wrapper.OrganizationMspWrapper
}

func NewBlockChainConfiguration() *BlockChainConfiguration {
	c := new(BlockChainConfiguration)
	c.Log = log.NewCommonBaseLoggerWithLog4go("config")
	return c
}

type ExecuteReq struct {
	MethodName     base.MethodName
	ChannelID      base.ChannelID
	OrganizationID base.OrganizationID
	ChainCodeID    base.ChainCodeID
	ReqData        interface{}

	Context context.Context
}

func (this *BlockChainConfiguration) GetChannelChainCodes(channelId string) []string {
	result := make([]string, 0)
	for _, ch := range this.Properties.Channels {
		if string(ch.ChannelID) != channelId {
			continue
		}

		for _, org := range ch.Organizations {
			for _, p := range org.Peer.EndorserPeers {
				for _, cc := range p.ChainCodes {
					result = append(result, string(cc.ChainCodeID))
				}
			}
		}
	}
	return result
}
func (this *BlockChainConfiguration) CallWithNewClient() {
}

func (this *BlockChainConfiguration) Execute(executeReq ExecuteReq) (base.ServiceLogicBaseResp, []byte, error3.IBaseError) {
	// name base.MethodName, id base.ChannelID, codeID ChainCodeID, req interface{}
	name := executeReq.MethodName
	id := executeReq.ChannelID
	codeID := executeReq.ChainCodeID
	req := executeReq.ReqData
	var (
		logicRes base.ServiceLogicBaseResp
		e        error3.IBaseError
	)
	response, baseError := this.defaultExecute(base.ChainBaseReq{
		MethodName:     name,
		ChannelID:      id,
		OrganizationID: executeReq.OrganizationID,
		ChainCodeID:    codeID,
		Context:        executeReq.Context,
	}, req)
	if nil != baseError {
		return logicRes, nil, baseError
	}
	resp, _ := HandleResponse(response)
	logicRes.LogicCode = int(converter.BigEndianBytes2Int64(resp.CodeBytes))
	logicRes.LogicMsg = string(resp.MsgBytes)
	logicRes.LogBytes = resp.LogBytes
	if resp.OtherBytes != nil && len(resp.OtherBytes) > 0 {
		this.Log.Debug("otherBytes的数据为:%s", string(resp.OtherBytes))
		if e := json.Unmarshal(resp.OtherBytes, &logicRes.CommAttribute); nil != e {
			return logicRes, nil, error3.NewJSONSerializeError(e, "CommAttribute反序列化失败")
		}
	}

	return logicRes, resp.DataBytes, e
}

func (this *BlockChainConfiguration) ExecuteWithPureBytes(executeReq ExecuteReq) (channel.Response, error3.IBaseError) {
	// name base.MethodName, id base.ChannelID, codeID ChainCodeID, req interface{}
	name := executeReq.MethodName
	id := executeReq.ChannelID
	codeID := executeReq.ChainCodeID
	req := executeReq.ReqData
	response, baseError := this.defaultExecute(base.ChainBaseReq{
		MethodName:     name,
		ChannelID:      id,
		OrganizationID: executeReq.OrganizationID,
		ChainCodeID:    codeID,
		Context:        executeReq.Context,
	}, req)
	return response, baseError
}

func (this *BlockChainConfiguration) defaultExecute(b base.ChainBaseReq, req interface{}) (channel.Response, error3.IBaseError) {
	var d interface{}
	switch req.(type) {
	case common.ICrypter:
		encrypt, e := req.(common.ICrypter).Encrypt(this.Version)
		if nil != e {
			return channel.Response{}, error3.NewArguError(e, "参数加密失败")
		}
		d = encrypt
	default:
		d = req
	}
	// response, baseError := this.execute(b, d)
	// if nil!=baseError{
	// 	return channel.Response{},baseError
	// }
	// return response,nil

	return this.execute(b, d)
}

func (this *BlockChainConfiguration) EnrollAdminUser(id base.ChannelID, admin, adminpw string) {
}

type Info struct {
	UserName string
}

func (this *BlockChainConfiguration) execute(b base.ChainBaseReq, data interface{}) (channel.Response, error3.IBaseError) {
	var args []string
	var result channel.Response
	args = append(args, string(b.MethodName))

	bytes, e := json.Marshal(data)
	if e != nil {
		return channel.Response{}, error3.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", data))
	}
	var resultBytes []byte
	var err error
	args = append(args, string(bytes))
	if b.Context == nil {
		adminClient := this.clientWrapper.clients[b.ChannelID].Clients[b.OrganizationID].Client
		result, err = adminClient.Execute(channel.Request{
			ChaincodeID: string(b.ChainCodeID),
			Fcn:         args[0],
			// 2020-01-08 update 为了与invokechaincode 一致,此处补齐一个string
			Args: [][]byte{[]byte(args[1]), []byte(strconv.Itoa(int(this.Version)))},
		}, channel.WithTimeout(fab.Execute, time.Second*60))
		resultBytes = result.Payload
	} else {
		info, ok := b.Context.Value("info").(*Info)
		if ok {
			userName := info.UserName
			gw, e := gateway.Connect(gateway.WithSDK(this.sdk), gateway.WithUser(userName))
			if nil != e {
				return channel.Response{}, error2.NewFabricError(e, "根据用户连接到区块链网络失败")
			}
			network, e := gw.GetNetwork(string(b.ChannelID))
			if nil != e {
				return channel.Response{}, error2.NewChannelError(e, "连接到network失败")
			}
			contract := network.GetContract(string(b.ChainCodeID))
			resultBytes, err = contract.SubmitTransaction(args[0], args[1], strconv.Itoa(int(this.Version)))
		} else {
			adminClient := this.clientWrapper.clients[b.ChannelID].Clients[b.OrganizationID].Client
			result, err = adminClient.Execute(channel.Request{
				ChaincodeID: string(b.ChainCodeID),
				Fcn:         args[0],
				// 2020-01-08 update 为了与invokechaincode 一致,此处补齐一个string
				Args: [][]byte{[]byte(args[1]), []byte(strconv.Itoa(int(this.Version)))},
			}, channel.WithTimeout(fab.Execute, time.Second*60))
			resultBytes = result.Payload
		}
	}

	if nil != err {
		return channel.Response{}, error2.NewFabricError(err, fmt.Sprintf("调用fabric失败,方法名称为:%s", b.MethodName))
	}

	result.Payload = resultBytes
	return result, nil
}

// FIXME 与上面的方法整合
func (this *BlockChainConfiguration) ExecuteWithClient(client *channel.Client, b base.ChainBaseReq, data interface{}) (channel.Response, error3.IBaseError) {
	var args []string
	args = append(args, string(b.MethodName))

	bytes, e := json.Marshal(data)
	if e != nil {
		return channel.Response{}, error3.NewJSONSerializeError(e, fmt.Sprintf("序列化data=[%v]", data))
	}
	args = append(args, string(bytes))
	response, err := client.Execute(channel.Request{
		ChaincodeID: string(b.ChainCodeID),
		Fcn:         args[0],
		// 2020-01-08 update 为了与invokechaincode 一致,此处补齐一个string
		Args: [][]byte{[]byte(args[1]), []byte(strconv.Itoa(int(this.Version)))},
	}, channel.WithTimeout(fab.Execute, time.Second*60))

	if nil != err {
		return response, error2.NewFabricError(err, fmt.Sprintf("调用fabric失败,方法名称为:%s", b.MethodName))
	}

	return response, nil
}

func (this *BlockChainConfiguration) Close() error {
	for _, eveAdapter := range this.events.Events {
		if nil != eveAdapter.EventClient {
			eveAdapter.blockStopEventFlagChan <- struct{}{}
			eveAdapter.EventClient.Unregister(eveAdapter.blockRegistration)
		}
	}
	return nil
}

func HandleResponse(response channel.Response) (base.BaseFabricResp, error3.IBaseError) {
	bytes := response.Payload
	var resp base.BaseFabricResp

	e := json.Unmarshal(bytes, &resp)
	if nil != e {
		return resp, error3.NewJSONSerializeError(e, "反序列化结构体失败")
	}

	return resp, nil
}

func (setUp *BlockChainConfiguration) Config(path string, configWrapper ConfigWrapper) error3.IBaseError {
	properties, baseError := setUp.loadYaml(path)
	if baseError != nil {
		return baseError
	}
	if e := setUp.doConfig(properties, configWrapper); nil != e {
		return e
	}
	if nil != configWrapper.AfterSetupOption {
		for _, after := range configWrapper.AfterSetupOption {
			if e := after(*properties); nil != e {
				return e
			}
		}
	}
	return nil
}

func (setUp *BlockChainConfiguration) loadYaml(path string) (*BlockChainProperties, error3.IBaseError) {
	var conf BlockChainProperties
	bytes, err := ioutil.ReadFile(path)
	if nil != err {
		return nil, error3.NewConfigError(err, "读取文件错误,path="+path)
	}
	err = yaml.Unmarshal(bytes, &conf)
	if nil != err {
		setUp.Log.Info("yaml unmarshal occur err:" + err.Error())
		return nil, error3.NewConfigError(err, "yaml反序列化错误")
	}
	return &conf, nil
}

func (setUp *BlockChainConfiguration) doConfig(p *BlockChainProperties, wrapper ConfigWrapper) error3.IBaseError {
	if e := setUp.initialize(*p); nil != e {
		return e
	}

	if e := setUp.InstallAndInstantiateCC(*p, wrapper); nil != e {
		return e
	}

	// 启动监听block
	RunTasks()

	setUp.Properties = p

	return nil
}

func (setUp *BlockChainConfiguration) initialize(p BlockChainProperties) error3.IBaseError {
	setUp.Log.Info("begin 初始化SDK")
	c := config.FromFile(p.ConfigPath)
	sdk, e := fabsdk.New(c)
	if nil != e {
		panic(e)
	}
	// defer sdk.Close()
	setUp.sdk = sdk

	setUp.Log.Info("end 初始化SDK")

	setUp.Log.Info("begin 初始化资源管理器")
	// 多个组织有多个不同的organization ,所以需要for 遍历多次初始化
	channelCreated := false
	for _, channel := range p.Channels {
		for _, organization := range channel.Organizations {
			setUp.Log.Info(fmt.Sprintf("begin 初始化组织为[%s]的资源管理器", organization.OrganizationID))
			setUp.Log.Info("信息为:", organization.String())
			resourceManagerClientContext := sdk.Context(fabsdk.WithOrg(string(organization.OrganizationID)), fabsdk.WithUser(organization.GetAdminUser()[0]))
			// resourceManagerClientContext := sdk.Context(fabsdk.WithOrg(string(organization.OrganizationID)))
			admin, e := resmgmt.New(resourceManagerClientContext)
			if nil != e {
				s := fmt.Sprintf("初始化组织为[%s]的资源管理器失败:%s", organization.OrganizationID, e.Error())
				return error3.NewConfigError(e, s)
			}
			if setUp.adminResourceWrapper == nil {
				w := new(ResourceAdminWrapper)
				w.admins = make(map[base.ChannelID]*ResourceAdminInfo)
				setUp.adminResourceWrapper = w
			}
			m := setUp.adminResourceWrapper.admins[channel.ChannelID]
			if m == nil {
				resourceMap := new(ResourceAdminInfo)
				resourceMap.admins = make(map[base.OrganizationID]*OrganizationResourceAdmin)
				setUp.adminResourceWrapper.admins[channel.ChannelID] = resourceMap
				m = resourceMap
			}
			m.admins[organization.OrganizationID] = &OrganizationResourceAdmin{
				Admin: admin,
			}
			setUp.Log.Info(fmt.Sprintf("end 初始化组织为[%s]的资源管理器", organization.OrganizationID))

			setUp.Log.Info("begin 开始初始化admin-mspclient")
			mspClient, e := mspclient.New(sdk.Context(), mspclient.WithOrg(string(organization.OrganizationID)))
			if nil != e {
				panic(e)
			}
			if setUp.msps == nil {
				setUp.msps = wrapper.NewOrganizationMspWrapper()
			}
			clientWrapper := wrapper.NewMspClientWrapper(mspClient)
			clientWrapper.CaInfo.CaName = organization.Ca.CaName
			setUp.msps.Clients[organization.OrganizationID] = clientWrapper
			setUp.Log.Info("begin 组装identites")
			identites := make([]msp.SigningIdentity, 0)
			identity, e := mspClient.GetSigningIdentity(organization.GetAdminUser()[0])
			if nil != e {
				panic(e)
			}
			identites = append(identites, identity)
			setUp.Log.Info("end 组装identites")
			setUp.Log.Info("end 初始化msp-client")

			setUp.Log.Info("begin 查询已经存在的channel")
			channelResp, e := admin.QueryChannels(resmgmt.WithTargetEndpoints(organization.Peer.EndorserPeers[0].Address))
			if nil != e {
				return error3.NewConfigError(e, fmt.Sprintf("查询anchorpeer=[%s]上的channel失败:%s", organization.Peer.EndorserPeers[0].Address, e.Error()))
			}
			setUp.Log.Info("end 查询已经存在的channel,长度为:%d", len(channelResp.Channels))

			setUp.Log.Info("begin 判断channel是否已经存在")
			if nil != channelResp {
				for _, c := range channelResp.Channels {
					setUp.Log.Info("存在channel名称为:" + c.ChannelId)
					if strings.EqualFold(c.ChannelId, string(channel.ChannelID)) {
						channelCreated = true
						break
					}
				}
			}

			// p := models.VlinkPeer{
			// 	ChannelName: string(channel.base.ChannelID),
			// 	Domain:      peer.AnchorPeers[0].Address,
			// 	Port:        peer.AnchorPeers[0].Port,
			// }
			if channelCreated {
				setUp.Log.Info(fmt.Sprintf("channel:[%s]已经存在\n", string(channel.ChannelID)))
			} else {
				setUp.Log.Info("begin 创建channel")
				saveChanReq := resmgmt.SaveChannelRequest{ChannelID: string(channel.ChannelID), ChannelConfigPath: channel.ChannelConfigPath, SigningIdentities: identites}
				// 获取某个order的keys 即可
				endPoints := make([]resmgmt.RequestOption, 0)
				orderP := channel.Orders[0]
				// endPoints = append(endPoints, resmgmt.WithOrdererEndpoint(orderP.OrdererAddress),
				// 	resmgmt.WithOrdererEndpoint("orderer2.vlink.link"),
				// 	resmgmt.WithOrdererEndpoint("orderer3.vlink.link"),
				// 	resmgmt.WithOrdererEndpoint("orderer4.vlink.link"),
				// 	resmgmt.WithOrdererEndpoint("orderer5.vlink.link"))
				endPoints = append(endPoints, resmgmt.WithOrdererEndpoint(orderP.OrdererAddress))
				endPoints = append(endPoints, channel.GetChannelAllPeersTarget())

				saveChanResp, e := admin.SaveChannel(saveChanReq, endPoints...)
				if nil != e || saveChanResp.TransactionID == "" {
					setUp.Log.Error("创建channel失败,可能是因为已经存在了channel,所以直接加入")

					tryJoinChannel := func() error3.IBaseError {
						setUp.Log.Info("尝试将节点加入channel")
						// resmgmt.WithTargetEndpoints()
						if e = admin.JoinChannel(string(channel.ChannelID)); nil != e {
							s := fmt.Sprintf("尝试channelId=[%s]加入通道失败:%s", channel.ChannelID, e.Error())
							return error3.NewConfigError(e, s)
						}
						return nil
					}
					if e := tryJoinChannel(); nil != e {
						return e
					}
				} else {
					setUp.Log.Debug("end 创建channel成功,sleep 10秒等待raft 选举完毕")
					time.Sleep(time.Second * 10)
					setUp.Log.Info("begin 将节点加入channel")
					// resmgmt.WithTargetEndpoints()
					if e = admin.JoinChannel(string(channel.ChannelID)); nil != e {
						s := fmt.Sprintf("channelId=[%s]加入通道失败:%s", channel.ChannelID, e.Error())
						return error3.NewConfigError(e, s)
					}
					setUp.Log.Info("end 将节点加入channel")
				}
			}
			setUp.Log.Info("begin 创建区块链账本相关")

			cCtx := make([]fabsdk.ContextOption, 0)
			cCtx = append(cCtx, fabsdk.WithOrg(string(organization.OrganizationID)))
			cCtx = append(cCtx, fabsdk.WithUser(organization.GetAdminUser()[0]))
			ledgerContext := sdk.ChannelContext(string(channel.ChannelID), cCtx...)
			ledgerClient, e := ledger.New(ledgerContext)
			if nil != e {
				panic(e)
			}
			if setUp.ledgerWrapper == nil {
				ledgerWrapper := new(ChannelLedgerWrapper)
				ledgerWrapper.ledgers = make(map[base.ChannelID]*ChannelLedgerInfo)
				setUp.ledgerWrapper = ledgerWrapper
			}
			setUp.ledgerWrapper.ledgers[channel.ChannelID] = &ChannelLedgerInfo{
				Ledger: ledgerClient,
			}
			setUp.Log.Info("end 创建区块链账本相关")

			channelCreated = false
		}
	}

	return nil
}

func (setUp *BlockChainConfiguration) InstallAndInstantiateCC(p BlockChainProperties, wrapper ConfigWrapper) error3.IBaseError {
	setUp.Log.Info("begin InstallAndInstantiateCC")
	for _, ccan := range p.Channels {
		for _, org := range ccan.Organizations {
			for _, peer := range org.Peer.EndorserPeers {
				for _, chaincode := range peer.ChainCodes {
					setUp.Log.Info(fmt.Sprintf("begin 创建chaincode package,chaincodeId=[%s]"), chaincode.ChainCodeID)
					ccPackage, e := packager.NewCCPackage(chaincode.ChainCodePath, p.GoPath)
					if nil != e {
						panic(e)
					}
					setUp.Log.Info("end 创建chaincodepackage")
					ccIsInstall := false

					setUp.Log.Info("begin 查询已经安装的chaincode")
					admin := setUp.adminResourceWrapper.admins[ccan.ChannelID].admins[org.OrganizationID].Admin
					endorserAddress := peer.Address
					queryInstallCcResp, e := admin.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(endorserAddress))
					if nil != e {
						panic(e)
					}

					fmt.Printf("begin 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)
					Versions := make([]int, 0)
					Versions = append(Versions, 0)
					for _, c := range queryInstallCcResp.Chaincodes {
						if strings.EqualFold(string(chaincode.ChainCodeID), c.Name) {
							ccIsInstall = true
							i, _ := strconv.Atoi(c.Version)
							Versions = append(Versions, i+1)
						}
					}

					fmt.Printf("end 判断是否已经安装了该[%s]cc \n", chaincode.ChainCodeID)
					if ccIsInstall {
						fmt.Printf("该cc[%s]已经安装\n", chaincode.ChainCodeID)
						if chaincode.NeedUpdate {
							setUp.Log.Info(fmt.Sprintf("链码[%s]需要升级,版本号为:%d", chaincode.ChainCodeID, Versions[len(Versions)-1]))
							newInstallCCReq := resmgmt.InstallCCRequest{Name: string(chaincode.ChainCodeID), Path: chaincode.ChainCodePath, Version: strconv.Itoa(Versions[len(Versions)-1]), Package: ccPackage}
							_, err := admin.InstallCC(newInstallCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
							if err != nil {
								return error3.NewConfigError(err, "failed to install chaincode")
							}
							setUp.Log.Info("Chaincode install successfully ,begin upgrade")
							request := resmgmt.UpgradeCCRequest{
								Name:    string(chaincode.ChainCodeID),
								Path:    chaincode.ChainCodePath,
								Version: strconv.Itoa(Versions[len(Versions)-1]),
							}
							member := cauthdsl.SignedByMspMember(string(org.OrganizationID))
							request.Policy = member
							response, err := admin.UpgradeCC(string(ccan.ChannelID), request)
							if nil != err {
								setUp.Log.Info("更新链码失败:", err.Error())
								return error3.NewConfigError(err, "更新链码失败")
							} else {
								setUp.Log.Info(response.TransactionID)
							}
							setUp.Log.Info(fmt.Sprintf("更新链码[%s]成功"), chaincode.ChainCodeID)
						}
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
					setUp.Log.Info("begin 查询已经实例化的链码")
					queryInstanReps, e := admin.QueryInstantiatedChaincodes(string(ccan.ChannelID), resmgmt.WithTargetEndpoints(endorserAddress))
					if nil != e {
						s := fmt.Sprintf("查询通道=[%s]上实例化的链码失败:%s", ccan.ChannelID)
						return error3.NewConfigError(e, s)
					}

					for _, c := range queryInstanReps.Chaincodes {
						if strings.EqualFold(c.Name, string(chaincode.ChainCodeID)) {
							ccIsInstance = true
							break
						}
					}
					setUp.Log.Info("end 查询已经实例化的链码")

					if ccIsInstance {
						fmt.Printf("该[%s]已经实例化\n", chaincode.ChainCodeID)
					} else {
						fmt.Printf("begin cc[%s]实例化\n", chaincode.ChainCodeID)
						ccPolicy := cauthdsl.SignedByMspMember(string(org.OrganizationID))
						instanReq := resmgmt.InstantiateCCRequest{
							Name:       string(chaincode.ChainCodeID),
							Path:       chaincode.ChainCodePath,
							Version:    "0",
							Args:       [][]byte{[]byte("init"), []byte("init")},
							Policy:     ccPolicy,
							CollConfig: nil,
						}
						instanResp, e := admin.InstantiateCC(string(ccan.ChannelID), instanReq,
							resmgmt.WithOrdererEndpoint(ccan.Orders[0].OrdererAddress),
							resmgmt.WithTargetEndpoints(endorserAddress))
						if nil != e || instanResp.TransactionID == "" {
							panic(e)
						}
						fmt.Printf("end cc[%s]实例化\n", chaincode.ChainCodeID)
					}
				}
			}

			setUp.Log.Info("begin 创建用于execute和query的channel client,基于channel->organization")
			// FIXME 这里需要确定,是通过channelId,还是org还是user 获取client信息
			// 若是org ,则 需要有一个 通过channelId获取org的map
			channelContext := setUp.sdk.ChannelContext(string(ccan.ChannelID), org.getEnrollUsers()...)
			channelClient, e := channel.New(channelContext)
			if nil != e {
				return error3.NewConfigError(e, "创建通道:"+string(ccan.ChannelID)+"失败")
			}

			if setUp.clientWrapper == nil {
				clientWrapper := new(ChannelClientWrapper)
				clientWrapper.clients = make(map[base.ChannelID]*ChannelClientInfo)
				setUp.clientWrapper = clientWrapper
			}
			clientOrganzationInfo := setUp.clientWrapper.clients[ccan.ChannelID]
			if clientOrganzationInfo == nil {
				clientOrganzationInfo = new(ChannelClientInfo)
				setUp.clientWrapper.clients[ccan.ChannelID] = clientOrganzationInfo
			}
			if clientOrganzationInfo.Clients == nil {
				clientOrganzationInfo = &ChannelClientInfo{
					Clients: map[base.OrganizationID]*OrganizationChannelClientInfo{org.OrganizationID: &OrganizationChannelClientInfo{
						Client: channelClient,
					}},
				}
				setUp.clientWrapper.clients[ccan.ChannelID] = clientOrganzationInfo
			} else if _, exist := clientOrganzationInfo.Clients[org.OrganizationID]; !exist {
				clientOrganzationInfo.Clients[org.OrganizationID] = &OrganizationChannelClientInfo{
					Client: channelClient,
				}
			} else {
				continue
			}

			setUp.Log.Info("end 创建用于execute和query的channel client")

			if ccan.NeedListOnBlockEvent {
				setUp.Log.Info("begin 创建event事件客户端")
				eveClient, e := event.New(channelContext, event.WithBlockEvents())
				if nil != e {
					return error3.NewConfigError(e, "创建通道事件:"+string(ccan.ChannelID)+"失败")
				}
				if setUp.events == nil {
					eventWrapper := &ChannelEventWrapper{
						Events: make(map[base.ChannelID]*ChannelEventInfo),
					}
					setUp.events = eventWrapper
				}
				eveAdapter := new(ChannelEventInfo)
				blockStopEventFlagChan := make(chan struct{}, 1)
				eveAdapter.blockStopEventFlagChan = blockStopEventFlagChan
				eveAdapter.EventClient = eveClient
				setUp.events.Events[ccan.ChannelID] = eveAdapter
				setUp.Log.Info("begin 创建event事件客户端")
				registration, events, e := eveClient.RegisterBlockEvent()
				if nil != e {
					return error3.NewSystemError(e, "监听block事件失败")
				}
				eveAdapter.blockRegistration = registration
				RegisterBlockEvent(ccan.ChannelID, wrapper.SetupBlockEventExecutor, ccan.GetInterestBlockEventChainCodes(), events, blockStopEventFlagChan)

				setUp.Log.Info("end InstallAndInstantiateCC")
			}
		}

	}

	return nil
}

type UserRegisterResp struct {
	EnrollSecret string
	Identity     msp.SigningIdentity

	Address   string
	PublicKey string
}

func (c *BlockChainConfiguration) RegisterAndEnroll(req model.UserRegisterReq) (UserRegisterResp, error3.IBaseError) {
	var (
		result UserRegisterResp
	)
	resp, baseError := c.Register(req)
	if nil != baseError {
		return result, baseError
	}
	if res, enroll := c.Enroll(model.UserEnrollReq{
		Oid:           req.Oid,
		UserUniqueKey: req.Name,
		UserPassword:  req.Secret,
		Profile:       "",
		Type:          req.Type,
	}); nil != enroll {
		return result, enroll
	} else {
		result.Identity = res.Identity
		publicKey, e := res.Identity.PrivateKey().PublicKey()
		if nil != e {
			return result, error3.UserRegistrationError(e, "用户注册失败,获取公钥失败")
		}
		bytes, e := publicKey.Bytes()
		if nil != e {
			return result, error3.UserRegistrationError(e, "用户注册失败,获取公钥失败")
		}
		result.Address = wallet.GenerateAddress(bytes)
		result.PublicKey = hex.EncodeToString(bytes)
	}
	result.EnrollSecret = resp.EnrollSecret
	return result, nil
}

func (c *BlockChainConfiguration) Register(req model.UserRegisterReq) (model.UserRegistrationResp, error3.IBaseError) {
	var (
		result model.UserRegistrationResp
	)
	mspClient, exist := c.msps.Clients[req.Oid]
	if !exist {
		return result, error3.OrganizationNotExistError(nil, "组织不存在")
	}
	s, e := mspClient.Client.Register(&mspclient.RegistrationRequest{
		Name:           req.Name,
		Type:           req.Type,
		MaxEnrollments: -1,
		Affiliation:    "",
		Attributes:     req.Attributes,
		CAName:         mspClient.CaInfo.CaName,
		Secret:         req.Secret,
	})
	if nil != e {
		return result, error3.UserRegistrationError(e, "用户注册失败")
	}

	result.EnrollSecret = s

	return result, nil
}

type EnrollResp struct {
	Identity msp.SigningIdentity
}

func (c *BlockChainConfiguration) Enroll(req model.UserEnrollReq) (EnrollResp, error3.IBaseError) {
	var (
		result EnrollResp
	)
	mspClient, exist := c.msps.Clients[req.Oid]
	if !exist {
		return result, error3.OrganizationNotExistError(nil, "组织不存在")
	}
	if e := mspClient.Client.Enroll(req.UserUniqueKey, req.BuildEnrollOptions()...); nil != e {
		return result, error2.NewFabricError(e, "enroll用户失败")
	}
	if identity, e := mspClient.Client.GetSigningIdentity(req.UserUniqueKey); nil != e {
		return result, error2.NewFabricError(e, "enroll用户成功,但是获取签名失败")
	} else {
		result.Identity = identity
	}

	return result, nil
}

func (c *BlockChainConfiguration) GetOrganizationAdminUsers(oid base.OrganizationID) []string {
	for _, ch := range c.Properties.Channels {
		for _, org := range ch.Organizations {
			if org.OrganizationID == oid {
				return org.GetAdminUser()
			}
		}
	}
	return nil
}

func (c *BlockChainConfiguration) NewChannelClient(cid base.ChannelID, oid base.OrganizationID, user string) (*channel.Client, error) {

	channelContext := c.sdk.ChannelContext(string(cid), fabsdk.WithOrg(string(oid)), fabsdk.WithUser(user))
	channelClient, e := channel.New(channelContext)
	return channelClient, e
}

func (c *BlockChainConfiguration) GetChannelLedgerWrapper(cid string) (*ChannelLedgerInfo, bool) {
	info, exist := c.ledgerWrapper.ledgers[base.ChannelID(cid)]
	return info, exist
}

func (c *BlockChainConfiguration) GetChannelClientByCid(cid string) *ChannelClientInfo {
	return c.clientWrapper.clients[base.ChannelID(cid)]
}

func (this *BlockChainConfiguration) GetTransactionByTxId(req model.TransactionGetByIdReq) (model.TransactionGetByIdResp, error3.IBaseError) {
	var (
		result model.TransactionGetByIdResp
	)
	ledgerWrapper, exist := this.GetChannelLedgerWrapper(req.ChannelId)
	channelLedger := ledgerWrapper.Ledger
	if !exist {
		return result, error3.NewArguError(nil, "参数channelId找不到匹配的账本")
	}
	block, e := channelLedger.QueryBlockByTxID(fab.TransactionID(req.TxID))
	if nil != e {
		this.Log.Error("账本通过交易id查询信息失败:" + e.Error())
		return result, error2.NewFabricError(e, "查询账本失败")
	}
	result.BlockHeight = block.Header.Number
	result.BlockHash = hex.EncodeToString(block.Header.DataHash)
	processedTransaction, err := channelLedger.QueryTransaction(fab.TransactionID(req.TxID))
	if nil != err {
		this.Log.Error("从 ledger中通过txID=%s查询 记录的时候失败:%s", req.TxID, err.Error())
		return result, error2.NewFabricError(err, "通过txId查询失败")
	} else if nil == processedTransaction {
		this.Log.Error("交易id对应的记录不存在")
		return result, error3.NewRecordNotExistError("该交易id不存在")
	}
	if processedTransaction.TransactionEnvelope == nil {
		this.Log.Error("txID=%s的envelop不存在", req.TxID)
		return result, error3.NewRecordNotExistError("对应的envelop不存在")
	}
	result.Signature = hex.EncodeToString(processedTransaction.TransactionEnvelope.Signature)
	return result, nil
}

func (this *BlockChainConfiguration) GetTransactionDetailData(req model.TransactionDetailGetByIdReq) (utils2.TransactionDetail, error3.IBaseError) {
	var (
		result utils2.TransactionDetail
	)
	ledgerWrapper, exist := this.GetChannelLedgerWrapper(req.ChannelId)
	if !exist {
		return result, error3.NewArguError(nil, "参数channelId找不到匹配的账本")
	}

	processedTransaction, err := ledgerWrapper.Ledger.QueryTransaction(fab.TransactionID(req.TxID))
	if nil != err {
		this.Log.Error("从 ledger中通过txID=%s查询 记录的时候失败:%s", req.TxID, err.Error())
		return result, error2.NewFabricError(err, "通过txId查询失败")
	} else if nil == processedTransaction {
		this.Log.Error("交易id对应的记录不存在")
		return result, error3.NewRecordNotExistError("该交易id不存在")
	}

	env := cb.Envelope{
		Payload:              processedTransaction.TransactionEnvelope.Payload,
		Signature:            processedTransaction.TransactionEnvelope.Signature,
		XXX_NoUnkeyedLiteral: processedTransaction.TransactionEnvelope.XXX_NoUnkeyedLiteral,
		XXX_unrecognized:     processedTransaction.TransactionEnvelope.XXX_unrecognized,
		XXX_sizecache:        processedTransaction.TransactionEnvelope.XXX_sizecache,
	}
	result, err = utils2.GetTransactionInfoFromEnvelop(&env, req.ChainCodeIdList, req.NeedArgs, req.DescriptionFunc)
	if nil != err {
		this.Log.Error("查询交易失败:" + err.Error())
		return result, error3.NewLedgerError(err, req.ChannelId, "获取交易失败")
	}
	result.Signature = hex.EncodeToString(env.Signature)
	return result, nil
}
