package bd09mc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLL2MC(t *testing.T) {
	lng, lat, err := LL2MC(108.95344, 34.265657)
	assert.Nil(t, err)
	assert.Equal(t, lng, 12128773.43)
	assert.Equal(t, lat, 4040249.01)
}

func TestMC2LL(t *testing.T) {
	lng, lat, err := MC2LL(12128773.43, 4040249.00)
	assert.Nil(t, err)
	assert.Equal(t, lng, 108.95344)
	assert.Equal(t, lat, 34.265657)
}

func TestParallel(t *testing.T) {
	const parallels = 10
	waits := make(chan int, parallels)
	for i := 0; i < parallels; i++ {
		go func() {
			for k := 0; k < 1e4; k++ {
				TestLL2MC(t)
				TestMC2LL(t)
			}
			waits <- 1
		}()
	}
	for i := 0; i < parallels; i++ {
		<-waits
	}
}

func BenchmarkLL2MC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = LL2MC(108.95344, 34.265657)
	}
}

func BenchmarkMC2LL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = MC2LL(12128773.43, 4040249.00)
	}
}
