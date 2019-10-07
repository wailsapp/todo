<template>
  <div>
    <h2 v-if="errorMessage.length > 0">{{errorMessage}}</h2>
    <section class="todoapp" v-cloak>
      <header class="header">
        <h1>todos</h1>
        <div class="buttons">
          <ul class="filters">
            <li>
              <a @click="saveAs">Save As</a>
            </li>
            <li>
              <a @click="loadNewList">Load</a>
            </li>
          </ul>
        </div>
        <input
          class="new-todo"
          autofocus
          autocomplete="off"
          placeholder="What needs to be done?"
          v-model="newTodo"
          @keyup.enter="addTodo"
        >
      </header>
      <section class="main" v-show="todos.length">
        <ul class="todo-list">
          <li
            class="todo"
            v-for="todo in todos"
            :key="todo.id"
            :class="{completed: todo.completed, editing: todo == editedTodo}"
          >
            <div class="view">
              <input class="toggle" type="checkbox" v-model="todo.completed">
              <label @dblclick="editTodo(todo)">{{todo.title}}</label>
              <button class="destroy" @click="removeTodo(todo)"></button>
            </div>
            <input
              class="edit"
              type="text"
              v-model.lazy="todo.title"
              @keyup.enter="doneEdit(todo)"
              @blur="doneEdit(todo)"
              @keyup.esc="cancelEdit(todo)"
              v-todo-focus="todo == editedTodo"
            >
          </li>
        </ul>
      </section>
    </section>
  </div>
</template>

<script>
import "./assets/css/base.css";
import "./assets/css/app.css";

import Wails from '@wailsapp/runtime';

export default {
  name: "app",
  data() {
    return {
      newTodo: "",
      editedTodo: null,
      errorMessage: "",
      loading: false,
      todos: []
    };
  },
  methods: {
    addTodo: function() {
      var value = this.newTodo && this.newTodo.trim();
      if (!value) {
        return;
      }
      this.todos.push({
        id: this.todos.length,
        title: value,
        completed: false
      });
      this.newTodo = "";
    },
    removeTodo: function(todo) {
      var index = this.todos.indexOf(todo);
      this.todos.splice(index, 1);

      for (var i = 0; i < this.todos.length; i++) {
        this.todos[i].id = i;
      }
    },

    editTodo: function(todo) {
      this.beforeEditCache = todo.title;
      this.editedTodo = todo;
    },

    doneEdit: function(todo) {
      if (!this.editedTodo) {
        return;
      }
      this.editedTodo = null;
      todo.title = todo.title.trim();
      if (!todo.title) {
        this.removeTodo(todo);
      }
    },
    cancelEdit: function(todo) {
      this.editedTodo = null;
      todo.title = this.beforeEditCache;
    },
    saveAs: function() {
      window.backend.Todos.SaveAs(JSON.stringify(this.todos, null, 2));
    },
    loadNewList: function() {
      window.backend.Todos.LoadNewList();
    },
    setErrorMessage: function(message) {
      this.errorMessage = message;
      setTimeout(() => {
        this.errorMessage = "";
      }, 3000);
    },
    loadList: function() {
      window.backend.Todos.LoadList()
        .then(list => {
          try {
            let todos = JSON.parse(list);
            this.loading = true;
            this.todos = todos;
          } catch (e) {
            this.setErrorMessage("Unable to load todo list");
          }
        })
        .catch(error => {
          this.setErrorMessage(error.message);
        });
    }
  },
  directives: {
    "todo-focus": function(el, binding) {
      if (binding.value) {
        el.focus();
      }
    }
  },
  watch: {
    todos: {
      handler: function(todos) {
        if (this.loading) {
          this.loading = false;
          return;
        }
        window.backend.Todos.SaveList(JSON.stringify(todos, null, 2));
      },
      deep: true
    }
  },
  mounted() {
    Wails.Events.On("filemodified", () => {
      this.loadList();
    });

    Wails.Events.On("error", (message, number) => {
      let result = number * 2;
      this.setErrorMessage(`${message}: ${result}`);
    });

    // Load the list at the start
    this.loadList();
  }
};
</script>
<style>
h2 {
  text-align: center;
  color: white;
  background-color: red;
  min-width: 230px;
  max-width: 550px;
  padding: 1rem;
  border-radius: 0.5rem;
}

.buttons {
  height: 20px;
  padding: 10px 20px;
  box-shadow: inset 0 -2px 1px rgba(0, 0, 0, 0.1);
  text-align: center;
  border-color: rgba(175, 47, 47, 0.2);
}

.buttons ul li a {
  margin: 10px;
}

.buttons li {
  border-color: rgba(175, 47, 47, 0.1);
}

.filters li a {
  color: inherit;
  margin: 3px;
  padding: 3px 7px;
  text-decoration: none;
  border: 1px solid transparent;
  border-radius: 3px;
  border-color: rgba(100, 100, 100, 0.1);
}
.filters li a:hover {
  border-color: rgba(255, 47, 47, 0.3);
  cursor: pointer;
}
</style>