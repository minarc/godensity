<template>
  <v-container fluid grid-list-lg class="white">
    <v-layout wrap column>
      <v-flex>
        <v-textarea counter solo v-model=text required @keyup.enter="submit()" hint="충분히 길어야 합니다." label="예문 입력" value=text single-line :disabled=input></v-textarea>
      </v-flex>
      <v-flex v-show="active">
        <v-progress-linear :active=active color="orange" :indeterminate="true"></v-progress-linear>
      </v-flex>
      <v-flex>
        <d3-network :net-nodes=nodes :net-links=links :options=options />
      </v-flex>
      <v-flex v-show="alert">
        <v-alert :value=alert color="error" icon="warning" transition="scale-transition"> {{ message }} </v-alert>
        <v-card>
          <v-chip v-for="item in textrank" :key=item color="green" text-color="white">
            <v-avatar class="green darken-4">{{ item.value.toFixed(2) }}</v-avatar>{{ item.key }}
          </v-chip>
        </v-card>
      </v-flex>
      <v-flex>
        <v-card>
          <v-chip v-for="item in nouns" :key=item color="orange accent-4" text-color="white"><v-icon left>label</v-icon>{{ item }}</v-chip>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios'
import D3Network from 'vue-d3-network'

export default {
  components: {
    D3Network
  },
  data () {
    return {
      text: '',
      active: false,
      input: false,
      alert: false,
      message: '',
      keywords: [],
      nouns: [],
      adjacency: [],
      nodes: [],
      links: [],
      options: {
        nodeSize: 40,
        canvas: false,
        force: 2500,
        linkLabels: true,
        nodeLabels: true,
        fontSize: 15,
        size: {
          h: 250
        },
        forces: {
          X: 0.7,
          Y: 1,
          Center: true
        }
      }
    }
  },
  methods: {
    submit (message, error) {
      this.active = this.input = true
      this.alert = false
      this.text = this.text.trim()
      this.keywords = []
      this.nodes = []
      this.links = []

      axios.post('/api/v2/textrank', {
        text: this.text
      }).then(response => {
        this.active = false
        this.input = false
        this.nouns = response.data['result']['linked_nouns']
        this.adjacency = response.data['result']['adjacency_list']
        this.keywords = response.data['result']['keywords']

        this.keywords.forEach(e => {
          this.nodes.push({name: e['key'], _color: 'orange'})
        })

        response.data['result']['adjacency_list'].forEach(e => {
          const sid = this.indexing(e['key'])
          e['value'].forEach(values => {
            this.links.push({sid: sid, tid: this.indexing(values['key']), name: values['value'].toFixed(2), _color: 'orange'})
          })
        })
      }).catch(error => {
        this.alert = true
        this.message = error['response']['status'] + ' ' + error['response']['statusText']
      })
    },
    indexing (word) {
      for (let i = 0; i < this.keywords.length; i++) {
        if (word === this.keywords[i]['key']) {
          return i
        }
      }
    }
  }
}
</script>

<style>
.node-label {
    fill: #e75325
}
.node {
    -webkit-transition: fill .5s ease;
    fill: #dcfaf3;
    stroke: rgba(240, 35, 35, 0.7);
    stroke-width: 1px;
    transition: fill .5s ease
}
.link:hover,
.node:hover {
    stroke-width: 2px
}
</style>
