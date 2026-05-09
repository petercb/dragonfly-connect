# PRD: Dragonfly Bedrock Connect

## Product Vision
A lightweight Bedrock Edition proxy that allows players to connect to various servers from a central menu, built using the Dragonfly server library.

## Goals & Success Criteria
- Provide a seamless server selection experience for Bedrock players.
- Ensure robust error handling for server startup and player transfers.
- Support flexible configuration via files and environment variables.

## User Flow
1. **Connect**: Player connects to the Dragonfly Connect proxy IP.
2. **Menu**: Player is immediately presented with a Bedrock form menu showing available servers.
3. **Select**: Player selects a server button.
4. **Transfer**: Player is transferred to the target server address.

## Features & Requirements
- **Server Menu**: Dynamic menu generated from configuration.
- **Server Transfer**: Reliable transfer logic with error reporting to the player.
- **Configuration**: Support for `servers.json` and environment variables via `go-envconfig`.
- **Logging**: Structured logging using `slog`.

## Out of Scope
- Custom game logic (this is a pure proxy/menu).
- User authentication (handled by target servers).

## Acceptance Criteria
- Proxy starts without errors if configured correctly.
- Player sees the menu upon join.
- Player is transferred to the correct address when a button is clicked.
- Environment variables override or supplement JSON configuration.
