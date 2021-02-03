protoc -I=. --go_out=. %1
protoc -I=../store --go_out=../store %1
protoc -I=. --java_out=D:\bidsun_server\bidsun-metadata\bidsun-metadata-common\src\main\java %1
copy %1 D:\bidsun_server\bidsun-metadata\bidsun-metadata-common\src\main\java\cn\bidsun\metadata\common\protodao