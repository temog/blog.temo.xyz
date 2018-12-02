import axios from 'axios'

class Util {
  static init (store) {
    console.log('util init')
    this.store = store
    axios.defaults.headers.common.Authorization = this.store.state.token
  }

  static overwriteToken (token) {
    axios.defaults.headers.common.Authorization = token
  }

  static test () {
    console.log('aaaaaaaaa')
    console.log(this.store)
  }

  /* api start */
  static async api (url, data) {
    url = process.env.apiBaseUrl + url
    const result = await this.post(url, data)
    return result
  }

  static async apiGet (url, data) {
    url = process.env.apiBaseUrl + url
    const result = await this.get(url, data)
    return result
  }
  /* api end */

  static async post (url, data) {
    this.store.commit('startLoading')
    try {
      const resp = await axios.post(url, data)
      console.warn(resp)
      this.store.commit('endLoading')
      return this.validateResponse(resp)
    } catch (e) {
      this.store.commit('endLoading')
      console.log('error catch................')
      console.log(e.response)
      if (e.response) {
        let message = e.response.status + ' ' + e.response.statusText
        if (e.response.data.message) {
          message += '<br>' + e.response.data.message
          if (e.response.data.message === 'invalid token') {
            this.store.commit('resetToken')
          }
        }
        this.errorModal(message)
      } else if (e.request) {
        this.errorModal(e.request)
      } else {
        this.errorModal(e.message)
      }
      return false
    }
  }

  static async get (url, data) {
    this.store.commit('startLoading')
    try {
      const resp = await axios.get(url, { params: data })
      console.warn(resp)
      this.store.commit('endLoading')
      return this.validateResponse(resp)
    } catch (e) {
      this.store.commit('endLoading')
      console.log('error catch................')
      console.log(e.response)
      if (e.response) {
        let message = e.response.status + ' ' + e.response.statusText
        if (e.response.data.message) {
          message += '<br>' + e.response.data.message
          if (e.response.data.message === 'invalid token') {
            this.store.commit('resetToken')
          }
        }
        this.errorModal(message)
      } else if (e.request) {
        this.errorModal(e.request)
      } else {
        this.errorModal(e.message)
      }
      return false
    }
  }

  static validateResponse (resp) {
    console.log('validation response')
    console.log(resp)
    if (!resp.data.status) {
      let message = resp.data.error ? resp.data.error : ''
      message += resp.data.message ? resp.data.message : ''
      this.errorModal(message)
      return false
    }
    return resp.data
  }

  static errorModal (error) {
    const data = {
      title: 'エラー'
    }
    data.body = error
    this.store.commit('modalOn', data)
  }

  static lang (code, param) {
    return code + param
  }

  static message (type, color, title, body) {
    this.store.commit('messageOn', {
      type: type,
      color: color,
      title: title,
      body: body
    })
  }

  static storageGet (key) {
    return localStorage.getItem(key) ? JSON.parse(localStorage.getItem(key)) : null
  }

  static eyecatchUrl (id, date) {
    const timestamp = +new Date(date)
    return process.env.imageBaseUrl + 'eyecatch/' + id + '.png?updated=' + timestamp
  }

  static date (dateObj) {
    return dateObj.getFullYear() + '-' +
    (dateObj.getMonth() + 1) + '-' +
    dateObj.getDate() + ' ' +
    dateObj.getHours() + ':' + dateObj.getMinutes()
  }

  // todo
  static articleLink (article) {
    const tagName = this.store.getters.tagName(article.tags[0])
    // const tagName = this.tagName(article.tags[0])
    return '/archives/' + tagName + '/' + article.alias
  }
}
export default Util
