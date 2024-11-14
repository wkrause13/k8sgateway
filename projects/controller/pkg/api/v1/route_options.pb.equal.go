// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/route_options.proto

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
func (m *RouteOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteOptions)
	if !ok {
		that2, ok := that.(RouteOptions)
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

	if h, ok := interface{}(m.GetTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTransformations(), target.GetTransformations()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetFaults()).(equality.Equalizer); ok {
		if !h.Equal(target.GetFaults()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetFaults(), target.GetFaults()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetPrefixRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPrefixRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPrefixRewrite(), target.GetPrefixRewrite()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRetries()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRetries()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRetries(), target.GetRetries()) {
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

	if h, ok := interface{}(m.GetTracing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTracing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTracing(), target.GetTracing()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetShadowing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetShadowing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetShadowing(), target.GetShadowing()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderManipulation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderManipulation(), target.GetHeaderManipulation()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAppendXForwardedHost()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAppendXForwardedHost()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAppendXForwardedHost(), target.GetAppendXForwardedHost()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCors()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCors()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCors(), target.GetCors()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetLbHash()).(equality.Equalizer); ok {
		if !h.Equal(target.GetLbHash()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetLbHash(), target.GetLbHash()) {
			return false
		}
	}

	if len(m.GetUpgrades()) != len(target.GetUpgrades()) {
		return false
	}
	for idx, v := range m.GetUpgrades() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpgrades()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUpgrades()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRatelimitBasic()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRatelimitBasic()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRatelimitBasic(), target.GetRatelimitBasic()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetWaf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetWaf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetWaf(), target.GetWaf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRbac()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRbac()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRbac(), target.GetRbac()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtauth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtauth(), target.GetExtauth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDlp()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDlp()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDlp(), target.GetDlp()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetBufferPerRoute()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBufferPerRoute()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBufferPerRoute(), target.GetBufferPerRoute()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCsrf()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCsrf(), target.GetCsrf()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStagedTransformations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStagedTransformations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStagedTransformations(), target.GetStagedTransformations()) {
			return false
		}
	}

	if len(m.GetEnvoyMetadata()) != len(target.GetEnvoyMetadata()) {
		return false
	}
	for k, v := range m.GetEnvoyMetadata() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEnvoyMetadata()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEnvoyMetadata()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetRegexRewrite()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegexRewrite()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegexRewrite(), target.GetRegexRewrite()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxStreamDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxStreamDuration(), target.GetMaxStreamDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetIdleTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetIdleTimeout(), target.GetIdleTimeout()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetExtProc()).(equality.Equalizer); ok {
		if !h.Equal(target.GetExtProc()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetExtProc(), target.GetExtProc()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAi()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAi()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAi(), target.GetAi()) {
			return false
		}
	}

	switch m.HostRewriteType.(type) {

	case *RouteOptions_HostRewrite:
		if _, ok := target.HostRewriteType.(*RouteOptions_HostRewrite); !ok {
			return false
		}

		if strings.Compare(m.GetHostRewrite(), target.GetHostRewrite()) != 0 {
			return false
		}

	case *RouteOptions_AutoHostRewrite:
		if _, ok := target.HostRewriteType.(*RouteOptions_AutoHostRewrite); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAutoHostRewrite()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAutoHostRewrite()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAutoHostRewrite(), target.GetAutoHostRewrite()) {
				return false
			}
		}

	case *RouteOptions_HostRewritePathRegex:
		if _, ok := target.HostRewriteType.(*RouteOptions_HostRewritePathRegex); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHostRewritePathRegex()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHostRewritePathRegex()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHostRewritePathRegex(), target.GetHostRewritePathRegex()) {
				return false
			}
		}

	case *RouteOptions_HostRewriteHeader:
		if _, ok := target.HostRewriteType.(*RouteOptions_HostRewriteHeader); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHostRewriteHeader()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHostRewriteHeader()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHostRewriteHeader(), target.GetHostRewriteHeader()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.HostRewriteType != target.HostRewriteType {
			return false
		}
	}

	switch m.RateLimitEarlyConfigType.(type) {

	case *RouteOptions_RatelimitEarly:
		if _, ok := target.RateLimitEarlyConfigType.(*RouteOptions_RatelimitEarly); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRatelimitEarly()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRatelimitEarly()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRatelimitEarly(), target.GetRatelimitEarly()) {
				return false
			}
		}

	case *RouteOptions_RateLimitEarlyConfigs:
		if _, ok := target.RateLimitEarlyConfigType.(*RouteOptions_RateLimitEarlyConfigs); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRateLimitEarlyConfigs()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimitEarlyConfigs()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRateLimitEarlyConfigs(), target.GetRateLimitEarlyConfigs()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RateLimitEarlyConfigType != target.RateLimitEarlyConfigType {
			return false
		}
	}

	switch m.RateLimitConfigType.(type) {

	case *RouteOptions_Ratelimit:
		if _, ok := target.RateLimitConfigType.(*RouteOptions_Ratelimit); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRatelimit()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRatelimit()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRatelimit(), target.GetRatelimit()) {
				return false
			}
		}

	case *RouteOptions_RateLimitConfigs:
		if _, ok := target.RateLimitConfigType.(*RouteOptions_RateLimitConfigs); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRateLimitConfigs()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimitConfigs()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRateLimitConfigs(), target.GetRateLimitConfigs()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RateLimitConfigType != target.RateLimitConfigType {
			return false
		}
	}

	switch m.RateLimitRegularConfigType.(type) {

	case *RouteOptions_RatelimitRegular:
		if _, ok := target.RateLimitRegularConfigType.(*RouteOptions_RatelimitRegular); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRatelimitRegular()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRatelimitRegular()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRatelimitRegular(), target.GetRatelimitRegular()) {
				return false
			}
		}

	case *RouteOptions_RateLimitRegularConfigs:
		if _, ok := target.RateLimitRegularConfigType.(*RouteOptions_RateLimitRegularConfigs); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRateLimitRegularConfigs()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRateLimitRegularConfigs()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRateLimitRegularConfigs(), target.GetRateLimitRegularConfigs()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.RateLimitRegularConfigType != target.RateLimitRegularConfigType {
			return false
		}
	}

	switch m.JwtConfig.(type) {

	case *RouteOptions_Jwt:
		if _, ok := target.JwtConfig.(*RouteOptions_Jwt); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwt()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwt()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwt(), target.GetJwt()) {
				return false
			}
		}

	case *RouteOptions_JwtStaged:
		if _, ok := target.JwtConfig.(*RouteOptions_JwtStaged); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwtStaged()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwtStaged()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwtStaged(), target.GetJwtStaged()) {
				return false
			}
		}

	case *RouteOptions_JwtProvidersStaged:
		if _, ok := target.JwtConfig.(*RouteOptions_JwtProvidersStaged); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwtProvidersStaged()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwtProvidersStaged()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwtProvidersStaged(), target.GetJwtProvidersStaged()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JwtConfig != target.JwtConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteOptions_MaxStreamDuration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteOptions_MaxStreamDuration)
	if !ok {
		that2, ok := that.(RouteOptions_MaxStreamDuration)
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

	if h, ok := interface{}(m.GetMaxStreamDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxStreamDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxStreamDuration(), target.GetMaxStreamDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGrpcTimeoutHeaderMax()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcTimeoutHeaderMax()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcTimeoutHeaderMax(), target.GetGrpcTimeoutHeaderMax()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetGrpcTimeoutHeaderOffset()).(equality.Equalizer); ok {
		if !h.Equal(target.GetGrpcTimeoutHeaderOffset()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetGrpcTimeoutHeaderOffset(), target.GetGrpcTimeoutHeaderOffset()) {
			return false
		}
	}

	return true
}
