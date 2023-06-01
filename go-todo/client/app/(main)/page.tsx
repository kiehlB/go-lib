"use client";

import Spinner from "@/components/spinner";
import TodoForm from "@/components/todo-form";
import TodoDataView from "@/components/todo-list";
import fetchTodo from "@/utils/fetch";

import React, { useEffect, useState } from "react";

interface TodoInfoProps {
  todoName: string;
}

interface Todo {}

function TodoInfo({ todoName }: TodoInfoProps) {
  const [state, setState] = useState<{
    status: "idle" | "pending" | "resolved" | "rejected";
    Todo: Todo | null;
    error: string | null;
  }>({
    status: todoName ? "pending" : "idle",
    Todo: null,
    error: null,
  });
  const { status, Todo, error } = state;

  useEffect(() => {
    if (!todoName) {
      return;
    }
    setState({ status: "pending", Todo: null, error: null });
    fetchTodo(todoName).then(
      (Todo) => {
        setState({ status: "resolved", Todo, error: null });
      },
      (error) => {
        setState({ status: "rejected", Todo: null, error });
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
    return <TodoDataView todo={Todo} />;
  }

  throw new Error("This should be impossible");
}

export default function Home() {
  const [todoName, setTodo] = React.useState("");

  const handleSubmit = (newtodoName) => {
    setTodo(newtodoName);
  };

  return (
    <main className='pt-4 col-span-5'>
      <div className='border p-4 rounded'>+ Add New Work</div>
      <TodoForm todoName={todoName} onSubmit={handleSubmit} />

      <TodoInfo todoName={todoName} />
    </main>
  );
}
