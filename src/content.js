const template = `
<search-box :submit="onSubmit"></search-box>
<status-label></status-label>
<link-status></link-status>
`

export default {
  template,
  methods: {
    onSubmit(url) {
      console.log(url)
    }
  }
}