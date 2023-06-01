import { useState, useEffect } from "react";

const formatDate = (date) =>
  `${date.getHours()}:${String(date.getMinutes()).padStart(2, "0")} ${String(
    date.getSeconds()
  ).padStart(2, "0")}.${String(date.getMilliseconds()).padStart(3, "0")}`;

export default async function fetchTodo(name) {
  try {
    const response = await fetch("", {
      method: "POST",
      headers: {
        "content-type": "application/json;charset=UTF-8",
      },
      body: JSON.stringify({
        query: "",
        variables: { name: name.toLowerCase() },
      }),
    });

    const { data } = await response.json();
    if (response.ok) {
      const todo = data?.todo;
      if (todo) {
        todo.fetchedAt = formatDate(new Date());
        return todo;
      } else {
        throw new Error(`No todo with the name "${name}"`);
      }
    } else {
      const error = {
        message: data?.errors?.map((e) => e.message).join("\n"),
      };
      throw error;
    }
  } catch (error) {
    throw error;
  }
}
