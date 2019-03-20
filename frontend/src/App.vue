<template>
  <div id="app">
    <section class="todoapp" v-cloak>
      <header class="header">
        <h1>todos</h1>
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
        <input id="toggle-all" class="toggle-all" type="checkbox" v-model="allDone">
        <label for="toggle-all">Mark all as complete</label>
        <ul class="todo-list">
          <li
            class="todo"
            v-for="todo in filteredTodos"
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
              v-model="todo.title"
              v-todo-focus="todo == editedTodo"
              @blur="doneEdit(todo)"
              @keyup.enter="doneEdit(todo)"
              @keyup.esc="cancelEdit(todo)"
            >
          </li>
        </ul>
      </section>
      <footer class="footer">
        <span class="todo-count">
          <strong v-text="remaining"></strong>
          {{pluralize('item', remaining)}} left
        </span>
        <button
          class="clear-completed"
          @click="load"
        >Load</button>
        <button
          class="clear-completed"
          @click="removeCompleted"
          v-show="todos.length > remaining"
        >Clear completed</button>
      <button
          class="clear-completed"
          @click="saveAs"
          v-show="todos.length > 0"
        >Save As</button>
      </footer>
    </section>
  </div>
</template>

<script>
import "./assets/css/todomvc-common.css";
import "./assets/css/todomvc-app.css";
import * as todoStorage from "./assets/js/store.js";
import EventBus from "./assets/js/eventbus";

// eslint-disable-next-line
var filters = {
  all: function(todos) {
    return todos;
  },
  active: function(todos) {
    return todos.filter(function(todo) {
      return !todo.completed;
    });
  },
  completed: function(todos) {
    return todos.filter(function(todo) {
      return todo.completed;
    });
  }
};

export default {
  name: "app",
  data() {
    return {
      newTodo: "",
      todos: [],
      editedTodo: null,
      visibility: "all"
    };
  },
  watch: {
    todos: {
      deep: true,
      handler: todoStorage.save
    }
  },
  directives: {
    "todo-focus": function(el, binding) {
      if (binding.value) {
        el.focus();
      }
    }
  },
  computed: {
    filteredTodos: function() {
      return filters[this.visibility](this.todos);
    },
    remaining: function() {
      return filters.active(this.todos).length;
    },
    allDone: {
      get: function() {
        return this.remaining === 0;
      },
      set: function(value) {
        this.todos.forEach(function(todo) {
          todo.completed = value;
        });
      }
    }
  },
  methods: {
    pluralize: function(word, count) {
      return word + (count === 1 ? "" : "s");
    },

    addTodo: function() {
      var value = this.newTodo && this.newTodo.trim();
      if (!value) {
        return;
      }
      this.todos.push({
        id: this.todos.length + 1,
        title: value,
        completed: false
      });
      this.newTodo = "";
    },

    removeTodo: function(todo) {
      var index = this.todos.indexOf(todo);
      this.todos.splice(index, 1);

      // Fix ids
      for(let i=0; i<this.todos.length; i++){ 
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

    removeCompleted: function() {
      this.todos = filters.active(this.todos);
    },

    saveAs: function() {
      todoStorage.saveAs(this.todos);
    },

    load: async function() {
      this.todos = await todoStorage.loadTodosFromFile();
    },
  },
  mounted() {
    var self = this;
    EventBus.$once("wails:ready", async function() {
      /* eslint-disable */
      try {
        self.todos = await todoStorage.fetch();
      } catch (e) {
        console.error(e);
      }
    });
  }
};
</script>
