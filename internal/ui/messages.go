package ui

type CreateKeyMsg struct{}

type KeyCreatedMsg struct{}

type SshFileEditedMsg struct{}

type ErrorMsg struct {
	Err error
}
