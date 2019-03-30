<template>
  <v-container fluid grid-list-lg>
    <v-layout row wrap>
      <v-flex xs12 v-for="item in items" :key="item['_id']['key']">
        <v-card class='elevation-4'>
          <v-layout row wrap>
            <v-flex xs12 sm4 md4>
              <v-img :src='item.image[0]' max-height='180px' position='top'></v-img>
            </v-flex>
            <v-flex xs12 sm8 md8>
              <v-card-title>
                <h4 class='text-sm-left'>{{ item.title }}</h4>
                <p class='text-sm-left'>{{ item.description }}</p>
              </v-card-title>
            </v-flex>
          </v-layout>
          <v-progress-linear color="warning" :active="item.discoveryTags.length == 0" height="5" :indeterminate="true"></v-progress-linear>
          <v-divider></v-divider>
          <v-card-text>
            <v-chip
              small
              disabled
              outline
              color='orange'
              text-color='orange'
              v-for='(key, index) in item.discoveryTags'
              :key='index + 100'
            >#{{ key }}</v-chip>
            <v-chip
              small
              disabled
              outline
              color='blue'
              text-color='blue'
              v-for='key in item.keyword'
              :key='key'
            >#{{ key }}</v-chip>
          </v-card-text>
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
      items: []
    }
  },
  created () {
    axios.get('/api/v1/news').then(response => {
      this.items = response.data['data']

      response.data['data'].forEach(element => {
        element['discoveryTags'] = []
        let key = element['_id']['key']
        axios.get('/api/v1/news/contents?key=' + key).then(response => {
          axios.post('/api/v2/textrank', { text: response.data }).then(response => {
            element['discoveryTags'] = response.data['result']['linked_nouns']
          })
        })
      })
    })
  }
}
</script>
