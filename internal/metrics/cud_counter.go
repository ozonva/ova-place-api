package metrics

// CudCounter keeps counters for successful create/update/delete actions
type CudCounter struct {
	SuccessfulCreates Counter
	SuccessfulUpdates Counter
	SuccessfulDeletes Counter
}

// NewCudCounter returns CudCounter
func NewCudCounter(SuccessfulCreates Counter, SuccessfulUpdates Counter, SuccessfulDeletes Counter) CudCounter {
	return CudCounter{
		SuccessfulCreates: SuccessfulCreates,
		SuccessfulUpdates: SuccessfulUpdates,
		SuccessfulDeletes: SuccessfulDeletes,
	}
}
