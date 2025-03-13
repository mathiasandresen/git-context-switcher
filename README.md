# gitc (Git Context Switcher)

A command-line tool that makes it easy to switch between different Git contexts, managing SSH keys and Git configurations automatically. Perfect for developers who need to maintain separate Git identities for different projects (e.g., work and personal).

## Features

- Seamlessly switch between different Git contexts
- Automatically manage SSH keys through symbolic links
- Configure Git email per context
- Automatic SSH key addition to SSH agent
- YAML-based configuration

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/gitc.git
cd gitc
```

2. Build the project:
```bash
go build
```

3. Optionally, move the binary to your PATH to use it from anywhere:
```bash
sudo mv gitc /usr/local/bin/
```

## Configuration

Create a configuration file with your contexts. Each context can specify:
- SSH private and public keys
- Git email
- Context name

Example configuration:
```yaml
current_context: personal
contexts:
  - name: personal
    private_key: ~/.ssh/personal_id_rsa
    public_key: ~/.ssh/personal_id_rsa.pub
    email: personal@example.com
  - name: work
    private_key: ~/.ssh/work_id_rsa
    public_key: ~/.ssh/work_id_rsa.pub
    email: work@company.com
```

## Usage

```bash
# Switch to a specific context
gitc <context-name>

# Without arguments, switches to the current_context from config
gitc

# List available contexts
gitc list

# Initialize configuration
gitc init
```

When you switch contexts, the tool will:
1. Create symbolic links for SSH keys
2. Add the SSH key to your SSH agent
3. Configure your Git email

## Project Structure

```
.
├── cmd/            # Command-line interface implementation
│   └── root.go     # Main CLI command definition
├── config/         # Configuration handling
│   └── config.go   # YAML config parsing and validation
├── context/        # Context switching logic
│   └── context.go  # Core context switching functionality
├── utils/          # Utility functions
│   ├── git.go      # Git-related operations
│   ├── ssh.go      # SSH key management
│   └── symlink.go  # Symbolic link operations
└── main.go        # Application entry point
```

## Dependencies

- github.com/spf13/cobra - CLI framework
- gopkg.in/yaml.v2 - YAML configuration parsing

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Security Note

This tool manages SSH keys and Git configurations. Always ensure you:
- Keep your SSH keys secure
- Review the configuration file permissions
- Verify symbolic link targets before switching contexts
