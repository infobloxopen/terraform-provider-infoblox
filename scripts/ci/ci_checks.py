#!/usr/bin/env python3
"""CI checks for the provider, run from the repo root.

Each check is a subcommand and exits non-zero when it finds a problem, so the
PR workflow can gate on it. Currently:

  python3 scripts/ci/ci_checks.py orphans   # unused generated model files
"""
from __future__ import annotations

import re
import sys
from pathlib import Path

REPO = Path(".")

# Identifier-like token, and exported (capitalised) top-level type/var/const
# declarations in both the single-line and grouped `type ( ... )` forms.
TOKEN_RE = re.compile(r"[A-Za-z_][A-Za-z0-9_]*")
DECL_RE = re.compile(r"^(?:type|var|const)\s+([A-Z][A-Za-z0-9_]*)", re.MULTILINE)
GROUP_DECL_RE = re.compile(r"^\s+([A-Z][A-Za-z0-9_]*)\s", re.MULTILINE)


def check_orphans() -> int:
    """Report generated model files whose exported names are used nowhere else.

    Generated model files (`model_*.go`) declare a Go type plus its AttrTypes /
    schema vars. If none of a file's exported names are referenced by any other
    file, the object was removed or restructured upstream and this file is dead
    output. The fix is to regenerate from the source definitions, not to
    hand-delete here.
    """
    # 1. Every Go file in the repo except vendored dependencies.
    go_files = [
        p for p in REPO.rglob("*.go")
        if "/vendor/" not in str(p) and not str(p).startswith("vendor/")
    ]

    # 2. Candidate files = generated models, mapped to the exported names each
    #    declares. These are the files we might flag as orphans.
    candidates: dict[str, set[str]] = {}
    for p in _model_files():
        names = _exported_names(p.read_text(encoding="utf-8", errors="replace"))
        if names:
            candidates[str(p)] = names

    # 3. For each declared name, find which files mention it (a reference). We
    #    only track the names candidates declare — nothing else matters here.
    wanted = set().union(*candidates.values()) if candidates else set()
    refs: dict[str, set[str]] = {name: set() for name in wanted}
    for p in go_files:
        text = p.read_text(encoding="utf-8", errors="replace")
        for tok in wanted.intersection(TOKEN_RE.findall(text)):
            refs[tok].add(str(p))

    # 4. A file is an orphan when none of its names are referenced by any *other*
    #    still-alive file. Removing an orphan can leave a nested model with no
    #    referrer, so repeat until nothing new drops out (a fixpoint). `level`
    #    records the round: 1 = orphan outright, 2+ = orphaned only after the
    #    files above it were removed.
    alive = {str(p) for p in go_files}
    orphans: list[tuple[int, str, set[str]]] = []
    level = 0
    while True:
        level += 1
        found = [
            f for f, names in candidates.items()
            if f in alive
            and not any(r != f and r in alive for n in names for r in refs[n])
        ]
        if not found:
            break
        for f in found:
            alive.discard(f)
            orphans.append((level, f, candidates[f]))

    if not orphans:
        print("No orphan model files found.")
        return 0

    print(f"Orphan model files ({len(orphans)}):\n")
    for level, f, names in sorted(orphans):
        tag = "" if level == 1 else f"  [cascade L{level}]"
        print(f"  {f}{tag}")
        print(f"      declares (unused): {', '.join(sorted(names))}")
    print("\nThe files above are orphaned. Delete them or regenerate from source to resolve.")
    return 1


def _model_files() -> list[Path]:
    """Generated model files: internal/service/**/model_*.go."""
    svc = REPO / "internal" / "service"
    if not svc.is_dir():
        return []
    return sorted(p for p in svc.rglob("model_*.go") if not p.name.endswith("_test.go"))


def _exported_names(text: str) -> set[str]:
    """Exported type/var/const names declared in a file (incl. grouped blocks)."""
    names = set(DECL_RE.findall(text))
    for block in re.findall(r"^(?:type|var|const)\s*\((.*?)^\)", text, re.MULTILINE | re.DOTALL):
        names |= set(GROUP_DECL_RE.findall(block))
    return names


# Register a new check by adding "name": function here (each returns an exit code).
CHECKS = {
    "orphans": check_orphans,
}


def main() -> None:
    if len(sys.argv) != 2 or sys.argv[1] not in CHECKS:
        sys.exit(f"usage: {sys.argv[0]} {{{'|'.join(CHECKS)}}}")
    sys.exit(CHECKS[sys.argv[1]]())


if __name__ == "__main__":
    main()
