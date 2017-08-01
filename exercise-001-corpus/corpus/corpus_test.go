package corpus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorpus(t *testing.T) {
	result := WordCount("../7oldsamr.txt")
	assert.Equal(t, len(result), 243)
	assert.Equal(t, result[0].Word, "the")
	assert.Equal(t, result[0].Count, 36)
	assert.Equal(t, result[215].Word, "trouble")
	assert.Equal(t, result[215].Count, 1)
}
