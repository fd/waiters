package waiters

import (
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestWaiter(t *testing.T) {
	t.Run("pre waited", func(t *testing.T) {
		var (
			w = &Waiter{}
			g errgroup.Group
		)

		for i := 0; i < 10; i++ {
			g.Go(func() error {
				<-w.WaitC()
				return nil
			})
		}

		g.Go(func() error {
			time.Sleep(100 * time.Millisecond)
			w.Trigger()
			return nil
		})

		g.Wait()
	})

	t.Run("pre triggered", func(t *testing.T) {
		var (
			w = &Waiter{}
			g errgroup.Group
		)

		// this should not block
		w.Trigger()

		for i := 0; i < 10; i++ {
			g.Go(func() error {
				<-w.WaitC()
				return nil
			})
		}

		g.Go(func() error {
			time.Sleep(100 * time.Millisecond)
			w.Trigger()
			return nil
		})

		g.Wait()
	})
}
