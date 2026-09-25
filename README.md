# Rolarte CLI

Rolarte CLI is a terminal application for standardizing and automating audiovisual production workflows used by Rolarte.

The initial version focuses on creating and maintaining standardized directory structures for projects using tools such as Adobe Audition, After Effects, Premiere Pro, and Character Animator.

## Goals

- Standardize audiovisual project directory structures.
- Reduce repetitive project setup.
- Provide an interactive terminal UI.
- Keep filesystem operations predictable and safe.
- Grow into a larger suite of Rolarte production tools.

## Initial Features

### Create Project

Creates a new System Project.

The user provides:

- Project name.
- Subprojects to include.
- Number of initial clips.

All subprojects are selected by default.

The default number of clips is:

```text
1
```

Project names are normalized using uppercase snake case.

Example:

```text
Outta Bubblegum
↓
OUTTA_BUBBLEGUM
```

Before modifying the filesystem, Rolarte CLI will show a preview and request confirmation.

### Create Clip

Adds one or more clips to an existing System Project.

Rolarte CLI will:

1. List existing projects.
2. Inspect the selected project's existing clips.
3. Determine the next available clip number.
4. Ask how many new clips should be created.
5. Preview the directories to create.
6. Request confirmation before modifying the filesystem.

Existing clips must never be overwritten.

### Config

Rolarte CLI will provide configuration for values such as the root directory where System Projects are stored.

Example:

```bash
rolarte config
```

## System Project Structure

A System Project can contain the following subprojects:

```text
PROJECT_NAME/
├── 1_MICROFICTION/
│   ├── 1.1_ASSETS/
│   ├── 1.2_AUDIO/
│   ├── 1.3_ANIMATION/
│   └── 1.4_EXPORTS/
│
├── 2_SYSTEM_OVERVIEW/
│   ├── 2.1_ASSETS/
│   ├── 2.2_AUDIO/
│   ├── 2.3_REEL/
│   └── 2.4_EXPORTS/
│
├── 3_PODCAST/
│   ├── 3.1_GAME_SESSION/
│   │   ├── 3.1.1_AUDIO/
│   │   ├── 3.1.2_VIDEO/
│   │   └── 3.1.3_EXPORTS/
│   │
│   └── 3.2_ROLAFTER/
│       ├── 3.2.1_AUDIO/
│       ├── 3.2.2_VIDEO/
│       └── 3.2.3_EXPORTS/
│
└── 4_CLIPS/
    └── 4.1_CLIP1/
        ├── 4.1.1_ASSETS/
        ├── 4.1.2_AUDIO/
        ├── 4.1.3_ANIMATION/
        └── 4.1.4_EXPORTS/
```

Additional clips follow the same structure:

```text
4.2_CLIP2/
4.3_CLIP3/
4.4_CLIP4/
...
```

## Project Creation Flow

```text
Main Menu
    ↓
Create New Project
    ↓
Enter Project Name
    ↓
Select Subprojects
    ↓
Select Initial Clip Count
    ↓
Preview Generation Plan
    ↓
Confirm
    ↓
Create Directories
```

Filesystem changes should only occur after confirmation.

## Development

Run the application:

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

Format source files:

```bash
gofmt -w .
```

Build:

```bash
go build -o rolarte ./cmd/rolarte
```

Run the compiled binary:

```bash
./rolarte
```

## Status

Early development.

The initial goal is to establish a reliable project scaffolding workflow before expanding Rolarte CLI into a broader production toolkit.