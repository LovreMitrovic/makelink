# makelink

A CLI tool that creates HTML shortcuts (`.html` files) that redirect to a given URL.

## Usage

```
makelink <name> <link>
```

**Example:**
```
makelink github https://github.com
# Creates github.html that redirects to https://github.com
```

## Build

Requires [Go](https://go.dev/dl/) 1.19 or later.

```bash
git clone https://github.com/LovreMitrovic/makelink.git
cd makelink
go build -o makelink .
```

## Add to PATH

Move the compiled binary to a directory in your `$PATH`:

```bash
sudo mv makelink /usr/local/bin/
```

Or install to your local bin (no sudo needed):

```bash
mkdir -p ~/.local/bin
mv makelink ~/.local/bin/

# If ~/.local/bin is not already in your PATH, add this to ~/.bashrc or ~/.zshrc:
export PATH="$HOME/.local/bin:$PATH"
```

Verify the install:

```bash
makelink --help
```

## Download

Linux binary: [Download](https://github.com/LovreMitrovic/makelink/raw/main/makelink)
