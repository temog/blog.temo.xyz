<template>
  <div>
    <div class="grid">
      <div v-for="(article, key) in articles" :key="key" class="item">
        <div @click="$router.push(articleLink(article))" class="wrap"></div>
        <div v-lazy:background-image="eyecatch(article)" class="image" />
        <div class="title">{{ article.title }}</div>
        <div class="tags">
          <div v-for="(tag, key) in article.tags" :key="key" @click="$router.push('/tag/' + tagName(tag))" class="tag is-dark">
            {{ tagName(tag) }}
          </div>
        </div>
        <div class="date">{{ article.date | date }}</div>
      </div>
    </div>

    <div v-if="articlesError" class="notification is-warning">
      記事がありません
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import Util from '@/plugins/util'
export default {
  props: ['type', 'keyword'],
  computed: {
    ...mapGetters(['loading', 'articles', 'articlesPage', 'tagName', 'articlesError'])
  },
  mounted () {
    console.warn(this.type)
    console.warn(this.keyword)

    if (this.type === 'latest') {
      this.search.keyword = this.keyword ? this.keyword : '_'
    } else if (this.type === 'tag') {
      this.search.keyword = this.keyword
    }

    this.searchArticle(this.search)
    window.addEventListener('scroll', this.watchScrollEnd)
    this.$nextTick(() => {
      this.watchScrollEnd()
    })
  },
  destroyed () {
    window.removeEventListener('scroll', this.watchScrollEnd)
  },
  data () {
    return {
      timer: null,
      search: {
        type: this.type,
        status: 'publish',
        keyword: '_',
        created: 'desc',
        page: '1',
        limit: '20'
      }
    }
  },
  methods: {
    ...mapActions(['searchArticle']),
    articleLink (article) {
      return Util.articleLink(article)
    },
    eyecatch (article) {
      return Util.eyecatchUrl(article._id, article.updated_at)
    },
    watchScrollEnd () {
      clearTimeout(this.timer)
      this.timer = setTimeout(() => {
        this.autoSearch()
      }, 500)
    },
    autoSearch () {
      if (this.loading || this.search.page >= this.articlesPage) {
        return
      }
      const innerHeight = window.innerHeight
      const height = document.documentElement.scrollHeight - 400
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop
      const current = scrollTop + innerHeight
      if (current < height) {
        return
      }
      console.log('auto search')
      this.search.page = String(parseInt(this.search.page) + 1)
      this.searchArticle(this.search)
    }
  }
}
</script>
