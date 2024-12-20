// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/matching/custom_matchers/server_name/v3/server_name_matcher.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_cncf_xds_go_xds_type_matcher_v3 "github.com/cncf/xds/go/xds/type/matcher/v3"
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
func (m *ServerNameMatcher) Clone() proto.Message {
	var target *ServerNameMatcher
	if m == nil {
		return target
	}
	target = &ServerNameMatcher{}

	if m.GetServerNameMatchers() != nil {
		target.ServerNameMatchers = make([]*ServerNameMatcher_ServerNameSetMatcher, len(m.GetServerNameMatchers()))
		for idx, v := range m.GetServerNameMatchers() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ServerNameMatchers[idx] = h.Clone().(*ServerNameMatcher_ServerNameSetMatcher)
			} else {
				target.ServerNameMatchers[idx] = proto.Clone(v).(*ServerNameMatcher_ServerNameSetMatcher)
			}

		}
	}

	return target
}

// Clone function
func (m *ServerNameMatcher_ServerNameSetMatcher) Clone() proto.Message {
	var target *ServerNameMatcher_ServerNameSetMatcher
	if m == nil {
		return target
	}
	target = &ServerNameMatcher_ServerNameSetMatcher{}

	if m.GetServerNames() != nil {
		target.ServerNames = make([]string, len(m.GetServerNames()))
		for idx, v := range m.GetServerNames() {

			target.ServerNames[idx] = v

		}
	}

	if h, ok := interface{}(m.GetOnMatch()).(clone.Cloner); ok {
		target.OnMatch = h.Clone().(*github_com_cncf_xds_go_xds_type_matcher_v3.Matcher_OnMatch)
	} else {
		target.OnMatch = proto.Clone(m.GetOnMatch()).(*github_com_cncf_xds_go_xds_type_matcher_v3.Matcher_OnMatch)
	}

	return target
}
