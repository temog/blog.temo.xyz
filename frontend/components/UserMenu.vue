<template>
  <aside class="menu">

    <p class="menu-label">
      Article
    </p>
    <ul class="menu-list">
      <li>
        <router-link to="/user/article/latest/_/desc/1"
        :class="isActive(['user-article-type-keyword-created-page','user-article-edit-articleId'])">All Articles</router-link>
      </li>
      <li><router-link to="/user/article/add" :class="isActive(['user-article-add'])">Create</router-link></li>
    </ul>

    <p class="menu-label">
      Tag
    </p>
    <ul class="menu-list">
      <li><router-link to="/user/tag" :class="isActive(['user-tag'])">All Tags</router-link></li>
      <li><router-link to="/user/tag/add" :class="isActive(['user-tag-add'])">Create</router-link></li>
    </ul>

    <p class="menu-label">
      Other
    </p>
    <ul class="menu-list">
      <li><a @click="resetToken">Sign out</a></li>
    </ul>

  </aside>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
export default {
  created () {
    if (!this.isSignIn) {
      this.$router.push({ name: 'user-signin' })
    }
  },
  mounted () {
  },
  watch: {
    isSignIn (value) {
      if (!value) {
        this.$router.push({ name: 'user-signin' })
      }
    }
  },
  computed: {
    ...mapGetters(['isSignIn'])
  },
  methods: {
    ...mapMutations(['resetToken']),
    isActive (routeNames) {
      for (let i = 0; i < routeNames.length; i++) {
        const name = routeNames[i]

        if (this.$route.name === name) {
          return 'is-active'
        }
      }
    }
  }
}
</script>
