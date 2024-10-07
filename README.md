# tfws - Terraform Workspace Selector

The Terraform Workspace Selector is a command-line tool that helps you interactively select a Terraform workspace from the available options.

## Functionality

The application provides the following basic functionality:

- Retrieves the list of Terraform workspaces using the `terraform workspace list` command.
- Displays an interactive prompt that allows you to select a workspace.
- Sets the selected workspace using the `terraform workspace select` command.

## Requirements

`tfws` calls `terraform` directly, so terraform will need to be resolvable from your PATH. Tools like [tenv](https://github.com/tofuutils/tenv) or [asdf](https://github.com/asdf-vm/asdf) should natively work.

## Install
The Terraform Workspace Selector can be installed on macOS, Linux, or Windows, and should work with most standard shells, including PowerShell.

To install Terraform Workspace Selector, you can do either of the following:

- Install using Homebrew: `brew install abyss/tools/tfws`
- Download a binary from the [GitHub releases page](https://github.com/abyss/tfws/releases) and put it in your PATH.

## Usage

To use the Terraform Workspace Selector, just run `tfws` in a Terraform project with multiple workspaces.

Upon running the application, it will retrieve the list of Terraform workspaces and display an interactive prompt. Use the arrow keys or fuzzy search to navigate the options and press Enter to select a workspace. The selected workspace will be set using the `terraform workspace select` command.

## Contributing
Contributions in the form of issues and pull requests are welcome. 😄

## License
This project is licensed under the [ISC License](LICENSE.md).

Copyright (c) 2024 Abyss
