<template>
<!-- <v-container> -->
  <v-layout>
    <v-flex>
      <!-- <v-card> -->
        <v-list two-line>
          <template v-for="(item, index) in items">
            <v-subheader v-if="item.header" :key="item.header"> {{ item.header }} </v-subheader>

            <v-divider v-else-if="item.divider" :inset="item.inset" :key="index"></v-divider>

            <v-list-tile v-else :key="item.title" avatar>
              <!-- <img v-if="item.image.length > 0" :src="item.image[0]" height="60" aspect-ratio=1.0 style="margin-right:10px"> -->

              <v-list-tile-content>
                <v-list-tile-title v-html="item.title"></v-list-tile-title>
                <v-list-tile-sub-title v-html="item.description"></v-list-tile-sub-title>
                <v-list-tile-sub-title>
                  <v-chip small disabled label outline color="orange" text-color="orange" v-for="key in item.discoveryTags" :key="key"> #{{ key }} </v-chip>
                </v-list-tile-sub-title>
                <v-list-tile-sub-title>
                  <v-chip small disabled label outline color="blue" text-color="blue" v-for="key in item.keyword" :key="key"> #{{ key }} </v-chip>
                </v-list-tile-sub-title>
              </v-list-tile-content>
            </v-list-tile>
          </template>
        </v-list>
      <!-- </v-card> -->
    </v-flex>
  </v-layout>
  <!-- </v-container> -->
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
        let key = element['_id']['key']
        axios.get('/api/v1/news/contents?key=' + key).then(response => {
          axios.post('/api/v2/textrank', {text: response.data}).then(response => {
            element['discoveryTags'] = response.data['result']['linked_nouns']
          })
        })
      })
    })
  }
}
</script>

<style>
</style>
