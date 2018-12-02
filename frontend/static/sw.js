importScripts('/_nuxt/workbox.dev.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/02e8992c04196e758682.js",
    "revision": "40d5972e2a98d40c9a5438cb763f4305"
  },
  {
    "url": "/_nuxt/084e42e2a532b51b24df.js",
    "revision": "27ecfef269f924de19823ecf9910cbcf"
  },
  {
    "url": "/_nuxt/182c694c6fa67aeb5f8e.js",
    "revision": "3ad5ae58fb3b42d389ab4e506fb477e9"
  },
  {
    "url": "/_nuxt/1841f9515727cac5bd1e.js",
    "revision": "4b823406cf7b21d9bb5e7f4ae18a7575"
  },
  {
    "url": "/_nuxt/205a84afb3588b24218a.js",
    "revision": "0f428885aa7b7a20b7db914dbb74f79c"
  },
  {
    "url": "/_nuxt/242ffada911700eb37df.js",
    "revision": "40198d592323de43cc775707c55ae7ce"
  },
  {
    "url": "/_nuxt/2997ed9e7b21c8212374.js",
    "revision": "a2fedc77ade0aa5f0b5ef382d99ec323"
  },
  {
    "url": "/_nuxt/35420c15ebcf19b29c51.js",
    "revision": "7ab00c055d82275bc047bace8b61a55d"
  },
  {
    "url": "/_nuxt/505ae3c1c545db07c307.js",
    "revision": "b64c62b839f835b8f5fffbeb3f0501d3"
  },
  {
    "url": "/_nuxt/5459f85fe762d291e154.js",
    "revision": "a3833ff44959cd8e3e5e05e15840435d"
  },
  {
    "url": "/_nuxt/69acc3be7fee93a3dbf3.js",
    "revision": "0764727302f321af71049aa9fde05608"
  },
  {
    "url": "/_nuxt/73bb6cff3c1033ca85b1.js",
    "revision": "f2ebe2fa9f407fecb0f03860a5a86bfe"
  },
  {
    "url": "/_nuxt/7a3717a8454b4cb09dde.js",
    "revision": "0478e19fb112fab68f34a3eaa2f9b5d2"
  },
  {
    "url": "/_nuxt/82e2bbe6a7839ed0d6ba.js",
    "revision": "3ca7310357b03007848bd048a6dc34af"
  },
  {
    "url": "/_nuxt/8b54690d1fa4eeec37e1.js",
    "revision": "11a0edbfc78cc0dac315ed5142ddfc49"
  },
  {
    "url": "/_nuxt/9815900ecaf5b2e987e3.js",
    "revision": "510a549874a73aecb0ebe9db5a90d1f3"
  },
  {
    "url": "/_nuxt/ad0572f1f5136208f96e.js",
    "revision": "06f6219b37619539d86914f3273eb9f8"
  },
  {
    "url": "/_nuxt/ae88df5b5fab60b1b274.js",
    "revision": "93bbdd9bd83ffaa87e98198a9f4d7988"
  },
  {
    "url": "/_nuxt/fd63c3ce10ee6210b114.js",
    "revision": "22a487e54f07b6a31e443e73ef0b1194"
  }
], {
  "cacheId": "blog.temo.xyz",
  "directoryIndex": "/",
  "cleanUrls": false
})

workbox.clientsClaim()
workbox.skipWaiting()

workbox.routing.registerRoute(new RegExp('/_nuxt/.*'), workbox.strategies.cacheFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('/.*'), workbox.strategies.networkFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('https://fonts.gstatic.com/*'), workbox.strategies.cacheFirst({}), 'GET')
