import { Elysia } from "elysia";
import { todo } from "./module/todo";

const app = new Elysia()
  .get("/", () => "Hello Elysia")
  .use(todo)
  .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`,
);
