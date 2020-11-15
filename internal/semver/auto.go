package semver

type AutoStrategy struct {
	GitCommitStrategy
	PatchStrategy
}

func NewAutoStrategy(gitCommitStrategy GitCommitStrategy) AutoStrategy {
	return AutoStrategy{GitCommitStrategy: gitCommitStrategy, PatchStrategy: PatchStrategy{}}
}

// GetLevel gets the level to increment using the AutoStrategy.
// It will attempt to determine the level with several strategies:
//		1. the GitCommitStrategy
// 		2. the PatchStrategy
// Returns the determined level or an error if anything went wrong.
func (s AutoStrategy) GetLevel() (level string, err error) {
	if level, err = s.GitCommitStrategy.GetLevel(); err != nil {
		return s.PatchStrategy.GetLevel()
	}
	return
}
