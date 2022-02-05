package gomime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders(t *testing.T) {
	assert.EqualValues(t, "Content-Type", HeaderContentType)
	assert.EqualValues(t, "User-Agent", HeaderUserAgent)
	assert.EqualValues(t, "application/json", ContentTypeJson)
	assert.EqualValues(t, "application/xml", ContentTypeXml)
}
