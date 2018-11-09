<template>
  <v-container fluid grid-list-lg>
    <d3-network ref='net' :net-nodes=nodes :net-links=links :options=options />
  </v-container>
</template>
<script>
import D3Network from 'vue-d3-network'

export default {
  components: {
    D3Network
  },
  data () {
    return {
      nodes: [
        { id: 1, name: 'my awesome node 1' },
        { id: 2, name: 'my node 2' },
        { id: 3, name: 'orange node' }
      ],
      links: [
        { sid: 1, tid: 2, name: '2.57' },
        { sid: 2, tid: 2 },
        { sid: 3, tid: 2 }
      ],
      options: {
        nodeSize: 40,
        canvas: false,
        force: 2500,
        linkLabels: true,
        nodeLabels: true,
        fontSize: 15
      }
    }
  }
}
</script>
