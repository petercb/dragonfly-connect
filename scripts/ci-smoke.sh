#!/usr/bin/env bash
# Bounded smoke test for CI: proves startup, logs readiness, listens on UDP
# Bedrock port, then stops the server cleanly (SIGTERM).

set -euo pipefail

readonly READY_MARK='BedrockConnect ready!'
SMOKE_LOG="${SMOKE_LOG:-smoke.log}"
readonly WAIT_SECONDS="${WAIT_SECONDS:-30}"

rm -f "$SMOKE_LOG"

./dragonfly-connect >"$SMOKE_LOG" 2>&1 &
pid=$!

# shellcheck disable=SC2329  # invoked via EXIT trap
cleanup() {
	kill -TERM "$pid" 2>/dev/null || true
	# Ignore non-zero exit when the server is signaled.
	wait "$pid" 2>/dev/null || true
}
trap cleanup EXIT

deadline=$(( $(date +%s) + WAIT_SECONDS ))

while true; do
	if grep -Fq "$READY_MARK" "$SMOKE_LOG" 2>/dev/null; then
		break
	fi
	if ! kill -0 "$pid" 2>/dev/null; then
		echo "dragonfly-connect exited before ready:"
		cat "$SMOKE_LOG"
		exit 1
	fi
	if [ "$(date +%s)" -ge "$deadline" ]; then
		echo "timeout waiting for readiness (${WAIT_SECONDS}s):"
		cat "$SMOKE_LOG"
		exit 1
	fi
	sleep 0.25
done

if command -v ss >/dev/null 2>&1; then
	if ! ss -uln | grep -q ':19132'; then
		echo "expected UDP listener on :19132 not found (ss):"
		ss -uln || true
		cat "$SMOKE_LOG"
		exit 1
	fi
elif ! grep -Eq 'addr=.*:19132' "$SMOKE_LOG"; then
	echo "expected listener on :19132 not found (no ss; checking logs):"
	cat "$SMOKE_LOG"
	exit 1
fi

echo "Smoke OK: readiness log line and UDP listener on :19132"
exit 0
