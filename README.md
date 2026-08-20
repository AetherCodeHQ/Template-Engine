# Template Engine

![CI](https://github.com/Qyroxen/Template-Engine/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/Template-Engine?style=social)

> Powerful template engine for text files

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Template-Engine?style=social)](https://github.com/Qyroxen/Template-Engine/stargazers)

## What is it?

Template Engine processes text templates with variables, conditionals, and loops.

## Why should you care?

Manual text processing is error-prone. This tool automates it.

## Demo

```bash
./template-engine process --template config.yaml.tmpl --var name=prod
```

**Output:**
```
Processed template:
  environment: production
  database: prod-db.example.com
  debug: false
```

## Features

- Variable substitution
- Conditionals
- Loops
- File inclusion
- Custom functions

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Template-Engine.git
cd Template-Engine
go build -o template-engine .

# Run
./template-engine process --template config.tmpl
```

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--template` | Template file | (required) |
| `--var` | Variables (key=value) | `none` |
| `--output` | Output file | stdout |

## Examples

# Process template
./template-engine process --template config.tmpl --var env=prod

# With multiple vars
./template-engine process --template config.tmpl --var env=prod --var version=1.0

# Output to file
./template-engine process --template config.tmpl --output config.yaml

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Template-Engine/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Template-Engine?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Template-Engine/network/members">
    <img src="https://img.shields.io/github/forks/Qyroxen/Template-Engine?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Template-Engine/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Template-Engine" alt="Issues">
  </a>
</p>
