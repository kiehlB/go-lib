export async function fetchTodo(name) {
  try {
    const response = await fetch("http://localhost:4000/api/v1/post", {
      method: "GET",
      headers: {
        "content-type": "application/json;charset=UTF-8",
      },
    });

    const data = await response.json();

    if (response.ok) {
      if (data) {
        return data;
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

export async function fetchWrieteTodo(content) {
  try {
    const response = await fetch("http://localhost:4000/api/v1/create", {
      method: "POST",
      headers: {
        "content-type": "application/json;charset=UTF-8",
      },
      body: JSON.stringify({
        Content: content,
        UserID: 1,
      }),
    });

    const data = await response.json();
    if (response.ok) {
      if (data) {
        return data;
      } else {
        throw new Error(`No todo with the name`);
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
