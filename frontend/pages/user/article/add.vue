<template>
  <section class="container admin">
    <div class="columns">
      <div class="column is-one-fifth">
        <user-menu />
      </div>

      <div class="column">
        <h1 class="title">記事</h1>

        <create-article />

        <div class="field mt-40">
          <div class="control">
            <button @click="preview" class="button">Preview</button>
            <button @click="add" class="button is-link">追加</button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
import Util from '@/plugins/util'
import CreateArticle from '@/components/CreateArticle'
import UserMenu from '@/components/UserMenu'

export default {
  components: {
    UserMenu,
    CreateArticle
  },
  computed: {
    ...mapGetters(['previewArticle'])
  },
  mounted () {
    // 作成画面リロード以外は preview reset
    if (Util.storageGet('preview.id') !== 'create') {
      this.resetPreviewArticle()
    }

    this.setPreviewArticle({
      key: 'id',
      value: 'create'
    })

    if (!Util.storageGet('preview.date')) {
      this.setPreviewArticle({
        key: 'date',
        value: Util.date(new Date())
      })
    }
  },
  methods: {
    ...mapActions(['addArticle']),
    ...mapMutations(['setPreviewError', 'setPreviewArticle', 'resetPreviewArticle']),
    preview () {
      open('/user/article/preview', 'preview', 'width=800,height=' + window.outerHeight)
    },
    add () {
      if (!this.validation()) {
        return
      }

      this.addArticle(this.previewArticle)
    },
    validation () {
      this.setPreviewError()

      const e = {}

      let validate = true
      const f = this.previewArticle
      for (const key in f) {
        if (key === 'updated_at') {
          continue
        }
        if (!f[key]) {
          e[key] = 'required'
          validate = false
        }
      }

      if (!f.tags.length) {
        e.tags = 'required'
        validate = false
      }

      this.setPreviewError(e)

      if (!validate) {
        Util.message('notify', 'warning', 'input error', 'input error')
      }

      return validate
    }
  }
}
</script>
