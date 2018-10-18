<template>
  <v-container fluid grid-list-lg class="orange lighten-5">
    <v-layout row wrap>
          <d3-network ref='net' :net-nodes="nodes" :net-links="links" />
    </v-layout>
  </v-container>
</template>
<script>
import D3Network from 'vue-d3-network'

export default {
  data () {
    return {
      components: {
        D3Network
      },
      network: '',
      nodes: [
        { id: 1 },
        { id: 2 },
        { id: 3 },
        { id: 4 },
        { id: 5 },
        { id: 6 },
        { id: 7 },
        { id: 8 },
        { id: 9 }
      ],
      links: [
        { sid: 1, tid: 2 },
        { sid: 2, tid: 8 },
        { sid: 3, tid: 4 },
        { sid: 4, tid: 5 },
        { sid: 5, tid: 6 },
        { sid: 7, tid: 8 },
        { sid: 5, tid: 8 },
        { sid: 3, tid: 8 },
        { sid: 7, tid: 9 }
      ],
      nodeSize: 20
    }
  },
  created () {
    this.network = D3Network
  }
}
</script>
