<template>
  <section class="container">
    <div class="tags has-addons is-marginless">
      <span class="tag is-white is-medium">Tag</span>
      <span class="tag is-dark is-medium">{{ tagName }}</span>
    </div>

    <articles-component type="tag" :keyword="keyword" />
  </section>
</template>

<script>
import { mapGetters } from 'vuex'
import ArticlesComponent from '@/components/Articles'
export default {
  components: {
    ArticlesComponent
  },
  computed: {
    ...mapGetters(['tags']),
    tagName () {
      return this.$route.params.tagName
    },
    keyword () {
      for (let i = 0; i < this.tags.length; i++) {
        if (this.tags[i].name === this.tagName) {
          return this.tags[i]._id
        }
      }
    }
  }
}
</script>
