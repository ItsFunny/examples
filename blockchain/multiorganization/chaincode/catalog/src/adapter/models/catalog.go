/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/1/27 15:02
# @File : catalog.go
# @Description :
# @Attention :
*/
package models

import (
	"bidchain/base/dataengine/gojsonschema"
	"bidchain/protocol/transport/catalog"
	"encoding/json"
	"errors"
	"strings"
)

var (
	catalogSchemal=`{
  "type": "object",
  "properties": {
    "catalogBasicInfo": {
      "type": "object",
      "properties": {
        "catalogId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "catalogName": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "uploadVersion": {
          "type": "integer",
          "minimum": 0
        },
        "showVersionId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "showVersion": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "catalogOwnerPlatformId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "stateEnumId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "level": {
          "type": "integer",
          "minimum": 0
        },
        "publishTime": {
          "type": "number"
        },
        "createDraftTime": {
          "type": "number"
        },
        "startPublicityTime": {
          "type": "number"
        },
        "endPublicityTime": {
          "type": "number"
        },
        "inheritDetailId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "parentCatalogId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "parentCatalogUploadVersion": {
          "type": "integer",
          "minimum": 0
        }
      },
      "required": ["catalogId","catalogName","uploadVersion","showVersionId","showVersion","catalogOwnerPlatformId","stateEnumId","level","publishTime","createDraftTime","startPublicityTime","endPublicityTime"]
    },
    "dataItemDefinitionList": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "dataItemDefinitionId": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "dataItemName": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "dataItemVersion": {
            "type": "number",
            "minimum": 0
          },
          "weight": {
            "type": "number",
            "minimum": 0
          },
          "publicFieldList": {
            "type": "array",
            "properties": {
              "dataFieldName": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "englishName": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "enumSecurityLevelId": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "weight": {
                "type": "number",
                "minimum": 0
              },
              "level": {
                "type": "number",
                "minimum": 0
              },
              "orderIndex": {
                "type": "number",
                "minimum": 0
              },
              "dataFieldEnumId": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "dataTypeDescriptor": {
                "type": "object",
                "properties": {
                  "maxLength": {
                    "type": "number",
                    "minimum": 0
                  },
                  "enumId": {
                    "type": "string",
                    "minLength": 0,
                    "maxLength": 200
                  }
                },
                "required": ["maxLength","enumId"]
              }
            },
            "required": ["dataFieldName","englishName","enumSecurityLevelId","weight","level","orderIndex","dataFieldEnumId","dataTypeDescriptor"]
          },
          "weight": {
            "type": "number",
            "minimum": 0
          },
          "mergeFieldList": {
            "type": "array",
            "properties": {
              "sourceFieldEnglishName": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "mergedFieldEnglishName": {
                "type": "string",
                "minLength": 0,
                "maxLength": 200
              },
              "mergeFieldLevel": {
                "type": "number",
                "minimum": 0
              }
            },
            "required": ["sourceFieldEnglishName","mergedFieldEnglishName","mergeFieldLevel"]
          }
        },
        "required": ["dataItemDefinitionId","dataItemName","dataItemVersion","publicFieldList","weight","mergeFieldList"]
      }
    },
    "glossaryList": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "dataType": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "enumId": {
            "type": "number",
            "minimum": 0
          },
          "enumEnglishName": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "enumChineseName": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "enumSourceId": {
            "type": "number",
            "minimum": 0
          },
          "enumValueList": {
            "type": "array",
            "items": {
              "type": "object",
              "properties": {
                "codeName": {
                  "type": "string",
                  "minLength": 0,
                  "maxLength": 200
                },
                "content": {
                  "type": "string",
                  "minLength": 0,
                  "maxLength": 200
                }
              }
            }
          }
        },
        "required": ["dataType","enumId","enumEnglishName","enumChineseName","enumSourceId","enumValueList"]
      }
    },
    "tagList": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "tagId":{
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "tagName":{
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "tagType":{
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "createTagNodeID":{
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "toWhichDataItemIdList": {
            "type": "array",
            "items": {
              "type": "string",
              "minLength": 0,
              "maxLength": 200
            }
          }
        },
        "required": ["tagId","tagName","tagType","createTagNodeID","toWhichDataItemIdList"]
      }
    },
    "constraintList": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "constraintId": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "leftTableId": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "leftTableFieldEnglishName": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "rightTableId": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          },
          "rightTableFieldEnglishName": {
            "type": "string",
            "minLength": 0,
            "maxLength": 200
          }
        },
        "required": ["constraintId","leftTableId","leftTableFieldEnglishName","rightTableId","rightTableFieldEnglishName"]
      }
    }
  }
}`

	mainDataSchemal=`
{
  "type": "object",
  "properties": {
    "dataID": {
      "type": "string",
      "minLength": 0,
      "maxLength": 200
    },
    "dataAuthDetailId": {
      "type": "string",
      "minLength": 0,
      "maxLength": 200
    },
    "dataItemDefinitionId": {
      "type": "string",
      "minLength": 0,
      "maxLength": 200
    },
    "catalogId": {
      "type": "string",
      "minLength": 0,
      "maxLength": 200
    },
    "catalogUploadVersion": {
      "type": "number",
      "minimum": 0
    },
    "dataItemDefinitionName": {
      "type": "string",
      "minLength": 0,
      "maxLength": 200
    },
    "dataItemDefinitionVersion": {
      "type": "number",
      "minimum": 0
    },
    "publicFieldList": {
      "type": "array",
      "properties": {
        "dataFieldName": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "englishName": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "enumSecurityLevelId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "weight": {
          "type": "number",
          "minimum": 0
        },
        "level": {
          "type": "number",
          "minimum": 0
        },
        "orderIndex": {
          "type": "number",
          "minimum": 0
        },
        "dataFieldEnumId": {
          "type": "string",
          "minLength": 0,
          "maxLength": 200
        },
        "dataTypeDescriptor": {
          "type": "object",
          "properties": {
            "maxLength": {
              "type": "number",
              "minimum": 0
            },
            "enumId": {
              "type": "string",
              "minLength": 0,
              "maxLength": 200
            }
          },
          "required": ["maxLength","enumId"]
        }
      },
      "required": ["dataFieldName","englishName","enumSecurityLevelId","weight","level","orderIndex","dataFieldEnumId","dataTypeDescriptor"]
    }
  },
  "required": ["dataID","dataAuthDetailId","dataItemDefinitionId","catalogId","catalogUploadVersion","dataItemDefinitionVersion","publicFieldList"]
}`
)

