# dragonfly-connect

<!-- markdownlint-disable MD013 -->
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/petercb/dragonfly-connect/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/petercb/dragonfly-connect/tree/main)
[![GitHub release (including prereleases)](https://img.shields.io/github/v/release/petercb/dragonfly-connect?include_prereleases&label=release)](https://github.com/petercb/dragonfly-connect/releases)
[![GHCR](https://img.shields.io/badge/ghcr.io-petercb%2Fdragonfly--connect-2496ED?logo=github)](https://github.com/petercb/dragonfly-connect/pkgs/container/dragonfly-connect)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Go](https://img.shields.io/github/go-mod/go-version/petercb/dragonfly-connect)](go.mod)
<!-- markdownlint-enable MD013 -->

**dragonfly-connect** is a small **Minecraft Bedrock Edition** proxy: players
connect to one address and immediately get an in-game **form menu** listing
configured servers; choosing an entry **transfers** them to that server. It is
built with the [Dragonfly](https://github.com/df-mc/dragonfly) server library
in Go.

The goal is a **lightweight hub** with **simple JSON configuration**—no custom
game logic and no built-in auth (target servers handle that).

## Inspiration: [BedrockConnect]

[BedrockConnect]: https://github.com/Pugmatt/BedrockConnect

[BedrockConnect][] exists so players—especially on consoles and other clients
that cannot type arbitrary server addresses—can still reach community servers
by joining through a **central menu** after their client is pointed at that menu
server (often via **DNS**). This project exists for the **same idea**: a single
entry point and a chooser instead of being locked to a fixed server list.

After discovering **Dragonfly**, the author wanted a **Go** implementation of
that pattern **on top of Dragonfly** rather than the original Java stack, for a
tight fit with the Dragonfly ecosystem and a small, hackable codebase.

## How it works

1. A Bedrock client connects to the proxy (default Dragonfly listen, typically
   **UDP `19132`**).
2. The proxy shows a **Bedrock form** (“Bedrock Connect”) with one button per
   configured server (optional images).
3. The player picks a server; the proxy calls **transfer** to the configured
   `host:port`.
4. Failures are **logged** and surfaced to the player (e.g. disconnect with an
   error message).

Configuration is **`servers.json`** in the process working directory (see the
example in this repo).

## Running with Docker

Images are published to **`ghcr.io/petercb/dragonfly-connect`** (tags match
releases, e.g. a version tag).

```bash
docker run -d --name dragonfly-connect \
  -p 19132:19132/udp \
  -v "$(pwd)/servers.json:/servers.json:ro" \
  ghcr.io/petercb/dragonfly-connect:<version>
```

Replace `<version>` with a tag from the
[packages page](https://github.com/petercb/dragonfly-connect/pkgs/container/dragonfly-connect)
or
[releases](https://github.com/petercb/dragonfly-connect/releases).
Ensure **`servers.json`** is mounted at **`/servers.json`** if you replace the
default file.

Bedrock uses **UDP**; the `-p` mapping must use **`19132/udp`** (or whichever
port you configure in Dragonfly).

## Running a release binary

1. Open [**Releases**](https://github.com/petercb/dragonfly-connect/releases)
   and download the archive for your platform (GoReleaser currently builds
   **Linux `amd64` and `arm64`** static binaries).
2. Extract the `dragonfly-connect` binary.
3. Place **`servers.json`** in the **same directory** you run from (the program
   loads `servers.json` from the current working directory).
4. Run:

   ```bash
   chmod +x dragonfly-connect
   ./dragonfly-connect
   ```

On **macOS** or **Windows**, or for other architectures, use **Docker** above or
build from source.

## Building from source

Requires [Go](https://go.dev/dl/) 1.26+ (see `go.mod`).

```bash
git clone https://github.com/petercb/dragonfly-connect.git
cd dragonfly-connect
# edit servers.json as needed
go run .
```

## DNS and Pi-hole (how clients reach your proxy)

Bedrock clients usually resolve **fixed hostnames** for featured or built-in
server entries. To use **dragonfly-connect** the same way as BedrockConnect-style
setups, run a **local DNS server** (for example
**[Pi-hole](https://pi-hole.net/)**) on your network and create **local DNS
records** (or custom host overrides) that point those hostnames at the **IP
address of the machine running this proxy**.

Then, when a console or other device uses that DNS, “joining” those servers
connects to **your** menu instead of the public address; the menu still
**transfers** players to the real destinations defined in `servers.json`.

Example hostnames taken from the sample **`servers.json`** (use only the
**hostname** part in DNS; ports stay part of the game transfer, not DNS):

| Server (label) | Example hostname from `address` |
| --- | --- |
| HiveGames | `geo.hivebedrock.network` |
| Cubecraft | `play.cubecraft.net` |
| Mineplex | `pe.mineplex.com` |

Adjust the list to match **your** `servers.json` and the **exact hostnames**
your clients use. On Pi-hole, use **Local DNS → DNS Records** (or your router’s
equivalent) to map each hostname to your proxy host’s LAN or public IP.

## Repository layout

- `servers.json` — menu entries: `name`, `address` (`host:port`), optional
  `image` URL.
- `docs/PRD.md` — product requirements and scope.

## License

GPL-3.0 — see [LICENSE](LICENSE).
