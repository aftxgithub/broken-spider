import PageHeader from './header.js';

const rootTemplate = `
<page-header></page-header>
`
const root = {
    template: rootTemplate
}

const app = Vue.createApp(root)
// register cmpts
app.component('page-header', PageHeader);
app.mount("#app")