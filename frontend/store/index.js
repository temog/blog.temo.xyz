import Vuex from 'vuex'
import Util from '@/plugins/util'

const createStore = () => {
  return new Vuex.Store({
    state: () => ({
      counter: 0,
      token: localStorage.getItem('token'),
      loading: false,
      showIndex: false,
      vimMode: localStorage.getItem('vimMode') === 'true',
      modal: {
        title: null,
        body: null
      },
      message: {
        type: 'notify',
        color: 'dark',
        title: null,
        body: null
      },
      tags: JSON.parse(localStorage.getItem('tags')),
      previewArticle: {
        id: Util.storageGet('preview.id'),
        title: Util.storageGet('preview.title'),
        alias: Util.storageGet('preview.alias'),
        date: Util.storageGet('preview.date'),
        status: Util.storageGet('preview.status'),
        tags: Util.storageGet('preview.tags') ? Util.storageGet('preview.tags') : [],
        body: Util.storageGet('preview.body'),
        eyecatch: Util.storageGet('preview.eyecatch'),
        eyecatchUpdate: Util.storageGet('preview.eyecatchUpdate')
      },
      previewScroll: 0,
      previewError: {
        title: null,
        alias: null,
        date: null,
        tags: null,
        body: null,
        eyecatch: null
      },
      image: {
        edit: false,
        tag: null
      },
      user: {
        articles: [],
        page: 0
      },
      /* トップページ用 */
      articles: [],
      articlesPage: 0,
      articlesError: null,
      /* 記事ページ */
      article: null,
      /* popular articles */
      popularArticles: []
    }),
    actions: {
      async signIn ({ commit }, data) {
        const result = await Util.api('user/auth', data)
        if (result === false) {
          return
        }
        commit('setToken', result)
      },
      async addTag ({ commit }, data) {
        const result = await Util.api('tag/add', data)
        if (result === false) {
          return
        }
        commit('messageOn', {
          color: 'success',
          title: 'Success',
          body: '追加しました'
        })
      },
      async getTags ({ commit }) {
        const result = await Util.apiGet('tag/getAll')
        if (result === false) {
          return
        }
        commit('setTags', result)
      },
      async getPopularArticles ({ commit }) {
        const result = await Util.apiGet('article/getPopularArticles')
        if (result === false) {
          return
        }
        console.warn(result)
        commit('setPopularArticles', result)
      },
      async updateTag ({ commit }, data) {
        const result = await Util.api('tag/update', data)
        if (result === false) {
          return
        }
        commit('messageOn', { color: 'success', title: 'Success', body: 'タグ更新' })
      },
      async addImage ({ commit }, data) {
        const result = await Util.api('image/add', data)
        if (result === false) {
          return
        }
        console.log(result)
        commit('setImageEdit', false)

        const url = process.env.imageBaseUrl + result.fileName
        commit('setImageTag', `![${result.fileName}](${url})`)
      },
      async addArticle ({ commit }, data) {
        const result = await Util.api('article/add', data)
        if (result === false) {
          return
        }
        commit('messageOn', { color: 'success', title: 'Success', body: '記事を投稿しました' })
      },
      async editArticle ({ commit }, data) {
        const result = await Util.api('article/edit', data)
        if (result === false) {
          return
        }
        commit('messageOn', { color: 'success', title: 'Success', body: '記事を編集しました' })
      },
      async searchArticle ({ commit }, data) {
        const result = await Util.api('article/search', data)
        if (result === false) {
          return
        }
        result.page = data.page
        commit('setArticles', result)
      },
      async searchArticleAdmin ({ commit }, data) {
        const result = await Util.api('article/search', data)
        if (result === false) {
          return
        }
        commit('setUserArticles', result)
      },
      async findAlias ({ commit }, alias) {
        const result = await Util.apiGet('article/alias/' + alias)
        if (result === false) {
          return
        }
        console.log(result)
        commit('setArticle', result.article)
      },
      async findId ({ commit }, id) {
        const result = await Util.apiGet('article/find/' + id)
        if (result === false) {
          return
        }
        console.log(result)

        const a = result.article

        const keys = ['title', 'alias', 'status', 'date', 'tags', 'body', 'eyecatch', 'eyecatchUpdate', 'updated_at']

        // 編集中の記事があったらそっちを使う
        const savedId = Util.storageGet('preview.id')
        if (savedId === id) {
          const latestEdited = new Date(Util.storageGet('preview.latestEdited'))
          const updatedAt = new Date(a.updated_at)
          if (latestEdited > updatedAt) {
            console.error('local storage を使う')
            keys.forEach(k => {
              commit('setPreviewArticle', { key: k, value: Util.storageGet('preview.' + k) })
            })
            commit('setPreviewArticle', { key: 'id', value: id })
            return
          }
        }
        console.error('local storage 使わない！')

        keys.forEach(k => {
          commit('setPreviewArticle', { key: k, value: a[k] ? a[k] : null })
        })

        commit('setPreviewArticle', { key: 'id', value: id })
      }
    },
    mutations: {
      startLoading (state) {
        state.loading = true
      },
      endLoading (state) {
        state.loading = false
      },
      toggleIndex (state) {
        if (state.showIndex) {
          state.showIndex = false
        } else {
          state.showIndex = true
        }
      },
      setToken (state, data) {
        state.token = data.token
        localStorage.setItem('token', state.token)
        Util.overwriteToken(state.token)
      },
      resetToken (state) {
        state.token = null
        localStorage.removeItem('token')
      },
      setTags (state, data) {
        state.tags = data.tags
        localStorage.setItem('tags', JSON.stringify(state.tags))
      },
      setPopularArticles (state, data) {
        state.popularArticles = data.articles
      },
      modalOn (state, data) {
        state.modal.title = data.title
        state.modal.body = data.body
      },
      modalOff (state) {
        state.modal.title = null
        state.modal.body = null
      },
      messageOn (state, data) {
        state.message.type = data.type ? data.type : 'notify'
        state.message.color = data.color
        state.message.title = data.title
        state.message.body = data.body
      },
      messageOff (state) {
        state.message.type = 'notify'
        state.message.color = 'dark'
        state.message.title = null
        state.message.body = null
      },
      setPreviewArticle (state, data) {
        state.previewArticle.latestEdited = new Date()
        state.previewArticle[data.key] = data.value
        localStorage.setItem('preview.' + data.key, JSON.stringify(data.value))
        localStorage.setItem('preview.latestEdited', JSON.stringify(state.previewArticle.latestEdited))
      },
      setPreviewScroll (state, percent) {
        state.previewScroll = percent
        localStorage.setItem('previewScroll', percent)
      },
      setPreviewError (state, error = {}) {
        for (const key in state.previewError) {
          state.previewError[key] = error.hasOwnProperty(key) ? error[key] : null
        }
      },
      resetPreviewArticle (state) {
        for (const key in state.previewArticle) {
          if (typeof state.previewArticle[key] === 'object') {
            state.previewArticle[key] = []
          } else {
            state.previewArticle[key] = null
          }
          localStorage.setItem('preview.' + key, '')
        }
      },
      setVimMode (state, bool) {
        state.vimMode = bool
        localStorage.setItem('vimMode', bool)
      },
      setImageEdit (state, bool) {
        state.image.edit = bool
      },
      setImageTag (state, tag) {
        state.image.tag = tag
      },
      setArticle (state, article) {
        state.article = article
      },
      setArticles (state, data) {
        if (data.page === '1') {
          state.articles = data.articles
        } else {
          state.articles = state.articles.concat(data.articles)
        }
        state.articlesPage = data.totalPage

        if (data.totalPage === 0) {
          state.articlesError = true
        } else {
          state.articlesError = false
        }
      },
      setUserArticles (state, data) {
        state.user.articles = data.articles
        state.user.page = data.totalPage
      }
    },
    getters: {
      isSignIn: state => state.token !== null,
      loading: state => state.loading,
      showIndex: state => state.showIndex,
      modal: state => state.modal,
      message: state => state.message,
      tags: state => state.tags,
      tagName: state => (id) => {
        for (let i = 0; i < state.tags.length; i++) {
          if (state.tags[i]._id === id) {
            return state.tags[i].name
          }
        }
      },
      previewArticle: state => state.previewArticle,
      popularArticles: state => state.popularArticles,
      previewEyecatch: state => state.previewArticle.eyecatch,
      previewScroll: state => state.previewScroll,
      previewError: state => state.previewError,
      vimMode: state => state.vimMode,
      image: state => state.image,
      userArticles: state => state.user.articles,
      userPage: state => state.user.page,
      articles: state => state.articles,
      articlesPage: state => state.articlesPage,
      articlesError: state => state.articlesError,
      article: state => state.article
    }
  })
}

export default createStore
