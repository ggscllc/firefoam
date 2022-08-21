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

func TestMultipleFireFoam(t *testing.T) {
	newRateLimit_1 := NewRateLimit(5, 500*time.Millisecond)
	newRateLimit_2 := NewRateLimit(2, 750*time.Millisecond)

	assert.Equal(t, true, newRateLimit_1.TakeItToTheLimit()) //1
	assert.Equal(t, true, newRateLimit_1.TakeItToTheLimit()) //2
	assert.Equal(t, true, newRateLimit_1.TakeItToTheLimit()) //3
	assert.Equal(t, true, newRateLimit_1.TakeItToTheLimit()) //4
	assert.Equal(t, true, newRateLimit_1.TakeItToTheLimit()) //5

	assert.Equal(t, true, newRateLimit_2.TakeItToTheLimit()) //1
	assert.Equal(t, true, newRateLimit_2.TakeItToTheLimit()) //2

	assert.Equal(t, false, newRateLimit_1.TakeItToTheLimit()) //6
	assert.Equal(t, false, newRateLimit_2.TakeItToTheLimit()) //3

	time.Sleep(6 * time.Second)

	assert.Equal(t, 0, newRateLimit_1.GetCurrentProcs())
	assert.Equal(t, 0, newRateLimit_2.GetCurrentProcs())

}
