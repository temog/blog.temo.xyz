import Util from '@/plugins/util'
export default function ({ store }) {
  Util.init(store)
  store.dispatch('getTags')
  store.dispatch('getPopularArticles')
}
