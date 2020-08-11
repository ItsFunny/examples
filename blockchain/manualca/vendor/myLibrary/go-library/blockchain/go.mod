module myLibrary/go-library/blockchain

go 1.14

require (
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20200511190512-bcfeb58dd83a // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20200506201313-25f6564b9ac4
	github.com/hyperledger/fabric-sdk-go v0.0.0
	github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric v0.0.0
	github.com/pkg/errors v0.9.1
	go.uber.org/atomic v1.6.0
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
	myLibrary/go-library/common v0.0.0
	myLibrary/go-library/go v0.0.0
)

replace github.com/hyperledger/fabric-sdk-go => /Users/joker/Desktop/go-dependency/github.com/hyperledger/fabric-sdk-go

replace github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric => /Users/joker/Desktop/go-dependency/github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric

replace code.google.com/log4go => /Users/joker/go/src/code.google.com/log4go

replace myLibrary/go-library/go => /Users/joker/go/src/myLibrary/go-library/go

replace myLibrary/go-library/common => /Users/joker/go/src/myLibrary/go-library/common
		replace github.com/FactomProject/basen => /Users/joker/Desktop/go-dependency/github.com/FactomProject/basen