// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/cors/cors.proto

package cors

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
func (m *CorsPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CorsPolicy)
	if !ok {
		that2, ok := that.(CorsPolicy)
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

	if len(m.GetAllowOrigin()) != len(target.GetAllowOrigin()) {
		return false
	}
	for idx, v := range m.GetAllowOrigin() {

		if strings.Compare(v, target.GetAllowOrigin()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowOriginRegex()) != len(target.GetAllowOriginRegex()) {
		return false
	}
	for idx, v := range m.GetAllowOriginRegex() {

		if strings.Compare(v, target.GetAllowOriginRegex()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowMethods()) != len(target.GetAllowMethods()) {
		return false
	}
	for idx, v := range m.GetAllowMethods() {

		if strings.Compare(v, target.GetAllowMethods()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAllowHeaders()) != len(target.GetAllowHeaders()) {
		return false
	}
	for idx, v := range m.GetAllowHeaders() {

		if strings.Compare(v, target.GetAllowHeaders()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetExposeHeaders()) != len(target.GetExposeHeaders()) {
		return false
	}
	for idx, v := range m.GetExposeHeaders() {

		if strings.Compare(v, target.GetExposeHeaders()[idx]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetMaxAge(), target.GetMaxAge()) != 0 {
		return false
	}

	if m.GetAllowCredentials() != target.GetAllowCredentials() {
		return false
	}

	if m.GetDisableForRoute() != target.GetDisableForRoute() {
		return false
	}

	return true
}

// Equal function
func (m *CorsPolicyMergeSettings) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*CorsPolicyMergeSettings)
	if !ok {
		that2, ok := that.(CorsPolicyMergeSettings)
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

	if m.GetExposeHeaders() != target.GetExposeHeaders() {
		return false
	}

	return true
}
