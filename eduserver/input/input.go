// Copyright 2017 Vector Creations Ltd
// Copyright 2017-2018 New Vector Ltd
// Copyright 2019-2020 The Matrix.org Foundation C.I.C.
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

package input

import (
	"github.com/Shopify/sarama"
	"github.com/matrix-org/dendrite/eduserver/cache"
	userapi "github.com/matrix-org/dendrite/userapi/api"
	"github.com/matrix-org/gomatrixserverlib"
)

// EDUServerInputAPI implements api.EDUServerInputAPI
type EDUServerInputAPI struct {
	// Cache to store the current typing members in each room.
	Cache *cache.EDUCache
	// The kafka topic to output new typing events to.
	OutputTypingEventTopic string
	// The kafka topic to output new send to device events to.
	OutputSendToDeviceEventTopic string
	// The kafka topic to output new receipt events to
	OutputReceiptEventTopic string
	// kafka producer
	Producer sarama.SyncProducer
	// Internal user query API
	UserAPI userapi.UserInternalAPI
	// our server name
	ServerName gomatrixserverlib.ServerName
}
