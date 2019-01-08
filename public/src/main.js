'use strict'
import Vue from 'vue'
import truoraRecipe from './components/TruoraRecipe.vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue);
new Vue({
  el: 'truora-recipe',
  created: function () {
    console.log('root instance was created')
  },
  components: {truoraRecipe},
  methods: {}
});