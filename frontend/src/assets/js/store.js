export async function fetch() {
  /* eslint-disable */
  try {
    let result = await backend.loadTodos()
    if( !result) return [];
    let todos = JSON.parse(result);
    for(let i=0; i<todos.length; i++) {
      todos[i].id = i;
    }
    return todos;
  } catch (e) {
    console.log(e)
  }
}

export async function save(todos) {
  await backend.save(JSON.stringify(todos))
}
