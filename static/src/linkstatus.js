const template = `
<div id="link-status">
    <a href="#">{{ url }}</a>
    <span class="status">{{ broken ? 'Bad' : 'Good' }}</span>
</div>
`
export default {
    template,
    props: ['url', 'broken']
}