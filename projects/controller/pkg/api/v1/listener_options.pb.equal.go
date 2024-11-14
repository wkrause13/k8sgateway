// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/listener_options.proto

package v1

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
func (m *ListenerOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ListenerOptions)
	if !ok {
		that2, ok := that.(ListenerOptions)
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

	if h, ok := interface{}(m.GetAccessLoggingService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAccessLoggingService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAccessLoggingService(), target.GetAccessLoggingService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtensions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtensions(), target.GetExtensions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPerConnectionBufferLimitBytes(), target.GetPerConnectionBufferLimitBytes()) {
			return false
		}
	}

	if len(m.GetSocketOptions()) != len(target.GetSocketOptions()) {
		return false
	}
	for idx, v := range m.GetSocketOptions() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSocketOptions()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSocketOptions()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetProxyProtocol()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProxyProtocol()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProxyProtocol(), target.GetProxyProtocol()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConnectionBalanceConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionBalanceConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionBalanceConfig(), target.GetConnectionBalanceConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetListenerAccessLoggingService()).(equality.Equalizer); ok {
		if !h.Equal(target.GetListenerAccessLoggingService()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetListenerAccessLoggingService(), target.GetListenerAccessLoggingService()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTcpStats()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTcpStats()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTcpStats(), target.GetTcpStats()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ConnectionBalanceConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionBalanceConfig)
	if !ok {
		that2, ok := that.(ConnectionBalanceConfig)
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

	if h, ok := interface{}(m.GetExactBalance()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExactBalance()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExactBalance(), target.GetExactBalance()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ConnectionBalanceConfig_ExactBalance) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnectionBalanceConfig_ExactBalance)
	if !ok {
		that2, ok := that.(ConnectionBalanceConfig_ExactBalance)
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

	return true
}
