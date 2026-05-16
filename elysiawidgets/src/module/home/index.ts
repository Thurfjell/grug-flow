import Elysia from "elysia";
import { html } from "@elysia/html";

export const home = new Elysia({ prefix: "/home" })
    .use(html())
    .get("/", async () => {
        return `
<div class="flex flex-col gap-4 p-2">

  <div class="text-center py-10">
    <div class="mb-4">
      <span class="inline-flex items-center gap-2 px-3 py-1 rounded-full text-xs font-medium bg-indigo-950 text-indigo-400 border border-indigo-900 animate-pulse">
        <span class="w-1.5 h-1.5 rounded-full bg-indigo-400"></span>
        live and absolutely cooked
      </span>
    </div>
    <h1 class="text-6xl font-black tracking-tighter text-white mb-3">
      grug<span class="text-indigo-400">flow</span>
    </h1>
    <p class="text-zinc-500 font-mono text-sm">widgets all the way down. no cap.</p>
    <div class="flex gap-2 justify-center flex-wrap mt-4">
      ${["go", "htmx", "bun", "elysia", "zig?"].map(t =>
            `<span class="bg-indigo-950 text-indigo-300 border border-indigo-900 rounded px-2 py-0.5 text-xs font-mono">${t}</span>`
        ).join("")}
    </div>
  </div>

  <div class="grid grid-cols-2 md:grid-cols-4 gap-3">
    ${[
                ["0ms", "nav load time", "text-indigo-400"],
                ["404", "ms to install pkgs", "text-emerald-400"],
                ["0", "shared services", "text-pink-400"],
                ["∞", "grug satisfaction", "text-orange-400"],
            ].map(([val, label, color]) => `
      <div class="bg-zinc-900 border border-zinc-800 rounded-xl p-4 text-center">
        <div class="text-3xl font-black ${color}">${val}</div>
        <div class="text-xs text-zinc-500 mt-1">${label}</div>
      </div>
    `).join("")}
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
    ${[
                ["widget isolation", "one widget dies. others don't care. blast radius: zero. sleep well."],
                ["vertical slices", "query what you need. skip the 3k LOC service wrapper boogaloo."],
                ["any language", "go. bun. zig. cobol if you want. just return html."],
            ].map(([title, desc]) => `
      <div class="bg-zinc-900 border border-zinc-800 rounded-xl p-5">
        <div class="text-white font-medium mb-2">${title}</div>
        <p class="text-zinc-500 text-sm leading-relaxed">${desc}</p>
      </div>
    `).join("")}
  </div>

  <p class="text-center text-zinc-700 font-mono text-xs py-4">
    GET /widgets/:name/ → text/html → done. that's the whole API.
  </p>

</div>
`
    })