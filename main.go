package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dheeraj-sn/go-resilient-runner/orchestrator"
	"github.com/dheeraj-sn/go-resilient-runner/task"
)

type ScratchCardInfo struct {
	ScratchCardID int64 `json:"scratch_card_id"`
}

type GetBalanceResponse struct {
	Balance         int             `json:"balance"`
	ScratchCardInfo ScratchCardInfo `json:"scratch_card_info"`
}

func assembleResponse(results []orchestrator.TaskResult) GetBalanceResponse {
	resp := GetBalanceResponse{
		Balance:         0,
		ScratchCardInfo: ScratchCardInfo{ScratchCardID: 0},
	}

	for _, r := range results {
		switch r.Name {
		case "get_balance":
			if val, ok := r.Value.(int); ok {
				resp.Balance = val
			}
		case "get_scratch_card_info":
			if val, ok := r.Value.(int64); ok {
				resp.ScratchCardInfo.ScratchCardID = val
			}
		}
	}
	return resp
}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 50*time.Millisecond)
	defer cancel()

	runner := &orchestrator.TaskRunner{}
	runner.AddTask(task.NewGetBalance(1234))
	runner.AddTask(task.NewGetScratchCardInfo(1234))

	results, err := runner.RunAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := assembleResponse(results)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/get_balance", GetBalanceHandler)
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
