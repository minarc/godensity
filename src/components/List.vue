<template>
<!-- <v-container> -->
  <v-layout row>
    <v-flex>
      <!-- <v-card> -->
        <v-list two-line>
          <template v-for="(item, index) in items">
            <v-subheader v-if="item.header" :key="item.header"> {{ item.header }} </v-subheader>

            <v-divider v-else-if="item.divider" :inset="item.inset" :key="index"></v-divider>

            <v-list-tile v-else :key="item.title" avatar>
                <img :src="item.image" height="50" aspect-ratio=1.0 style="margin-right:10px">

              <v-list-tile-content>
                <v-list-tile-title v-html="item.title"></v-list-tile-title>
                <v-list-tile-sub-title v-html="item.description"></v-list-tile-sub-title>
              </v-list-tile-content>
              <v-chip label outline color="orange" text-color="orange">#키워드 복합 명사</v-chip>
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
    })
  }
}
</script>
