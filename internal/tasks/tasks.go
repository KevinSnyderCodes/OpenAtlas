package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/api"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/x/id"
	"github.com/hibiken/asynq"
)

const (
	TypeRunProcess = "run:process"
)

type RunProcessPayload struct {
	ID string
}

func NewRunProcessTask(id string) (*asynq.Task, error) {
	payload, err := json.Marshal(RunProcessPayload{ID: id})
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %w", err)
	}

	return asynq.NewTask(TypeRunProcess, payload), nil
}

type Handler struct {
	queries *db.Queries
}

func NewHandler(queries *db.Queries) *Handler {
	return &Handler{
		queries: queries,
	}
}

func (o *Handler) HandleRunProcessTask(ctx context.Context, task *asynq.Task) error {
	fmt.Println("Processing run:", string(task.Payload()))

	var payload RunProcessPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("error unmarshaling payload: %w", err)
	}

	// TODO: Implement

	runID, err := id.NewRunIDFromExternalID(payload.ID)
	if err != nil {
		return fmt.Errorf("error creating run id from external id: %w", err)
	}

	runInternalID := runID.InternalID()

	{
		arg := db.CreateTFEPlanParams{
			ID:         api.GenerateID(),
			RunID:      runInternalID,
			Status:     db.TfePlanStatusErrored,
			LogReadUrl: "", // TODO: Populate
		}
		if _, err := o.queries.CreateTFEPlan(ctx, arg); err != nil {
			return fmt.Errorf("error creating tfe plan: %w", err)
		}
	}

	{
		arg := db.UpdateTFERunStatusParams{
			ID:     runInternalID,
			Status: db.TfeRunStatusErrored,
		}
		if _, err := o.queries.UpdateTFERunStatus(ctx, arg); err != nil {
			return fmt.Errorf("error updating tfe run status: %w", err)
		}
	}

	return nil
}
