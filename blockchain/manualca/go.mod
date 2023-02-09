module examples/blockchain/manualca

require (
	github.com/gin-gonic/gin v1.7.7
	myLibrary/go-library/blockchain v0.0.0
)

replace myLibrary/go-library/blockchain => /Users/joker/go/src/myLibrary/go-library/blockchain

replace github.com/hyperledger/fabric-sdk-go => /Users/joker/Desktop/go-dependency/github.com/hyperledger/fabric-sdk-go

replace github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric => /Users/joker/Desktop/go-dependency/github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric

replace code.google.com/log4go => /Users/joker/go/src/code.google.com/log4go

replace myLibrary/go-library/go => /Users/joker/go/src/myLibrary/go-library/go

replace myLibrary/go-library/common => /Users/joker/go/src/myLibrary/go-library/common

replace github.com/FactomProject/basen => /Users/joker/Desktop/go-dependency/github.com/FactomProject/basen
