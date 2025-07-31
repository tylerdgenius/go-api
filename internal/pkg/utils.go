package pkg

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HandleSignals(ctx context.Context, ctxCancel context.CancelFunc, callback func()) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	go func() {
		<-quit

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Panic(
					"Shutdown timeout exceeded, forcing exit",
				)
			}
		}()

		callback()
		cancel()
		ctxCancel()

	}()
}
