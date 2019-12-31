package main

// ArgumentError is a extensions of the Go Error that is thrown when a method argument is invalid.
type ArgumentError struct{}

func (argError *ArgumentError) Error() string {
	return "Invalid argument error"
}

// TeamDatabaseRepository encapsulates all methods for interacting with the team database.
type TeamDatabaseRepository interface {
	Create(string) string
}

func createTeam(teamService TeamDatabaseRepository, teamName string) (string, error) {
	response := ""

	if len(teamName) == 0 {
		return response, &ArgumentError{}
	}

	response = teamService.Create(teamName)

	return response, nil
}
