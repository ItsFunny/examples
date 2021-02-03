package jsonutils

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func Marshal(pb proto.Message) (data []byte, err error) {
	var buf bytes.Buffer
	m := jsonpb.Marshaler{}
	err = m.Marshal(&buf, pb)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(data string, pb proto.Message) error {
	return jsonpb.UnmarshalString(data, pb)
}

func ProtoMarshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

func ProtoUnmarshal(buf []byte, pb proto.Message) error {
	return proto.Unmarshal(buf, pb)
}
