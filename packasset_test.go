package packasset_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelfAsset(t *testing.T) {
	// Do it twice, just to fetch the cached version
	for i := 0; i < 2; i++ {
		t.Run(fmt.Sprintf("run%d", i+1), func(t *testing.T) {
			buf, err := Asset("packasset_test.go")
			if !assert.NoError(t, err, "Asset should succeed") {
				return
			}

			golden, err := ioutil.ReadFile("packasset_test.go")
			if !assert.NoError(t, err, "ioutil.ReadFile should succeed") {
				return
			}

			if !assert.Equal(t, golden, buf, "contents should match") {
				return
			}
		})
	}
}
