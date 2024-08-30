package takefirstn

import "context"

func TakeFirstN(ctx context.Context, chIn <-chan interface{}, n int) <-chan interface{} {
	chResult := make(chan interface{})

	go func() {
		defer close(chResult)
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-chIn:
				if !ok {
					return
				}
				chResult <- val

			case <-ctx.Done():
				return
			}
		}
	}()

	return chResult
}
