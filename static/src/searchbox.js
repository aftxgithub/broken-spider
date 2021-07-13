const template = `
<input v-model="url" @keyup.enter="onSubmit" id="searchbox" type="text" placeholder="Enter a URL" />
`
export default {
  template,
  emits: ["submit"],
  data() {
    return {
      url: null
    }
  },
  methods: {
    onSubmit(e) {
      this.$emit("submit", this.url)
    }
  }
}