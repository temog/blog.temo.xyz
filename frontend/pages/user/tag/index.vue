<template>
  <section class="container admin">
    <div class="columns">
      <div class="column is-one-fifth">
        <user-menu />
      </div>
      <div class="column">
        <h1 class="title">タグ一覧</h1>

        <table class="table table is-striped">
          <thead>
            <tr><th>タグ名</th><th>記事数</th><th>action</th></tr>
          </thead>
          <tbody>
            <tr v-for="(tag, key) in tags" :key="key">
              <td>
                <input class="input" type="text" v-model="tag.name">
              </td>
              <td>{{ tag.articles }}</td>
              <td><a @click="saveTag(tag)" class="button">編集</a></td>
            </tr>
          </tbody>
        </table>

      </div>
    </div>
  </section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import UserMenu from '@/components/UserMenu'
export default {
  components: {
    UserMenu
  },
  data () {
    return {
      form: {
        tags: {}
      }
    }
  },
  computed: {
    ...mapGetters(['tags'])
  },
  mounted () {
    this.getTags()
  },
  methods: {
    ...mapActions(['getTags', 'updateTag']),
    saveTag (tag) {
      console.warn(tag)
      this.updateTag({ id: tag._id, name: tag.name })
    }
  }
}
</script>
