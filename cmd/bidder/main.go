package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	rtbHTTP "gitlab.com/radix-ads/rtb/internal/adapter/http"
	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/engine"
	"gitlab.com/radix-ads/rtb/internal/engine/pipeline"
	"gitlab.com/radix-ads/rtb/internal/metrics"
	"gitlab.com/radix-ads/rtb/internal/middleware"
	"gitlab.com/radix-ads/rtb/internal/pacing"
	"gitlab.com/radix-ads/rtb/internal/store"
)

func main() {
	fmt.Println("RTB bidder starting...")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	loader := store.NewStaticLoader()
	campaigns, err := loader.LoadCampaigns()
	if err != nil {
		log.Fatalf("failed to load campaigns: %v", err)
	}

	snapshot := store.BuildSnapshot(campaigns)
	atomicSnapshot := store.NewAtomicSnapshot(snapshot)

	store.StartReloadLoop(ctx, atomicSnapshot, loader, 30*time.Second)

	budgetStates := budget.BuildStates(campaigns)

	resetter := &budget.DailyResetter{
		States: budgetStates,
	}
	resetter.Start(ctx)

	budgetService := budget.NewService(budgetStates)
	budgetChecker := budget.NewChecker(atomicSnapshot, budgetStates)

	pacer := &pacing.Pacer{
		Snapshot:       snapshot,
		States:         budgetStates,
		Clock:          &pacing.SystemClock{},
		MaxBidFraction: 0.05,
		Tolerance:      1.05,
	}

	pl := pipeline.New(campaigns)
	eng := &engine.Engine{
		Pipeline: pl,
		Snapshot: snapshot,
		Checker:  budgetChecker,
		Pacer:    pacer,
	}

	metrics.RegisterBid()
	metrics.RegisterWin()

	bidHandler := &rtbHTTP.BidHandler{
		Engine: eng,
	}

	winHandler := &rtbHTTP.WinNotifyHandler{
		Budget: budgetService,
		Secret: os.Getenv("WIN_NOTIFY_SECRET"),
	}

	drain := &middleware.Drain{}

	mux := http.NewServeMux()
	mux.Handle("/bid", drain.Middleware(bidHandler))
	mux.Handle("/win", drain.Middleware(winHandler))
	mux.Handle("/metrics", metrics.Handler())

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Listening on ", server.Addr)
	go func() {
		if err = server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown initiated")

	drain.Close()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		log.Printf("server shutdown failed: %v\n", err)
	}

	drain.Wait()

	log.Println("shutdown complete")
}
