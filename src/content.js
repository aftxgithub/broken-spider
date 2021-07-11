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
  }
}