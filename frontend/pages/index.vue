<template>
  <section class="container">
    <articles-component type="latest" />
  </section>
</template>

<script>
import ArticlesComponent from '@/components/Articles'
export default {
  components: {
    ArticlesComponent
  }
}
</script>
