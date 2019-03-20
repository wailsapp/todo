export async function fetch() {
  /* eslint-disable */
  try {
    let result = await backend.Todos.LoadTodos()
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
  await backend.Todos.Save(JSON.stringify(todos))
}

export async function saveAs(todos) {
  await backend.Todos.SaveAs(JSON.stringify(todos))
}

export async function loadTodosFromFile() {
  /* eslint-disable */
  try {
    let result = await backend.Todos.LoadTodosFromFile()
    console.log("RESULT = " + JSON.stringify(result))
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
