<template>
  <div id="app">
    <div class="container">
      <div class="row col-md-6 offset-md-3">
        <create-recipe></create-recipe>
        <recipes></recipes>
      </div>
    </div>
  </div>
</template>


<script>
  import createRecipe from './createRecipe.vue';
  import recipes from './recipes.vue';
  export default {
    name: 'truoraRecipe',
    data() {
      return {}
    },
    components: {createRecipe,  recipes},
  }
</script>

<style>
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity .5s
  }
  .fade-enter,
  .fade-leave-active {
    opacity: 0
  }
</style>