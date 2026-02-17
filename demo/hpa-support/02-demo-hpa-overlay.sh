#!/usr/bin/env bash
# 02-demo-hpa-overlay.sh — alias for the HPA overlay demo runner.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
exec "${SCRIPT_DIR}/02-demo-omit-replicas.sh" "$@"
