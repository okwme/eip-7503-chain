// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package cachekv

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/polaris/eth/core/vm"
)

var _ storetypes.KVStore = (*ReadOnlyStore)(nil)

// ReadOnlyStore is a wrapper around cachekv.Store that panics on any write operation.
type ReadOnlyStore struct {
	storetypes.KVStore
}

func NewReadOnlyStoreFor(cacheKVStore storetypes.KVStore) *ReadOnlyStore {
	return &ReadOnlyStore{cacheKVStore}
}

func (s *ReadOnlyStore) Set(_, _ []byte) {
	panic(vm.ErrWriteProtection)
}

func (s *ReadOnlyStore) Delete(_ []byte) {
	panic(vm.ErrWriteProtection)
}
