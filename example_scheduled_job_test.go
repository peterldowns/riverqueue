package river_test

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/riverqueue/river"
	"github.com/riverqueue/river/internal/riverinternaltest"
	"github.com/riverqueue/river/internal/util/slogutil"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type ScheduledArgs struct {
	Message string `json:"message"`
}

func (ScheduledArgs) Kind() string { return "scheduled" }

type ScheduledWorker struct {
	river.WorkerDefaults[ScheduledArgs]
}

func (w *ScheduledWorker) Work(ctx context.Context, job *river.Job[ScheduledArgs]) error {
	fmt.Printf("Message: %s\n", job.Args.Message)
	return nil
}

// Example_scheduledJob demonstrates how to schedule a job to be worked in the
// future.
func Example_scheduledJob() {
	ctx := context.Background()

	// Required for purposes of our example here, but in reality t will be the
	// *testing.T that comes from a test's argument.
	t := &testing.T{}
	dbPool := riverinternaltest.TestDB(ctx, t)
	defer dbPool.Close()

	// Required for the purpose of this test, but not necessary in real usage.
	if err := riverinternaltest.TruncateRiverTables(ctx, dbPool); err != nil {
		panic(err)
	}

	workers := river.NewWorkers()
	river.AddWorker(workers, &ScheduledWorker{})

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Logger: slog.New(&slogutil.SlogMessageOnlyHandler{Level: slog.LevelWarn}),
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers: workers,
	})
	if err != nil {
		panic(err)
	}

	if err := riverClient.Start(ctx); err != nil {
		panic(err)
	}

	_, err = riverClient.Insert(ctx,
		ScheduledArgs{
			Message: "hello from the future",
		},
		&river.InsertOpts{
			// Schedule the job to be worked in three hours.
			ScheduledAt: time.Now().Add(3 * time.Hour),
		})
	if err != nil {
		panic(err)
	}

	// Unlike most other examples, we don't wait for the job to be worked since
	// doing so would require making the job's scheduled time contrived, and the
	// example therefore less realistic/useful.

	if err := riverClient.Stop(ctx); err != nil {
		panic(err)
	}

	// Output:
}
