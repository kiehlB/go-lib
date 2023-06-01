function TodoDataView({ todo }) {
  return (
    <div>
      <div className='todo-info__img-wrapper'>
        <img src={todo.image} alt={todo.name} />
      </div>
      <section>
        <h2>
          {todo.name}
          <sup>{todo.number}</sup>
        </h2>
      </section>
      <section>
        <ul>
          {todo.attacks.special.map((attack) => (
            <li key={attack.name}>
              <label>{attack.name}</label>:{" "}
              <span>
                {attack.damage} <small>({attack.type})</small>
              </span>
            </li>
          ))}
        </ul>
      </section>
      <small className='todo-info__fetch-time'>{todo.fetchedAt}</small>
    </div>
  );
}

export default TodoDataView;
