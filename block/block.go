// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Package block defines rootchain block interface.
package block

import (
	"time"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

// Block - Rootchain block interface.
type Block interface {
	// ID - Block index.
	ID() uint64

	// CID - Operation content identifier.
	CID() cid.Cid

	// Prev - Content ID of previous block in chain.
	Prev() cid.Cid

	// State - Content identifier of state.
	State() cid.Cid

	// Exec - Content identifier of exec.
	Exec() cid.Cid

	// Time - Block creation timestamp.
	Time() time.Time
}
