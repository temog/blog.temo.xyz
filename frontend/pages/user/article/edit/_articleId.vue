<template>
  <section class="container admin">
    <div class="columns">
      <div class="column is-one-fifth">
        <user-menu />
      </div>

      <div class="column">
        <h1 class="title">Edit</h1>

        <create-article />

        <div class="field mt-40">
          <div class="control">
            <button @click="preview" class="button">Preview</button>
            <button @click="update" class="button is-link">更新</button>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import CreateArticle from '@/components/CreateArticle'
import UserMenu from '@/components/UserMenu'
export default {
  components: {
    CreateArticle,
    UserMenu
  },
  data () {
    return {
      error: {
        title: null,
        alias: null,
        date: null,
        tags: null,
        body: null
      },
      pickrConfig: {
        enableTime: true,
        dataFormat: 'Y-m-d H:i',
        time_24hr: true,
        disableMobile: true
      },
      md: null,
      textarea: null
    }
  },
  computed: {
    ...mapGetters(['previewArticle'])
  },
  mounted () {
    this.findId(this.$route.params.articleId)
  },
  destroyed () {
  },
  methods: {
    ...mapActions(['editArticle', 'findId']),
    preview () {
      open('/user/article/preview', 'preview', 'width=800,height=' + window.outerHeight)
    },
    update () {
      console.log(this.previewArticle)
      this.editArticle(this.previewArticle)
    }
  }
}
</script>
