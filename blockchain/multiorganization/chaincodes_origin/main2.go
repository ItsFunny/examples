package main

import (
	"bidchain/chaincode/catalog/gen/command"
	"bidchain/fabric/chaincode_mock"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/http_framework/protocol"
	"bidchain/micro/catalog/bo"
	"bidchain/micro/common/dto"
	"bidchain/protocol/transport/catalog"
	"encoding/json"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"math/rand"
	"time"
)

/*
	测试用例
	1. 参数问题校验
	2. 逻辑问题校验
		2.1
		2.2 目录上链包含继承关系但是父目录未上链
		2.3 目录上链,但是上一个版本不存在
		2.4 目录上链,正常上链
		2.5 目录上链,查询链上信息
		2.6 目录上链,递归查询链上所有信息
*/
// func main() {
// 	cc := mockCC()
// 	ctx := cc.GetContext()
// 	// test_2_2_catalogUpload(cc,ctx)
// 	// test_2_3_catalogUpload(cc,ctx)
// 	// test_2_4_catalogUpload(cc,ctx)
// 	test_2_5_catalogUpload(cc, ctx)
// }

//  目录上链包含继承关系但是父目录未上链
func test_2_2_catalogUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	_, _, result := testUploadCatalog(buildCatalog(), cc, ctx, func(c *catalog.Catalog) {
		c.CatalogBasicInfo.UploadVersion = 1
	})
	fmt.Print(result)
}

var RESULT_DTO_HOOK = func(bytes []byte) (string, error) {
	var res dto.ResultDTO
	err := json.Unmarshal(bytes, &res)
	if nil != err {
		return "", err
	}
	if len(res.Data) == 0 {
		return string(res.Msg), nil
	}
	return string(res.Data), nil
}

// 目录上链,但是上一个版本不存在
func test_2_3_catalogUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	_, _, result := testUploadCatalog(buildCatalog(), cc, ctx, nil)
	fmt.Print(result)
}

// 目录上链,正常上链
func test_2_4_catalogUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	_, _, result := testUploadCatalog(nil, cc, ctx, func(c *catalog.Catalog) {
		c.CatalogBasicInfo.UploadVersion = 1
	})
	fmt.Print(result)
}

// 目录上链,查询链上信息
func test_2_5_catalogUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	parent, _, result := testUploadCatalog(nil, cc, ctx, func(c *catalog.Catalog) {
		c.CatalogBasicInfo.UploadVersion = 1
	})
	fmt.Println(result)

	ctx.SetFunctionName("GetCatalogInfoById")
	req := bo.GetCatalogInfoByIdReq{
		CatalogId: parent.CatalogBasicInfo.CatalogId,
	}
	bytes, _ := json.Marshal(req)

	request := &command.AddOrUpdateCatalogRequest{}
	res, err := chaincode_mock.InvokeCommandWithHookInternal(cc, ctx, request, func(ctx context.IBidchainContext) {
		ctx.SetParameters([]string{string(bytes)})
	}, RESULT_DTO_HOOK)
	fmt.Println(res)
	fmt.Println(err)
}

// 目录上链,递归查询链上所有信息
func test_2_6_catalogUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	parent, _, result := testUploadCatalog(nil, cc, ctx, func(c *catalog.Catalog) {
		c.CatalogBasicInfo.UploadVersion = 1
	})
	fmt.Println(result)

	fmt.Println("上链子目录")
	childParent, _, result := testUploadCatalog(parent, cc, ctx, func(c *catalog.Catalog) {
		c.CatalogBasicInfo.UploadVersion = 1
	})
	fmt.Println(result)
	ctx.SetFunctionName("CheckIsDesendatantCatalog")
	req := bo.CheckIsDesendatantCatalogReq{
		ParentCatalogId: parent.CatalogBasicInfo.CatalogId,
		ChildCatalogId:  childParent.CatalogBasicInfo.CatalogId,
	}
	bytes, _ := json.Marshal(req)

	request := &command.AddOrUpdateCatalogRequest{}
	res, err := chaincode_mock.InvokeCommandWithHookInternal(cc, ctx, request, func(ctx context.IBidchainContext) {
		ctx.SetParameters([]string{string(bytes)})
	}, RESULT_DTO_HOOK)
	fmt.Println(res)
	fmt.Println(err)
}

func testParentUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	testUploadCatalog(nil, cc, ctx, nil)
}

// 测试父目录上链之后,继承关系上链
func testParentUploadAndInterUpload(cc *CatalogContract, ctx context.IBidchainContext) {
	parent, _, _ := testUploadCatalog(nil, cc, ctx, nil)
	testUploadCatalog(parent, cc, ctx, nil)
}