// 继承关系
type CatalogInheritanceUploadReq struct {
	CatalogId                  string
	CatalogUploadVersion       int32
	ParentCatalogId            string
	ParentCatalogUploadVersion int32
	NodeId                     string
	InheritDetailId            string
}

func (c CatalogInheritanceUploadReq) Validation() error {
	if len(c.CatalogId) == 0 {
		return errors.New("catalogId不可为空")
	}
	if c.CatalogUploadVersion <= 0 {
		return errors.New("CatalogUploadVersion为空")
	}
	if len(c.ParentCatalogId) == 0 {
		return errors.New("父ID不可为空")
	}
	if c.ParentCatalogUploadVersion <= 0 {
		return errors.New("ParentUploadVersion不可为空")
	}
	if len(c.NodeId) == 0 {
		return errors.New("参数nodeId不可为空")
	}
	if len(c.InheritDetailId) == 0 {
		return errors.New("参数nodeId不可为空")
	}
	return nil
}

type CatalogInheritanceUploadResp struct {
}

// 目录上链
type CatalogUploadReq struct {
	Req *catalog.Catalog
}

func (c CatalogUploadReq) Validation() error {
	schemalLoader:=gojsonschema.NewStringLoader(catalogSchemal)
	marshal, _ := json.Marshal(c.Req)
	dataLoader:=gojsonschema.NewStringLoader(string(marshal))
	res, e := gojsonschema.Validate(schemalLoader, dataLoader)
	if nil!=e{
		return errors.New("schemal校验内置错误:"+ e.Error())
	}
	errorMsg:=strings.Builder{
	}
	if !res.Valid() {
		for _, err := range res.Errors() {
			errorMsg.WriteString(err.String()+",")
		}
	}
	if errorMsg.Len()>0{
		return errors.New(errorMsg.String())
	}
	return nil
	// basicInfo := c.Req.CatalogBasicInfo
	// if len(basicInfo.CatalogOwnerPlatformId) == 0 {
	// 	return errors.New("节点id不可为空")
	// }
	// if basicInfo.UploadVersion <= 0 {
	// 	return errors.New("uploadVersion不可小于0")
	// }
	// return nil
}

type CatalogUploadResp struct {
}

// 主数据
type CatalogMainDataUploadReq struct {
	Req *catalog.Data
}

func (c CatalogMainDataUploadReq) Validation() error {
	schemalLoader:=gojsonschema.NewStringLoader(mainDataSchemal)
	marshal, _ := json.Marshal(c.Req)
	dataLoader:=gojsonschema.NewStringLoader(string(marshal))
	res, e := gojsonschema.Validate(schemalLoader, dataLoader)
	if nil!=e{
		return errors.New("schemal校验内置错误:"+ e.Error())
	}
	errorMsg:=strings.Builder{
	}
	if !res.Valid() {
		for _, err := range res.Errors() {
			errorMsg.WriteString(err.String()+",")
		}
	}
	if errorMsg.Len()>0{
		return errors.New(errorMsg.String())
	}
	return nil
}

type CatalogMainDataUploadResp struct {
}
