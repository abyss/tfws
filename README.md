# Terraform Workspace Selector

The Terraform Workspace Selector is a command-line tool that helps you interactively select a Terraform workspace from the available options.

## Functionality

The application provides the following basic functionality:

- Retrieves the list of Terraform workspaces using the `terraform workspace list` command.
- Displays an interactive prompt that allows you to select a workspace.
- Sets the selected workspace using the `terraform workspace select` command.

## Install

To install Terraform Workspace Selector, you can do either of the following:
- Run `go install github.com/abyss/tfws@latest`

- Download a binary from GitHub and put it in your PATH

`tfws` calls `terraform` directly, so terraform will need to be resolvable from your path. Tools like [tenv](https://github.com/tofuutils/tenv) or [asdf](https://github.com/asdf-vm/asdf) should natively work.

## Usage

To use the Terraform Workspace Selector, just run `tfws` in a Terraform project with multiple workspaces.

Upon running the application, it will retrieve the list of Terraform workspaces and display an interactive prompt. Use the arrow keys or fuzzy search to navigate the options and press Enter to select a workspace. The selected workspace will be set using the `terraform workspace select` command.


## License

This project is licensed under the [MIT License](LICENSE).
