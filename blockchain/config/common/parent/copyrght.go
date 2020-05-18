/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-02-19 13:48 
# @File : copyrght.go
# @Description : 
# @Attention : 
*/
package parent

type ItemCopyrightInfoParent struct {
	ItemUpload2BlockChainReqParent
	// 签名
	Signature string `json:"signature"`
}
type ItemUpload2BlockChainReqParent struct {
	// tsa
	// tsa 给予的时间戳日期
	// private String timeStamp;
	TimeStamp string `json:"timeStamp"`
	// // tsa给予的序列号信息
	// private String serialNo;
	SerialNo string `json:"serialNo"`
	//
	//
	// // 作品本身
	// // 作品的MD5文件码,既itemCode
	// private String itemCode;
	ItemCode string `json:"itemCode"`

	// // 作品数据库ID
	// private long itemId;
	ItemId int `json:"itemId,string"`
	// // 作品名称
	// private String itemName;
	ItemName string `json:"itemName"`
	// // 作品类型,是音乐还是文学等
	// private int itemCategory;
	ItemCategory int `json:"itemCategory"`
	// // 作品描述
	// private String itemDescription;
	ItemDescription string `json:"itemDescription"`
	//
	// // 作品上链时候的创作类别: 连载中还是已完结
	// private int itemType;
	ItemType int `json:"itemType"`
	// // 作品具体分类,如文学: 散文,小说等 ,图片: 漫画 插画 相片 等等
	// private int itemSpecificCategory;
	ItemSpecificCategory int `json:"itemSpecificCategory"`
	// // 作品风格
	// private int itemStyle;
	ItemStyle int `json:"itemStyle"`
	// // 作品上链时候的状态,是已完结还是更新中
	// private int itemStatus;
	ItemStatus int `json:"itemStatus"`
	//
	//
	// // 不同类型的作品链上存储结构也不同,因而采用Object来代替
	// private Object itemSpecificData;
	ItemSpecificData interface{} `json:"itemSpecificData"`
	//
	//
	// // 用户信息
	// // 作者名称
	// private String itemAuthorName;
	ItemAuthorName string `json:"itemAuthorName"`
	// // 上传日期
	// private long createDate;
	CreateDate int64 `json:"createDate,string"`

	// // 上传用户
	// private String createUser;
	CreateUser string `json:"createUser"`
	// // 上传用户ID
	// private Long createUserId;
	CreateUserId int `json:"createUserId,string"`
	// // 最近一次更新日期
	// private long lastUpdateDate;
	LastUpdateDate int64 `json:"lastUpdateDate,string"`
	// // 最近一次更新用户
	// private String lastUpdateUser;
	LastUpdateUser string `json:"lastUpdateUser"`
	// // 最近一次更新的用户ID
	// private Long lastUpdateUserId;
	LastUpdateUserId int `json:"lastUpdateUserId,string"`


	TxID string `json:"txId"`
}
type ItemUpload2BlockChainRespParent struct {
	Dna string `json:"dna"`
	TxId string `json:"txid"`
}


// 获取版权信息
type GetItemCopyrightReqParent struct {
	ItemId int `json:"itemId,string"`
	UserId int `json:"userId,string"`
}

type GetItemCopyrightRespParent struct {
	ItemCopyrightInfoParent
	Dna string `json:"dna"`
	PublicKey string `json:"publicKey"`
}