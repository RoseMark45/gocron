func (s *Scheduler) runJob(j *Job) {
	s.mu.Lock()
	// ... existing logic to prepare job execution ...
	s.mu.Unlock()

	// Execute the job
	j.run()

	// Execute listeners without holding the lock
	for _, listener := range j.afterJobRunsListeners {
		listener(j.name)
	}
}