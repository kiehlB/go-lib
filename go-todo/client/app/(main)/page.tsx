"use client";

import Spinner from "@/components/spinner";
import TodoForm from "@/components/todo-form";
import TodoDataView from "@/components/todo-list";
import { fetchTodo, fetchWrieteTodo } from "@/utils/fetch";

import React, { useEffect, useState } from "react";

interface TodoInfoProps {
  todoName: string;
}

interface Todo {}

function TodoInfo({ todoName }: TodoInfoProps) {
  const [state, setState] = useState<{
    status: "idle" | "pending" | "resolved" | "rejected";
    todo: Todo | null;
    error: string | null;
  }>({
    status: todoName ? "pending" : "idle",
    todo: null,
    error: null,
  });
  const { status, todo, error } = state;

  useEffect(() => {
    setState({ status: "pending", todo: null, error: null });
    fetchTodo(todoName).then(
      (todo) => {
        setState({ status: "resolved", todo, error: null });
      },
      (error) => {
        setState({ status: "rejected", todo: null, error });
      }
    );
  }, [todoName]);

  if (status === "idle") {
    return <div>Submit a Todo</div>;
  } else if (status === "pending") {
    return <Spinner />;
  } else if (status === "rejected") {
    throw error;
  } else if (status === "resolved") {
    return <TodoDataView todo={todo} />;
  }

  throw new Error("This should be impossible");
}

export default function Home() {
  const [todoName, setTodo] = React.useState("");

  function handleSubmit(todo) {
    setTodo(todo);
  }

  return (
    <main className='pt-4 col-span-5'>
      <div className='border p-4 rounded'>+ Add New Work</div>
      <TodoForm todoName={todoName} onSubmit={handleSubmit} />

      <TodoInfo todoName={todoName} />
    </main>
  );
}
