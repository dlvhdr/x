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
	// CommitStateUnknown indicates that state is not known.
	CommitStateUnknown CommitState = "UNKNOWN"
)

// CheckRunState is the possible states of a check run in a status rollup.
type CheckRunState string

const (
	// CheckRunStateActionRequired indicates the check run requires action.
	CheckRunStateActionRequired CheckRunState = "ACTION_REQUIRED"

	// CheckRunStateCancelled indicates the check run has been cancelled.
	CheckRunStateCancelled CheckRunState = "CANCELLED"

	// CheckRunStateCompleted indicates the check run has been completed.
	CheckRunStateCompleted CheckRunState = "COMPLETED"

	// CheckRunStateFailure indicates the check run has failed.
	CheckRunStateFailure CheckRunState = "FAILURE"

	// CheckRunStateInProgress indicates the check run is in progress.
	CheckRunStateInProgress CheckRunState = "IN_PROGRESS"

	// CheckRunStateNeutral indicates the check run was neutral.
	CheckRunStateNeutral CheckRunState = "NEUTRAL"

	// CheckRunStatePending indicates the check run is in pending state.
	CheckRunStatePending CheckRunState = "PENDING"

	// CheckRunStateQueued indicates the check run has been queued.
	CheckRunStateQueued CheckRunState = "QUEUED"

	// CheckRunStateSkipped indicates the check run was skipped.
	CheckRunStateSkipped CheckRunState = "SKIPPED"

	// CheckRunStateStale indicates the check run was marked stale by GitHub. Only GitHub can use this conclusion.
	CheckRunStateStale CheckRunState = "STALE"

	// CheckRunStateStartupFailure indicates the check run has failed at startup.
	CheckRunStateStartupFailure CheckRunState = "STARTUP_FAILURE"

	// CheckRunStateSuccess indicates the check run has succeeded.
	CheckRunStateSuccess CheckRunState = "SUCCESS"

	// CheckRunStateTimedOut indicates the check run has timed out.
	CheckRunStateTimedOut CheckRunState = "TIMED_OUT"

	// CheckRunStateWaiting indicates the check run is in waiting state.
	CheckRunStateWaiting CheckRunState = "WAITING"
)

// IsStatusWaiting checks if the status is pending/queued/in progress/waiting.
func IsStatusWaiting(status string) bool {
	return status == string(CheckRunStatePending) ||
		status == string(CheckRunStateQueued) ||
		status == string(CheckRunStateInProgress) ||
		status == string(CheckRunStateWaiting)
}

// IsConclusionASkip checks is the status is skipped.
func IsConclusionASkip(conclusion string) bool {
	return conclusion == string(CheckRunStateSkipped)
}

// IsConclusionAFailure checks is the status is a failure/timed out/startup failure.
func IsConclusionAFailure(conclusion string) bool {
	return conclusion == string(CheckRunStateFailure) ||
		conclusion == string(CheckRunStateTimedOut) ||
		conclusion == string(CheckRunStateStartupFailure)
}

// IsConclusionASuccess checks if the status is a success.
func IsConclusionASuccess(conclusion string) bool {
	return conclusion == string(CheckRunStateSuccess)
}

// IsConclusionNeutral checks if the status is neutral.
func IsConclusionNeutral(conclusion string) bool {
	return conclusion == string(CheckRunStateNeutral)
}
