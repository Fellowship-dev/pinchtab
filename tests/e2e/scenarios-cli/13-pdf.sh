#!/bin/bash
# 13-pdf.sh — CLI PDF export command

source "$(dirname "$0")/common.sh"

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab pdf"

pt_ok nav "${FIXTURES_URL}/form.html"

# PDF outputs binary data to stdout by default
pt pdf
if [ "$PT_CODE" -eq 0 ] && [ -n "$PT_OUT" ]; then
  echo -e "  ${GREEN}✓${NC} pdf export succeeded"
  ((ASSERTIONS_PASSED++)) || true
else
  echo -e "  ${RED}✗${NC} pdf export failed or empty"
  ((ASSERTIONS_FAILED++)) || true
fi

end_test

# ─────────────────────────────────────────────────────────────────
start_test "pinchtab pdf -o <file>"

TMPFILE="/tmp/test-export-$$.pdf"
pt_ok pdf -o "$TMPFILE"

if [ -f "$TMPFILE" ] && [ -s "$TMPFILE" ]; then
  echo -e "  ${GREEN}✓${NC} pdf saved to file"
  ((ASSERTIONS_PASSED++)) || true
  rm -f "$TMPFILE"
else
  echo -e "  ${RED}✗${NC} pdf file not created or empty"
  ((ASSERTIONS_FAILED++)) || true
fi

end_test
