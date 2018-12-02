<template>
  <div v-if="loading" class="loading">
    <fa icon="spinner" class="fa-spin" />
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  computed: {
    ...mapGetters(['loading'])
  }
}
</script>
<style lang="scss" scoped>
.loading {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
  &:before {
    content:"";
    display: block;
    position: absolute;
    background-color: rgba(0, 0, 0, 0.5);
    width: 100%;
    height: 100%;
    z-index: 0;
  }
  svg {
    font-size: 60px;
    color: #fff;
  }
}
@media (max-width: 1087px) {
}
</style>
