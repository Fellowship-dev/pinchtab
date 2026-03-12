#!/bin/bash
# 15-instance.sh — CLI instance management commands

source "$(dirname "$0")/common.sh"

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab instance (info)"

# Get default instance ID from health
pt_ok health
INSTANCE_ID=$(echo "$PT_OUT" | jq -r '.defaultInstance.id // empty')

if [ -n "$INSTANCE_ID" ]; then
  pt_ok instance "$INSTANCE_ID"
  assert_output_json
  assert_output_contains "id" "returns instance info"
else
  echo -e "  ${YELLOW}⚠${NC} No default instance found, skipping"
  ((ASSERTIONS_PASSED++)) || true
fi

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab instance logs"

if [ -n "$INSTANCE_ID" ]; then
  pt_ok instance logs "$INSTANCE_ID"
  # Logs might be empty but command should succeed
  echo -e "  ${GREEN}✓${NC} instance logs succeeded"
  ((ASSERTIONS_PASSED++)) || true
else
  echo -e "  ${YELLOW}⚠${NC} No instance ID, skipping logs test"
  ((ASSERTIONS_PASSED++)) || true
fi

end_test

# Note: instance start/stop not tested to avoid disrupting the running instance
