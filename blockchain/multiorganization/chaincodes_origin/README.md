

# 目录链



### 继承关系

- 前置过滤:
  - 要判断下是否已经有继承关系,如果有的话,需要判断下是否相同





---





# 接口:

- 每个接口返回值都是

  - ```
    {
    	"code":0,
    	"msg":"success",
    	"data":{xxxx}的形式
    }
    ```

- code 为0 默认代表成功

- 图懒,省略code,msg 等

## 目录

- 获取目录的拥有者  

  - ```
    API: GetCatalogInfoById
    Content-Type: json
    REQ:{
    	CatalogId string
    }
    RESP:{
    	code: 0
    	msg: "success"
    	data:type GetCatalogInfoByIdResp struct {
    	CatalogId              string `json:"catalogId"`
    	CatalogName            string `json:"catalogName"`
    	UploadVersion          string `json:"uploadVersion"`
    	ShowVersionId          string `json:"showVersionId"`
    	ShowVersion            string `json:"showVersion"`
    	CatalogOwnerPlatformId string `json:"catalogOwnerPlatformId"`
    	StateEnumId            string `json:"stateEnumId"`
    	PublishTime            int64 `json:"publishTime"`
    	CreateDraftTime        int64 `json:"createDraftTime"`
    	StartPublicityTime     int64 `json:"startPublicityTime"`
    	EndPublicityTime       int64 `json:"endPublicityTime"`
    }
    }
    
    ```

- 判断是否是 后代:

  - ```
    API: CheckIsDesendatantCatalog
    Content-Type: json
    REQ:{
    	ParentCatalogId string `json:"parentCatalogId"`
    	ChildCatalogId  string `json:"childCatalogId"`
    }
    RESP:{
    	"desendatant":true
    }
    ```

- 继承关系  删除表关系

  - ```
    API: RevokeCatalogRelation
    Content-Type: json
    REQ:{
    	ParentCatalogId string
    	ChildCatalogId  string
    }
    RESP:{
    	"code":0,
    	"msg":"success",
    	"data":{
    		
    	}
    }
    ```

  - 