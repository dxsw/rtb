package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	rtbHTTP "gitlab.com/radix-ads/rtb/internal/adapter/http"
	"gitlab.com/radix-ads/rtb/internal/budget"
	"gitlab.com/radix-ads/rtb/internal/engine"
	"gitlab.com/radix-ads/rtb/internal/engine/pipeline"
	"gitlab.com/radix-ads/rtb/internal/metrics"
	"gitlab.com/radix-ads/rtb/internal/pacing"
	"gitlab.com/radix-ads/rtb/internal/store"
)

func main() {
	fmt.Println("RTB bidder starting...")

	loader := store.NewStaticLoader()
	campaigns, err := loader.LoadCampaigns()
	if err != nil {
		log.Fatalf("failed to load campaigns: %v", err)
	}

	snapshot := store.BuildSnapshot(campaigns)
	atomicSnapshot := store.NewAtomicSnapshot(snapshot)

	budgetStates := budget.BuildStates(campaigns)

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

	mux := http.NewServeMux()
	mux.Handle("/bid", bidHandler)
	mux.Handle("/win", winHandler)
	mux.Handle("/metrics", metrics.Handler())

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Listening on ", server.Addr)
	log.Fatal(server.ListenAndServe())
}
