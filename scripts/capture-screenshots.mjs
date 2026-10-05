import { chromium } from "playwright";
import { mkdir } from "node:fs/promises";

const baseURL = "http://127.0.0.1:8080";
const outDir = "docs/screenshots";
await mkdir(outDir, { recursive: true });

const browser = await chromium.launch({ headless: true });
try {
  const page = await browser.newPage({ viewport: { width: 1440, height: 1100 }, deviceScaleFactor: 1 });
  await page.goto(`${baseURL}/?year=2027`, { waitUntil: "networkidle" });
  await page.screenshot({ path: `${outDir}/month-geometry-matrix.png`, fullPage: true });

  const winter = page.locator("article").nth(5);
  await winter.screenshot({ path: `${outDir}/winter-rollover-suite.png` });

  const response = await page.request.get(`${baseURL}/v1/month/2027/2?weekStart=monday`);
  if (!response.ok()) throw new Error(`Month API returned ${response.status()}`);
  const result = await response.json();
  if (!result.invariants || Object.values(result.invariants).some((pass) => pass !== true)) {
    throw new Error("Month API invariants did not all pass");
  }
  const escape = (value) => String(value).replaceAll("&", "&amp;").replaceAll("<", "&lt;").replaceAll(">", "&gt;");
  const checks = Object.entries(result.invariants)
    .map(([name, pass]) => `<li><b class="${pass ? "pass" : "fail"}">${pass ? "PASS" : "FAIL"}</b> · ${escape(name)}</li>`)
    .join("");
  await page.setContent(`<!doctype html><meta charset="utf-8"><style>
    *{box-sizing:border-box}body{margin:0;background:#0b1020;color:#e9f0fb;font:16px/1.55 ui-sans-serif,system-ui;padding:44px}
    main{max-width:1150px;margin:auto}h1{font-size:30px;margin:0 0 8px}.sub{color:#9aacc7;margin-bottom:26px}
    section{background:#131b2f;border:1px solid #273550;border-radius:14px;padding:22px;margin:18px 0}
    ul{display:flex;gap:24px;list-style:none;padding:0}.pass{color:#55d68b}.fail{color:#ff7777}
    pre{white-space:pre-wrap;overflow-wrap:anywhere;color:#bfe9e4;font:13px/1.55 ui-monospace,monospace}
  </style><main><div class="sub">ENGINEERING DIAGNOSTICS · HTTP 200</div><h1>Month API · ${escape(result.monthName)} ${result.year}</h1>
  <div class="sub">${escape(result.weekStart)}-first · ${escape(result.naturalRows)} rows · ${escape(result.daysInMonth)} civil dates</div>
  <section><h2>Invariant results</h2><ul>${checks}</ul></section>
  <section><h2>GET /v1/month/2027/2?weekStart=monday</h2><pre>${escape(JSON.stringify(result, null, 2))}</pre></section></main>`);
  await page.screenshot({ path: `${outDir}/api-invariant-results.png`, fullPage: true });
} finally {
  await browser.close();
}
