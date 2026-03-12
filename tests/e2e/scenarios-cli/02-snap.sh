#!/bin/bash
# 02-snap.sh — CLI snapshot commands

source "$(dirname "$0")/common.sh"

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab snap (JSON format)"

pt_ok nav "${FIXTURES_URL}/form.html"
pt_ok snap
assert_output_json
assert_output_contains "nodes" "returns nodes array"
assert_output_contains "title" "returns page title"

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab snap --format text"

pt_ok snap --format text
assert_output_contains "[e" "contains element refs"
assert_output_not_contains "{" "not JSON format"

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab snap --filter interactive"

pt_ok snap --filter interactive
assert_output_json
assert_output_contains "textbox" "contains form inputs"
assert_output_contains "button" "contains buttons"

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab snap --tab <id>"

pt_ok nav "${FIXTURES_URL}/buttons.html"
TAB_ID=$(echo "$PT_OUT" | jq -r '.tabId')

pt_ok snap --tab "$TAB_ID"
assert_output_contains "buttons.html" "snapshot from correct tab"

end_test
