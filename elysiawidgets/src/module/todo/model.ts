import { t, UnwrapSchema } from "elysia";

const Todo = t.Object({
  id: t.String({ format: "uuid" }),
  name: t.String(),
  createdAt: t.String({ format: "date" }),
  completedAt: t.Optional(t.String({ format: "date" })),
});

export type Todo = UnwrapSchema<typeof Todo>

export const TodoModel = {
  createBody: t.Object({
    name: t.String({ minLength: 2 }),
  }),
  createReponse: Todo,
  getAllResponse: t.Array(Todo),
} as const;

export type TodoModel = {
  [k in keyof typeof TodoModel]: UnwrapSchema<(typeof TodoModel)[k]>;
};
