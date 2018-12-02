<template>
<div v-if="showVimMode" class="vimMode">

  <button v-if="vimMode" @click="setVimMode(false)"
    class="button is-small is-info">Vim Mode Off</button>

  <button v-else @click="setVimMode(true)"
    class="button is-small">Vim Mode On</button>

  <div v-if="debug" class="debug">
    <h5>vim mode debug</h5>
    使うか: {{ vimMode }}<br>
    mode: {{ mode }}<br>
    clip: {{ vimClip }}<br>
    lastInput: {{ normalLastInput }}<br>
    shiftOn: {{ shiftOn }}<br>
  </div>

</div>
</template>

<script>
import { mapMutations, mapGetters } from 'vuex'

export default {
  props: ['id'],
  data () {
    return {
      debug: false, // デバッグモード
      showVimMode: true,
      text: null,
      mode: 'normal',
      normalLastInput: null,
      vimClip: null, // vimのclipboard
      shiftOn: false,
      lineHeight: 0
    }
  },
  components: {
  },
  computed: {
    ...mapGetters({ storeVimMode: 'vimMode' }),
    vimMode () {
      return this.storeVimMode
    }
  },
  mounted () {
    this.$nextTick(function () {
      this.init()
    })
  },
  methods: {
    ...mapMutations(['setVimMode']),
    init () {
      // 渡されたid の データがない
      const text = document.getElementById(this.id)
      if (!text) {
        return
      }

      const ua = navigator.userAgent
      if (ua.indexOf('iPhone') > 0 || ua.indexOf('iPod') > 0 || (ua.indexOf('Android') > 0 && ua.indexOf('Mobile') > 0)) {
        this.showVimMode = false
        return
      } else if (ua.indexOf('iPad') > 0 || ua.indexOf('Android') > 0) {
        this.showVimMode = false
        return
      }

      this.text = text

      // css line-height
      const win = this.text.ownerDocument.defaultView
      this.lineHeight = win.getComputedStyle(this.text, null).lineHeight
      this.lineHeight = parseInt(this.lineHeight.replace('px', ''))

      // event 設定
      this.setEvent()

      this.showVimMode = true

      // テストデータ
      if (this.debug) {
        this.text.value = '1234567890\n1234567890\n1234567890\n1234567890\n1234567890\n1234567890\n'
      }
    },
    setEvent () {
      const self = this

      // keydown
      this.text.addEventListener('keydown', (e) => {
        // vim mode off はスルー
        if (!self.vimMode) {
          return
        }

        const code = e.keyCode

        if (this.mode === 'normal') {
          // 入力Off
          e.preventDefault()
          self.normalMode(code)
        } else {
          self.insertMode(code)
        }
      })

      // keyup
      this.text.addEventListener('keyup', (e) => {
        // vim mode off はスルー
        if (!self.vimMode) {
          return
        }

        const code = e.keyCode

        // shiftkey
        if (code === 16) {
          this.shiftOn = false
        }
      })
    },
    // ノーマルモードの処理
    normalMode (code) {
      // insert モードへ
      if (this.changeInsertMode(code)) {
        return
      }

      switch (code) {
        // shiftキー
        case 16:
          this.shiftOn = true
          break
        // j or ↓ カーソル移動下
        case 74:
        case 40:
          this.cursorDown()
          break
        // k or ↑ カーソル移動上
        case 75:
        case 38:
          this.cursorUp()
          break
        // l or → カーソル移動右
        case 76:
        case 39:
          this.cursorRight()
          break
        // h or ← カーソル移動左
        case 72:
        case 37:
          this.cursorLeft()
          break
        // x 1文字削除
        case 88:
          this.deleteRight()
          break
        // dd 行削除
        case 68:
          code = this.deleteLine(code)
          break
        // 行コピー
        case 89:
          code = this.lineCopy(code)
          break
        // p 貼り付け
        case 80:
          this.paste()
          break
        // 先頭に移動 or 最後に移動
        case 71:
          code = this.moveTopOrBottom(code)
          break
        // 行末に移動
        case 52:
          code = this.moveEndOfLine(code)
          break
        // 行頭に移動
        case 187:
          this.moveBeginOfLine()
          break
      }

      // 最後にノーマルモードで入力したキー
      this.normalLastInput = code
    },
    // insert mode への移行
    changeInsertMode (code) {
      const insertCodes = [65, 73, 79]
      if (insertCodes.indexOf(code) === -1) {
        return false
      }

      this.mode = 'insert'
      this.normalLastInput = null

      // o だったら改行追加
      if (code === 79) {
        this.insertLine()
      }

      // a だったら一文字移動
      if (code === 65) {
        this.cursorRight(1)
      }

      return true
    },
    insertMode (code) {
      // normal mode へ
      if (code === 27) {
        this.mode = 'normal'
      }
    },
    insertLine () {
      const text = this.text
      const position = text.selectionStart
      const val = text.value

      const line = val.indexOf('\n', position)
      let afterPosition = 0
      if (line !== -1) {
        text.value = val.substr(0, line) + '\n' + val.substr(line)
        afterPosition = line + 1
      } else {
        text.value = val + '\n'
        afterPosition = text.value.lastIndexOf('\n') + 1
      }

      text.setSelectionRange(afterPosition, afterPosition)
    },
    // カーソル位置の後に改行があったらカーソル移動。まじめにやると死ぬので手抜き
    cursorDown () {
      const text = this.text
      const position = text.selectionStart
      const val = text.value

      // 改行があるか探す
      let line = val.indexOf('\n', position)
      if (line === -1) {
        return
      }

      line++
      text.setSelectionRange(line, line)
      this.autoScroll(line)
    },
    // カーソル位置の前に改行があったらカーソル移動。まじめにやると死ぬので手抜き
    cursorUp () {
      const text = this.text
      const position = text.selectionStart
      const val = text.value

      let line = val.lastIndexOf('\n', position - 1)

      if (line === -1) {
        return
      }
      text.setSelectionRange(line, line)
      this.autoScroll(line)
    },
    cursorRight: function (step) {
      const text = this.text
      const position = text.selectionStart
      const val = text.value

      let line = val.substr(position, 1)
      if (line === '\n') {
        return
      }

      const afterPosition = step ? position + step : position + 1
      text.setSelectionRange(afterPosition, afterPosition)
    },
    cursorLeft () {
      const text = this.text
      const position = text.selectionStart
      const val = text.value

      let line = val.substr(position - 1, 1)
      if (line === '\n') {
        return
      }
      const afterPosition = position - 1
      text.setSelectionRange(afterPosition, afterPosition)
    },
    deleteRight () {
      const text = this.text
      let position = text.selectionStart
      const val = text.value

      // カーソルが行末 + 左が文字列
      if (val.substr(position, 1) === '\n' && val.substr(position - 1, 1) !== '\n') {
        text.value = val.substr(0, position - 1) + val.substr(position)
        position--
      } else if (!val.substr(position, 1)) {
        // カーソルが文字最後
        text.value = val.substr(0, position - 1)
      } else if (val.substr(position, 1) === '\n' && val.substr(position - 1, 1) === '\n') {
        // カーソルが行末 + 行頭
        return
      } else {
        text.value = val.substr(0, position) + val.substr(position + 1)
      }

      text.setSelectionRange(position, position)
    },
    deleteLine (code) {
      // 2回連続 d じゃなかったらスルー
      if (this.normalLastInput !== 68) {
        return code
      }

      const text = this.text
      let position = text.selectionStart
      const val = text.value

      // 行頭
      let start = val.lastIndexOf('\n', position) === -1
        ? 0 : val.lastIndexOf('\n', position)

      // 行末
      const end = val.indexOf('\n', position)

      // カーソルが行末
      if (start === end) {
        start = val.lastIndexOf('\n', position - 1)
      }

      if (end === -1) {
        text.value = val.substr(0, start)
        this.vimClip = val.substr(start)
      } else {
        text.value = val.substr(0, start) + val.substr(end)
        this.vimClip = val.substr(start, end - start)
      }

      text.setSelectionRange(position, position)
      return null
    },
    lineCopy (code) {
      // 2回連続 y じゃなかったらスルー
      if (this.normalLastInput !== 89) {
        return code
      }

      const text = this.text
      let position = text.selectionStart
      const val = text.value

      // 行頭
      let start = val.lastIndexOf('\n', position) === -1
        ? 0 : val.lastIndexOf('\n', position)

      // 行末
      const end = val.indexOf('\n', position)

      // カーソルが行末
      if (start === end) {
        start = val.lastIndexOf('\n', position - 1)
      }

      if (end === -1) {
        this.vimClip = val.substr(start)
      } else {
        this.vimClip = val.substr(start, end - start)
      }
      return null
    },
    paste () {
      const text = this.text
      let position = text.selectionStart
      const val = text.value

      // 行頭
      let start = val.lastIndexOf('\n', position) === -1
        ? 0 : val.lastIndexOf('\n', position)

      // 行末
      const end = val.indexOf('\n', position)

      // カーソルが行末
      if (start === end) {
        start = val.lastIndexOf('\n', position - 1)
      }

      if (end === -1) {
        text.value = val + '\n' + this.vimClip
        position = val.length + 1
      } else {
        text.value = val.substr(0, end) + '\n' + this.vimClip + val.substr(end)
        position = val.substr(0, end).length + 1
      }

      text.setSelectionRange(position, position)
    },
    // カーソルを先頭 or 最後
    moveTopOrBottom (code) {
      const text = this.text
      const val = text.value

      // shift + g で最後
      if (this.shiftOn) {
        text.setSelectionRange(val.length, val.length)
        text.scrollTop = 99999
        return null
      }

      // 2回連続 g で先頭
      if (this.normalLastInput === 71) {
        text.setSelectionRange(0, 0)
        text.scrollTop = 0
        return null
      }

      return code
    },
    moveEndOfLine (code) {
      if (!this.shiftOn) {
        return code
      }

      const text = this.text
      const val = text.value
      let position = text.selectionStart

      // 改行があるか探す
      let line = val.indexOf('\n', position)
      if (line === -1) {
        line = text.value.length
      }

      text.setSelectionRange(line, line)
      return null
    },
    moveBeginOfLine () {
      const text = this.text
      const val = text.value
      let position = text.selectionStart

      // 改行があるか探す
      let line = val.lastIndexOf('\n', position - 1)
      if (line === -1) {
        line = 0
      } else {
        line++
      }

      text.setSelectionRange(line, line)
      return null
    },
    autoScroll (line) {
      const text = this.text
      const val = text.value
      const halfHeight = text.clientHeight / 2

      const lineArray = val.substr(0, line).match(/\n/g)
      if (!lineArray) {
        text.scrollTop = 0
        return
      }

      const lineNum = lineArray.length
      const cursorHeight = lineNum * this.lineHeight
      if (cursorHeight > halfHeight) {
        text.scrollTop = cursorHeight - halfHeight
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.vimMode {
display: inline-block;
margin: 0 2px;
}
.debug {
  h5 {
    color: #fff;
  }
  position:absolute;
  background-color:black;
  color: #fff;
  padding: 5px;
  font-size: 12px;
  width: 100%;
  text-align:left;
  z-index: 99;
}
</style>
