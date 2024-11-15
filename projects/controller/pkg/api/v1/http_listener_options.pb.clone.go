// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/http_listener_options.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_filter_http_gzip_v2 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/filter/http/gzip/v2"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_buffer_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/filters/http/buffer/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_csrf_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/filters/http/csrf/v3"

	github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_proxylatency "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/proxylatency"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_caching "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/caching"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_dlp "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/dlp"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extauth_v1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extauth/v1"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extproc "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_ratelimit "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/ratelimit"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_stateful_session "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/stateful_session"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_waf "github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/waf"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_connection_limit "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/connection_limit"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_dynamic_forward_proxy "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/dynamic_forward_proxy"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_json "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc_json"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_web "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc_web"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_hcm "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/hcm"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_header_validation "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/header_validation"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_healthcheck "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/healthcheck"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_local_ratelimit "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/local_ratelimit"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_router "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/router"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_tap "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/tap"

	github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_wasm "github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/wasm"

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
func (m *HttpListenerOptions) Clone() proto.Message {
	var target *HttpListenerOptions
	if m == nil {
		return target
	}
	target = &HttpListenerOptions{}

	if h, ok := interface{}(m.GetGrpcWeb()).(clone.Cloner); ok {
		target.GrpcWeb = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_web.GrpcWeb)
	} else {
		target.GrpcWeb = proto.Clone(m.GetGrpcWeb()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_web.GrpcWeb)
	}

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(clone.Cloner); ok {
		target.HttpConnectionManagerSettings = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	} else {
		target.HttpConnectionManagerSettings = proto.Clone(m.GetHttpConnectionManagerSettings()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_hcm.HttpConnectionManagerSettings)
	}

	if h, ok := interface{}(m.GetHealthCheck()).(clone.Cloner); ok {
		target.HealthCheck = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_healthcheck.HealthCheck)
	} else {
		target.HealthCheck = proto.Clone(m.GetHealthCheck()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_healthcheck.HealthCheck)
	}

	if h, ok := interface{}(m.GetExtensions()).(clone.Cloner); ok {
		target.Extensions = h.Clone().(*Extensions)
	} else {
		target.Extensions = proto.Clone(m.GetExtensions()).(*Extensions)
	}

	if h, ok := interface{}(m.GetWaf()).(clone.Cloner); ok {
		target.Waf = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_waf.Settings)
	} else {
		target.Waf = proto.Clone(m.GetWaf()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_waf.Settings)
	}

	if h, ok := interface{}(m.GetDlp()).(clone.Cloner); ok {
		target.Dlp = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_dlp.FilterConfig)
	} else {
		target.Dlp = proto.Clone(m.GetDlp()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_dlp.FilterConfig)
	}

	if h, ok := interface{}(m.GetWasm()).(clone.Cloner); ok {
		target.Wasm = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_wasm.PluginSource)
	} else {
		target.Wasm = proto.Clone(m.GetWasm()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_wasm.PluginSource)
	}

	if h, ok := interface{}(m.GetExtauth()).(clone.Cloner); ok {
		target.Extauth = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	} else {
		target.Extauth = proto.Clone(m.GetExtauth()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extauth_v1.Settings)
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(clone.Cloner); ok {
		target.RatelimitServer = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_ratelimit.Settings)
	} else {
		target.RatelimitServer = proto.Clone(m.GetRatelimitServer()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_ratelimit.Settings)
	}

	if h, ok := interface{}(m.GetCaching()).(clone.Cloner); ok {
		target.Caching = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_caching.Settings)
	} else {
		target.Caching = proto.Clone(m.GetCaching()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_caching.Settings)
	}

	if h, ok := interface{}(m.GetGzip()).(clone.Cloner); ok {
		target.Gzip = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_filter_http_gzip_v2.Gzip)
	} else {
		target.Gzip = proto.Clone(m.GetGzip()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_config_filter_http_gzip_v2.Gzip)
	}

	if h, ok := interface{}(m.GetProxyLatency()).(clone.Cloner); ok {
		target.ProxyLatency = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_proxylatency.ProxyLatency)
	} else {
		target.ProxyLatency = proto.Clone(m.GetProxyLatency()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_proxylatency.ProxyLatency)
	}

	if h, ok := interface{}(m.GetBuffer()).(clone.Cloner); ok {
		target.Buffer = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_buffer_v3.Buffer)
	} else {
		target.Buffer = proto.Clone(m.GetBuffer()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_buffer_v3.Buffer)
	}

	if h, ok := interface{}(m.GetCsrf()).(clone.Cloner); ok {
		target.Csrf = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	} else {
		target.Csrf = proto.Clone(m.GetCsrf()).(*github_com_solo_io_gloo_projects_controller_pkg_api_external_envoy_extensions_filters_http_csrf_v3.CsrfPolicy)
	}

	if h, ok := interface{}(m.GetGrpcJsonTranscoder()).(clone.Cloner); ok {
		target.GrpcJsonTranscoder = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_json.GrpcJsonTranscoder)
	} else {
		target.GrpcJsonTranscoder = proto.Clone(m.GetGrpcJsonTranscoder()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_grpc_json.GrpcJsonTranscoder)
	}

	if h, ok := interface{}(m.GetSanitizeClusterHeader()).(clone.Cloner); ok {
		target.SanitizeClusterHeader = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.SanitizeClusterHeader = proto.Clone(m.GetSanitizeClusterHeader()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetLeftmostXffAddress()).(clone.Cloner); ok {
		target.LeftmostXffAddress = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.LeftmostXffAddress = proto.Clone(m.GetLeftmostXffAddress()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetDynamicForwardProxy()).(clone.Cloner); ok {
		target.DynamicForwardProxy = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_dynamic_forward_proxy.FilterConfig)
	} else {
		target.DynamicForwardProxy = proto.Clone(m.GetDynamicForwardProxy()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_dynamic_forward_proxy.FilterConfig)
	}

	if h, ok := interface{}(m.GetConnectionLimit()).(clone.Cloner); ok {
		target.ConnectionLimit = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_connection_limit.ConnectionLimit)
	} else {
		target.ConnectionLimit = proto.Clone(m.GetConnectionLimit()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_connection_limit.ConnectionLimit)
	}

	if h, ok := interface{}(m.GetNetworkLocalRatelimit()).(clone.Cloner); ok {
		target.NetworkLocalRatelimit = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_local_ratelimit.TokenBucket)
	} else {
		target.NetworkLocalRatelimit = proto.Clone(m.GetNetworkLocalRatelimit()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_local_ratelimit.TokenBucket)
	}

	if h, ok := interface{}(m.GetHttpLocalRatelimit()).(clone.Cloner); ok {
		target.HttpLocalRatelimit = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_local_ratelimit.Settings)
	} else {
		target.HttpLocalRatelimit = proto.Clone(m.GetHttpLocalRatelimit()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_local_ratelimit.Settings)
	}

	if h, ok := interface{}(m.GetRouter()).(clone.Cloner); ok {
		target.Router = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_router.Router)
	} else {
		target.Router = proto.Clone(m.GetRouter()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_router.Router)
	}

	if h, ok := interface{}(m.GetTap()).(clone.Cloner); ok {
		target.Tap = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_tap.Tap)
	} else {
		target.Tap = proto.Clone(m.GetTap()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_tap.Tap)
	}

	if h, ok := interface{}(m.GetStatefulSession()).(clone.Cloner); ok {
		target.StatefulSession = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_stateful_session.StatefulSession)
	} else {
		target.StatefulSession = proto.Clone(m.GetStatefulSession()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_stateful_session.StatefulSession)
	}

	if h, ok := interface{}(m.GetHeaderValidationSettings()).(clone.Cloner); ok {
		target.HeaderValidationSettings = h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_header_validation.HeaderValidationSettings)
	} else {
		target.HeaderValidationSettings = proto.Clone(m.GetHeaderValidationSettings()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_options_header_validation.HeaderValidationSettings)
	}

	switch m.ExtProcConfig.(type) {

	case *HttpListenerOptions_DisableExtProc:

		if h, ok := interface{}(m.GetDisableExtProc()).(clone.Cloner); ok {
			target.ExtProcConfig = &HttpListenerOptions_DisableExtProc{
				DisableExtProc: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.ExtProcConfig = &HttpListenerOptions_DisableExtProc{
				DisableExtProc: proto.Clone(m.GetDisableExtProc()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	case *HttpListenerOptions_ExtProc:

		if h, ok := interface{}(m.GetExtProc()).(clone.Cloner); ok {
			target.ExtProcConfig = &HttpListenerOptions_ExtProc{
				ExtProc: h.Clone().(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extproc.Settings),
			}
		} else {
			target.ExtProcConfig = &HttpListenerOptions_ExtProc{
				ExtProc: proto.Clone(m.GetExtProc()).(*github_com_solo_io_gloo_projects_controller_pkg_api_v1_enterprise_options_extproc.Settings),
			}
		}

	}

	return target
}
