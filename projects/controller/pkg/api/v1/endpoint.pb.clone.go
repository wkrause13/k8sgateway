// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/endpoint.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

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
func (m *Endpoint) Clone() proto.Message {
	var target *Endpoint
	if m == nil {
		return target
	}
	target = &Endpoint{}

	if m.GetUpstreams() != nil {
		target.Upstreams = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetUpstreams()))
		for idx, v := range m.GetUpstreams() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upstreams[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.Upstreams[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	target.Address = m.GetAddress()

	target.Port = m.GetPort()

	target.Hostname = m.GetHostname()

	if h, ok := interface{}(m.GetHealthCheck()).(clone.Cloner); ok {
		target.HealthCheck = h.Clone().(*HealthCheckConfig)
	} else {
		target.HealthCheck = proto.Clone(m.GetHealthCheck()).(*HealthCheckConfig)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	return target
}

// Clone function
func (m *HealthCheckConfig) Clone() proto.Message {
	var target *HealthCheckConfig
	if m == nil {
		return target
	}
	target = &HealthCheckConfig{}

	target.Hostname = m.GetHostname()

	return target
}
