// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/rest/rest.proto

package rest

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

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
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ServiceSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rest.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/rest.ServiceSpec")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetTransformations() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetSwaggerInfo()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SwaggerInfo")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSwaggerInfo(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SwaggerInfo")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *DestinationSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rest.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/rest.DestinationSpec")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("FunctionName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetFunctionName())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetParameters()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Parameters")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetParameters(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Parameters")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponseTransformation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ResponseTransformation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ServiceSpec_SwaggerInfo) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("rest.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/rest.ServiceSpec_SwaggerInfo")); err != nil {
		return 0, err
	}

	switch m.SwaggerSpec.(type) {

	case *ServiceSpec_SwaggerInfo_Url:

		if _, err = hasher.Write([]byte("Url")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetUrl())); err != nil {
			return 0, err
		}

	case *ServiceSpec_SwaggerInfo_Inline:

		if _, err = hasher.Write([]byte("Inline")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetInline())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}
