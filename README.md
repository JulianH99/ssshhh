# What is this
---
ssshhh is an terminal app or tui that aims to help managing available
configurations found in the `~/.ssh/config` file. It's mainly built for personal
purposes since I have the need to work on several repositories all with
different ssh keys and domains and I didn't want to keep modifying the config
file manually.

## features
There's not much as of now, the app:

- List available configurations
- Create a new ssh key with the builtin `ssh-key` command
- Appents the config entry to the `~/.ssh/config` file

Planned features:

- See a detail of the selected configuration
- Execute the `ssh-key` command in the background without showing it's output to
	the user
- Have subcommands so that showing a list does not involve starting all the app,
	and the same for all other currently available features
- Add the key to the ssh-agent too (not hard, just haven't done it yet)
- Better ui, some styles were already refactored but I doubt the colors, maybe
in future iterations I will just use a colorscheme like gruvbox or kanagawa

## Missing stuff (extras)

- Still have to write some (most) unit tests
- Some code needs to be refactored
- No distribution mechanism (yet)

