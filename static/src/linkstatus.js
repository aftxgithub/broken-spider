const template = `
<div id="link-status">
    <a href="#">{{ link }}</a>
    <span class="status">{{ broken ? 'Bad' : 'Good' }}</span>
</div>
`
export default {
    template,
    props: ['link', 'broken']
}