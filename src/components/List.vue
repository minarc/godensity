<template>
      <!-- <v-list two-line> -->
        <v-container fluid grid-list-lg>
          <v-layout row wrap>
            <v-flex xs12 v-for="item in items" :key="item">
              <v-card class="elevation-4">
                <v-layout>
                  <v-flex xs7>
                  <v-card-title>
                      <h2 class="text-sm-left">{{ item.title }}</h2>
                      <p class="text-sm-left"> {{ item.description.slice(0, 100) }} ... </p>
                  </v-card-title>
                </v-flex>
                  <v-flex xs5>
                    <v-img :src="item.image[0]" max-height="100%" position="top"></v-img>
                  </v-flex>
              </v-layout>
              <v-divider></v-divider>
                <v-card-text>
                  <v-chip small disabled outline color="orange" text-color="orange" v-for="key in item.discoveryTags" :key="key"> #{{ key }} </v-chip>
                  <v-chip small disabled outline color="blue" text-color="blue" v-for="key in item.keyword" :key="key"> #{{ key }} </v-chip>
                </v-card-text>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
        <!-- <template v-for="(item, index) in items">
            <v-subheader v-if="item.header" :key="item.header"> {{ item.header }} </v-subheader>

            <v-divider v-else-if="item.divider" :inset="item.inset" :key="index"></v-divider>

            <v-list-tile v-else :key="item.title" avatar style="height: 120px">
              <img v-if="item.image.length > 0" :src="item.image[0]" height="70" aspect-ratio=1.0 style="margin-right:10px">

              <v-list-tile-content style="height: 100px">
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
        </template>-->
      <!-- </v-list> -->
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
          axios
            .post('/api/v2/textrank', { text: response.data })
            .then(response => {
              element['discoveryTags'] =
                response.data['result']['linked_nouns']
            })
        })
      })
    })
  }
}
</script>

<style>
</style>
