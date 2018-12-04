// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package blockinterface defines rootchain block interface.
package blockinterface

import "time"

// Hash - Rootchain hash interface.
type Hash interface {
	// Bytes - Raw bytes of custom chain hash.
	Bytes() []byte
}

// Operation - Rootchain operation interface.
type Operation interface {
	// Hash - Operation hash identifier.
	Hash() Hash

	// Bytes - Contents of operation script.
	Bytes() []byte
}

// Block - Rootchain block interface.
type Block interface {
	// Index - Block height.
	Index() uint64

	// Hash - Operation hash identifier.
	Hash() Hash

	// Prev - Hash of previous block in chain.
	Prev() Hash

	// Bytes - Contents of operation script.
	Bytes() []byte

	// Operations - List of operations in block.
	// Known as transactions list in currencies.
	Operations() []Operation

	// Timestamp - Block creation timestamp.
	Timestamp() time.Time

	// Metadata - Custom operation metadata.
	Metadata() map[string]interface{}
}
