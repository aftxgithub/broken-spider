import PageHeader from './header.js';
import SearchBox from './searchbox.js'

const rootTemplate = `
<page-header></page-header>
<search-box></search-box>
`
const root = {
    template: rootTemplate
}

const app = Vue.createApp(root)
// register cmpts
app.component('page-header', PageHeader)
app.component('search-box', SearchBox)
app.mount("#app")