// 父目录上链
func testUploadCatalog(parent *catalog.Catalog, cc *CatalogContract, ctx context.IBidchainContext, decorator func(c *catalog.Catalog)) (*catalog.Catalog, protocol.ICommand, *chaincode_mock.ErrResult) {
	ctx.SetFunctionName("addOrUpdateCatalog")
	request := &command.AddOrUpdateCatalogRequest{}
	reqs := make([]*catalog.Catalog, 0)
	reqs = append(reqs, &catalog.Catalog{
		CatalogBasicInfo:       buildBasicInfoWithRelation(parent),
		DataItemDefinitionList: buildDataItemDefinitionList(),
		GlossaryList:           buildDataTypeDescriptor(),
		TagList:                buildTags(),
		ConstraintList:         buildForeignKeyConstraint(),
	})
	if decorator != nil {
		decorator(reqs[0])
	}
	request.SetReq(reqs)

	iCommand, result := chaincode_mock.InvokeCommand(cc, ctx, request)
	fmt.Println(result)
	fmt.Println(iCommand)

	time.Sleep(time.Second * 4)
	return reqs[0], iCommand, result
}
func buildCatalog() *catalog.Catalog {
	return &catalog.Catalog{
		CatalogBasicInfo:       buildBasicInfoWithRelation(nil),
		DataItemDefinitionList: buildDataItemDefinitionList(),
		GlossaryList:           buildDataTypeDescriptor(),
		TagList:                buildTags(),
		ConstraintList:         buildForeignKeyConstraint(),
	}
}
func buildBasicInfo() *catalog.CatalogBasicInfo {
	return &catalog.CatalogBasicInfo{
		CatalogId:              randomString(),
		CatalogName:            randomString(),
		UploadVersion:          randint32(),
		ShowVersionId:          randomString(),
		ShowVersion:            randomString(),
		CatalogOwnerPlatformId: randomString(),
		StateEnumId:            randomString(),
		PublishTime:            1111,
		CreateDraftTime:        111,
		StartPublicityTime:     01111,
		EndPublicityTime:       1111,
	}
}

func buildBasicInfoWithRelation(parent *catalog.Catalog) *catalog.CatalogBasicInfo {
	pid := ""
	pver := int32(0)
	detailId := ""
	if parent != nil {
		pid = parent.CatalogBasicInfo.CatalogId
		pver = parent.CatalogBasicInfo.UploadVersion
		detailId = randomString()
	}
	return &catalog.CatalogBasicInfo{
		CatalogId:                  randomString(),
		CatalogName:                randomString(),
		UploadVersion:              randint32(),
		ShowVersionId:              randomString(),
		ShowVersion:                randomString(),
		CatalogOwnerPlatformId:     randomString(),
		StateEnumId:                randomString(),
		Level:                      randint32(),
		PublishTime:                1111,
		CreateDraftTime:            111,
		StartPublicityTime:         01111,
		EndPublicityTime:           1111,
		InheritDetailId:            detailId,
		ParentCatalogId:            pid,
		ParentCatalogUploadVersion: pver,
	}
}
func buildDataItemDefinitionList() []*catalog.DataItemDefinition {
	res := make([]*catalog.DataItemDefinition, 0)
	res = append(res, buildDataItemDefinition())
	return res
}
func buildDataItemDefinition() *catalog.DataItemDefinition {
	return &catalog.DataItemDefinition{
		DataItemDefinitionId: randomString(),
		DataItemName:         randomString(),
		DataItemVersion:      randint32(),
		PublicFieldList:      buildDataFieldDefinitionList(),
		Weight:               1,
		MergeFieldList:       buildDataItemMergedFieldData(),
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
}
func buildDataFieldDefinitionList() []*catalog.DataFieldDefinition {
	res := make([]*catalog.DataFieldDefinition, 0)
	res = append(res, &catalog.DataFieldDefinition{
		DataFieldName:       randomString(),
		EnglishName:         randomString(),
		IsRequired:          false,
		IsEncrypted:         false,
		EnumSecurityLevelId: randint32(),
		Weight:              1,
		Level:               1,
		OrderIndex:          1,
		DataFieldEnumId:     1,
	})
	return res
}
func buildDataItemMergedFieldData() []*catalog.DataItemMergedFieldData {
	res := make([]*catalog.DataItemMergedFieldData, 0)
	res = append(res, &catalog.DataItemMergedFieldData{
		SourceFieldEnglishName: randomString(),
		MergedFieldEnglishName: randomString(),
		MergeFieldLevel:        1,
	})

	return res
}
func buildDataTypeDescriptor() []*catalog.DataTypeDescriptor {
	res := make([]*catalog.DataTypeDescriptor, 0)
	res = append(res, &catalog.DataTypeDescriptor{
		DataType:        randomString(),
		EnumId:          randint32(),
		EnumEnglishName: randomString(),
		EnumChineseName: randomString(),
		EnumSourceId:    randint32(),
		EnumValueList:   buildCodeInfo(),
	})
	return res
}

func buildTags() []*catalog.Tag {
	res := make([]*catalog.Tag, 0)
	res = append(res, &catalog.Tag{
		TagId:                 randomString(),
		TagName:               randomString(),
		TagType:               randomString(),
		CreateTagNodeID:       randomString(),
		ToWhichDataItemIdList: []string{randomString()},
	})

	return res
}
func buildCodeInfo() []*catalog.CodeInfo {
	res := make([]*catalog.CodeInfo, 0)
	res = append(res, &catalog.CodeInfo{
		CodeName: randomString(),
		Content:  randomString(),
	})

	return res
}

func buildForeignKeyConstraint() []*catalog.ForeignKeyConstraint {
	res := make([]*catalog.ForeignKeyConstraint, 0)
	res = append(res, &catalog.ForeignKeyConstraint{
		ConstraintId:               randomString(),
		LeftTableId:                randomString(),
		LeftTableFieldEnglishName:  randomString(),
		RightTableId:               randomString(),
		RightTableFieldEnglishName: randomString(),
	})

	return res
}
func mockCC() *CatalogContract {
	log.SetLevel(l4g.DEBUG)
	cc := new(CatalogContract)
	cc.Child = cc
	// 注册cmd
	for _, cmd := range cmdList {
		protocol.RegisterCommand(cmd)
	}

	stub := shim.NewMockStub("catalog", cc)
	cc.Init(stub)
	// ctx := cc.GetContext()
	return cc
}

func randomString() string {
	return uuid.New().String()
}
func randBytes() []byte {
	return []byte(randomString())
}

func randint32() int32 {
	return rand.Int31()
}
