<template>
<article v-if="message.title" class="message" :class="colorCss">
  <div class="message-header">
    <p>{{ message.title }}</p>
    <button @click="messageOff" class="delete" aria-label="delete"></button>
  </div>
  <div class="message-body">
    {{ message.body }}
  </div>
</article>
</template>
<script>
import { mapGetters, mapMutations } from 'vuex'
export default {
  data () {
    return {
      colorCss: ''
    }
  },
  watch: {
    'message.title' () {
      this.setColor()
    }
  },
  computed: {
    ...mapGetters(['message'])
  },
  methods: {
    ...mapMutations(['messageOff']),
    setColor () {
      if (!this.message.title) {
        return
      }

      if (this.message.type === 'notify') {
        // slide in
        this.colorCss = 'is-' + this.message.color + ' notify slidein'
        // 5秒後 slide out を開始し
        setTimeout(() => {
          this.colorCss = 'is-' + this.message.color + ' notify slideout'
          // 0.5秒後(slide out 完了)に message を初期化
          setTimeout(() => {
            this.messageOff()
          }, 500)
        }, 5000)
        return
      }
      this.colorCss = 'is-' + this.message.color
    }
  }
}
</script>
<style lang="scss" scoped>
.message {
  position: fixed;
  top: 20%;
  left: 20%;
  width: 60%;
  height: 60%;
  z-index: 100;
  box-shadow: 0 0 14px 2px rgba(0, 0, 0, 0.2);
  &.notify {
    top: 2%;
    left: auto;
    width: 300px;
    height: auto;
    .delete {
      display: none;
    }
  }
}
.slidein {
  right: 5px;
  animation-duration: 0.5s;
  animation-name: slidein;
}
.slideout {
  right: -350px;
  animation-duration: 0.5s;
  animation-name: slideout;
}
@media (max-width: 1087px) {
  .message {
    top: 4%;
    left: 4%;
    width: 92%;
    height: 92%;
  }
}
@keyframes slidein {
  0% {
    display: block;
    right: -350px;
  }
  1% {
    display: block;
    right: -350px;
  }
  100% {
    display: block;
    right: 5px;
  }
}
@keyframes slideout {
  0% {
    display: block;
    right: 5px;
  }
  100% {
    display: none;
    right: -350px;
  }
}
</style>
