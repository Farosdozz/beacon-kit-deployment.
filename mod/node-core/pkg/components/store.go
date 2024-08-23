// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
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

package components

import (
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb"
	"github.com/berachain/beacon-kit/mod/storage/pkg/encoding"
)

// KVStoreInput is the input for the ProvideKVStore function.
type KVStoreInput struct {
	depinject.In
	KVStoreService store.KVStoreService
}

// ProvideKVStore is the depinject provider that returns a beacon KV store.
func ProvideKVStore[
	BeaconBlockHeaderT BeaconBlockHeader[BeaconBlockHeaderT],
](in KVStoreInput) *beacondb.KVStore[
	BeaconBlockHeaderT,
	*Eth1Data,
	*ExecutionPayloadHeader,
	*Fork,
	*Validator,
	Validators,
] {
	payloadCodec := &encoding.SSZInterfaceCodec[*ExecutionPayloadHeader]{}
	return beacondb.New[
		BeaconBlockHeaderT,
		*Eth1Data,
		*ExecutionPayloadHeader,
		*Fork,
		*Validator,
		Validators,
	](in.KVStoreService, payloadCodec)
}