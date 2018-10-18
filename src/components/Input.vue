<template>
  <v-container fluid grid-list-lg class="white">
    <v-layout wrap column>
      <v-flex>
        <v-textarea counter solo v-model=text required @keyup.enter="submit()" hint="충분히 길어야 합니다." label="예문 입력" value=text single-line :disabled=input></v-textarea>
      </v-flex>
      <v-flex>
        <v-progress-linear :active=active color="orange" :indeterminate="true"></v-progress-linear>
      </v-flex>
      <v-flex>
        <v-alert :value=alert color="error" icon="warning" transition="scale-transition"> {{ message }} </v-alert>
        <v-card>
          <v-chip outline v-for="item in textrank" :key=item color="green" text-color="green">
            <v-avatar class="green darken-4">{{ item.value.toFixed(2) }}</v-avatar>{{ item.key }}
          </v-chip>
        </v-card>
      </v-flex>
      <v-flex>
        <v-card>
          <v-chip v-for="item in keywords" :key=item color="orange accent-4" text-color="white"><v-icon left>label</v-icon>{{ item }}</v-chip>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      text: '',
      active: false,
      input: false,
      alert: false,
      message: '',
      keywords: [],
      adjacency: [],
      textrank: []
    }
  },
  methods: {
    submit (message, error) {
      this.active = this.input = true
      this.alert = false
      this.text = this.text.trim()

      axios.post('/api/v2/textrank', {
        text: this.text
      }).then(response => {
        this.active = false
        this.input = false
        this.keywords = []
        this.keywords = response.data['result']['linked_nouns']
        this.adjacency = response.data['result']['adjacency_list']
      }).catch(error => {
        this.alert = true
        this.message = error['response']['status'] + ' ' + error['response']['statusText']
      })

      axios.post('/api/v1/textrank', {
        text: this.text
      }).then(response => {
        this.textrank = []
        this.textrank = response.data['result']
      }).catch(error => {
        this.alert = true
        this.message = error['response']['status'] + ' ' + error['response']['statusText']
      })
    }
  }
}
</script>
