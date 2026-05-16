import Elysia from "elysia";
import { html } from "@elysia/html";
import { Service } from "./service";
import { TodoModel } from "./model";

const escape = (str: string) =>
  str.replace(/&/g, "&amp;")
     .replace(/</g, "&lt;")
     .replace(/>/g, "&gt;")
     .replace(/"/g, "&quot;")

export const todo = new Elysia({ prefix: "/todos" })
  .use(html())
  .get("/", async () => {
    const todos = Service.getAll();
    return `
    <div hx-get="/widgets/todos/" 
        hx-trigger="todo-added from:body" 
        hx-swap="outerHTML transition:true"
        class="bg-zinc-900 border border-zinc-800 rounded-xl overflow-hidden shadow-2xl">
        
        <div class="px-4 py-3 border-b border-zinc-800 bg-zinc-900/50 flex justify-between items-center">
            <h3 class="text-sm font-medium text-zinc-400 uppercase tracking-wider">Live Tasks</h3>
            <span class="px-2 py-0.5 text-xs bg-indigo-500/10 text-indigo-400 border border-indigo-500/20 rounded-full animate-pulse">
                Live
            </span>
        </div>

        <ul class="divide-y divide-zinc-800">
            ${
              todos.length > 0
                ? todos
                    .map(
                      (todo) => `
                <li id="${todo.id}" class="px-4 py-3 hover:bg-zinc-800/50 transition-colors group flex items-center justify-between">
                    <span class="text-zinc-200 group-hover:text-white">${escape(todo.name)}</span>
                    <div class="flex gap-2">
                      <button class="text-zinc-600 hover:text-red-400 text-xs uppercase font-bold tracking-tighter">Delete</button>
                    </div>
                </li>
            `,
                    )
                    .join("")
                : `
                <li class="px-4 py-8 text-center text-zinc-500 italic">
                    No tasks found. Add one!
                </li>
            `
            }
        </ul>
    </div>
    `;
  })
  .get(
    "/form",
    () => `
    <div class="bg-zinc-900 border border-zinc-800 rounded-xl p-4 shadow-2xl">
        <h3 class="text-sm font-medium text-zinc-400 uppercase tracking-wider mb-4">Add Task</h3>
        <form hx-post="/widgets/todos/" 
              hx-target="this" 
              hx-swap="none"
              hx-on:htmx:after:request="this.reset()"
              class="flex gap-2">
            <input type="text" 
                   name="name" 
                   placeholder="What needs doink?" 
                   class="flex-1 bg-zinc-950 border border-zinc-800 rounded-lg px-3 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-indigo-500 transition-all">
            <button type="submit" 
                    class="bg-indigo-600 hover:bg-indigo-500 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
                Add
            </button>
        </form>
    </div>
  `,
  )
  .post(
    "/",
    async ({ body, set }) => {
      set.status = 204 // no content
      set.headers["HX-Trigger"] = "todo-added";
      Service.create(body);
    },
    {
      body: TodoModel.createBody,
    },
  );
