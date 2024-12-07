package ginkit

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Serve(ctx context.Context, server *http.Server, opt ServeOptions) error {
	if server == nil {
		return fmt.Errorf("server (*http.Server) is nil")
	}

	opt.InitNilHandler()

	go func(server *http.Server) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", opt.GetPort()))
		if err != nil {
			opt.OnListenFailed(err)
			return
		}

		opt.BeforeServe()

		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			opt.OnServeFailed(err)
		}
	}(server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(ctx, opt.GetShutdownWaitTime())
	defer cancel()

	opt.BeforeShutdown()

	if err := server.Shutdown(ctx); err != nil {
		opt.OnShutdownFailed(err)
	}

	opt.AfterShutdown()

	return nil
}

type ServeOptions struct {
	Port             int
	ShutdownWaitTime time.Duration
	OnListenFailed   func(err error)
	BeforeServe      func()
	OnServeFailed    func(err error)
	BeforeShutdown   func()
	OnShutdownFailed func(err error)
	AfterShutdown    func()
}

func (s *ServeOptions) GetPort() int {
	if s.Port == 0 {
		return 8080
	}
	return s.Port
}

func (s *ServeOptions) GetShutdownWaitTime() time.Duration {
	if s.ShutdownWaitTime == 0 {
		return 10 * time.Second
	}
	return s.ShutdownWaitTime
}

func (s *ServeOptions) InitNilHandler() {
	if s.OnListenFailed == nil {
		s.OnListenFailed = func(err error) {}
	}
	if s.BeforeServe == nil {
		s.BeforeServe = func() {}
	}
	if s.OnServeFailed == nil {
		s.OnServeFailed = func(err error) {}
	}
	if s.BeforeShutdown == nil {
		s.BeforeShutdown = func() {}
	}
	if s.OnShutdownFailed == nil {
		s.OnShutdownFailed = func(err error) {}
	}
	if s.AfterShutdown == nil {
		s.AfterShutdown = func() {}
	}
}
