<template>
  <div class="tags has-addons">
    <span class="tag is-light is-rounded" :class="addClass">{{ tagName }}</span>
    <span class="tag is-gray is-rounded" :class="addClass">{{ tagArticles }}</span>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  props: ['tagId', 'tag', 'addClass'],
  computed: {
    ...mapGetters(['tags']),
    tagName () {
      if (this.tag) {
        return this.tag.name
      }

      for (let i = 0; i < this.tags.length; i++) {
        const t = this.tags[i]
        if (t._id === this.tagId) {
          return t.name
        }
      }
    },
    tagArticles () {
      if (this.tag) {
        return this.tag.articles
      }

      for (let i = 0; i < this.tags.length; i++) {
        const t = this.tags[i]
        if (t._id === this.tagId) {
          return t.articles
        }
      }
    }
  }
}
</script>
