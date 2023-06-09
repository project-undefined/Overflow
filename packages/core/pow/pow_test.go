package pow

import (
	"context"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	_ "golang.org/x/crypto/blake2b" // required by crypto.BLAKE2b_512
)

const (
	workers = 2
	target  = 10
)

var testWorker = New(workers)

func TestWorker_Work(t *testing.T) {
	nonce, err := testWorker.Mine(context.Background(), nil, target)
	require.NoError(t, err)
	difficulty, err := testWorker.LeadingZerosWithNonce(nil, nonce)
	require.GreaterOrEqual(t, difficulty, target)
	require.NoError(t, err)
}

func TestWorker_Validate(t *testing.T) {
	tests := []*struct {
		blk             []byte
		nonce           uint64
		expLeadingZeros int
		expErr          error
	}{
		{blk: nil, nonce: 0, expLeadingZeros: 1, expErr: nil},
		{blk: nil, nonce: 4611686018451317632, expLeadingZeros: 28, expErr: nil},
		{blk: make([]byte, 10240), nonce: 0, expLeadingZeros: 1, expErr: nil},
	}

	w := &Worker{}
	for _, tt := range tests {
		zeros, err := w.LeadingZerosWithNonce(tt.blk, tt.nonce)
		require.Equal(t, tt.expLeadingZeros, zeros)
		require.Equal(t, tt.expErr, err)
	}
}

func TestWorker_Cancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	go func() {
		defer wg.Done()
		_, err = testWorker.Mine(ctx, nil, math.MaxInt32)
	}()
	time.Sleep(10 * time.Millisecond)
	cancel()

	wg.Wait()

	require.ErrorIs(t, err, ErrCancelled)
}

func BenchmarkWorker(b *testing.B) {
	var (
		buf     = make([]byte, 1024)
		done    uint32
		counter uint64
	)
	go func() {
		_, _ = testWorker.worker(buf, 0, math.MaxInt32, &done, &counter)
	}()
	b.ResetTimer()
	for atomic.LoadUint64(&counter) < uint64(b.N) {
	}
	atomic.StoreUint32(&done, 1)
}
