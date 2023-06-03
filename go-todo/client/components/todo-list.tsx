function TodoDataView({ todo }) {
  console.log(todo);
  return (
    <div>
      <section>
        <h2>
          {todo.map((e) => {
            return <div key={e.id}>{e.content}</div>;
          })}
        </h2>
      </section>
      <section>
        {/* <ul>
          <li>
            <label>User ID:</label> <span>{todo.user.id}</span>
          </li>
          <li>
            <label>User Name:</label> <span>{todo.user.name}</span>
          </li>
        </ul> */}
      </section>
      <small className='todo-info__fetch-time'>{todo.updatedAt}</small>
    </div>
  );
}

export default TodoDataView;
