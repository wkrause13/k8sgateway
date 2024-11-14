// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/extproc/extproc.proto

package extproc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *Settings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extproc.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc.Settings")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGrpcService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcService")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcService")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetFilterStage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FilterStage")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFilterStage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FilterStage")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetFailureModeAllow()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FailureModeAllow")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFailureModeAllow(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FailureModeAllow")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetProcessingMode()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ProcessingMode")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProcessingMode(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ProcessingMode")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAsyncMode()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AsyncMode")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAsyncMode(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AsyncMode")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetRequestAttributes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetResponseAttributes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetMessageTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MessageTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMessageTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MessageTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatPrefix()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatPrefix(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatPrefix")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMutationRules()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MutationRules")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMutationRules(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MutationRules")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMaxMessageTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MaxMessageTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMaxMessageTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MaxMessageTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisableClearRouteCache()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisableClearRouteCache")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisableClearRouteCache(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisableClearRouteCache")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetForwardRules()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ForwardRules")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetForwardRules(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ForwardRules")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetFilterMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("FilterMetadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetFilterMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("FilterMetadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAllowModeOverride()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowModeOverride")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowModeOverride(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowModeOverride")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetMetadataContextNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetTypedMetadataContextNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *RouteSettings) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extproc.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc.RouteSettings")); err != nil {
		return 0, err
	}

	switch m.Override.(type) {

	case *RouteSettings_Disabled:

		if h, ok := interface{}(m.GetDisabled()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Disabled")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDisabled(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Disabled")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *RouteSettings_Overrides:

		if h, ok := interface{}(m.GetOverrides()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Overrides")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetOverrides(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Overrides")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *GrpcService) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extproc.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc.GrpcService")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetExtProcServerRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ExtProcServerRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtProcServerRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ExtProcServerRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAuthority()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Authority")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAuthority(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Authority")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRetryPolicy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryPolicy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Timeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Timeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetInitialMetadata() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *Overrides) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extproc.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc.Overrides")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetProcessingMode()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ProcessingMode")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProcessingMode(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ProcessingMode")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAsyncMode()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AsyncMode")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAsyncMode(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AsyncMode")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetRequestAttributes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetResponseAttributes() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetGrpcService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcService")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcService")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetMetadataContextNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	for _, v := range m.GetTypedMetadataContextNamespaces() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *HeaderForwardingRules) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("extproc.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/extproc.HeaderForwardingRules")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetAllowedHeaders()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AllowedHeaders")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAllowedHeaders(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AllowedHeaders")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDisallowedHeaders()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DisallowedHeaders")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDisallowedHeaders(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DisallowedHeaders")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}
