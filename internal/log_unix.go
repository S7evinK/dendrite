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

package internal

import (
	"io"

	"github.com/MFAshby/stdemuxerhook"
	"github.com/matrix-org/dendrite/setup/config"
	"github.com/sirupsen/logrus"
)

// SetupHookLogging configures the logging hooks defined in the configuration.
// If something fails here it means that the logging was improperly configured,
// so we just exit with the error
func SetupHookLogging(hooks []config.LogrusHook) {
	levelLogAddedMu.Lock()
	defer levelLogAddedMu.Unlock()
	for _, hook := range hooks {
		// Check we received a proper logging level
		level, err := logrus.ParseLevel(hook.Level)
		if err != nil {
			logrus.Fatalf("Unrecognised logging level %s: %q", hook.Level, err)
		}

		// Perform a first filter on the logs according to the lowest level of all
		// (Eg: If we have hook for info and above, prevent logrus from processing debug logs)
		if logrus.GetLevel() < level {
			logrus.SetLevel(level)
		}

		switch hook.Type {
		case "std":
			setupStdLogHook(level)
		default:
			logrus.Fatalf("Unrecognised logging hook type: %s", hook.Type)
		}
	}
	setupStdLogHook(logrus.InfoLevel)
	// Hooks are now configured for stdout/err, so throw away the default logger output
	logrus.SetOutput(io.Discard)
}

func setupStdLogHook(level logrus.Level) {
	if stdLevelLogAdded[level] {
		return
	}
	logrus.AddHook(&logLevelHook{level, stdemuxerhook.New(logrus.StandardLogger())})
	stdLevelLogAdded[level] = true
}
