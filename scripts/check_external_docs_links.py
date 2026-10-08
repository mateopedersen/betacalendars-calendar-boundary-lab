"""Check unique HTTPS links in Markdown documentation using the standard library."""

from pathlib import Path
import re
import sys
import time
import urllib.error
import urllib.request


ROOT = Path(__file__).resolve().parents[1]
LINK_PATTERN = re.compile(r"\[[^\]]*\]\((https?://[^)\s]+)(?:\s+[^)]*)?\)")
FENCE_PATTERN = re.compile(r"(?ms)^```.*?^```\s*")


def main() -> int:
    links: dict[str, set[str]] = {}
    for path in (ROOT / "README.md", *(ROOT / "docs").rglob("*.md")):
        if not path.exists():
            continue
        content = FENCE_PATTERN.sub("", path.read_text(encoding="utf-8"))
        for url in LINK_PATTERN.findall(content):
            links.setdefault(url.rstrip(".,"), set()).add(str(path.relative_to(ROOT)))

    failures = []
    for url, files in sorted(links.items()):
        for attempt in range(3):
            request = urllib.request.Request(
                url,
                headers={"User-Agent": "CalendarBoundaryLab-doc-link-check/1.0"},
            )
            try:
                with urllib.request.urlopen(request, timeout=20) as response:
                    status = response.status
                if status < 400:
                    print(f"{status} {url}")
                    break
                error = f"HTTP {status}"
            except urllib.error.HTTPError as exc:
                error = f"HTTP {exc.code}"
                if exc.code not in (429, 500, 502, 503, 504) or attempt == 2:
                    failures.append((url, files, error))
                    break
            except (TimeoutError, urllib.error.URLError) as exc:
                error = str(exc)
                if attempt == 2:
                    failures.append((url, files, error))
                    break
            time.sleep(attempt + 1)
        else:
            failures.append((url, files, error))

    for url, files, error in failures:
        print(f"FAIL {error} {url} ({', '.join(sorted(files))})", file=sys.stderr)
    return int(bool(failures))


if __name__ == "__main__":
    raise SystemExit(main())
