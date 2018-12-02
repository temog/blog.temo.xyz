<template>
  <section class="container admin">

    <div class="columns">
      <div class="column is-one-fifth">
        <user-menu />
      </div>

      <div class="column">
        <h1 class="title">All Articles</h1>

        <div class="field is-grouped is-grouped-multiline">
          <div class="control">
            <div @click="$router.push('/user/article/latest/_/desc/1')" class="tags has-addons">
              <span class="tag is-white" :class="activeLatest">Latest</span>
              <span class="tag is-dark"><fa icon="search" /></span>
            </div>
          </div>

          <div class="control">
            <div class="tags has-addons">
              <span class="tag is-white" :class="activeTag">Tag</span>
              <span class="tag is-white">
                <div class="select no-border">
                  <select v-model="search.tag">
                    <option v-for="(t, key) in tags" :key="key"
                      :value="t._id">
                      {{ t.name }}
                    </option>
                  </select>
                </div>
              </span>
              <span @click="tagSearch(search.tag)" class="tag is-dark"><fa icon="search" /></span>
            </div>
          </div>

          <div class="control">
            <div class="tags has-addons">
              <span class="tag is-white">Limit</span>
              <span class="tag is-white">
                <div class="select no-border">
                  <select v-model="searchLimit">
                    <option value="1">1</option>
                    <option value="2">2</option>
                    <option value="20">20</option>
                    <option value="50">50</option>
                  </select>
                </div>
              </span>
            </div>
          </div>

        </div>

        <table class="table is-striped">
          <thead>
            <tr>
              <th>Title</th>
              <th>Status</th>
              <th>Date</th>
              <th>Tags</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(a, key) in userArticles" :key="key">
              <td>
                {{ a.title }}
                <br>
                <router-link :to="'/user/article/edit/' + a._id" class="button is-link is-small">Edit</router-link>
              </td>
              <td><span class="tag" :class="statusCss(a.status)">{{ a.status }}</span></td>
              <td><div class="date">{{ a.date | date }}</div></td>
              <td>
                <div class="field is-grouped is-grouped-multiline">
                  <div v-for="(tag, key) in a.tags" :key="key" @click="tagSearch(tag)" class="control">
                    <tag-component :tagId="tag" />
                  </div>
                </div>
              </td>
            </tr>
          </tbody>
        </table>

        <nav v-if="userPage" class="pagination" role="navigation" aria-label="pagination">
          <a @click="push('prev')" class="pagination-previous" :disabled="page === 1">前のページ</a>
          <a @click="push('next')" class="pagination-next" :disabled="page >= userPage">次のページ</a>
          <ul class="pagination-list">
            <li v-for="p in userPage" :key="p">
              <a @click="push(p)" class="pagination-link" :class="{'is-current': page === p}">{{ p }}</a>
            </li>
          </ul>
        </nav>

      </div>
    </div>

  </section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import UserMenu from '@/components/UserMenu'
import TagComponent from '@/components/Tag'
export default {
  components: {
    UserMenu,
    TagComponent
  },
  data () {
    return {
      search: {
        tag: null,
        limit: '1'
      }
    }
  },
  computed: {
    ...mapGetters(['userPage', 'userArticles', 'tags']),
    activeLatest () {
      if (this.$route.params.type === 'latest') {
        return 'is-link'
      }
    },
    activeTag () {
      if (this.$route.params.type === 'tag') {
        return 'is-link'
      }
    },
    searchLimit: {
      get () {
        return this.search.limit
      },
      set (value) {
        this.search.limit = value
        localStorage.setItem('searchLimit', value)
      }
    },
    page () {
      return parseInt(this.$route.params.page)
    }
  },
  mounted () {
    const searchLimit = localStorage.getItem('searchLimit')
    this.search.limit = searchLimit || '1'

    const p = this.$route.params
    const param = {
      type: p.type, // tag
      status: 'all',
      keyword: p.keyword,
      created: p.created,
      page: p.page,
      limit: this.search.limit
    }

    if (p.type === 'tag') {
      this.search.tag = p.keyword
    }

    this.searchArticleAdmin(param)
  },
  methods: {
    ...mapActions(['searchArticleAdmin']),
    statusCss (status) {
      if (status === 'private') {
        return 'is-dark'
      }
      return 'is-success'
    },
    tagSearch (tagId) {
      this.$router.push('/user/article/tag/' + tagId + '/desc/1')
    },
    push (page) {
      const p = this.$route.params
      if (page === 'prev') {
        page = parseInt(p.page) - 1
      } else if (page === 'next') {
        page = parseInt(p.page) + 1
      }

      if (p.type === 'latest') {
        this.$router.push('/user/article/latest/_/' + p.created + '/' + page)
      } else if (p.type === 'tag') {
        this.$router.push('/user/article/tag/' + p.keyword + '/' + p.created + '/' + page)
      }
      /*
      const param = {
        type: p.type, // tag
        keyword: p.keyword,
        created: p.created,
        page: p.page,
        limit: this.search.limit
      }
      */
    }
  }
}
</script>

<style lang="scss" scoped>
.tag.is-white {
  overflow: hidden;
}
.no-border {
  select {
    border: 0;
  }
}
</style>
