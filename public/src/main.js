'use strict'
import Vue from 'vue'
import truoraRecipe from './components/TruoraRecipe.vue'

new Vue({
  el: 'truora-recipe',
  created: function () {
    console.log('root instance was created')
  },
  components: {truoraRecipe},
  methods: {}
})