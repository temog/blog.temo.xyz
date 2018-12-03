import Util from '@/plugins/util'
export default function ({ store }) {
  store.commit('setTitle', document.title)
  Util.init(store)
  store.dispatch('getTags')
  store.dispatch('getPopularArticles')
}
