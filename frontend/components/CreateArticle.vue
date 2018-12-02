<template>
  <div>
    <div class="columns">
      <div class="column">
        <div class="field">
          <label class="label">タイトル</label>
          <div class="control">
            <input v-model="title" class="input" type="text">
          </div>
          <input-error :msg="previewError.title" />
        </div>

        <div class="field">
          <label class="label">Alias</label>
          <div class="control">
            <input v-model="alias" class="input" type="text">
          </div>
          <input-error :msg="previewError.alias" />
        </div>
      </div>
      <div class="column is-one-third">
        <div class="field">
          <label class="label">Eye catch</label>
          <div class="control">
            <img v-if="previewEyecatch" :src="previewEyecatch">
            <img v-if="!previewEyecatch && previewArticle.id !== 'create'" :src="savedEyecatch()">
          </div>
          <input-error :msg="previewError.eyecatch" />
        </div>
      </div>
    </div>

    <div class="field">
      <label class="label">タグ</label>
      <div class="field is-grouped is-grouped-multiline">
        <div v-for="(tag, key) in tags" :key="key" class="control">
          <span @click="toggleTag(tag)" class="tags has-addons">
            <span class="tag" :class="isTagActive(tag)">{{ tag.name }}</span>
            <span class="tag is-dark">{{ tag.articles }}</span>
          </span>
        </div>
      </div>
      <input-error :msg="previewError.tags" />
    </div>

    <div class="field">
      <label class="label">公開日時</label>
      <div class="control">
        <flat-pickr v-model="date" :config="pickrConfig" class="input" />
      </div>
      <input-error :msg="previewError.date" />
    </div>

    <div class="field">
      <label class="label">状態</label>
      <div class="control">
        <div class="select">
          <select v-model="status">
            <option value="publish">公開</option>
            <option value="private">非公開</option>
          </select>
        </div>
      </div>
    </div>

    <div class="field">
      <div class="is-pulled-right">
        <upload-image />
        <vim-mode id="body" />
      </div>
      <label class="label">
        テキスト
      </label>
      <div class="control">
        <textarea v-model="body" id="body" class="textarea"></textarea>
      </div>
      <input-error :msg="previewError.body" />
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from 'vuex'
import Util from '@/plugins/util'
import FlatPickr from 'vue-flatpickr-component'
import 'flatpickr/dist/flatpickr.css'
import VimMode from '@/components/VimMode'
import UploadImage from '@/components/UploadImage'
import InputError from '@/components/InputError'
import UserMenu from '@/components/UserMenu'
export default {
  components: {
    FlatPickr,
    VimMode,
    UploadImage,
    UserMenu,
    InputError
  },
  data () {
    return {
      pickrConfig: {
        enableTime: true,
        dataFormat: 'Y-m-d H:i',
        time_24hr: true,
        disableMobile: true
      },
      md: null,
      textarea: null
    }
  },
  watch: {
    'image.tag' (tag) {
      const body = document.getElementById('body')
      let text = body.value
      const cursorPosition = body.selectionStart
      const before = text.substr(0, cursorPosition)
      const after = text.substr(cursorPosition)
      body.value = before + tag + after
    }
  },
  computed: {
    ...mapGetters(['tags', 'previewArticle', 'previewError', 'previewEyecatch', 'image']),
    title: {
      get (value) {
        return this.previewArticle.title
      },
      set (value) {
        this.setPreview('title', value)
      }
    },
    alias: {
      get (value) {
        return this.previewArticle.alias
      },
      set (value) {
        this.setPreview('alias', value)
      }
    },
    date: {
      get (value) {
        return this.previewArticle.date
      },
      set (value) {
        this.setPreview('date', value)
      }
    },
    status: {
      get (value) {
        return this.previewArticle.status
      },
      set (value) {
        this.setPreview('status', value)
      }
    },
    body: {
      get (value) {
        return this.previewArticle.body
      },
      set (value) {
        this.setPreview('body', value)
      }
    }
  },
  mounted () {
    this.$nextTick(() => {
      this.textarea = document.querySelector('#body')
      this.textarea.addEventListener('keyup', this.previewAutoScroll)
    })
  },
  destroyed () {
    this.textarea.removeEventListener('keyup', this.previewAutoScroll)
  },
  methods: {
    ...mapActions(['findId']),
    ...mapMutations(['setPreviewArticle', 'setPreviewScroll']),
    previewAutoScroll (e) {
      const value = this.textarea.value

      // すべての改行数を取得
      const allNewLine = value.match(/\n/g)
        ? value.match(/\n/g).length : 0

      // カーソル位置の改行がどこか取得
      const ss = this.textarea.selectionStart
      const currnetNewLine = value.substr(0, ss).match(/\n/g)
        ? value.substr(0, ss).match(/\n/g).length : 0
      // 現在位置の％を取得
      this.setPreviewScroll(currnetNewLine / allNewLine * 100)
    },
    setPreview (key, value) {
      if (!value) {
        return
      }

      this.setPreviewArticle({
        key: key,
        value: value
      })
    },
    toggleTag (tag) {
      const tags = JSON.parse(JSON.stringify(this.previewArticle.tags))
      const i = tags.indexOf(tag._id)
      if (i === -1) {
        tags.push(tag._id)
      } else {
        tags.splice(i, 1)
      }
      console.log(tags)
      this.setPreviewArticle({
        key: 'tags',
        value: tags
      })
    },
    isTagActive (tag) {
      return this.previewArticle.tags.indexOf(tag._id) !== -1 ? 'is-primary' : 'is-white'
    },
    savedEyecatch () {
      return Util.eyecatchUrl(this.previewArticle.id, this.previewArticle.updated_at)
    }
  }
}
</script>
