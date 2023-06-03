"use client";

import { fetchWrieteTodo } from "@/utils/fetch";
import React, { useState, useEffect } from "react";

interface TodoFormProps {
  todoName: string;
  initialTodoName?: string;
  onSubmit(e: string): void;
}

function TodoForm({
  todoName: externalTodoName,
  initialTodoName = externalTodoName || "",
  onSubmit,
}: TodoFormProps) {
  const [TodoName, setTodoName] = useState(initialTodoName);

  useEffect(() => {
    if (typeof externalTodoName === "string") {
      setTodoName(externalTodoName);
    }
  }, [externalTodoName]);

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    setTodoName(e.target.value);
  }

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    fetchWrieteTodo(TodoName);
    onSubmit(TodoName);
  }

  return (
    <form onSubmit={handleSubmit} className='py-4'>
      <label htmlFor='TodoName-input'>Add Todo Name</label>

      <div className='pt-2'>
        <input
          className='border-b py-2 mr-4 outline-none'
          id='TodoName-input'
          name='TodoName'
          placeholder='Write Todo...'
          value={TodoName}
          onChange={handleChange}
        />
        <button
          type='submit'
          className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'
          disabled={!TodoName.length}
        >
          Submit
        </button>
      </div>
    </form>
  );
}

export default TodoForm;
