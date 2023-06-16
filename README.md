# Terraform Workspace Selector

The Terraform Workspace Selector is a command-line tool that helps you interactively select a Terraform workspace from the available options.

I generated this entirely using ChatGPT and don't know what I'm doing, so there might be problems or inefficiencies. Feel free to open issues.


## Functionality

The application provides the following basic functionality:

- Retrieves the list of Terraform workspaces using the `terraform workspace list` command.
- Displays an interactive prompt that allows you to select a workspace.
- Sets the selected workspace using the `terraform workspace select` command.


## Usage

To use the Terraform Workspace Selector, follow these steps:

1. Run `task build` to create a binary (located in `bin/`)
2. Copy that binary somewhere in your path.
3. Run `tfws` in your Terraform project.

Upon running the application, it will retrieve the list of Terraform workspaces and display an interactive prompt. Use the arrow keys or fuzzy search to navigate the options and press Enter to select a workspace. The selected workspace will be set using the `terraform workspace select` command.


## License

This project is licensed under the [MIT License](LICENSE).
