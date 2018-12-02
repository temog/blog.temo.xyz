<template>
  <section class="container">
    <div class="tags has-addons is-marginless">
      <span class="tag is-white is-medium">Search</span>
      <span class="tag is-dark is-medium">{{ keyword }}</span>
    </div>

    <articles-component type="latest" :keyword="keyword" />
  </section>
</template>

<script>
import ArticlesComponent from '@/components/Articles'
export default {
  components: {
    ArticlesComponent
  },
  computed: {
    keyword () {
      return this.$route.params.keyword
    }
  }
}
</script>
