// Copyright 2020 The Matrix.org Foundation C.I.C.
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

package storage

import (
	"context"
	"time"

	"github.com/matrix-org/gomatrixserverlib"
	"github.com/matrix-org/gomatrixserverlib/spec"

	"github.com/matrix-org/dendrite/federationapi/storage/shared/receipt"
	"github.com/matrix-org/dendrite/federationapi/types"
	rstypes "github.com/matrix-org/dendrite/roomserver/types"
)

type Database interface {
	gomatrixserverlib.KeyDatabase

	UpdateRoom(ctx context.Context, roomID string, addHosts []types.JoinedHost, removeHosts []string, purgeRoomFirst bool) (joinedHosts []types.JoinedHost, err error)

	GetJoinedHosts(ctx context.Context, roomID string) ([]types.JoinedHost, error)
	GetAllJoinedHosts(ctx context.Context) ([]spec.ServerName, error)
	// GetJoinedHostsForRooms returns the complete set of servers in the rooms given.
	GetJoinedHostsForRooms(ctx context.Context, roomIDs []string, excludeSelf, excludeBlacklisted bool) ([]spec.ServerName, error)

	StoreJSON(ctx context.Context, js string) (*receipt.Receipt, error)

	GetPendingPDUs(ctx context.Context, serverName spec.ServerName, limit int) (pdus map[*receipt.Receipt]*rstypes.HeaderedEvent, err error)
	GetPendingEDUs(ctx context.Context, serverName spec.ServerName, limit int) (edus map[*receipt.Receipt]*gomatrixserverlib.EDU, err error)

	AssociatePDUWithDestinations(ctx context.Context, destinations map[spec.ServerName]struct{}, dbReceipt *receipt.Receipt) error
	AssociateEDUWithDestinations(ctx context.Context, destinations map[spec.ServerName]struct{}, dbReceipt *receipt.Receipt, eduType string, expireEDUTypes map[string]time.Duration) error

	CleanPDUs(ctx context.Context, serverName spec.ServerName, receipts []*receipt.Receipt) error
	CleanEDUs(ctx context.Context, serverName spec.ServerName, receipts []*receipt.Receipt) error

	GetPendingPDUServerNames(ctx context.Context) ([]spec.ServerName, error)
	GetPendingEDUServerNames(ctx context.Context) ([]spec.ServerName, error)

	// these don't have contexts passed in as we want things to happen regardless of the request context
	AddServerToBlacklist(serverName spec.ServerName) error
	RemoveServerFromBlacklist(serverName spec.ServerName) error
	RemoveAllServersFromBlacklist() error
	IsServerBlacklisted(serverName spec.ServerName) (bool, error)

	// Update the notary with the given server keys from the given server name.
	UpdateNotaryKeys(ctx context.Context, serverName spec.ServerName, serverKeys gomatrixserverlib.ServerKeys) error
	// Query the notary for the server keys for the given server. If `optKeyIDs` is not empty, multiple server keys may be returned (between 1 - len(optKeyIDs))
	// such that the combination of all server keys will include all the `optKeyIDs`.
	GetNotaryKeys(ctx context.Context, serverName spec.ServerName, optKeyIDs []gomatrixserverlib.KeyID) ([]gomatrixserverlib.ServerKeys, error)
	// DeleteExpiredEDUs cleans up expired EDUs
	DeleteExpiredEDUs(ctx context.Context) error

	PurgeRoom(ctx context.Context, roomID string) error
}
