'use strict'
import Vue from 'vue'
import 'bootstrap/dist/css/bootstrap.min.css';
import truoraRecipe from './components/truoraRecipe.vue'

new Vue({
  el: 'app',
  created: function () {
    console.log('root instance was created')
  },
  components: {truoraRecipe},
  methods: {}
})