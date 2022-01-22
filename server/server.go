package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tangwh1206/twitter/core"
	"github.com/tangwh1206/twitter/handlers"
)

type Server struct {
	HttpServer *http.Server
	Config     *core.ServerConfig
}

var (
	svr  *Server
	once sync.Once
)

func New(config *core.ServerConfig) *Server {
	once.Do(
		func() {
			svr = &Server{
				HttpServer: nil,
				Config:     config,
			}
		},
	)
	return svr
}

func (s *Server) Run() error {
	gin.SetMode(s.Config.Mode)
	router := gin.Default()

	handlers.InitRouter(router)
	addr := fmt.Sprintf(":%s", s.Config.ServerPort)
	readTimeout := 10 * time.Second
	if s.Config.ReadTimeout > 0 {
		readTimeout = s.Config.ReadTimeout
	}
	writeTimeout := 10 * time.Second
	if s.Config.WriteTimeout > 0 {
		writeTimeout = s.Config.WriteTimeout
	}
	serv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	s.HttpServer = serv

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	errCh := make(chan error, 1)
	go func() {
		log.Printf("Run in %s mode\n", s.Config.Mode)
		errCh <- s.HttpServer.Serve(listener)
	}()
	// start debug server for pprof
	startDebugServer(s.Config)

	// gracefully shutdown
	if err = waitSignal(errCh); err != nil {
		return err
	}
	return nil
}

func waitSignal(errCh <-chan error) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	// define your code(before shutdown)
	for {
		select {
		case sig := <-ch:
			return errors.New(sig.String())
		case err := <-errCh:
			return err
		}
	}
}
func startDebugServer(config *core.ServerConfig) {
	if !config.EnablePprof {
		log.Printf("Debug server not enable.")
		return
	}
	if len(config.DebugPort) == 0 {
		log.Printf("Debug port is not specified.")
		return
	}
	go func() {
		debugMux := http.NewServeMux()
		debugMux.HandleFunc("/debug/pprof/", pprof.Index)
		debugMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		debugMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		debugMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		debugMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		debugAddr := ":" + config.DebugPort
		listener, err := net.Listen("tcp", debugAddr)
		if err != nil {
			log.Printf("create debug server listener failed: %v\n", err)
			panic(err)
		}
		log.Fatal(http.Serve(listener, debugMux))
	}()
}
