<script>
import AppTask from "./AppTask.vue";

export default {
  name: "AppTasks",
  props: {
    tasks: Array,
  },
  components: {
    AppTask,
  },
  emits: ["delete-task", "toggle-reminder"],
};
</script>


<template>
  <div v-for="task in tasks" :key="task.id">
    <AppTask
      @toggle-reminder="$emit('toggle-reminder', task.id)"
      @delete-task="$emit('delete-task', task.id)"
      :task="task"
    ></AppTask>
  </div>
</template>