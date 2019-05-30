importScripts('/_nuxt/workbox.4c4f5ca6.js')

workbox.precaching.precacheAndRoute([
  {
    "url": "/_nuxt/1ee022d689b45df95842.js",
    "revision": "d9aed54c8fb695555069105b49cf48b8"
  },
  {
    "url": "/_nuxt/4b92555c685b24ac4885.js",
    "revision": "f9eb73ab53f6941279592203179d599f"
  },
  {
    "url": "/_nuxt/57c43d2615438f0af440.js",
    "revision": "c2f8316d553c85b928ac39c67d2a305f"
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
    "url": "/_nuxt/670cbfe4ad5060b769cf.js",
    "revision": "09589a0dbf6ebacaa08c4e78e27fe2f0"
  },
  {
    "url": "/_nuxt/6a74dd6d8762d28bedbc.js",
    "revision": "4e70cd62099b9a85a0952b81d67de504"
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
    "url": "/_nuxt/772aa63e7e5997507295.js",
    "revision": "5c20a749aa0c11828eeac56af728dac7"
  },
  {
    "url": "/_nuxt/88c9ab9b94c45cee764a.js",
    "revision": "0129bef9ca6e8d5e4fc5dfbd05e73084"
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
    "url": "/_nuxt/cf48c446953caf0e0943.js",
    "revision": "44c67c807bf1861cbc376ac6c9da5b26"
  },
  {
    "url": "/_nuxt/dde7cb00b5baf76b08cb.js",
    "revision": "d400b215d3913bf6b4aa82d202384224"
  },
  {
    "url": "/_nuxt/e239c8e4657910f712d7.js",
    "revision": "26e2a85674cdccb6bf7bb8e120a979c9"
  },
  {
    "url": "/_nuxt/e5dd184d08d3d2f33273.js",
    "revision": "06de25f100bc60e2c341616055d0daa4"
  },
  {
    "url": "/_nuxt/e9c42f774717a8b876ae.js",
    "revision": "10cef785085c046611658928028ca189"
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
