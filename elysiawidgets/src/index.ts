import { Elysia } from "elysia";
import { todo } from "./module/todo";
import { home } from "./module/home";

const app = new Elysia()
  .get("/", () => "Hello Elysia")
  .use(todo)
  .use(home)
  .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);
