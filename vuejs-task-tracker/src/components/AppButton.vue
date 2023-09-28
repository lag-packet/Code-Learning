<script>
export default {
  name: "AppButton",
  props: {
    text: String,
    color: String,
  },
  methods: {
    onClick() {
      console.log("click");
    },
  },
};
</script>


<template>
  <button @click="onClick()" :style="{ backgroundColor: color }" class="btn">
    {{ text }}
  </button>
</template>