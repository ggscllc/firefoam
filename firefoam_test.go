package firefoam

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFireFoam(t *testing.T) {
	newRateLimit := NewRateLimit(5, 500*time.Millisecond)
	assert.Equal(t, true, newRateLimit.TakeItToTheLimit())  //1
	assert.Equal(t, true, newRateLimit.TakeItToTheLimit())  //2
	assert.Equal(t, true, newRateLimit.TakeItToTheLimit())  //3
	assert.Equal(t, true, newRateLimit.TakeItToTheLimit())  //4
	assert.Equal(t, true, newRateLimit.TakeItToTheLimit())  //5
	assert.Equal(t, false, newRateLimit.TakeItToTheLimit()) // 6
	time.Sleep(6 * time.Second)
	assert.Equal(t, 0, newRateLimit.GetCurrentProcs())
}
