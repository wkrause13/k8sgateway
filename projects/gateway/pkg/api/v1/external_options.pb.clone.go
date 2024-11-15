// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/external_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1"

	github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *VirtualHostOption) Clone() proto.Message {
	var target *VirtualHostOption
	if m == nil {
		return target
	}
	target = &VirtualHostOption{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.VirtualHostOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.VirtualHostOptions)
	}

	if m.GetTargetRefs() != nil {
		target.TargetRefs = make([]*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName, len(m.GetTargetRefs()))
		for idx, v := range m.GetTargetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TargetRefs[idx] = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			} else {
				target.TargetRefs[idx] = proto.Clone(v).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			}

		}
	}

	return target
}

// Clone function
func (m *RouteOption) Clone() proto.Message {
	var target *RouteOption
	if m == nil {
		return target
	}
	target = &RouteOption{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.RouteOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.RouteOptions)
	}

	if m.GetTargetRefs() != nil {
		target.TargetRefs = make([]*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReference, len(m.GetTargetRefs()))
		for idx, v := range m.GetTargetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TargetRefs[idx] = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReference)
			} else {
				target.TargetRefs[idx] = proto.Clone(v).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReference)
			}

		}
	}

	return target
}

// Clone function
func (m *ListenerOption) Clone() proto.Message {
	var target *ListenerOption
	if m == nil {
		return target
	}
	target = &ListenerOption{}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.ListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.ListenerOptions)
	}

	if m.GetTargetRefs() != nil {
		target.TargetRefs = make([]*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName, len(m.GetTargetRefs()))
		for idx, v := range m.GetTargetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TargetRefs[idx] = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			} else {
				target.TargetRefs[idx] = proto.Clone(v).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			}

		}
	}

	return target
}

// Clone function
func (m *HttpListenerOption) Clone() proto.Message {
	var target *HttpListenerOption
	if m == nil {
		return target
	}
	target = &HttpListenerOption{}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetOptions()).(clone.Cloner); ok {
		target.Options = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.HttpListenerOptions)
	} else {
		target.Options = proto.Clone(m.GetOptions()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1.HttpListenerOptions)
	}

	if m.GetTargetRefs() != nil {
		target.TargetRefs = make([]*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName, len(m.GetTargetRefs()))
		for idx, v := range m.GetTargetRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TargetRefs[idx] = h.Clone().(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			} else {
				target.TargetRefs[idx] = proto.Clone(v).(*github_com_solo_io_skv2_pkg_api_core_skv2_solo_io_v1.PolicyTargetReferenceWithSectionName)
			}

		}
	}

	return target
}
