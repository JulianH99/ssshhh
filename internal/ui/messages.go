package ui

type CreateKeyMsg struct{}

type KeyCreatedMsg struct{ Key string }

type SshFileEditedMsg struct{}

type ErrorMsg struct {
	Err error
}
