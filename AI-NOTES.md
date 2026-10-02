# AI created notes — long-job-scheduler

## Topics to revisit (chapter 6-7: interfaces + errors)

Struggled most with these while wiring up the `Task` interface and switching to real `error` returns. Worth re-drilling later.

- **Struct literal syntax.** Mixed up "declaring a shape" with "filling in values" — tried `field: value` pairs directly on an anonymous struct without the second `{ }` block. Core pattern to remember: `TypeName{ field: value, ... }`, always two parts (type, then values).
- **Field vs method confusion.** Wrote `Task: Job.Run()` — tried to put a method _call_ where a field expected a _value_. A struct field holds data, not an action to perform.
- **Nested struct values.** Put a bare string directly as the `Task` field value instead of wrapping it in a full `PrintTask{printString: "..."}`. A field's type can itself be a whole other struct — you have to construct that inner value fully, not just hand over a raw piece of it.
- **Bool → error migration bugs (recurring, hit 3 times).** When switching functions from `(string, bool)` to `(string, error)`, kept checking `err != nil` for the _success_ path instead of `== nil`. Mental model didn't fully flip: bool's "true = good" is the opposite polarity of error's "non-nil = bad". Re-drill this specifically.
- **Side effects vs return values.** Called `Run()` twice — once to check the error, once again just to print it — not realizing `Run()` _does something real_ every time it's called (prints, sleeps, hits a URL), separate from whatever its return value is used for.
- **Redundant/overlapping conditions.** In `DescribeJobDuration`, had two `if` blocks where the first already fully covered a condition the second was still re-checking — dead logic that could never trigger. Watch for this after copy-pasting a condition.
- **Generic error messages.** First instinct was one broad message covering multiple distinct failure reasons (e.g. HTTP task's "empty URL or invalid method" combined into one string). Habit to build: one specific, contextual message per specific failure — include the actual bad value in the message via `fmt.Errorf("%v", ...)`.

### Queued for later (came up, not yet learned)

- Sentinel errors — declaring a named `var ErrX = errors.New(...)` at package level so callers can check identity via `errors.Is` instead of comparing message text. Colleague asked about this; need to actually learn it, not just used it.

## Topics to revisit (chapter 6-7: interfaces + errors)

Struggled most with these while wiring up `Task` interface and switching to real `error` returns. Worth re-drilling later.

- **Struct literal syntax.** Mixed up "declaring a shape" with "filling in values" — tried `field: value` pairs directly on an anonymous struct without the second `{ }` block. Core pattern to remember: `TypeName{ field: value, ... }`, always two parts (type, then values).
- **Field vs method confusion.** Wrote `Task: Job.Run()` — tried to put a method _call_ where a field expected a _value_. A struct field holds data, not an action to perform.
- **Nested struct values.** Put a bare string directly as the `Task` field value instead of wrapping it in a full `PrintTask{printString:
