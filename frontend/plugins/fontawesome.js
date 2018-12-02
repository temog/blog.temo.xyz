import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import {
  faSearch,
  faMapMarkerAlt,
  faLink,
  faSpinner,
  faUndo,
  faRedo,
  faTimesCircle,
  faCheck,
  faFileImage,
  faSearchPlus,
  faSearchMinus
} from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faGithub, faTwitter } from '@fortawesome/free-brands-svg-icons'

library.add(
  faGithub,
  faTwitter,
  faSearch,
  faMapMarkerAlt,
  faLink,
  faSpinner,
  faUndo,
  faRedo,
  faTimesCircle,
  faCheck,
  faFileImage,
  faSearchPlus,
  faSearchMinus
)
Vue.component('fa', FontAwesomeIcon)
