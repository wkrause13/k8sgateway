// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/aws/filter.proto

package aws

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *AWSLambdaPerRoute) Clone() proto.Message {
	var target *AWSLambdaPerRoute
	if m == nil {
		return target
	}
	target = &AWSLambdaPerRoute{}

	target.Name = m.GetName()

	target.Qualifier = m.GetQualifier()

	target.Async = m.GetAsync()

	if h, ok := interface{}(m.GetEmptyBodyOverride()).(clone.Cloner); ok {
		target.EmptyBodyOverride = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.EmptyBodyOverride = proto.Clone(m.GetEmptyBodyOverride()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	target.UnwrapAsAlb = m.GetUnwrapAsAlb()

	if h, ok := interface{}(m.GetTransformerConfig()).(clone.Cloner); ok {
		target.TransformerConfig = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.TypedExtensionConfig)
	} else {
		target.TransformerConfig = proto.Clone(m.GetTransformerConfig()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.TypedExtensionConfig)
	}

	if h, ok := interface{}(m.GetRequestTransformerConfig()).(clone.Cloner); ok {
		target.RequestTransformerConfig = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.TypedExtensionConfig)
	} else {
		target.RequestTransformerConfig = proto.Clone(m.GetRequestTransformerConfig()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_core_v3.TypedExtensionConfig)
	}

	return target
}

// Clone function
func (m *AWSLambdaProtocolExtension) Clone() proto.Message {
	var target *AWSLambdaProtocolExtension
	if m == nil {
		return target
	}
	target = &AWSLambdaProtocolExtension{}

	target.Host = m.GetHost()

	target.Region = m.GetRegion()

	target.AccessKey = m.GetAccessKey()

	target.SecretKey = m.GetSecretKey()

	target.SessionToken = m.GetSessionToken()

	target.RoleArn = m.GetRoleArn()

	target.DisableRoleChaining = m.GetDisableRoleChaining()

	return target
}

// Clone function
func (m *AWSLambdaConfig) Clone() proto.Message {
	var target *AWSLambdaConfig
	if m == nil {
		return target
	}
	target = &AWSLambdaConfig{}

	target.PropagateOriginalRouting = m.GetPropagateOriginalRouting()

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(clone.Cloner); ok {
		target.CredentialRefreshDelay = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.CredentialRefreshDelay = proto.Clone(m.GetCredentialRefreshDelay()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	switch m.CredentialsFetcher.(type) {

	case *AWSLambdaConfig_UseDefaultCredentials:

		if h, ok := interface{}(m.GetUseDefaultCredentials()).(clone.Cloner); ok {
			target.CredentialsFetcher = &AWSLambdaConfig_UseDefaultCredentials{
				UseDefaultCredentials: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.CredentialsFetcher = &AWSLambdaConfig_UseDefaultCredentials{
				UseDefaultCredentials: proto.Clone(m.GetUseDefaultCredentials()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	case *AWSLambdaConfig_ServiceAccountCredentials_:

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(clone.Cloner); ok {
			target.CredentialsFetcher = &AWSLambdaConfig_ServiceAccountCredentials_{
				ServiceAccountCredentials: h.Clone().(*AWSLambdaConfig_ServiceAccountCredentials),
			}
		} else {
			target.CredentialsFetcher = &AWSLambdaConfig_ServiceAccountCredentials_{
				ServiceAccountCredentials: proto.Clone(m.GetServiceAccountCredentials()).(*AWSLambdaConfig_ServiceAccountCredentials),
			}
		}

	}

	return target
}

// Clone function
func (m *ApiGatewayTransformation) Clone() proto.Message {
	var target *ApiGatewayTransformation
	if m == nil {
		return target
	}
	target = &ApiGatewayTransformation{}

	return target
}

// Clone function
func (m *AWSLambdaConfig_ServiceAccountCredentials) Clone() proto.Message {
	var target *AWSLambdaConfig_ServiceAccountCredentials
	if m == nil {
		return target
	}
	target = &AWSLambdaConfig_ServiceAccountCredentials{}

	target.Cluster = m.GetCluster()

	target.Uri = m.GetUri()

	if h, ok := interface{}(m.GetTimeout()).(clone.Cloner); ok {
		target.Timeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Timeout = proto.Clone(m.GetTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	target.Region = m.GetRegion()

	return target
}
