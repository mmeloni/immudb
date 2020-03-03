/*
Copyright 2019-2020 vChain, Inc.

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

package store

import (
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreReference(t *testing.T) {
	st, closer := makeStore()
	defer closer()

	_, err := st.Set(schema.KeyValue{Key: []byte(`firstKey`), Value: []byte(`firstValue`)})
	//st.tree.flush()
	assert.NoError(t, err)

	refOpts := &schema.ReferenceOptions{
		Reference: &schema.Key{Key: []byte(`myTag`)},
		Key:       &schema.Key{Key: []byte(`firstKey`)},
	}

	reference, err := st.Reference(refOpts)
	//st.tree.flush()
	assert.NoError(t, err)
	assert.NotEmptyf(t, reference, "Should not be empty")

	firstItemRet, err := st.Get(schema.Key{Key: []byte(`myTag`)})

	assert.NoError(t, err)
	assert.NotEmptyf(t, firstItemRet, "Should not be empty")
	assert.Equal(t, firstItemRet.Value, []byte(`firstValue`), "Should have referenced item value")
}
