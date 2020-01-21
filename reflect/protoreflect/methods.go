// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protoreflect

import (
	"google.golang.org/protobuf/internal/pragma"
)

// The following types are used by the fast-path Message.ProtoMethods method.
//
// To avoid polluting the public protoreflect API with types used only by
// low-level implementations, the canonical definitions of these types are
// in the runtime/protoiface package. The definitions here and in protoiface
// must be kept in sync.
type (
	methods = struct {
		pragma.NoUnkeyedLiterals
		Flags         supportFlags
		Size          func(m Message, opts marshalOptions) int
		Marshal       func(m Message, in marshalInput) (marshalOutput, error)
		Unmarshal     func(m Message, in unmarshalInput) (unmarshalOutput, error)
		IsInitialized func(m Message) error
	}
	supportFlags = uint64
	marshalInput = struct {
		pragma.NoUnkeyedLiterals
		Buf     []byte
		Options marshalOptions
	}
	marshalOutput = struct {
		pragma.NoUnkeyedLiterals
		Buf []byte
	}
	marshalOptions = struct {
		pragma.NoUnkeyedLiterals
		AllowPartial  bool
		Deterministic bool
		UseCachedSize bool
	}
	unmarshalInput = struct {
		pragma.NoUnkeyedLiterals
		Buf     []byte
		Options unmarshalOptions
	}
	unmarshalOutput = struct {
		pragma.NoUnkeyedLiterals
	}
	unmarshalOptions = struct {
		pragma.NoUnkeyedLiterals
		Merge          bool
		AllowPartial   bool
		DiscardUnknown bool
		Resolver       interface {
			FindExtensionByName(field FullName) (ExtensionType, error)
			FindExtensionByNumber(message FullName, field FieldNumber) (ExtensionType, error)
		}
	}
)