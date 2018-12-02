<template>
  <div class="uploadImage">

    <a @click="selectFile" class="button is-small">
      <fa icon="file-image" />
      Image
    </a>

    <div v-if="image.edit" id="cropperEditor">
      <div class="controller">
        <button @click="rotateLeft" class="icon"><fa icon="undo" /></button>
        <button @click="rotateRight" class="icon"><fa icon="redo" /></button>
        <button @click="zoomOut" class="icon"><fa icon="search-minus" /></button>
        <button @click="zoomIn" class="icon"><fa icon="search-plus" /></button>
        <button @click="closeEditor" class="icon"><fa icon="times-circle" /></button>
        <button @click="uploadImage" class="icon"><fa icon="check" /></button>
      </div>
      <img id="cropperImg">
    </div>

    <div class="modal" :class="{'is-active': modal}">
      <div @click="modalOff" class="modal-background"></div>
      <div class="modal-content">
        <article class="message">
          <div class="message-header">
            <p>Eye catch or Image</p>
            <a @click="modalOff" class="delete" aria-label="delete"></a>
          </div>
          <div class="message-body">
            <a @click="eyecatchStart" class="button">Eye catch</a>
            <a @click="imageStart" class="button">Image</a>
          </div>
          <footer class="modal-card-foot">
            <a @click="modalOff" class="button">Cancel</a>
          </footer>
        </article>
      </div>
    </div>

    <input @change="previewImage" type="file" id="previewImage">
  </div>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
import Cropper from 'cropperjs'
import 'cropperjs/dist/cropper.min.css'

export default {
  data () {
    return {
      cropper: undefined,
      modal: false,
      eyecatch: false
    }
  },
  computed: {
    ...mapGetters(['image', 'previewArticle'])
  },
  methods: {
    ...mapActions(['addImage']),
    ...mapMutations(['setImageEdit', 'setPreviewArticle', 'setImageEdit']),
    eyecatchStart () {
      this.modalOff()
      this.eyecatch = true
      document.getElementById('previewImage').click()
    },
    imageStart () {
      this.modalOff()
      this.eyecatch = false
      document.getElementById('previewImage').click()
    },
    selectFile () {
      if (this.image.edit) {
        return
      }
      console.log('selectfile')
      this.modal = true
    },
    modalOff () {
      console.log('modal off')
      this.modal = false
    },
    previewImage (e) {
      const file = e.target.files[0]
      const reader = new FileReader()
      this.setImageEdit(true)
      reader.onload = (event) => {
        const img = document.getElementById('cropperImg')
        img.src = event.target.result
        // 画像読み込み完了後
        img.onload = () => {
          // set cropper
          this.cropper = new Cropper(img, {
            zoomable: true,
            zoomOnWheel: false
          })
        }
      }
      reader.readAsDataURL(file)
    },
    rotateLeft () {
      this.cropper.rotate(-15)
    },
    rotateRight () {
      this.cropper.rotate(15)
    },
    zoomIn () {
      this.cropper.zoom(0.5)
    },
    zoomOut () {
      this.cropper.zoom(-0.5)
    },
    closeEditor () {
      document.getElementById('previewImage').value = null
      this.cropper.destroy()
      this.setImageEdit(false)
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
    uploadImage () {
      // 画像幅決定 eyecatch 600, other 800
      const maxWidth = this.eyecatch ? 600 : 800
      const info = this.cropper.getData()
      let outputWidth = info.width
      if (outputWidth > maxWidth) {
        outputWidth = maxWidth
      }

      // 画像文字列生成
      const dataUrl = this.cropper.getCroppedCanvas({ width: outputWidth }).toDataURL()
      if (this.eyecatch) {
        this.setPreview('eyecatchUpdate', new Date())
        this.setPreview('eyecatch', dataUrl)
        this.setImageEdit(false)
        return
      }

      this.addImage({ dataUrl: dataUrl })
    }
  }
}
</script>

<style lang="scss" scoped>
#previewImage {
  display: none;
}

.button {
  font-weight: normal;
  svg {
    margin-right: 5px;
  }
}

#cropperEditor {
  .icon {
    background: transparent;
    border: 0;
    color:#fff;
    cursor: pointer;
  }
  * {
    box-sizing: border-box;
  }
  display: block;
  box-sizing: border-box;
  position: fixed;
  width: 92%;
  top: 1%;
  left: 4%;
  background-color: #5B5B5B;
  box-shadow: 0 0 12px 7px rgba(0, 0, 0, 0.3);
  border-radius: 6px;
  z-index: 100;
  animation: fadeIn 1s ease 0s 1 normal;
  #cropperImg {
    max-width: 92%;
  }
  .controller {
    padding: 3px 10px;
    .btn {
      display: inline-block;
      color: #fff;
      width: 24px;
      height: 24px;
      text-align: center;
      cursor: pointer;
    }
  }
}
.uploadImage {
  display: inline-block;
  margin: 0 2px;
  .uk-button {
    text-transform: inherit;
    line-height: 20px;
    padding: 0 5px;
    font-size: 12px;
    background-color: #616161;
    color: #fff;
    border: 0;
    border-radius: 3px;
  }
}
h3 {
  margin-bottom: 5px;
}
@keyframes fadeIn {
  0% {opacity:0}
  100% {opacity:1}
}
</style>
