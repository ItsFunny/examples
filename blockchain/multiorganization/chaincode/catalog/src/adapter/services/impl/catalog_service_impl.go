/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 14:55
# @File : catalog_service_impl.go
# @Description :
# @Attention :
*/
package impl

import (
	"bidchain/base/jsonutils"
	"bidchain/base/services/impl"
	"bidchain/base/utils"
	"bidchain/chaincode/catalog/src/adapter/bcso"
	"bidchain/chaincode/catalog/src/adapter/constants"
	"bidchain/chaincode/catalog/src/adapter/models"
	"bidchain/chaincode/catalog/src/adapter/services"
	"bidchain/fabric/bserror"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/chaincode"
	"bidchain/micro/catalog/bo"
	"bidchain/micro/common/dto"
	"bidchain/protocol/transport/catalog"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

var _ services.ICatalogService = &cataLogServiceImpl{}

type cataLogServiceImpl struct {
	*baseImpl.ContractAdapterImpl
}

func NewCatalogServiceImpl(mediator *chaincode.BasisContract, reqId string) services.ICatalogService {
	c := &cataLogServiceImpl{
		ContractAdapterImpl: baseImpl.NewContractAdapterImpl(mediator, reqId),
	}
	return c
}

func (c *cataLogServiceImpl) CheckIsDesendatantCatalog(req bo.CheckIsDesendatantCatalogReq) (dto.ResultDTO, bserror.IBSError) {
	c.BeforeStart("CheckIsDesendatantCatalog")
	defer c.AfterEnd()
	var dataResp bo.CheckIsChildCatalogResp
	// 递归一点一点往下查
	b, err := c.loopTillFindTarget(req.ParentCatalogId, req.ChildCatalogId)
	if nil != err {
		c.Error("递归查询失败:" + err.Error())
		return dto.ResultDTO{}, err
	}
	dataResp.Desendatant = b

	return dto.Success(dataResp), nil
}

func (c *cataLogServiceImpl) loopTillFindTarget(parentId, childCatalogId string) (bool, bserror.IBSError) {
	// 递归一点一点往下查
	_, infoObj, e := c.getCatalogRelations(parentId)
	if nil != e {
		c.Error("查询数据失败:" + e.Error())
		return false, e
	}
	if len(infoObj.Nodes) == 0 {
		return false, nil
	}
	searchNodes := make([]string, 0)
	for _, node := range infoObj.Nodes {
		for _, v := range node.Nodes {
			searchNodes = append(searchNodes, v.CatalogId)
			if v.CatalogId == childCatalogId {
				return true, nil
			}
		}
	}
	for _, id := range searchNodes {
		b, ibsError := c.loopTillFindTarget(id, childCatalogId)
		if nil != ibsError {
			c.Error("递归查询发生错误:" + ibsError.Error())
			return false, ibsError
		} else if b {
			c.Info("找寻到匹配记录,该cataLogId:[" + childCatalogId + "],在:[" + id + "]中")
			return true, nil
		}
	}
	return false, nil
}

func (c *cataLogServiceImpl) RevokeCatalogRelation(req bo.RevokeCatalogRelationReq) (dto.ResultDTO, bserror.IBSError) {
	c.BeforeStart("RevokeCatalogRelation")
	c.AfterEnd()

	key, infoObj, e := c.getCatalogRelations(req.ParentCatalogId)
	if nil != e {
		c.Error("查询数据失败:" + e.Error())
		return dto.ResultDTO{}, e
	}
	if len(infoObj.Nodes) == 0 {
		c.Info("关系记录为空")
		return dto.Success(nil), nil
	}

	removeIndex := -1
	for _, node := range infoObj.Nodes {
		for j := 0; j < len(node.Nodes); j++ {
			v := node.Nodes[j]
			if v.CatalogId == req.ChildCatalogId {
				removeIndex = j
				break
			}
		}
		if removeIndex >= 0 {
			sList, _ := utils.DeleteSlice(node.Nodes, removeIndex)
			node.Nodes = sList.([]*catalog.CatalogInheriRelationInfoNode)
		}
		removeIndex = 1
	}
	if e := c.PutByKey(key, infoObj); nil != e {
		c.Error("更新失败:" + e.Error())
		return dto.ResultDTO{}, e
	}

	return dto.EmptySuccess(), nil
}
func (c *cataLogServiceImpl) getCatalogRelations(cid string) (string, bcso.BCSCatalogInheriRelationInfo, bserror.IBSError) {
	var infoObj bcso.BCSCatalogInheriRelationInfo
	key, e := c.GetByBuildKey(constants.OT_RELEATION_CHILD_INFO_LIST, func(bytes []byte) error {
		if len(bytes) == 0 {
			return nil
		}
		return jsonutils.ProtoUnmarshal(bytes, infoObj)
	}, []string{cid})
	if nil != e {
		c.Error("查询数据失败:" + e.Error())
		return key, infoObj, e
	}
	return key, infoObj, nil
}
func (c *cataLogServiceImpl) GetCatalogInfoById(req bo.GetCatalogInfoByIdReq) (dto.ResultDTO, bserror.IBSError) {
	c.BeforeStart("GetCatelogInfoById")
	defer c.AfterEnd()

	var res dto.ResultDTO
	var dataRes bo.GetCatalogInfoByIdResp
	oneKey := ""
	if e := c.GetByPrefix(constants.OT_CATALOG_BASICINFO, func(iteratorInterface shim.StateQueryIteratorInterface) error {
		if iteratorInterface == nil {
			return nil
		}
		if iteratorInterface.HasNext() {
			next, err := iteratorInterface.Next()
			if nil != err {
				return err
			}
			oneKey = next.Key
		}
		return nil
	}, []string{req.CatalogId}); nil != e {
		c.Error("通过前缀查询失败:" + e.Error())
		return res, e
	}

	if len(oneKey) == 0 {
		res.Code = bserror.DATA_NOT_UPLOADED.Code
		res.Msg = "数据链上不存在"
		return res, nil
	}
	 basicInfo:=&bcso.BCSCatalogBasicInfo{
		 CatalogBasicInfo: &catalog.CatalogBasicInfo{},
	 }
	if e := c.GetByKey(oneKey, func(bytes []byte) error {
		if len(bytes) == 0 {
			return nil
		}
		return jsonutils.ProtoUnmarshal(bytes, basicInfo)
	}); nil != e {
		c.Error("查询数据失败:" + e.Error())
		return res, e
	}
	if len(basicInfo.CatalogId) == 0 {
		res.Code = bserror.DATA_NOT_UPLOADED.Code
		res.Msg = "数据链上不存在"
		return res, nil
	}
	dataRes.CatalogId,dataRes.UploadVersion,dataRes.CatalogOwnerPlatformId=basicInfo.CatalogId,basicInfo.UploadVersion,basicInfo.CatalogOwnerPlatformId
	dataRes.ShowVersionId=basicInfo.ShowVersionId
	dataRes.ShowVersion=basicInfo.ShowVersion
	dataRes.StateEnumId=basicInfo.StateEnumId
	dataRes.PublishTime=basicInfo.PublishTime
	dataRes.CreateDraftTime=basicInfo.CreateDraftTime
	dataRes.StartPublicityTime=basicInfo.StartPublicityTime
	dataRes.EndPublicityTime=basicInfo.EndPublicityTime
	return dto.Success(dataRes), nil
}

// 1. 继承关系上链
func (c *cataLogServiceImpl) AddCatalogInheritDetail(req models.CatalogInheritanceUploadReq) (models.CatalogInheritanceUploadResp, bserror.IBSError) {
	c.BeforeStart("CatalogInheritanceUpload#继承关系上链")
	defer c.AfterEnd()

	var (
		res models.CatalogInheritanceUploadResp
	)
	c.Info("校验当前节点:" + req.NodeId + ",对该父目录" + req.ParentCatalogId + "是否具有继承权")
	if permission, ibsError := c.checkCatalogPermission(req.NodeId, req.ParentCatalogId, 0); nil != ibsError {
		c.Error("跨链查询校验权限失败:" + ibsError.Error())
		return res, ibsError
	} else if !permission {
		c.Error("当前节点无继承权")
		return res, bserror.NewBSError(bserror.BAD_REQUEST_NO_AUTH, "未授权,不具有继承权", bsmodule.ALL_MODULE)
	}

	c.Info("校验父目录是否已经上链,父目录ID为:" + req.ParentCatalogId)

	_, exist, ibsError := c.CheckExist(constants.OT_CATALOG_BASICINFO, []string{req.ParentCatalogId, strconv.Itoa(int(req.ParentCatalogUploadVersion))})
	if nil != ibsError {
		return models.CatalogInheritanceUploadResp{}, ibsError
	} else if !exist {
		c.Error("该父目录未上链,父目录iD:" + req.ParentCatalogId + ",父目录版本号:" + strconv.Itoa(int(req.ParentCatalogUploadVersion)))
		return models.CatalogInheritanceUploadResp{}, bserror.WithBSError(bserror.DATA_NOT_UPLOADED, "该父目录未上链,父目录iD:"+req.ParentCatalogId+",父目录版本号:"+strconv.Itoa(int(req.ParentCatalogUploadVersion)))
	}
	if e := c.updateParentRelationInfos(req); nil != e {
		return models.CatalogInheritanceUploadResp{}, bserror.WithBSError(e, "维护继承关系表失败")
	}

	return res, nil
}

// 目录上链
func (c *cataLogServiceImpl) AddOrUpdateCatalog(req models.CatalogUploadReq) (models.CatalogUploadResp, bserror.IBSError) {
	c.BeforeStart("CatalogUpload#目录上链")
	defer c.AfterEnd()

	var (
		resp models.CatalogUploadResp
	)

	c.Info("判断该节点是否有上传权的权限,该节点ID为:" + req.Req.CatalogBasicInfo.CatalogOwnerPlatformId)
	if permission, ibsError := c.checkCatalogPermission(req.Req.CatalogBasicInfo.CatalogOwnerPlatformId, req.Req.CatalogBasicInfo.CatalogId, 0); nil != ibsError {
		return models.CatalogUploadResp{}, ibsError
	} else if !permission {
		return models.CatalogUploadResp{}, bserror.WithCodeMsg(bserror.BAD_REQUEST_NO_AUTH, "无权上传该目录")
	}

	if e := c.uploadCatalogBaseInfo(req); nil != e {
		return models.CatalogUploadResp{}, e
	}
	if e := c.uploadCatalogEncoder(req.Req); nil != e {
		return models.CatalogUploadResp{}, e
	}
	if e := c.uploadCatalogTag(req.Req); nil != e {
		return models.CatalogUploadResp{}, e
	}
	if e := c.uploadTableConstraints(req.Req); nil != e {
		return models.CatalogUploadResp{}, e
	}

	return resp, nil
}

// 主数据上链
func (c *cataLogServiceImpl) AddMainData(req models.CatalogMainDataUploadReq) (models.CatalogMainDataUploadResp, bserror.IBSError) {
	c.BeforeStart("CatalogMainDataUpload#主数据上链")
	defer c.AfterEnd()

	// 1. 判断目录是否已经上链
	// 2. 判断目录是否包含该数据项

	var (
		res models.CatalogMainDataUploadResp
	)

	basicInfo := new(bcso.BCSCatalogBasicInfo)
	if _, e := c.GetByBuildKey(constants.OT_CATALOG_BASICINFO, func(bytes []byte) error {
		if e := jsonutils.ProtoUnmarshal(bytes, basicInfo); nil != e {
			return e
		}
		return nil
	}, []string{req.Req.CatalogId, strconv.Itoa(int(req.Req.CatalogUploadVersion))}); nil != e {
		return models.CatalogMainDataUploadResp{}, bserror.NewInternalServerError(e, "查询目录失败")
	}
	if basicInfo.CatalogId == "" {
		c.Error("所属目录尚未上链,目录id为:" + req.Req.CatalogId + ",version为:" + strconv.Itoa(int(req.Req.CatalogUploadVersion)))
		return models.CatalogMainDataUploadResp{}, bserror.NewBSError(bserror.BAD_REQUEST_ERROR_CODE, "所属目录尚未上链", bsmodule.ALL_MODULE)
	}

	c.Info("查询catalogId为:" + req.Req.CatalogId + ",的主数据存储相关信息记录")
	obj := new(catalog.BCSCatalogMainData)
	key, e := c.GetByBuildKey(constants.OT_MAINDATA, func(bytes []byte) error {
		if len(bytes) == 0 {
			return nil
		}
		return jsonutils.ProtoUnmarshal(bytes, obj)
	}, []string{req.Req.CatalogId})
	if nil != e {
		c.Error("查询数据失败:" + e.Error())
		return models.CatalogMainDataUploadResp{}, e
	}

	for _, node := range obj.Nodes {
		if node.DataID == req.Req.DataID {
			c.Error("重复上链,dataId为:" + req.Req.DataID + ",已经上链")
			return models.CatalogMainDataUploadResp{}, bserror.DUPLICATE_KEY_ERROR
		}
	}
	publicNodes := make([]*catalog.BCSCatalogMainDataDefinitionNode, 0)
	for _, node := range req.Req.PublicFieldList {
		publicNodes = append(publicNodes, &catalog.BCSCatalogMainDataDefinitionNode{
			FileDefinition: node,
		})
	}
	obj.Nodes = append(obj.Nodes, &catalog.BCSCatalogMainDataNode{
		DataID:                   req.Req.DataID,
		DataAuthDetailId:         req.Req.DataAuthDetailId,
		DataItemDefinitionId:     req.Req.DataItemDefinitionId,
		CatalogId:                req.Req.CatalogId,
		CatalogUploadVersion:     req.Req.CatalogUploadVersion,
		DataItemDefinitionName:   req.Req.DataItemDefinitionName,
		DataItemDefinitioVersion: req.Req.DataItemDefinitionVersion,
		Nodes:                    publicNodes,
		CryptoMessage:            req.Req.CryptoMessage,
	})

	if e := c.PutByKey(key, obj); nil != e {
		c.Error("更新主数据失败:" + e.Error())
		return models.CatalogMainDataUploadResp{}, e
	}

	return res, nil
}

func (c *cataLogServiceImpl) updateParentRelationInfos(detail models.CatalogInheritanceUploadReq) bserror.IBSError {
	c.BeforeStart("updateParentRelationInfos")
	defer c.AfterEnd()

	var infoObejct *bcso.BCSCatalogInheriRelationInfo
	// 维护所有的父目录信息
	key, err := c.GetByBuildKey(constants.OT_RELEATION_CHILD_INFO_LIST, func(bytes []byte) error {
		if len(bytes) == 0 {
			return nil
		}
		return jsonutils.ProtoUnmarshal(bytes, infoObejct)
	}, detail.ParentCatalogId)
	if nil != err {
		c.Error("查询数据失败:" + err.Error())
		return err
	}
	if infoObejct == nil {
		infoObejct = &bcso.BCSCatalogInheriRelationInfo{
			CatalogInheriRelationInfo: &catalog.CatalogInheriRelationInfo{
				Nodes: make([]*catalog.CatalogInheriRelationVersionNode, 0),
			},
		}
	}
	var versionNode *catalog.CatalogInheriRelationVersionNode
	if len(infoObejct.Nodes) != 0 {
		// 判断对应的的父版本
		for _, node := range infoObejct.Nodes {
			if node.Version == detail.ParentCatalogUploadVersion {
				versionNode = node
				break
			}
		}
	}

	if versionNode == nil {
		versionNode = &catalog.CatalogInheriRelationVersionNode{
			Version: detail.ParentCatalogUploadVersion,
			Nodes:   make([]*catalog.CatalogInheriRelationInfoNode, 0),
		}
	}

	// 判断是否重复继承
	for _, node := range versionNode.Nodes {
		if node.CatalogId == detail.CatalogId {
			return bserror.WithCodeMsg(bserror.BAD_REQUEST_ERROR_CODE, "不可以重复继承,该目录已经继承了该父目录")
		}
	}
	versionNode.Nodes = append(versionNode.Nodes, &catalog.CatalogInheriRelationInfoNode{
		CatalogId: detail.CatalogId,
	})
	return c.PutByKey(key, infoObejct)
}

// 目录基本信息上链
func (c *cataLogServiceImpl) uploadCatalogBaseInfo(req models.CatalogUploadReq) bserror.IBSError {
	c.BeforeStart("uploadCatalogBaseInfo#目录基本信息上链")
	defer c.AfterEnd()

	if req.Req.CatalogBasicInfo.UploadVersion > 1 {
		c.Info("当前目录上链的版本号大于1,为:" + strconv.Itoa(int(req.Req.CatalogBasicInfo.UploadVersion)) + ",需要判断上一个版本是否存在")
		if _, exist, e := c.CheckExist(constants.OT_CATALOG_BASICINFO, []string{req.Req.CatalogBasicInfo.CatalogId, strconv.Itoa(int(req.Req.CatalogBasicInfo.UploadVersion - 1))}); nil != e {
			c.Error("查询失败:" + e.Error())
			return e
		} else if !exist {
			c.Error("上一个版本不存在,预期不符,上一个版本必须存在")
			return bserror.NewBSError(bserror.BAD_REQUEST_ERROR_CODE, "判断不可中途跳跃,上一个版本必须先上链,当前版本"+strconv.Itoa(int(req.Req.CatalogBasicInfo.UploadVersion))+","+
				"的上一个版本:"+strconv.Itoa(int(req.Req.CatalogBasicInfo.UploadVersion-1))+",链上不存在", bsmodule.ALL_MODULE)
		}
	}
	storeInfo := bcso.BCSCatalogBasicInfo{
		CatalogBasicInfo: req.Req.CatalogBasicInfo,
	}
	if e := c.PutInterfaceData(constants.OT_CATALOG_BASICINFO, storeInfo, []string{storeInfo.CatalogId, strconv.Itoa(int(storeInfo.UploadVersion))}); nil != e {
		c.Error("目录上链,保存基本信息失败:" + e.Error())
		return e
	}
	if len(req.Req.CatalogBasicInfo.ParentCatalogId) > 0 {
		c.Info("父目录信息不为空,父目录的ID为:" + req.Req.CatalogBasicInfo.ParentCatalogId + ",父目录版本号为:" + strconv.Itoa(int(req.Req.CatalogBasicInfo.ParentCatalogUploadVersion)))
		if _, e := c.AddCatalogInheritDetail(models.CatalogInheritanceUploadReq{
			CatalogId:                  req.Req.CatalogBasicInfo.CatalogId,
			CatalogUploadVersion:       req.Req.CatalogBasicInfo.UploadVersion,
			ParentCatalogId:            req.Req.CatalogBasicInfo.ParentCatalogId,
			ParentCatalogUploadVersion: req.Req.CatalogBasicInfo.ParentCatalogUploadVersion,
			NodeId:                     req.Req.CatalogBasicInfo.CatalogOwnerPlatformId,
			InheritDetailId:            req.Req.CatalogBasicInfo.InheritDetailId,
		}); nil != e {
			c.Error("目录继承关系更新失败:" + e.Error())
			return e
		}
	}

	return nil
}

// 编码集上链
func (c *cataLogServiceImpl) uploadCatalogEncoder(cg *catalog.Catalog) bserror.IBSError {
	c.BeforeStart("uploadCatalogEncoder#编码集上链")
	defer c.AfterEnd()

	req := cg.GlossaryList
	nodes := make([]*catalog.CatalogDataTypeInfoNode, 0)

	for _, r := range req {
		codeList := make([]*catalog.CatalogDataTypeInfoNodeCodeInfo, 0)
		for _, v := range r.EnumValueList {
			codeList = append(codeList, &catalog.CatalogDataTypeInfoNodeCodeInfo{
				CodeName: v.CodeName,
				Content:  v.Content,
			})
		}
		nodes = append(nodes, &catalog.CatalogDataTypeInfoNode{
			DataType:            r.DataType,
			EnumId:          r.EnumId,
			EnumEnglishName: r.EnumEnglishName,
			EnumChineseName: r.EnumChineseName,
			EnumSourceId:    r.EnumSourceId,
			EnumValueList:   codeList,
		})
	}
	obj := &bcso.BCSCatalogDataTypeInfo{
		CatalogDataTypeInfo: &catalog.CatalogDataTypeInfo{
			Nodes: nodes,
		},
	}
	if e := c.PutInterfaceData(constants.OT_CATALOG_DATATYPE, obj, []string{cg.CatalogBasicInfo.CatalogId, strconv.Itoa(int(cg.CatalogBasicInfo.UploadVersion))}); nil != e {
		return e
	}
	return nil
}

// 标签上链
func (c *cataLogServiceImpl) uploadCatalogTag(cg *catalog.Catalog) bserror.IBSError {
	c.BeforeStart("uploadTag#标签上链")
	defer c.AfterEnd()

	req := cg.TagList

	nodes := make([]*catalog.BCSTagNode, 0)
	for _, v := range req {
		node := &catalog.BCSTagNode{
			TagId:                 v.TagId,
			TagName:               v.TagName,
			TagType:               v.TagType,
			CreateTagNodeID:       v.CreateTagNodeID,
			ToWhichDataItemIdList: v.ToWhichDataItemIdList,
		}
		nodes = append(nodes, node)
	}
	obj := &catalog.BCSTag{
		Nodes: nodes,
	}
	if e := c.PutInterfaceData(constants.OT_CATALOG_TAG, obj, []string{cg.CatalogBasicInfo.CatalogId, strconv.Itoa(int(cg.CatalogBasicInfo.UploadVersion))}); nil != e {
		c.Error("上链标签失败:" + e.Error())
		return e
	}

	return nil
}

// 表关系上链
func (c *cataLogServiceImpl) uploadTableConstraints(cg *catalog.Catalog) bserror.IBSError {
	c.BeforeStart("uploadTableConstraints#表关系上链")
	defer c.AfterEnd()

	req := cg.ConstraintList
	obj := &catalog.BCSForeignKeyConstraint{}
	nodes := make([]*catalog.BCSForeignKeyConstraintNode, 0)
	for _, v := range req {
		node := &catalog.BCSForeignKeyConstraintNode{
			ConstraintId:               v.ConstraintId,
			LeftTableId:                v.LeftTableId,
			LeftTableFieldEnglishName:  v.LeftTableFieldEnglishName,
			RightTableId:               v.RightTableId,
			RightTableFieldEnglishName: v.RightTableFieldEnglishName,
		}
		nodes = append(nodes, node)
	}
	obj.Nodes = nodes

	if e := c.PutInterfaceData(constants.OT_CATALOG_TABLECONSTRAINT, obj, []string{cg.CatalogBasicInfo.CatalogId, strconv.Itoa(int(cg.CatalogBasicInfo.UploadVersion))}); nil != e {
		c.Error("表关系上链失败:" + e.Error())
		return e
	}
	return nil
}

func (c *cataLogServiceImpl) checkCatalogPermission(ownerId string, catLogId string, authValue int) (bool, bserror.IBSError) {
	c.BeforeStart("checkCatalogPermission")
	defer c.AfterEnd()

	// FIXME ,补齐参数
	catalogAuthReq := bo.CatalogAuthenticationReq{
		CatalogOwnerPlatformId: ownerId,
		Authentication:         authValue,
	}
	args, err := json.Marshal(catalogAuthReq)
	if nil != err {
		return false, bserror.NewInternalServerError(err, "序列化跨链请求对象失败")
	}
	resBytes, bsError := c.CallOverChainWithArgument(constants.OC_CATALOG_AUTH, args)
	if nil != bsError {
		return false, bsError
	}
	res := bo.CatalogAuthenticationResp{}
	if err = json.Unmarshal(resBytes, &res); nil != err {
		return false, bserror.NewInternalServerError(err, "反序列化跨链返回值对象失败,数据为:"+string(resBytes))
	}

	return res.Allowed, nil
}
