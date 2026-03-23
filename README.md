# envault

A CLI tool that encrypts `.env` files and backs them up to a private GitHub repository.

## Install

```bash
go install github.com/winnerx0/envault/cmd/envault@latest
```

Or build from source:

```bash
git clone https://github.com/winnerx0/envault.git
cd envault
go build -o envault ./cmd/envault
```

## Quick start

```bash
# 1. Login — set your password, GitHub token, and create a private backup repo
envault login -p <password> -t <github_token>

# 2. Initialize — run this in your project root
envault init

# 3. Backup — encrypt and push .env files to the private repo
envault backup

# 4. Recover — download and decrypt .env files on a new machine
envault recover
```

## Commands

### `envault login`

Sets your encryption password and GitHub token. Creates a private GitHub repo for storing encrypted backups and saves credentials to `~/.envault/config.yaml`.

```bash
envault login -p <password> -t <github_token>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--password` | `-p` | Encryption passphrase (required) |
| `--token` | `-t` | GitHub personal access token with `repo` scope |
| `--repo` | `-r` | Use an existing repo name instead of generating one |

### `envault init`

Creates an `envault.json` config file in the current directory, marking it as the project root.

```bash
envault init
```

### `envault backup`

Finds all `.env*` files from the project root, encrypts them, and pushes the encrypted data to the private GitHub repo.

```bash
envault backup
```

Files are stored under the project folder name in the repo (e.g. `myapp/.env.enc`, `myapp/config/.env.local.enc`), so multiple projects share one backup repo without conflicts.

### `envault recover`

Downloads encrypted `.env` files from the private GitHub repo and decrypts them back into the project directory.

```bash
envault recover
```

This fetches all `.enc` files for the current project from the backup repo, decrypts them using the passphrase in `~/.envault/config.yaml`, and writes the original `.env` files back to their original paths.

## Configuration

### Global config — `~/.envault/config.yaml`

Created by `envault login`. Stores credentials used across all projects.

```yaml
token: ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
passphrase: your-secret-password
repo: username/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6
```

### Project config — `envault.json`

Created by `envault init` in the project root.

```json
{
  "name": "myapp",
  "version": "1.0"
}
```

## How it works

1. **Login** saves your passphrase and GitHub token to `~/.envault/config.yaml` and creates a private repo via the GitHub API
2. **Init** creates `envault.json` in the current directory, marking it as the project root
3. **Backup** walks the project root for `.env*` files, encrypts each one with AES-256-GCM (key derived via Argon2), and pushes the encrypted bytes to GitHub using the Git Data API
4. **Recover** fetches the encrypted files from the GitHub repo using the Contents API, decrypts them, and restores the original `.env` files

## Encryption

- **Key derivation:** Argon2id (3 iterations, 32 MB memory, 4 threads)
- **Cipher:** AES-256-GCM
- Each file gets a unique random 16-byte salt and nonce

## Future features

- Backup and recover specific files (e.g. `envault backup .env.local`, `envault recover .env.production`)
