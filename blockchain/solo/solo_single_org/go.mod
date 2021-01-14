module examples/blockchain/solo/solo_single_org

go 1.14

require myLibrary/go-library/go v0.0.1


replace myLibrary/go-library/go => /Users/joker/go/src/myLibrary/go-library/go

replace github.com/hyperledger/fabric-sdk-go => /Users/joker/Desktop/hyperledger/fabric-sdk-go

replace github.com/tyler-smith/go-bip32 => /Users/joker/go/src/github.com/tyler-smith/go-bip32

replace code.google.com/log4go => /Users/joker/go/src/code.google.com/log4go

replace github.com/FactomProject/basen => /Users/joker/Desktop/go-dependency/github.com/FactomProject/basen

replace github.com/FactomProject/btcutilecc => /Users/joker/Desktop/go-dependency/github.com/FactomProject/btcutilecc

exclude  github.com/go-kit/kit v0.9.0