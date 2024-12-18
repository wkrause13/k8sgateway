// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/socket_option.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *SocketOption) Clone() proto.Message {
	var target *SocketOption
	if m == nil {
		return target
	}
	target = &SocketOption{}

	target.Description = m.GetDescription()

	target.Level = m.GetLevel()

	target.Name = m.GetName()

	target.State = m.GetState()

	switch m.Value.(type) {

	case *SocketOption_IntValue:

		target.Value = &SocketOption_IntValue{
			IntValue: m.GetIntValue(),
		}

	case *SocketOption_BufValue:

		if m.GetBufValue() != nil {
			newArr := make([]byte, len(m.GetBufValue()))
			copy(newArr, m.GetBufValue())
			target.Value = &SocketOption_BufValue{
				BufValue: newArr,
			}
		} else {
			target.Value = &SocketOption_BufValue{
				BufValue: nil,
			}
		}

	}

	return target
}
