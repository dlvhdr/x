// Package checks provides utlities for GitHub Actions checks
package checks

// CommitState is the status check rollup of a commit on GitHub.
type CommitState string

const (
	// CommitStateExpected indicates we are still waiting for the checks to start.
	CommitStateExpected CommitState = "EXPECTED"
	// CommitStateError indicates some error occurred when running the checks.
	CommitStateError CommitState = "ERROR"
	// CommitStateFailure indicates a failure.
	CommitStateFailure CommitState = "FAILURE"
	// CommitStatePending indicates some check is still pending.
	CommitStatePending CommitState = "PENDING"
	// CommitStateSuccess indicates that overall the checks have been successful.
	CommitStateSuccess CommitState = "SUCCESS"
	// CommitStateUnknown indicates that that the state is not known.
	CommitStateUnknown CommitState = "UNKNOWN"
)
