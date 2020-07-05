/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-12 13:56 
# @File : main.go
# @Description : 
# @Attention : 
*/
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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
	"strings"
)

var (
	configFilePath = "/Users/joker/go/src/examples/blockchain/oneorg_several_chaincodes/config.yaml"

	ChannelID           = "demo"
	ChannelConfigTxPath = "/Users/joker/go/src/examples/blockchain/oneorg_several_chaincodes/artifacts/demo.tx"

	OrderID           = "orderer0.demo.com"
	ChainCodeIDList   = []string{"usercc", "paycc", "itemcc"}
	ChainCodePathList = []string{"examples/blockchain/oneorg_several_chaincodes/chaincode/usercc",
		"examples/blockchain/oneorg_several_chaincodes/chaincode/paycc",
		"examples/blockchain/oneorg_several_chaincodes/chaincode/itemcc"}
	ChainCodeGoPath = "/Users/joker/go"

	orgName  = "Org1"
	orgAdmin = "Admin"
	OrgUser  = "User1"
)

type BlockChainSetup struct {
	client *channel.Client
	admin  *resmgmt.Client
	sdk    *fabsdk.FabricSDK
	event  *event.Client
	ledger *ledger.Client
}

func main() {

	setUp := new(BlockChainSetup)
	//  18600029374
	fmt.Println("demo程序启动")

	fmt.Println("begin 初始化SDK")
	sdk, e := fabsdk.New(config.FromFile(configFilePath))
	if nil != e {
		panic(e)
	}
	defer sdk.Close()
	setUp.sdk = sdk
	fmt.Println("end 初始化SDK")

	fmt.Println("begin 初始化资源管理器")
	resourceManagerClientContext := sdk.Context(fabsdk.WithOrg(orgName), fabsdk.WithUser(orgAdmin))
	resMClient, e := resmgmt.New(resourceManagerClientContext)
	if nil != e {
		panic(e)
	}
	setUp.admin = resMClient
	fmt.Println("end 初始化资源管理器")

	fmt.Println("begin 初始Admin化msp-client")
	mspClient, e := mspclient.New(sdk.Context(), mspclient.WithOrg(orgName))
	if nil != e {
		panic(e)
	}
	fmt.Println("begin 组装identites")
	identites := make([]msp.SigningIdentity, 0)
	identity, e := mspClient.GetSigningIdentity(orgAdmin)
	if nil != e {
		panic(e)
	}
	identites = append(identites, identity)
	fmt.Println("end 组装identites")
	fmt.Println("end 初始化msp-client")

	fmt.Println("begin 查询已经存在的channel")
	channelResp, e := setUp.admin.QueryChannels(resmgmt.WithTargetEndpoints("peer0.org1.demo.com"))
	if nil != e {
		panic(e)
	}
	fmt.Println("end 查询已经存在的channel")

	fmt.Println("begin 判断channel是否已经存在")
	isChannelExist := false
	if nil != channelResp {
		for _, c := range channelResp.Channels {
			if strings.EqualFold(c.ChannelId, ChannelID) {
				isChannelExist = true
				break
			}
		}
	}

	if isChannelExist {
		fmt.Println("[channel已经存在]")
	} else {
		fmt.Println("begin 创建channel")
		saveChanReq := resmgmt.SaveChannelRequest{ChannelID: ChannelID, ChannelConfigPath: ChannelConfigTxPath, SigningIdentities: identites}
		saveChanResp, e := setUp.admin.SaveChannel(saveChanReq, resmgmt.WithOrdererEndpoint(OrderID))
		if nil != e || saveChanResp.TransactionID == "" {
			panic(e)
		}
		fmt.Println("end 创建channel,channel创建成功")

		fmt.Println("begin 将节点加入channel")
		if e = resMClient.JoinChannel(ChannelID, resmgmt.WithOrdererEndpoint(OrderID), resmgmt.WithOrdererEndpoint("orderer0.demo.com")); nil != e {
			panic(e)
		}
		fmt.Println("end 将节点加入channel")
	}

	fmt.Println("begin 创建区块链账本相关")

	ledgerContext := sdk.ChannelContext(ChannelID, fabsdk.WithOrg(orgName), fabsdk.WithUser(OrgUser))
	ledgerClient, e := ledger.New(ledgerContext)
	if nil != e {
		panic(e)
	}
	setUp.ledger = ledgerClient
	fmt.Println("end 创建区块链账本相关")

	setUp.InstallAndInstantiateCC()

	beego.Get("/query", func(c *context.Context) {
		m := c.Input.Query("method")
		if m == "" {
			m = "fromCC"
		}
		p := c.Input.Query("param")
		if p == "" {
			p = "这是参数"
		}
		cc := c.Input.Query("cc")
		if cc == "" {
			cc = "usercc"
		}
		args := make([][]byte, 0)
		// methodBytes:=[]byte(m)
		// args=append(args,methodBytes)
		paramBytes := []byte(p)
		args = append(args, paramBytes)
		msg := ""
		go func() {
			resp, e := setUp.client.Query(channel.Request{
				ChaincodeID: cc,
				Fcn:         m,
				Args:        args,
			})
			if nil != e {
				msg = e.Error()
			} else {
				msg = string(resp.Payload)
			}
		}()

		r, events, e := setUp.event.RegisterTxStatusEvent("123")
		defer setUp.event.Unregister(r)

		if nil != e {
			c.WriteString(e.Error())
		} else {
			c.WriteString(string(resp.Payload))
		}
	})
	beego.Run(":9000")
}

func (setup *BlockChainSetup) InstallAndInstantiateCC() {
	fmt.Println("begin InstallAndInstantiateCC")
	l := len(ChainCodePathList)
	for i := 0; i < l; i++ {
		fmt.Println("begin 创建chaincode package")
		ccPackage, e := packager.NewCCPackage(ChainCodePathList[i], ChainCodeGoPath)
		if nil != e {
			panic(e)
		}
		fmt.Println("end 创建chaincodepackage")
		ccIsInstall := false

		fmt.Println("begin 查询已经安装的chaincode")
		queryInstallCcResp, e := setup.admin.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints("peer0.org1.demo.com"))
		if nil != e {
			panic(e)
		}
		fmt.Println("begin 查询已经安装的chaincode")

		fmt.Printf("begin 判断是否已经安装了该[%s]cc \n", ChainCodeIDList[i])
		for _, c := range queryInstallCcResp.Chaincodes {
			if strings.EqualFold(ChainCodeIDList[i], c.Name) {
				ccIsInstall = true
				break
			}
		}
		fmt.Printf("end 判断是否已经安装了该[%s]cc \n", ChainCodeIDList[i])

		if ccIsInstall {
			fmt.Printf("该cc[%s]已经安装\n", ChainCodeIDList[i])
		} else {
			fmt.Printf("begin 安装[%s]链码\n", ChainCodeIDList[i])
			installCcReq := resmgmt.InstallCCRequest{
				Name:    ChainCodeIDList[i],
				Path:    ChainCodePathList[i],
				Version: "0",
				Package: ccPackage,
			}
			responses, e := setup.admin.InstallCC(installCcReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
			if nil != e || responses == nil {
				panic(e)
			}
			fmt.Printf("end 安装[%s]链码\n", ChainCodeIDList)
		}

		ccIsInstance := false
		fmt.Println("begin 查询已经实例化的链码")
		queryInstanReps, e := setup.admin.QueryInstantiatedChaincodes(ChannelID, resmgmt.WithTargetEndpoints("peer0.org1.demo.com"))
		if nil != e {
			panic(e)
		}
		fmt.Println("end 查询已经实例化的链码")

		for _, c := range queryInstanReps.Chaincodes {
			if strings.EqualFold(c.Name, ChainCodeIDList[i]) {
				ccIsInstance = true
				break
			}
		}

		if ccIsInstance {
			fmt.Printf("该[%s]已经实例化\n", ChainCodeIDList[i])
		} else {
			fmt.Printf("begin cc[%s]实例化\n", ChainCodeIDList[i])
			ccPolicy := cauthdsl.SignedByMspMember("Org1MSP")
			instanReq := resmgmt.InstantiateCCRequest{
				Name:       ChainCodeIDList[i],
				Path:       ChainCodePathList[i],
				Version:    "0",
				Args:       [][]byte{[]byte("init"), []byte("init")},
				Policy:     ccPolicy,
				CollConfig: nil,
			}
			instanResp, e := setup.admin.InstantiateCC(ChannelID, instanReq, resmgmt.WithTargetEndpoints("peer0.org1.demo.com"))
			if nil != e || instanResp.TransactionID == "" {
				panic(e)
			}
			fmt.Printf("end cc[%s]实例化\n", ChainCodeIDList[i])
		}

	}
	fmt.Println("begin 创建用于execute和query的channel client")
	channelContext := setup.sdk.ChannelContext(ChannelID, fabsdk.WithUser(OrgUser))
	channelClient, e := channel.New(channelContext)
	if nil != e {
		panic(e)
	}
	setup.client = channelClient
	fmt.Println("end 创建用于execute和query的channel client")

	fmt.Println("begin 创建event事件客户端")
	eveClient, e := event.New(channelContext, event.WithBlockEvents())
	if nil != e {
		panic(e)
	}
	setup.event = eveClient
	fmt.Println("begin 创建event事件客户端")

	fmt.Println("end InstallAndInstantiateCC")
}

func (setup *BlockChainSetup) QueryTest() {
}
