import PageHeader from './header.js';
import PageContent from './content.js'
import SearchBox from './searchbox.js'
import StatusLabel from './statuslabel.js'

const rootTemplate = `
<page-header></page-header>
<page-content></page-content>
`
const root = {
    template: rootTemplate
}

const app = Vue.createApp(root)

// register cmpts
app.component('page-header', PageHeader)

app.component('page-content', PageContent)
app.component('search-box', SearchBox)
app.component('status-label', StatusLabel)

app.mount("#app")