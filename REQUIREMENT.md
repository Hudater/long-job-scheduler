# long-job-scheduler

A job scheduler built one chapter at a time alongside boot.dev's Learn Go course. The list below is the whole requirement.

## The 16 increments

1. **Variables.** Hardcode a fake job: name, interval (int seconds), status (string). Print it with fmt.Println. No functions yet, just var declarations and zero values.
2. **Constants + formatting.** Define job status as constants (StatusPending, StatusRunning, StatusDone, as strings or ints, your call). Use fmt.Printf/Sprintf to format the status line instead of default printing.
3. **Conditionals.** Decide what to print based on status: pending prints "waiting", running prints "in progress", and so on. Validate interval > 0, else print an error message. No error type yet, just a string check.
4. **Functions.** Extract job creation and status printing into functions. NewJobDescription(name string, interval int) string or similar. Practice multiple return values here, maybe (description string, valid bool).
5. **Structs.** The real model starts. Job struct { Name string; Interval int; Status string } with methods like (j Job) Describe() string. This replaces the loose variables from chapters 1-4.
6. **Interfaces.** Define Task interface { Run() error } and two or three concrete types: PrintTask, SleepTask, maybe a fake HTTPTask that just prints a URL. Job now holds a Task field. This is what lets job types vary later.
7. **Errors.** Task.Run() returns real errors now. Wrap them with fmt.Errorf("task %s failed: %w", name, err). Add a custom error type if useful, say TaskFailedError with the job name embedded. Practice errors.Is/As in whatever runs the job.
8. **Loops.** Write the scheduler tick as a for loop that walks through jobs once and runs whichever are due. Practice range, break and continue for skip conditions like disabled jobs.
9. **Slices.** Replace hardcoded jobs with []Job. Add and remove jobs from the slice. Practice append, slicing tricks, maybe sorting by interval with sort.Slice.
10. **Maps.** Index jobs by ID or name (map[string]Job) so lookup and cancel-by-name are O(1). Keep the slice too as the insertion-order list. Good excuse to reconcile both.
11. **Pointers.** Jobs need mutable state, status changes after running, so switch to map[string]*Job or []*Job. This is the chapter where value vs pointer semantics actually bite.
12. **Packages and modules.** Split into real packages: task/, job/, scheduler/, cmd/. Set up go.mod properly. This is where it stops being a single main.go. JSON persistence lands here too: jobs.json via encoding/json plus os.ReadFile/os.WriteFile, loaded at boot and saved on change, so jobs survive restarts.
13. **Channels.** Replace the single-threaded loop-tick scheduler with goroutines. A dispatcher sends due jobs on a channel, worker goroutines consume and run them, results and errors come back on another channel.
14. **Mutexes.** The job map is now touched by multiple goroutines, the scheduler reads while workers write status back. Add a sync.Mutex or sync.RWMutex around the shared map. Run with -race first and watch it catch the race before you fix it.
15. **Generics.** Write a generic Queue[T any] or Result[T any] instead of the channel just being chan Job. Or a generic retry-with-backoff function that works over any function returning (T, error).
16. **Enums.** Go back and redo Status with Go's iota enum pattern instead of the chapter 2 string constants. Small satisfying refactor that shows why the earlier hack was a hack.

## Ground rules

- Stdlib only. No external dependencies for the whole project.
- Tasks are simulated and stay in memory. PrintTask prints, SleepTask sleeps, HTTPTask prints a URL. Nothing touches the network and nothing spawns processes.
- One main.go, rewritten in place, until chapter 12. No main_chN.go files.
- The time package is fair game where the scheduler needs it (Now, Duration, Ticker). Everything else in an increment comes from its chapter.
- Each increment has to run before the next one starts.
