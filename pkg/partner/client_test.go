package partner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	apiVersion := "version"
	client, err := New(apiVersion)
	assert.NoError(t, err)
	assert.Equal(t, apiVersion, client.APIVersion)
}

func TestClient_GetOffer(t *testing.T) {

}
