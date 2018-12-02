<template>
  <section class="container">

    <user-menu />

  </section>
</template>

<script>
// import { mapGetters, mapActions, mapMutations } from 'vuex'
import UserMenu from '@/components/UserMenu'
export default {
  components: {
    UserMenu
  },
  data () {
    return {
      account: null,
      password: null
    }
  },
  computed: {
  },
  methods: {
  }
}
</script>
