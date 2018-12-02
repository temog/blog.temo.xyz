<template>
  <div>

    <div @click="toggleIndex" class="nuko"></div>

    <nav class="navbar is-fixed-top">
      <div class="navbar-brand">
        <a @click="toggleSideMenu" class="navbar-burger">
          <span aria-hidden="true" />
          <span aria-hidden="true" />
          <span aria-hidden="true" />
        </a>
        <router-link to="/" class="navbar-item brand">
          てもぐ
        </router-link>
      </div>
    </nav>
    <div v-if="showSideMenu" class="sidemenu">
      <div @click="toggleSideMenu" class="bg"></div>
      <aside :class="{on: showMenu}" class="menu">
        <p class="menu-label">Search</p>
        <div class="field has-addons">
          <div class="control">
            <input v-model="keyword" class="input" type="text">
          </div>
          <div class="control">
            <a @click="search" class="button is-info">
              <span class="icon">
              <fa icon="search" /></span>
              <span>Search</span>
            </a>
          </div>
        </div>

        <p class="menu-label">Popular Articles</p>
        <ul class="menu-list popular">
          <li v-for="pa in popularArticles" :key="pa._id">
            <router-link :to="articleLink(pa)" @click.native="toggleSideMenu">
              {{ pa.title }}
              <span class="date tag is-gray">{{ pa.date | date }}</span>
            </router-link>
          </li>
        </ul>

        <p class="menu-label">Tags</p>
        <div class="field is-grouped is-grouped-multiline menuTags">
          <div v-for="(tag, key) in tags" :key="key"
            @click="toggleSideMenu(); $router.push('/tag/' + tag.name)" class="control">
            <tag-component :tag="tag" :addClass="'is-medium'" />
          </div>
        </div>

        <p class="menu-label">About Me</p>
        <div class="card profile">
          <div class="card-content">
            <div class="media">
              <div class="media-left">
                <figure class="image is-48x48">
                  <img class="is-rounded" src="https://pbs.twimg.com/profile_images/713621737418670080/nyAp3HeN_400x400.jpg">
                </figure>
              </div>
              <div class="media-content">
                <p class="title is-4">ても</p>
                <p class="subtitle is-7">渋谷の某ベンチャー企業で Web Developer</p>
              </div>
            </div>

            <div class="content">
              <div v-for="(data, key) in myTags" :key="key" class="tags has-addons">
                <span class="tag is-dark">{{ data.name }}</span>
                <span v-for="(tag, key2) in data.tags" :key="key2" class="tag">{{ tag }}</span>
              </div>

              <fa icon="map-marker-alt" /> Tokyo<br>

              <fa :icon="['fab', 'github']" />
              <a href="https://github.com/temog" target="_blank">https://github.com/temog</a><br>

              <fa :icon="['fab', 'twitter']" />
              <a href="https://twitter.com/temo_g" target="_blank">https://twitter.com/temo_g</a><br>
            </div>
          </div>
        </div>

      </aside>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import Util from '@/plugins/util'
import TagComponent from '@/components/Tag'
export default {
  components: {
    TagComponent
  },
  data () {
    return {
      keyword: null,
      showSideMenu: false,
      showMenu: false,
      myTags: [{
        name: 'Age',
        tags: ['おっさん']
      },
      {
        name: 'Gender',
        tags: ['♂']
      },
      {
        name: 'Job',
        tags: ['Web Developer']
      },
      {
        name: 'Favorite',
        tags: ['Programming', 'PHP', 'Go', 'Python', 'Vue.js', 'Vim', 'Atom', 'ONE PIECE', 'スラムダンク', 'LOVE PSYCHEDELICO', '桑田佳祐さま', '小雪さま', 'チャリンコ', '(ΦωΦ)']
      },
      {
        name: 'Skills',
        tags: ['PHPそれなり', 'Perlそれなり', 'JavaScriptまーまー', 'Python', 'Go ちょっとだけ', 'Java微妙', 'androidアプリ1個作った']
      }
      ]
    }
  },
  computed: {
    ...mapGetters(['tags', 'popularArticles'])
  },
  mounted () {
    console.log('Hoge')
  },
  methods: {
    ...mapMutations(['toggleIndex']),
    toggleSideMenu () {
      console.log('toggle')
      if (this.showSideMenu) {
        this.showMenu = false
        setTimeout(() => {
          this.showSideMenu = false
        }, 300)
      } else {
        this.showSideMenu = true
        setTimeout(() => {
          this.showMenu = true
        }, 100)
      }
    },
    search () {
      this.$router.push('/keyword/' + this.keyword)
      this.toggleSideMenu()
    },
    articleLink (article) {
      return Util.articleLink(article)
    }
  }
}
</script>

<style lang="scss" scoped>
.sidemenu {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 100;
  .bg {
    content: "";
    display: block;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,0.5);
    z-index: 1;
  }
  .menu-label {
    margin-top: 2.5em;
  }
}

.menu {
  width: 80%;
  height: 100%;
  padding: 10px;
  position: fixed;
  top: 0;
  left: -80%;
  z-index: 2;
  background-color: #fff;
  overflow-y: auto;
  transition: all 300ms 0s ease;
  &.on {
    left: 0;
  }
}

.profile {
  .is-dark {
    width: 72px;
  }
  .media {
    margin-bottom: 1.5rem;
  }
}

.popular {
  font-size: 0.8rem;
  .date {
    font-size: 0.6rem;
  }
}
</style>
