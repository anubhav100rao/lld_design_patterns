package prototype

import "time"

// JobConfig holds all parameters for an ETL run
type JobConfig struct {
	Name       string
	SourceDSN  string
	Table      string
	TargetDSN  string
	Transform  string
	Schedule   time.Duration
	RetryCount int
}

// Clone deep‑copies a JobConfig
func (j *JobConfig) Clone() *JobConfig {
	copy := *j
	return &copy
}

// Usage: set up prototypes at startup
var (
	hourlyJob = &JobConfig{
		SourceDSN:  "user:pass@tcp(src)/db",
		TargetDSN:  "user:pass@tcp(dst)/db",
		Transform:  "normalize_dates",
		Schedule:   time.Hour,
		RetryCount: 3,
	}
	dailyJob = &JobConfig{
		SourceDSN:  hourlyJob.SourceDSN,
		TargetDSN:  hourlyJob.TargetDSN,
		Transform:  "aggregate_daily",
		Schedule:   24 * time.Hour,
		RetryCount: 1,
	}
)

func ScheduleJobs() {
	for _, proto := range []*JobConfig{hourlyJob, dailyJob} {
		cfg := proto.Clone()
		cfg.Name = proto.Transform + "_job"
		launchCron(cfg.Schedule, func() {
			runETL(cfg)
		})
	}
}
func runETL(cfg *JobConfig) {
	// Simulate ETL run
	println("Running ETL job:", cfg.Name)
	println("Source:", cfg.SourceDSN)
	println("Target:", cfg.TargetDSN)
	println("Transform:", cfg.Transform)
	println("Retry Count:", cfg.RetryCount)
	// Actual ETL logic would go here
}
func launchCron(interval time.Duration, task func()) {
	// Simulate a cron job that runs every interval
	go func() {
		for {
			task()
			time.Sleep(interval)
		}
	}()
}
