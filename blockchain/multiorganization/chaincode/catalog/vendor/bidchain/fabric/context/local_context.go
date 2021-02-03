package context

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	lb "github.com/syndtr/goleveldb/leveldb"
	"runtime"
	"sync"
)

/*
已经废弃使用， 扩展使用不便
用于本地测试使用
实例化LocalContext对象后，必须:
1. 在存储和查询数据前指定通道名称，这里指定数据库库名称和通道名称相同。
2. 设置chaincode名称， 用于区分相同通道下不同chaincode

注意: 本地测试完数据后 需要及时手动删除临时生成的数据库
 */

const (
	DEFAULT_CHANNEL_ID      string = "LOCAL"
	DEFULAT_DATABASE_PREFIX        = ""
)

var (
	DATABASE_PREFIX = DEFULAT_DATABASE_PREFIX
)

type LocalContext struct {
	sync.Mutex
	db            *lb.DB
	stub          shim.ChaincodeStubInterface
	funcName      string
	params        []string
	channelID     string
	chaincodeName string
}

//func NewLocalContext(stub shim.ChaincodeStubInterface) IBidchainContext {
//	ctx := &LocalContext{
//		stub:      stub,
//		channelID: DEFAULT_CHANNEL_ID,
//	}
//	return ctx
//}

// 这里假定链码的名称不为空字符串
func (lc *LocalContext) GetState(key string) ([]byte, error) {
	if lc.db == nil {
		lc.openDB()
	}
	prefix := lc.GetChaincodeName()
	key = prefix + "/" + key

	lc.Lock()
	data, err := lc.db.Get([]byte(key), nil)
	lc.Unlock()
	if err == lb.ErrNotFound {
		err = nil
	}
	return data, err
}

func (lc *LocalContext) PutState(key string, value []byte) error {
	if lc.db == nil {
		lc.openDB()
	}

	prefix := lc.GetChaincodeName()
	key = prefix + "/" + key

	lc.Lock()
	err := lc.db.Put([]byte(key), value, nil)
	lc.Unlock()
	return err
}

func (lc *LocalContext) DelState(key string) error {
	if lc.db == nil {
		lc.openDB()
	}
	prefix := lc.GetChaincodeName()
	key = prefix + "/" + key

	lc.Lock()
	err := lc.db.Delete([]byte(key), nil)
	lc.Unlock()
	return err
}

func (lc *LocalContext) GetFunctionAndParameters() (function string, params []string) {
	if lc.funcName == "" {
		panic("本次测试需要先调用SetFunctionName和SetParameters，再获取对应的函数名和方法名称")
	}
	return lc.funcName, lc.params
}

func (lc *LocalContext) SetFunctionName(name string) {
	lc.funcName = name
}

func (lc *LocalContext) SetParameters(params []string) {
	lc.params = params
}

func (lc *LocalContext) GetFunctionName() string {
	return lc.funcName
}

func (lc *LocalContext) GetParameters() []string {
	return lc.params
}

func (lc *LocalContext) SetChannelID(cid string) {
	if cid != "" && lc.channelID != cid {
		if lc.db != nil {
			panic("SetChannelID error: 已加载数据库文件")
		}
	}
	lc.channelID = cid
	if lc.db != nil {
		panic("SetChannelID error: 已加载数据库文件")
	}
	lc.openDB()
}

func (lc *LocalContext) openDB() {
	// 一个channel对应一个账本，这里指定数据库名称和channel名称相同
	dbName := DATABASE_PREFIX + lc.GetChannelID() + ".db"
	fmt.Printf("Creating database: %s\n", dbName)
	db, err := lb.OpenFile(dbName, nil)
	if err != nil {
		panic(fmt.Sprintf("LocalContext openDB %s error: %s", dbName, err.Error()))
	}
	lc.db = db
}

func (lc *LocalContext) GetChannelID() string {
	return lc.channelID
}

func (lc *LocalContext) SetChaincodeName(cname string) {
	lc.chaincodeName = cname
}

func (lc *LocalContext) GetChaincodeName() string {
	return lc.chaincodeName
}

func init() {
	if runtime.GOOS == "windows" {
		DATABASE_PREFIX = "D:/"
	}
}
