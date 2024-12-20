// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/http_uri.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *HttpUri) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpUri)
	if !ok {
		that2, ok := that.(HttpUri)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetUri(), target.GetUri()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	switch m.HttpUpstreamType.(type) {

	case *HttpUri_Cluster:
		if _, ok := target.HttpUpstreamType.(*HttpUri_Cluster); !ok {
			return false
		}

		if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.HttpUpstreamType != target.HttpUpstreamType {
			return false
		}
	}

	return true
}
