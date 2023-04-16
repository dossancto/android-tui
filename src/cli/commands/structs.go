package commands

type CommandFunc func(arg []string) error

type Command struct {
	Name        string
	Action          CommandFunc
	Description string
}
