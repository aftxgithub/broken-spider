import searchBox from './searchbox.js'
import statusLabel from './statuslabel.js'
import linkStatus from './linkstatus.js'

const template = `
<search-box :submit="onSubmit"></search-box>
<status-label :workstatus="workstatus"></status-label>
<link-status v-for="linkStatus in linkStatuses" :link="linkStatus.link" :broken="linkStatus.broken" />
`

export default {
  template,
  data() {
    return {
      workstatus: null,
      linkStatuses: []
    }
  },
  methods: {
    onSubmit(url) {
      this.workstatus = "working"
    }
  },
  components: {
    searchBox,
    statusLabel,
    linkStatus
  }
}