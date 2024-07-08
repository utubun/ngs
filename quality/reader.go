package quality

import (
	"bufio"
	"context"
	"io"
)

func Reader(ctx context.Context, r io.Reader) <-chan []string {
	out := make(chan []string)

	scanner := bufio.NewScanner(r)
	var batch []string

	go func() {
		defer close(out)

		for {
			scanned := scanner.Scan()

			select {
			case <-ctx.Done():
				return
			default:
				line := scanner.Text()

				if len(batch) == 3 {
					out <- batch
					batch = []string{}
				}
				batch = append(batch, line)
			}
			if !scanned {
				if len(batch) > 0 {
					out <- batch
				}
				return
			}
		}
	}()
	return out
}
