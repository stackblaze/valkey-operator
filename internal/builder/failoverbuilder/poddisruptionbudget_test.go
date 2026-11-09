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

package failoverbuilder

import (
	"testing"

	v1alpha1 "github.com/chideat/valkey-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewPodDisruptionBudgetForCR_MaxUnavailable(t *testing.T) {
	for _, tc := range []struct {
		replicas int32
		want     int
	}{
		// One pod: drains may evict it rather than hang.
		{replicas: 1, want: 1},
		// A master and replicas: all but one may go.
		{replicas: 2, want: 1},
		{replicas: 3, want: 2},
	} {
		rf := &v1alpha1.Failover{
			ObjectMeta: metav1.ObjectMeta{Name: "rf-test", Namespace: "default"},
			Spec:       v1alpha1.FailoverSpec{Replicas: tc.replicas},
		}
		pdb := NewPodDisruptionBudgetForCR(rf)
		assert.Equal(t, tc.want, pdb.Spec.MaxUnavailable.IntValue(), "replicas=%d", tc.replicas)
		assert.Nil(t, pdb.Spec.MinAvailable)
	}
}
