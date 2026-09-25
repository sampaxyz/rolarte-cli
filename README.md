# Rolarte CLI

[![Release](https://img.shields.io/github/v/release/sampaxyz/rolarte-cli?style=flat-square&sort=semver)](https://github.com/sampaxyz/rolarte-cli/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/sampaxyz/rolarte-cli?style=flat-square&logo=go)](https://github.com/sampaxyz/rolarte-cli/blob/master/go.mod)
[![CI](https://github.com/sampaxyz/rolarte-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/sampaxyz/rolarte-cli/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sampaxyz/rolarte-cli)](https://goreportcard.com/report/github.com/sampaxyz/rolarte-cli)
[![Homebrew](https://img.shields.io/badge/homebrew-sampaxyz%2Ftap-orange?style=flat-square&logo=homebrew)](https://github.com/sampaxyz/homebrew-tap)
[![License](https://img.shields.io/github/license/sampaxyz/rolarte-cli?style=flat-square)](LICENSE)

A terminal toolkit for managing Rolarte audiovisual production workspaces.

Built with Go, Bubble Tea, Lip Gloss, Cobra, GoReleaser and Homebrew.

<p align="center">
  <img src="https://raw.githubusercontent.com/sampaxyz/rolarte-cli/master/record-1.gif" alt="Rolarte CLI demo" width="500">
</p>

---

## Installation

### Homebrew

```bash
brew tap sampaxyz/tap
brew install --cask sampaxyz/tap/rolarte
```

Then run:

```bash
rolarte
```

### GitHub Releases

Prebuilt macOS binaries are also available from:

https://github.com/sampaxyz/rolarte-cli/releases

---

## Usage

Launch the interactive TUI:

```bash
rolarte
```

The dashboard provides access to:

```text
Create Project
Create Clip
Configuration
Exit
```

You can also run workflows directly from the command line.

### Create a project

```bash
rolarte create project
```

The project workflow lets you:

```text
Enter project name
Select production modules
Choose initial clip count
Preview generated directories
Confirm creation
```

Project names are normalized automatically.

```text
Outta Bubblegum
↓
OUTTA_BUBBLEGUM
```

### Create clips

```bash
rolarte create clip
```

Select an existing System Project and create one or more additional clips.

Rolarte automatically discovers the next available clip number.

Example:

```text
4.1_CLIP1
4.2_CLIP2

+ 3 clips

↓

4.3_CLIP3
4.4_CLIP4
4.5_CLIP5
```

Existing clips are never overwritten.

### Configuration

```bash
rolarte config
```

Configures the root directory where System Projects are stored.

The same configuration can also be edited from the TUI dashboard.

### Version

```bash
rolarte version
```

Example:

```text
Rolarte CLI 1.1.3
commit: abc1234
built: 2026-09-25T...
```

---

## Keyboard Controls

### Dashboard

```text
↑ / ↓     Navigate
Enter     Select
Q         Quit
```

### Project modules

```text
↑ / ↓     Navigate
Space     Enable / disable
Enter     Continue
Esc       Back
```

### Preview

```text
← / →     Change page
Enter     Confirm
Esc       Back
```

### Forms

```text
Enter     Continue
Esc       Back
Ctrl+C    Quit
```

---

## Workspace Structure

A System Project can include production areas such as:

```text
PROJECT_NAME/
├── 1_MICROFICTION/
├── 2_SYSTEM_OVERVIEW/
├── 3_PODCAST/
│   ├── 3.1_GAME_SESSION/
│   └── 3.2_ROLAFTER/
└── 4_CLIPS/
    ├── 4.1_CLIP1/
    ├── 4.2_CLIP2/
    └── ...
```

Each production area contains a standardized audiovisual workspace for assets, audio, animation or video, and exports.

---

## Updating

If Rolarte was installed with Homebrew:

```bash
brew update
brew upgrade --cask rolarte
```

Check the installed version with:

```bash
rolarte version
```

---

## Development

Clone the repository:

```bash
git clone https://github.com/sampaxyz/rolarte-cli.git
cd rolarte-cli
```

Run locally:

```bash
go run ./cmd/rolarte
```

Run tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

Format:

```bash
gofmt -w .
```

Build:

```bash
go build -o rolarte ./cmd/rolarte
```

---

## Releases

Releases are built automatically with GoReleaser when a version tag is pushed.

Example:

```bash
git tag -a v1.2.0 -m "Rolarte CLI v1.2.0"
git push origin v1.2.0
```

The release pipeline publishes:

```text
GitHub Release
macOS universal binary
SHA-256 checksums
Homebrew Cask
```

---

## Tech Stack

- Go
- Bubble Tea
- Bubbles
- Lip Gloss
- Cobra
- GoReleaser
- GitHub Actions
- Homebrew

---

## License

Rolarte CLI is released under the [MIT License](LICENSE).
