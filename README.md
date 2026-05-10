# GrugFlow

A reference architecture for distributed UI composition using HTML-over-the-wire (HTMX).

GrugFlow demonstrates a server-driven UI system where:

- Core composes pages
- Layouts define structure and spacing
- Widgets are independent HTTP services returning HTML fragments
- HTMX assembles the final UI in the browser

---

## Core Idea

Layouts own:

- spacing
- grids
- visual consistency
- slot structure

Widgets own:

- content only

Core owns:

- which widgets go into which slots

This keeps layout concerns centralized and prevents design drift across pages.

---

## Architecture

Core server:

- chooses layout
- assigns widgets to slots
- serves page shell

Layouts:

- predefined page structures
- slot-based composition
- consistent HTML/CSS patterns

Widgets:

- standalone services
- return HTML fragments
- layout agnostic
- independently deployable

Browser:

- HTMX fetches widget fragments asynchronously
- page progressively assembles itself

---

## Example Flow

1. Core renders `TitleGrid`
2. Layout defines slots (`main`, `sidebar`, etc.)
3. HTMX requests widget endpoints
4. Widget services return HTML fragments
5. Browser swaps fragments into the page

---

## Why

Traditional component systems often distribute layout responsibility across components, leading to:

- inconsistent spacing
- wrapper hell
- layout drift
- duplicated styling decisions

GrugFlow instead treats layouts as strict UI contracts.

Pages choose layouts.
Widgets fill slots.
Layouts enforce consistency.

---

## Widget Contract

Each widget is simply:

GET /widgets/<name>

→ returns HTML fragment

No JSON APIs required.
No frontend framework required.
No shared runtime required.

---

## Goals

- Server-driven UI composition
- Independent widget services
- Layout-enforced consistency
- Minimal frontend complexity
- Failure isolation
- Async fragment loading via HTMX

---

## Status

Experimental / reference architecture.

This project is intended as:

- a demo
- a template
- an exploration of distributed HTML composition

Not a production framework (yet).

---

## Planned Features

- Widget caching
- Timeout handling
- Failure fallbacks
- Header propagation
- Widget registry
- Dockerized local environment

---

## Tech

- Go
- HTMX
- HTML templates
- Independent widget services (Go/Bun/etc.)

---

## Philosophy

Layout = shape  
Slots = holes  
Core = placement  
Widgets = content
