<template>
  <div class="container article">

    <div class="columns jumbotron">
      <div class="column is-three-fifths left">
        <div class="columnInner">
          <h1 class="title">{{ title }}</h1>

          <div class="date">{{ date | date }}</div>

          <div class="field is-grouped is-grouped-multiline">
            <div v-for="(t, key) in tags" :key="key" class="control">
              <span class="tags has-addons">
                <span class="tag">{{ tagName(t) }}</span>
                <span class="tag is-dark">{{ tagArticles(t) }}</span>
              </span>
            </div>
          </div>
        </div>
      </div>

      <div class="column right">
        <div class="columnInner">
          <img v-if="eyecatch" :src="eyecatch">
        </div>
      </div>
    </div>

    <div class="columns body">
      <div id="body" class="column has-background-white markdown-body" v-html="body"></div>
      <article-index v-if="body" :id="'body'" :body="body" />
    </div>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Util from '@/plugins/util'
import Markdown from 'markdown-it'
import Emoji from 'markdown-it-emoji'
import '@/assets/css/markdown.css'
import markdownItHighlight from 'markdown-it-highlight'
import 'markdown-it-highlight/dist/index.css'
import jump from 'jump.js'
import ArticleIndex from '@/components/ArticleIndex'

export default {
  components: {
    ArticleIndex
  },
  data () {
    return {
      title: null,
      body: null,
      tags: null,
      date: null,
      md: null,
      timer: null,
      scroll: null,
      eyecatch: null,
      latestEdited: 0,
      eyecatchUpdate: new Date('1999-01-01')
    }
  },
  computed: {
    ...mapGetters(['tagName', 'previewScroll', 'previewEyecatch']),
    markdown () {
      if (!this.body) {
        return
      }
      return this.md.render(this.body)
    }
  },
  mounted () {
    this.md = new Markdown()
    this.md.use(markdownItHighlight)
    this.md.use(Emoji)

    this.watchArticle()

    const id = Util.storageGet('preview.id')
    if (id !== 'create') {
      this.eyecatch = Util.eyecatchUrl(id)
    }
  },
  destroyed () {
    clearTimeout(this.timer)
  },
  methods: {
    watchArticle () {
      const scroll = localStorage.getItem('previewScroll')
        ? parseInt(localStorage.getItem('previewScroll')) : 0

      if (this.scroll !== scroll) {
        const clientHeight = document.querySelector('.container').clientHeight
        jump(
          (scroll / 100 * clientHeight) - document.documentElement.scrollTop,
          { duration: 500 }
        )
      }
      this.scroll = scroll

      const latestEdited = Util.storageGet('preview.latestEdited')
      if (!latestEdited || latestEdited === this.latestEdited) {
        this.watchLoop()
        return
      }
      this.latestEdited = latestEdited
      console.log('watch update')

      this.title = Util.storageGet('preview.title')
      this.tags = Util.storageGet('preview.tags') ? Util.storageGet('preview.tags') : []
      this.date = Util.storageGet('preview.date')

      const eyecatchUpdate = new Date(Util.storageGet('preview.eyecatchUpdate'))
      if (eyecatchUpdate && eyecatchUpdate > this.eyecatchUpdate) {
        this.eyecatch = Util.storageGet('preview.eyecatch')
        this.eyecatchUpdate = eyecatchUpdate
      }

      const body = Util.storageGet('preview.body')
      if (body) {
        this.body = this.md.render(body)
      }

      this.watchLoop()
    },
    watchLoop () {
      this.timer = setTimeout(() => {
        this.watchArticle()
      }, 3000)
    },
    tagArticles (tagId) {
      for (let i = 0; i < this.tags.length; i++) {
        const tag = this.tags[i]
        if (tag._id === tagId) {
          return tag.articles
        }
      }
    }
  }
}
</script>

<style>
.topTags {
  height: 100%;
}
</style>
