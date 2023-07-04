<script>
import AppHeader from "./components/AppHeader";
import AppTasks from "./components/AppTasks.vue";
import AddTask from "./components/AddTask";

export default {
  name: "App",
  components: {
    AppHeader,
    AppTasks,
    AddTask,
  },
  data() {
    return {
      tasks: [],
    };
  },
  methods: {
    deleteTask(id) {
      if (confirm("Are you sure?")) {
        this.tasks = this.tasks.filter((task) => task.id !== id);
      }
    },
    toggleReminder(id) {
      this.tasks = this.tasks.map((task) =>
        task.id === id ? { ...task, reminder: !task.reminder } : task
      );
      /*this.tasks = this.tasks.map((task) => {
        if (task.id === id) {
          return { ...task, reminder: !task.reminder };
        } else {
          return task;
        }
      });*/
    },
    addTask(newTask) {
      this.tasks = [...this.tasks, newTask];
    },
  },
  created() {
    this.tasks = [
      {
        id: 1,
        text: "Go jim",
        day: "March 1st at 2:30pm",
        reminder: true,
      },
      {
        id: 2,
        text: "Party time",
        day: "March 2nd at 4:30pm",
        reminder: true,
      },
      {
        id: 3,
        text: "Sleep in",
        day: "March 3rd at 9:00pm",
        reminder: false,
      },
    ];
  },
};
</script>

<template>
  <div class="container">
    <AppHeader title="Task Tracker" />
    <AddTask @add-task="addTask" />
    <AppTasks
      @toggle-reminder="toggleReminder"
      @delete-task="deleteTask"
      :tasks="tasks"
    />
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400&display=swap");

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: "Poppins", sans-serif;
}

.container {
  max-width: 500px;
  margin: 30px auto;
  overflow: auto;
  min-height: 300px;
  border: 1px solid steelblue;
  padding: 30px;
  border-radius: 5px;
}

.btn {
  display: inline-block;
  background: #000;
  color: #fff;
  border: none;
  padding: 10px 20px;
  margin: 5px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
}

.btn:focus {
  outline: none;
}

.btn:active {
  transform: scale(0.98);
}

.btn-block {
  display: block;
  width: 100%;
}
</style>
