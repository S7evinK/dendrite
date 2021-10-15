// Copyright 2021 The Matrix.org Foundation C.I.C.
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

package consumers

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/getsentry/sentry-go"
	"github.com/matrix-org/dendrite/eduserver/api"
	"github.com/matrix-org/dendrite/internal"
	"github.com/matrix-org/dendrite/setup/config"
	"github.com/matrix-org/dendrite/setup/process"
	"github.com/matrix-org/dendrite/syncapi/notifier"
	"github.com/matrix-org/dendrite/syncapi/storage"
	"github.com/matrix-org/dendrite/syncapi/types"
	log "github.com/sirupsen/logrus"
)

// OutputPresenceDataConsumer consumes events that originated in the EDU server.
type OutputPresenceDataConsumer struct {
	presenceConsumer *internal.ContinualConsumer
	db               storage.Database
	stream           types.StreamProvider
	notifier         *notifier.Notifier
}

// NewOutputPresenceDataConsumer creates a new OutputPresenceDataConsumer.
// Call Start() to begin consuming from the EDU server.
func NewOutputPresenceDataConsumer(
	process *process.ProcessContext,
	cfg *config.SyncAPI,
	kafkaConsumer sarama.Consumer,
	store storage.Database,
	notifier *notifier.Notifier,
	stream types.StreamProvider,
) *OutputPresenceDataConsumer {

	consumer := internal.ContinualConsumer{
		Process:        process,
		ComponentName:  "syncapi/eduserver/presence",
		Topic:          cfg.Matrix.Kafka.TopicFor(config.TopicOutputPresenceData),
		Consumer:       kafkaConsumer,
		PartitionStore: store,
	}

	s := &OutputPresenceDataConsumer{
		presenceConsumer: &consumer,
		db:               store,
		notifier:         notifier,
		stream:           stream,
	}

	consumer.ProcessMessage = s.onMessage

	return s
}

// Start consuming from EDU api
func (s *OutputPresenceDataConsumer) Start() error {
	return s.presenceConsumer.Start()
}

func (s *OutputPresenceDataConsumer) onMessage(msg *sarama.ConsumerMessage) error {
	var output api.OutputPresenceData
	if err := json.Unmarshal(msg.Value, &output); err != nil {
		// If the message was invalid, log it and move on to the next message in the stream
		log.WithError(err).Errorf("EDU server output log: message parse failure")
		sentry.CaptureException(err)
		return nil
	}
	log.Debugf("presence received by sync api! %+v", output)

	s.stream.Advance(output.StreamPos)
	s.notifier.OnNewPresence(types.StreamingToken{PresenceDataPosition: output.StreamPos}, output.UserID)

	return nil
}