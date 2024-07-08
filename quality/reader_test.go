package quality

import (
	"context"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("test reader", func(t *testing.T) {
		f, err := os.Open("util.go")
		if err != nil {
			t.Errorf("Cann not open the file: %s", err)
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		out := Reader(ctx, f)
		m := <-out
		t.Logf("%+v", m)

	})
}
