module democc

go 1.13

require (
	github.com/hyperledger/fabric v2.0.1+incompatible
	github.com/hyperledger/fabric-sdk-go v1.0.0-beta1 // indirect
	github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric v0.0.0-20190524192706-bfae339c63bf // indirect
	vlink.com/v2/vlink-common v0.0.3 // indirect
)

replace vlink.com/v2/vlink-common => /Users/joker/go/src/vlink.com/v2/vlink-common

replace myLibrary/go-library/go => /Users/joker/go/src/myLibrary/go-library/go

replace google.golang.org/grpc => /Users/joker/go/src/google.golang.org/grpc

replace github.com/hyperledger/fabric => /Users/joker/go/src/github.com/hyperledger/fabric

replace code.google.com/log4go => /Users/joker/go/src/code.google.com/log4go

replace github.com/milagro-crypto/amcl/version3/go/amcl => /Users/joker/go/src/github.com/milagro-crypto/amcl/version3/go/amcl

replace github.com/tyler-smith/go-bip32 => /Users/joker/go/src/github.com/tyler-smith/go-bip32

replace github.com/docker/docker => /Users/joker/go/src/github.com/docker/docker
