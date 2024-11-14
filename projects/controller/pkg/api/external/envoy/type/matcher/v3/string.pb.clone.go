// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/type/matcher/v3/string.proto

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
func (m *StringMatcher) Clone() proto.Message {
	var target *StringMatcher
	if m == nil {
		return target
	}
	target = &StringMatcher{}

	target.IgnoreCase = m.GetIgnoreCase()

	switch m.MatchPattern.(type) {

	case *StringMatcher_Exact:

		target.MatchPattern = &StringMatcher_Exact{
			Exact: m.GetExact(),
		}

	case *StringMatcher_Prefix:

		target.MatchPattern = &StringMatcher_Prefix{
			Prefix: m.GetPrefix(),
		}

	case *StringMatcher_Suffix:

		target.MatchPattern = &StringMatcher_Suffix{
			Suffix: m.GetSuffix(),
		}

	case *StringMatcher_SafeRegex:

		if h, ok := interface{}(m.GetSafeRegex()).(clone.Cloner); ok {
			target.MatchPattern = &StringMatcher_SafeRegex{
				SafeRegex: h.Clone().(*RegexMatcher),
			}
		} else {
			target.MatchPattern = &StringMatcher_SafeRegex{
				SafeRegex: proto.Clone(m.GetSafeRegex()).(*RegexMatcher),
			}
		}

	}

	return target
}

// Clone function
func (m *ListStringMatcher) Clone() proto.Message {
	var target *ListStringMatcher
	if m == nil {
		return target
	}
	target = &ListStringMatcher{}

	if m.GetPatterns() != nil {
		target.Patterns = make([]*StringMatcher, len(m.GetPatterns()))
		for idx, v := range m.GetPatterns() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Patterns[idx] = h.Clone().(*StringMatcher)
			} else {
				target.Patterns[idx] = proto.Clone(v).(*StringMatcher)
			}

		}
	}

	return target
}
