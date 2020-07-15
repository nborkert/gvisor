// Copyright 2018 The containerd Authors.
// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package shim

import (
	"github.com/containerd/containerd/api/events"
)

type TaskCreate = events.TaskCreate
type TaskStart = events.TaskStart
type TaskOOM = events.TaskOOM
type TaskExit = events.TaskExit
type TaskDelete = events.TaskDelete
type TaskExecAdded = events.TaskExecAdded
type TaskExecStarted = events.TaskExecStarted