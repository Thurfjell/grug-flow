import { randomUUIDv7 } from "bun";
import { Todo, TodoModel } from "./model";

export const Service = (function () {
  const todos: Todo[] = [];

  return {
    getAll() {
      return todos;
    },
    create({ name }: TodoModel["createBody"]) {
      const newTodo: Todo = {
        name,
        id: randomUUIDv7(),
        createdAt: new Date().toISOString(),
      };
      todos.push(newTodo);
      return newTodo;
    },
  } as const;
})();
