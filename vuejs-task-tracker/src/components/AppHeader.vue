<script>
import AppButton from "./AppButton.vue";

export default {
  name: "AppHeader",
  props: {
    title: String,
  },
  components: {
    AppButton,
  },
};
</script>

<template>
  <header>
    <h1>{{ title }}</h1>
    <AppButton text="Add Task" color="green" />
  </header>
</template>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
</style>
