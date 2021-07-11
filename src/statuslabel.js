const template = `
<p v-if="workstatus" id="status-label" :style="{color: statusColor}">{{status}}</p>
`
export default {
    template,
    props: ['workstatus'],
    computed: {
        status() {
            let status = this.workstatus.toLowerCase()
            return status === 'working' ? 'Working...' : status.charAt(0).toUpperCase() + status.slice(1)
        },
        statusColor() {
            let status = this.workstatus.toLowerCase()
            switch (status) {
                case 'working':
                    return '#000000'
                case 'success':
                    return '#00ff00'
                case 'failure':
                    return '#ff0000'
            }
        }
    }
}