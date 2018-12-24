<template>
  <v-app>
    <v-content>
      <v-container fluid fill-height grid-list class="white">
        <v-layout row wrap align-start justify-center>
          <v-flex sm4 class="elevation-5">
            <v-tabs animated bounce infinite fixed-tabs color="orange darken-1" grow dark show-arrows>
              <v-tabs-slider color="white"></v-tabs-slider>
              <v-tab><v-icon>keyboard</v-icon></v-tab>
              <v-tab><v-icon>view_list</v-icon></v-tab>
              <v-tabs-items v-model="currentItem">
                <v-tab-item>
                  <component :is="tab1"></component>
                </v-tab-item>
                <v-tab-item>
                  <component :is="tab2"></component>
                </v-tab-item>
              </v-tabs-items>
            </v-tabs>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import input from './Input.vue'
import list from './List.vue'

export default {
  name: 'Tabs',
  data () {
    return {
      tab1: '',
      tab2: '',
      cards: ''
    }
  },
  created () {
    this.tab1 = input
    this.tab2 = list
  }
}
</script>
