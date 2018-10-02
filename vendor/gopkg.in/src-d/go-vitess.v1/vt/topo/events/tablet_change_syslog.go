/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreedto in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import (
	"fmt"
	"log/syslog"

	"gopkg.in/src-d/go-vitess.v1/event/syslogger"
	"gopkg.in/src-d/go-vitess.v1/vt/topo/topoproto"
)

// Syslog writes the event to syslog.
func (tc *TabletChange) Syslog() (syslog.Priority, string) {
	return syslog.LOG_INFO, fmt.Sprintf("%s/%s/%s [tablet] %s",
		tc.Tablet.Keyspace, tc.Tablet.Shard, topoproto.TabletAliasString(tc.Tablet.Alias), tc.Status)
}

var _ syslogger.Syslogger = (*TabletChange)(nil) // compile-time interface check