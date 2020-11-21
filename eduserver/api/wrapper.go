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

package api

import (
	"context"
	"encoding/json"

	"github.com/matrix-org/dendrite/eduserver/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/matrix-org/gomatrixserverlib"
)

// SendTyping sends a typing event to EDU server
func SendTyping(
	ctx context.Context, eduAPI EDUServerInputAPI, userID, roomID string,
	typing bool, timeoutMS int64,
) error {
	_, err := eduAPI.SendTypingEvent(ctx, &proto.TypingEvent{
		UserID:         userID,
		RoomID:         roomID,
		Typing:         typing,
		Timeout:        uint64(timeoutMS),
		OriginServerTS: timestamppb.Now(),
	})

	return err
}

// SendToDevice sends a typing event to EDU server
func SendToDevice(
	ctx context.Context, eduAPI EDUServerInputAPI, sender, userID, deviceID, eventType string,
	message interface{},
) error {
	js, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = eduAPI.SendToDevice(ctx, &proto.SendToDeviceEvent{
		UserID:   userID,
		DeviceID: deviceID,
		Sender:   sender,
		Type:     eventType,
		Content:  js,
	})
	return err
}

// SendReceipt sends a receipt event to EDU Server
func SendReceipt(
	ctx context.Context,
	eduAPI EDUServerInputAPI, userID, roomID, eventID, receiptType string,
	timestamp gomatrixserverlib.Timestamp,
) error {
	_, err := eduAPI.SendReceiptEvent(ctx, &proto.ReceiptEvent{
		UserID:    userID,
		RoomID:    roomID,
		EventID:   eventID,
		Type:      receiptType,
		Timestamp: timestamppb.New(timestamp.Time()),
	})

	return err
}
