importScripts('/_nuxt/workbox.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/09fc4e92ac9b397eeee4.js",
    "revision": "e6921d726c0c9c0d1a5fdb28ecc01d30"
  },
  {
    "url": "/_nuxt/386bfcd2e0f174c258c1.js",
    "revision": "b5cc011e6273c7cfba24a696002e0dff"
  },
  {
    "url": "/_nuxt/4ad53c7721325003520f.js",
    "revision": "052d37ea13d78aca54957db46c3549c5"
  },
  {
    "url": "/_nuxt/5a1c85b8805d7d51620b.js",
    "revision": "537cf11e8335fdae0039352524610688"
  },
  {
    "url": "/_nuxt/5e4dbd5c45836ac81aa5.js",
    "revision": "349d336655cbebe926755826b4d22bc1"
  },
  {
    "url": "/_nuxt/5ea8ef74a598d6b3ff9a.js",
    "revision": "74f310c2e53f2620d0637599b5b25cd4"
  },
  {
    "url": "/_nuxt/606a1c9c323fd9fd542b.js",
    "revision": "40a4e7fac9287e52849be21abf773a33"
  },
  {
    "url": "/_nuxt/66fb60f94495eaadceeb.js",
    "revision": "87904b248550aa90ffc33e6922c7d789"
  },
  {
    "url": "/_nuxt/6b3145d944986ec474c1.js",
    "revision": "cc580390c7b0aa41729f2d9ab2890e9d"
  },
  {
    "url": "/_nuxt/6fd99a1ccc0d8c410e80.js",
    "revision": "acfe95a00285700ff796ebe608f36e59"
  },
  {
    "url": "/_nuxt/70eb1d2411dc324ebef3.js",
    "revision": "0ac9efb73907d1fa34e163cb7b537d8a"
  },
  {
    "url": "/_nuxt/7c8f875e13981ab4640e.js",
    "revision": "f3bd3c7c51a3044ebe963f96080ffbd1"
  },
  {
    "url": "/_nuxt/88c9ab9b94c45cee764a.js",
    "revision": "0129bef9ca6e8d5e4fc5dfbd05e73084"
  },
  {
    "url": "/_nuxt/a91bf701674c37f40f58.js",
    "revision": "0e47187e6713f642cc78c08ed9f4e76b"
  },
  {
    "url": "/_nuxt/adfe75e22c54ced9d64e.js",
    "revision": "1d126e73f2a65ea9dce12b94100f4cb9"
  },
  {
    "url": "/_nuxt/c87d699d7ef7bbd529c6.js",
    "revision": "9ec5eddddfedd95ef71dda1935733eef"
  },
  {
    "url": "/_nuxt/dd9c1738c7c4a8a8f47f.js",
    "revision": "4b80f82e87242499927543742a7e7004"
  },
  {
    "url": "/_nuxt/dde7cb00b5baf76b08cb.js",
    "revision": "d400b215d3913bf6b4aa82d202384224"
  },
  {
    "url": "/_nuxt/e9c42f774717a8b876ae.js",
    "revision": "10cef785085c046611658928028ca189"
  }
], {
  "cacheId": "blog",
  "directoryIndex": "/",
  "cleanUrls": false
})

workbox.clientsClaim()
workbox.skipWaiting()

workbox.routing.registerRoute(new RegExp('/_nuxt/.*'), workbox.strategies.cacheFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('/.*'), workbox.strategies.networkFirst({}), 'GET')

workbox.routing.registerRoute(new RegExp('https://fonts.gstatic.com/*'), workbox.strategies.cacheFirst({}), 'GET')
