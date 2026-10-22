/*
Copyright 2024 chideat.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

// MaxConcurrentReconciles is the number of concurrent reconcile workers used by
// the Cluster, Failover, Sentinel and User controllers. The default of 8 keeps
// upstream behaviour; cmd/main.go overrides it from --max-concurrent-reconciles
// (or MAX_CONCURRENT_RECONCILES) before the controllers are set up.
var MaxConcurrentReconciles = 8
