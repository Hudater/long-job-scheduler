# long-job-scheduler

A local cron clone, built one chapter at a time alongside boot.dev's Learn Go course. Single process. No network, no isolation, no remote job fetch. Not a GitHub Actions runner, just cron with fake scripts.

## What a job is

One entity: `Job`. No separate "worker" concept until chapter 13 where a goroutine pulls jobs off a channel.

A Job has:

- Name
- Interval, how often it runs
- MaxDuration, how long it's allowed to run before it counts as failed
- Status. Comes out of a state machine driven by elapsed time against Interval and MaxDuration: pending before it starts, running while under MaxDuration, then done or failed depending on outcome.
- A Task it executes, arrives in chapter 6 through an interface

Status is a fixed, closed set of values. Treat it as an enum from the start, even before chapter 16 formalizes it with `iota`.

## Chapters: concepts, not signatures

Each chapter proves out one Go concept against the Job model above. Function and method names aren't contracts, the concept is what matters. Every increment has to compile and run before the next one starts.

1. Variables. One hardcoded job as loose variables: name, interval, status. Print it.
2. Constants and formatting. Status values become constants. `Printf`/`Sprintf` for output.
3. Conditionals. Branch on status for what message to print. Check interval and MaxDuration with plain `if` statements, no error type yet.
4. Functions. Pull job description and status logic into functions that take the job's data as arguments. Practice multiple return values, something like `(string, bool)` for a validity flag.
5. Structs. The loose data becomes a real `Job` struct. Methods replace free functions where that fits, `(j Job) Describe() string`.
6. Interfaces. `Task interface { Run() error }`, with a few implementations: PrintTask, SleepTask, a fake HTTPTask. Job holds a Task.
7. Errors. `Task.Run()` returns real errors. Wrap with `%w`, add a custom error type if it earns its keep, practice `errors.Is`/`errors.As`.
8. Loops. A scheduler tick: loop over jobs, run whichever are due. `range`, `break`/`continue` for skip conditions.
9. Slices. Jobs live in `[]Job`. Add and remove, sort by interval.
10. Maps. `map[string]Job` for lookup by name, alongside or instead of the slice.
11. Pointers. Status has to mutate after a run, so jobs become `[]*Job` or `map[string]*Job`. Value versus pointer semantics stop being theoretical here.
12. Packages and modules. Split into `task/`, `job/`, `scheduler/`, `cmd/`, with a proper `go.mod`. Add JSON persistence, `jobs.json` through `encoding/json` and `os.ReadFile`/`os.WriteFile`, so jobs survive a restart.
13. Channels. The worker concept shows up for real: a dispatcher goroutine sends due jobs on a channel, worker goroutines consume and run them, results and errors come back on another channel.
14. Mutexes. Multiple goroutines now touch the shared job map. `sync.Mutex` or `sync.RWMutex`. Run with `-race` first and watch it catch the problem before fixing it.
15. Generics. A generic `Queue[T any]` or `Result[T any]`, or a generic retry-with-backoff over any `func() (T, error)`.
16. Enums. Status gets rebuilt with Go's `iota` pattern, replacing the chapter 2 string constants.

## Ground rules

- Stdlib only, no external dependencies.
- Tasks are simulated and stay in memory. PrintTask prints, SleepTask sleeps, HTTPTask prints a URL. No real network calls, no spawned processes.
- One `main.go`, rewritten in place, until the chapter 12 package split. No `main_chN.go` files.
- `time` is fair game wherever the scheduler needs it (`Now`, `Duration`, `Ticker`), regardless of chapter.
- Each increment has to compile and run before the next one starts.
