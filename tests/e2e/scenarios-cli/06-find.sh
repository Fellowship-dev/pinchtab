#!/bin/bash
# 06-find.sh — CLI semantic find command

source "$(dirname "$0")/common.sh"

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab find <query>"

pt_ok nav "${FIXTURES_URL}/form.html"
pt_ok find "submit button"
assert_output_json
assert_output_contains "matches" "returns matches"

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab find (input field)"

pt_ok nav "${FIXTURES_URL}/form.html"
pt_ok find "username input"
assert_output_contains "textbox" "finds input element"

end_test
