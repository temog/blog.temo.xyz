<template>
  <div v-if="article" class="container article">

    <div class="columns jumbotron">
      <div class="column is-three-fifths left">
        <div class="columnInner">
          <h1 class="title">{{ article.title }}</h1>

          <div class="date">{{ article.date | date }}</div>

          <div class="field is-grouped is-grouped-multiline">
            <div v-for="(t, key) in article.tags" :key="key" class="control">
              <span @click="$router.push('/tag/' + tagName(t))" class="tags has-addons">
                <span class="tag is-light">{{ tagName(t) }}</span>
                <span class="tag is-dark">{{ tagArticles(t) }}</span>
              </span>
            </div>
          </div>

          <div v-if="isSignIn">
            <router-link :to="'/user/article/edit/' + article._id" class="button is-link">Edit</router-link>
          </div>
        </div>
      </div>

      <div class="column right">
        <div class="columnInner">
          <img :src="eyecatch">
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
import { mapGetters, mapActions } from 'vuex'
import Util from '@/plugins/util'
import Markdown from 'markdown-it'
import Emoji from 'markdown-it-emoji'
import '@/assets/css/markdown.css'
import markdownItHighlight from 'markdown-it-highlight'
import 'markdown-it-highlight/dist/index.css'
import ArticleIndex from '@/components/ArticleIndex'

export default {
  components: {
    ArticleIndex
  },
  data () {
    return {
      md: null,
      lazyImages: []
    }
  },
  computed: {
    ...mapGetters(['article', 'tagName', 'tags', 'isSignIn']),
    eyecatch () {
      return Util.eyecatchUrl(this.article._id)
    },
    body () {
      if (!this.md) {
        return
      }
      let md = this.article.old ? this.article.body.replace(/\r\n\r\n/g, '<br>') : this.md.render(this.article.body)
      md = md.replace(/img src/g, 'img lazy="loading" src="data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7" data-src')
      return md
    }
  },
  watch: {
    article () {
      setTimeout(() => {
        document.querySelectorAll('[lazy]').forEach(img => {
          this.lazyImages.push(img)
        })
        console.error(this.lazyImages.length)
        if (this.lazyImages.length) {
          window.addEventListener('scroll', this.lazyArticleImage)
          this.lazyArticleImage()
        }
      }, 500)
    }
  },
  mounted () {
    this.findAlias(this.$route.params.alias)
    this.md = new Markdown()
    this.md.use(markdownItHighlight)
    this.md.use(Emoji)
  },
  destroyed () {
    window.removeEventListener('scroll', this.lazyArticleImage)
  },
  methods: {
    ...mapActions(['findAlias']),
    tagArticles (tagId) {
      for (let i = 0; i < this.tags.length; i++) {
        const tag = this.tags[i]
        if (tag._id === tagId) {
          return tag.articles
        }
      }
    },
    // markdown 変換後に v-lazy が動いてくれないので自前で実装
    lazyArticleImage () {
      if (!this.lazyImages.length) {
        window.removeEventListener('scroll', this.lazyArticleImage)
      }
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      const height = window.innerHeight
      const position = scrollTop + height + 50
      for (let i = 0; i < this.lazyImages.length; i++) {
        const img = this.lazyImages[i]
        if (img.offsetTop > position) {
          continue
        }

        img.src = img.getAttribute('data-src')
        this.lazyImages.splice(i, 1)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.tag {
  cursor: pointer;
}
.tag.is-light {
  &:hover {
    background-color:#ccc;
  }
}
.columns.body {
  min-height: 84%;
  background-color: #fff;
}
.container.article {
  height: 100%;
}
#body {
  border-bottom-left-radius: 5px;
  padding-bottom: 60px;
  margin-bottom: 10px;
}
</style>
