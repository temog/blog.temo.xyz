<template>
  <section class="container admin">
    <div class="columns">
      <div class="column is-one-fifth">
        <user-menu />
      </div>
      <div class="column">
        <h1 class="title">タグ追加</h1>

        <div class="field">
          <label class="label">タグ名</label>
          <div class="control">
            <input v-model="name" class="input" type="text">
          </div>
        </div>

        <div class="field">
          <div class="control">
            <button @click="addTag({name: name})" class="button is-link">追加</button>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>

<script>
import { mapActions } from 'vuex'
import UserMenu from '@/components/UserMenu'
export default {
  components: {
    UserMenu
  },
  data () {
    return {
      name: null
    }
  },
  methods: {
    ...mapActions(['addTag'])
  }
}
</script>
