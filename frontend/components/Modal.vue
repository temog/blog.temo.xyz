<template>
<div class="modal" :class="{'is-active': modal.body}">
  <div @click="modalOff" class="modal-background"></div>
  <div class="modal-content">
    <article class="message is-warning">
      <div class="message-header">
        <p>{{ modal.title }}</p>
        <button @click="modalOff" class="delete" aria-label="delete"></button>
      </div>
      <div class="message-body" v-html="modal.body">
      </div>
    </article>
  </div>
  <button @click="modalOff" class="modal-close is-large" aria-label="close"></button>
</div>
</template>
<script>
import { mapGetters, mapMutations } from 'vuex'
export default {
  computed: {
    ...mapGetters(['modal'])
  },
  methods: {
    ...mapMutations(['modalOff'])
  }
}
</script>
<style lang="scss" scoped>
</style>
