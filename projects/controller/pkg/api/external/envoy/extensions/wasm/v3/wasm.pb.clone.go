// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/wasm/v3/wasm.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	google_golang_org_protobuf_types_known_anypb "google.golang.org/protobuf/types/known/anypb"
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
func (m *VmConfig) Clone() proto.Message {
	var target *VmConfig
	if m == nil {
		return target
	}
	target = &VmConfig{}

	target.VmId = m.GetVmId()

	target.Runtime = m.GetRuntime()

	if h, ok := interface{}(m.GetCode()).(clone.Cloner); ok {
		target.Code = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.AsyncDataSource)
	} else {
		target.Code = proto.Clone(m.GetCode()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.AsyncDataSource)
	}

	if h, ok := interface{}(m.GetConfiguration()).(clone.Cloner); ok {
		target.Configuration = h.Clone().(*google_golang_org_protobuf_types_known_anypb.Any)
	} else {
		target.Configuration = proto.Clone(m.GetConfiguration()).(*google_golang_org_protobuf_types_known_anypb.Any)
	}

	target.AllowPrecompiled = m.GetAllowPrecompiled()

	target.NackOnCodeCacheMiss = m.GetNackOnCodeCacheMiss()

	return target
}

// Clone function
func (m *PluginConfig) Clone() proto.Message {
	var target *PluginConfig
	if m == nil {
		return target
	}
	target = &PluginConfig{}

	target.Name = m.GetName()

	target.RootId = m.GetRootId()

	if h, ok := interface{}(m.GetConfiguration()).(clone.Cloner); ok {
		target.Configuration = h.Clone().(*google_golang_org_protobuf_types_known_anypb.Any)
	} else {
		target.Configuration = proto.Clone(m.GetConfiguration()).(*google_golang_org_protobuf_types_known_anypb.Any)
	}

	target.FailOpen = m.GetFailOpen()

	switch m.Vm.(type) {

	case *PluginConfig_VmConfig:

		if h, ok := interface{}(m.GetVmConfig()).(clone.Cloner); ok {
			target.Vm = &PluginConfig_VmConfig{
				VmConfig: h.Clone().(*VmConfig),
			}
		} else {
			target.Vm = &PluginConfig_VmConfig{
				VmConfig: proto.Clone(m.GetVmConfig()).(*VmConfig),
			}
		}

	}

	return target
}

// Clone function
func (m *WasmService) Clone() proto.Message {
	var target *WasmService
	if m == nil {
		return target
	}
	target = &WasmService{}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*PluginConfig)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*PluginConfig)
	}

	target.Singleton = m.GetSingleton()

	return target
}
