// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/stats/stats.proto

package stats

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
func (m *Stats) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Stats)
	if !ok {
		that2, ok := that.(Stats)
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

	if len(m.GetVirtualClusters()) != len(target.GetVirtualClusters()) {
		return false
	}
	for idx, v := range m.GetVirtualClusters() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualClusters()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualClusters()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualCluster) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualCluster)
	if !ok {
		that2, ok := that.(VirtualCluster)
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

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetPattern(), target.GetPattern()) != 0 {
		return false
	}

	if strings.Compare(m.GetMethod(), target.GetMethod()) != 0 {
		return false
	}

	return true
}
