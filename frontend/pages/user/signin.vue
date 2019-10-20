<template>
  <section class="container">
    <div class="field">
      <label class="label">Account</label>
      <div class="control">
        <input v-model="account" class="input" type="text">
      </div>
    </div>

    <div class="field">
      <label class="label">Password</label>
      <div class="control">
        <input v-model="password" class="input" type="password">
      </div>
    </div>

    <div class="field is-grouped">
      <div class="control">
        <button @click="signIn({account: account, password: password})" class="button is-link">Sign in</button>
      </div>
    </div>
  </section>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
export default {
  data () {
    return {
      account: null,
      password: null
    }
  },
  created () {
    if (this.isSignIn) {
      this.$router.push('/user/article/latest/_/desc/1')
    }
  },
  mounted () {
  },
  watch: {
    isSignIn (value) {
      if (value) {
        this.$router.push('/user/article/latest/_/desc/1')
      }
    }
  },
  computed: {
    ...mapGetters(['isSignIn'])
  },
  methods: {
    ...mapActions(['signIn'])
  }
}
</script>
