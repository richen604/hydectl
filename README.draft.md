# hydectl

[![Go Version](https://img.shields.io/github/go-mod/go-version/HyDE-Project/hydectl?style=for-the-badge)](https://go.dev/)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/HyDE-Project/hydectl?style=for-the-badge)](https://github.com/HyDE-Project/hydectl/releases)
[![AUR (git)](https://img.shields.io/aur/version/hydectl-git?style=for-the-badge)](https://aur.archlinux.org/packages/hydectl-git)
[![AUR (bin)](https://img.shields.io/aur/version/hydectl-bin?style=for-the-badge)](https://aur.archlinux.org/packages/hydectl-bin)
[![License](https://img.shields.io/github/license/HyDE-Project/hydectl?style=for-the-badge)](./LICENSE)

**hydectl** is the official command-line interface for the HyDE project, designed to be a powerful and extensible companion for the Hyprland compositor. It simplifies managing your desktop environment by providing a rich set of commands to control everything from wallpapers and themes to window layouts and system configuration, while also offering a powerful plugin system for ultimate customization.

## Description

- **What was your motivation?** Many desktop customizations rely on a series of disconnected shell scripts. `hydectl` was built to provide a unified, powerful, and easy-to-use tool to manage a Hyprland-based desktop environment, replacing scattered scripts with a robust, fast, and extensible Go application.
- **What problem does it solve?** It provides a single, consistent interface for common desktop actions like changing themes, managing wallpapers, and even complex window manipulations like creating tabbed container layouts. Its plugin system allows users to extend its functionality infinitely.
- **What did you learn?** This project is a deep dive into creating a feature-rich CLI with Go, using libraries like Cobra and Bubble Tea to create a polished user experience. It also explores the intricacies of interacting with Hyprland's IPC and designing a flexible plugin architecture.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Theme Management](#theme-management)
  - [Wallpaper Management](#wallpaper-management)
  - [Interactive Config Editor](#interactive-config-editor)
  - [Window Tabs](#window-tabs)
  - [Screen Zoom](#screen-zoom)
  - [Extending with Plugins](#extending-with-plugins)
- [Configuration](#configuration)
- [How to Contribute](#how-to-contribute)
- [License](#license)

## Features

- **Theming:** Interactively select, import, and switch between themes.
- **Wallpapers:** A full-featured wallpaper manager with support for multiple backends (swww, mpvpaper), interactive selection (rofi), and multiple monitors.
- **Interactive Config Editor:** A beautiful TUI for editing your configuration files, defined in a central registry.
- **i3/Sway-like Tabs:** Group all windows on a workspace into a tabbed layout.
- **Screen Zoom:** Smoothly zoom in and out of your screen.
- **Plugin System:** Extend `hydectl` with your own scripts. `hydectl` can automatically discover them and even generate CLI commands with flags and help text.
- **Pickers:** Includes `rofi`-based pickers for emojis and glyphs.
- **Hyprland Integration:** Built on the `hyprland-go` library for robust communication with the Hyprland compositor.

## Installation

There are several ways to install `hydectl`.

### Arch Linux (AUR)

`hydectl` is pushed to the archlinux AUR under [`hydectl-bin`](https://aur.archlinux.org/packages/hydectl-bin) and [`hydectl-git`](https://aur.archlinux.org/packages/hydectl-git).

- For the latest stable release (pre-compiled):

  ```sh
  yay -S hydectl-bin
  ```

- For the latest development version:

  ```sh
  yay -S hydectl-git
  ```

### Using `go install`

If you have a Go environment set up, you can install `hydectl` directly.

```sh
go install github.com/HyDE-Project/hydectl@latest
```

### Build from Source

You can also clone the repository and build it yourself.

```sh
# Make sure you have Go and Make installed
git clone https://github.com/HyDE-Project/hydectl.git
cd hydectl
make all      # Build the binary
sudo make install # Install the binary to /usr/local/bin
```

## Usage

`hydectl` provides several commands to manage your environment.

```sh
hydectl [command] --help
```

### Theme Management

The `theme` command is a powerful tool for managing your system's look and feel.

```sh
# Interactively select a theme
hydectl theme select

# Switch to the next/previous theme
hydectl theme next
hydectl theme prev

# Set a specific theme
hydectl theme set "Catppuccin-Mocha"

# Import themes from the hyde-gallery or a custom URL
hydectl theme import --name "My Awesome Theme" --url "https://github.com/user/my-awesome-theme"
```

### Wallpaper Management

The `wallpaper` command simplifies wallpaper handling.

```sh
# Interactively select a wallpaper (uses rofi)
hydectl wallpaper select

# Set the next/previous/random wallpaper
hydectl wallpaper next
hydectl wallpaper previous
hydectl wallpaper random

# Set a specific wallpaper
hydectl wallpaper set /path/to/your/image.png

# Specify a backend (default is swww)
hydectl wallpaper --backend mpvpaper next
```

### Interactive Config Editor

The `config` command launches a TUI to help you edit configuration files.

```sh
hydectl config
```

This command reads a `config-registry.toml` file to know which applications and files it can edit.

### Window Tabs

Group all windows in the current workspace into a single tabbed container.

```sh
hydectl tabs
```

Run the command again on a tabbed container to ungroup the windows.

### Screen Zoom

Zoom in, out, or reset the screen magnification.

```sh
# Zoom in with a specific intensity
hydectl zoom --in --intensity 0.2

# Zoom out with gradual steps for a smooth animation
hydectl zoom --out --step 0.05

# Reset zoom
hydectl zoom --reset
```

### Extending with Plugins

`hydectl`'s most powerful feature is its script-based plugin system. You can add any executable script to one of the script paths, and `hydectl` will make it available as a command.

**Script Paths:**

- `$XDG_CONFIG_HOME/lib/hydectl/scripts`
- `$HOME/.local/lib/hydectl/scripts`
- `/usr/local/lib/hydectl/scripts`
- `/usr/lib/hydectl/scripts`

**Example:**
Create a script at `~/.local/lib/hydectl/scripts/hello`:

```bash
#!/bin/sh
echo "Hello, from a hydectl plugin!"
```

Make it executable: `chmod +x ~/.local/lib/hydectl/scripts/hello`.

Now you can run it:

```sh
hydectl hello
# or
hydectl dispatch hello
```

For more advanced plugins, you can provide a JSON output for the `__usage__` argument to have `hydectl` generate a native command with flags and help text.

## Configuration

For the `hydectl config` command, you need to create a `config-registry.toml` file in `$XDG_CONFIG_HOME/hydectl/`.

Here is an example structure:

```toml
# $XDG_CONFIG_HOME/hydectl/config-registry.toml

# Order in which applications appear in the TUI
apps_order = ["Hyprland", "Kitty", "Rofi"]

[apps.Hyprland]
description = "Main configuration for the Hyprland compositor."
[apps.Hyprland.files.hyprland.conf]
path = "~/.config/hypr/hyprland.conf"
pre_hook = "echo 'About to edit Hyprland config...'"
post_hook = "hyprctl reload"

[apps.Kitty]
description = "Terminal emulator configuration."
[apps.Kitty.files.kitty.conf]
path = "~/.config/kitty/kitty.conf"
```

## How to Contribute

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## Contributing

Contributions are welcome! Please see our [CONTRIBUTING.md](./CONTRIBUTING.md) file for details on:

- Reporting bugs
- Suggesting enhancements
- Submitting code changes
- Commit message guidelines
- Pull request process

## License

This project is licensed under the terms of the [GNU General Public License v3.0](./LICENSE).
