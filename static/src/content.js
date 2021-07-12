import searchBox from './searchbox.js'
import statusLabel from './statuslabel.js'
import linkStatus from './linkstatus.js'

const template = `
<search-box :submit="onSubmit"></search-box>
<status-label :workstatus="workstatus"></status-label>
<link-status v-for="linkStatus in linkStatuses" :link="linkStatus.url" :broken="linkStatus.broken" />
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
      this.populateLinkStatuses(url)
    },
    
    populateLinkStatuses(url) {
      let ws = new WebSocket(`ws://${window.location.host}/spider?url=${url}`)

      ws.onopen = () => {
        this.workstatus = "working"
      }

      ws.onmessage = (e) => {
        let data = JSON.parse(e.data)
        // insert to beginning of array
        this.linkStatuses.unshift(data)
      }

      ws.onclose = () => {
        this.workstatus = "success"
      }

      ws.onerror = (e) => {
        this.workstatus = "failure"
      }
    }
  },
  components: {
    searchBox,
    statusLabel,
    linkStatus
  }
}