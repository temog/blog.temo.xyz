<template>
  <div id="indexes" class="column index is-one-fifth" :class="addClass">
    <div class="inner" :class="{fixed: fixed}">
      <p v-for="(i, key) in indexes" :key="key" @click="jumpArticle(i.index)" :class="i.class">{{ i.label }}</p>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import jump from 'jump.js'
export default {
  props: ['id', 'body'],
  data () {
    return {
      indexes: [],
      domIndexes: null,
      indexOffsetTop: 0,
      fixed: false
    }
  },
  watch: {
    body () {
      console.log('watch body')
      this.createIndex()
    }
  },
  computed: {
    ...mapGetters(['showIndex']),
    addClass () {
      let css = ''
      /*
      if (this.fixed) {
        css += 'fixed '
      }
      */
      if (this.showIndex) {
        css += 'active '
      }
      console.warn(css)
      return css
    }
  },
  mounted () {
    console.log('article index mounted')

    this.createIndex()

    this.$nextTick(() => {
      this.domIndexes = document.querySelector('#indexes')
      setTimeout(() => {
        this.indexOffsetTop = this.domIndexes.offsetTop
        window.addEventListener('scroll', this.stickyIndexes)
        window.addEventListener('resize', this.setIndexOffsetTop)
      }, 1000)
    })
  },
  destroyed () {
    window.removeEventListener('scroll', this.stickyIndexes)
  },
  methods: {
    ...mapMutations(['toggleIndex']),
    createIndex () {
      this.indexes = []
      let i = 1
      document.querySelectorAll('#' + this.id + ' h1,h2,h3').forEach(index => {
        console.log(index.tagName)
        index.id = 'article-index-' + i
        this.indexes.push({
          label: index.textContent,
          index: index.id,
          class: index.tagName.toLowerCase()
        })
        i++
      })
    },
    jumpArticle (index) {
      jump('#' + index, {
        offset: -60,
        duration: 500
      })
      this.toggleIndex()
    },
    stickyIndexes () {
      if (document.documentElement.scrollTop >= this.indexOffsetTop) {
        this.fixed = true
      } else {
        this.fixed = false
      }
    },
    setIndexOffsetTop () {
      this.indexOffsetTop = this.domIndexes.offsetTop
    }
  }
}
</script>
