package git

import "github.com/restechnica/backbone-cli/internal/commands"

const DefaultTag = "0.0.0"

type CLIService struct {
	commander commands.Commander
}

func NewCLIService(commander commands.Commander) CLIService {
	return CLIService{commander: commander}
}

// CreateTag creates an annotated git tag.
// Returns an error if the command failed.
func (service CLIService) CreateTag(tag string) (err error) {
	return service.commander.Run("git", "tag", "-a", tag, "-m", tag)
}

// GetLatestCommitMessage gets the latest commit message from git.
// Returns the commit message.
func (service CLIService) GetLatestCommitMessage() (message string, err error) {
	return service.commander.Output("git", "show", "-s", "--format='%s'")
}

// GetTag gets the latest annotated git tag.
// It returns the latest annotated git tag or the default "0.0.0" tag if the git command fails.
func (service CLIService) GetTag() (output string) {
	var err error

	if output, err = service.commander.Output("git", "describe"); err != nil {
		return DefaultTag
	}

	return
}
