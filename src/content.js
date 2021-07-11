const template = `
<search-box :submit="onSubmit"></search-box>
<status-label :workstatus="workstatus"></status-label>
<link-status></link-status>
`

export default {
  template,
  data() {
    return {
      workstatus: null
    }
  },
  methods: {
    onSubmit(url) {
      this.workstatus = "working"
    }
  }
}