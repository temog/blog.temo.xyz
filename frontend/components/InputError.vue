<template>
  <p v-if="msg" class="help tag is-danger">
    {{ msg }}
  </p>
</template>

<script>
export default {
  props: ['msg']
}
</script>

<style scoped>
.tag {
  position: absolute;
  margin-top: -10px;
}
</style>
