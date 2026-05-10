# GrugFlow TODO — Widget System + Core Demo

## GOAL

Build a distributed HTMX-driven UI system:

- Core server composes pages
- Layout defines structure + HTMX wiring
- Widgets are independent HTTP services returning HTML
- Browser assembles UI via HTMX

---

## PHASE 1 — WIDGET SERVICES (NOW)

Each widget is a standalone HTTP service:

GET /widgets/url/path
→ returns HTML fragment

### Widget ideas to implement

- [ ] stats widget (Go or Bun)
  - CPU / memory mock data
  - simulate latency (100–800ms)

- [ ] activity widget
  - list of fake events
  - optionally randomized output

- [ ] deployments widget
  - simple service status list

### Requirements

- [ ] Content-Type: text/html
- [ ] No dependency on core server
- [ ] Runnable independently
- [ ] Simulate latency for realism

---

## PHASE 2 — CORE WIRING (DONE / NEAR DONE)

- [x] PageSpec defines widgets as URLs
- [x] Layout renders HTMX placeholders
- [x] HTMX loads widgets on page load
- [ ] Remove any local WidgetRenderer logic (if still present)

---

## PHASE 3 — LAYOUT UX IMPROVEMENTS

- [ ] TitleGrid layout uses proper grid container
- [ ] Add loading placeholders per widget
- [ ] Basic styling (Tailwind or minimal CSS)

---

## PHASE 4 — PRODUCTION REALISM LAYER

### CACHING

- [ ] Per-widget cache in core
  - TTL: 2–10 seconds
  - Key: widget URL

### FAILURE HANDLING

- [ ] If widget request fails:
  - Render fallback HTML block
  - Show “widget unavailable”
  - Do not break page rendering

### TIMEOUTS

- [ ] Per-widget HTTP timeout (1–2s max)

---

## PHASE 5 — DEMO POLISH

- [ ] Add slow widget (2–3s delay)
- [ ] Add failing widget (random 500 errors)
- [ ] Show HTMX async loading behavior
- [ ] Demonstrate resilience of dashboard

---

## ARCHITECTURE RULES (DO NOT BREAK)

- Core never renders widget HTML directly
- Widgets never know about layouts
- Layout never calls widget services directly (HTMX only)
- Widgets are fully independent services

---

## STRETCH GOALS (OPTIONAL)

- [ ] Widget registry (start static map, DB later)
- [ ] Per-user widget composition
- [ ] Auth header forwarding (core → widget)
- [ ] Docker compose for full local system

---

## DEFINITION OF DONE (MVP)

- /dashboard loads instantly
- widgets load asynchronously via HTMX
- at least one widget fails gracefully
- at least one widget is slow but non-blocking
- core contains zero widget rendering logic
