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
  import CreateRecipe from './createRecipe.vue';
  import Recipes from './recipes.vue';
  export default {
    name: 'truoraRecipe',
    data() {
      return {}
    },
    components: {CreateRecipe,  Recipes},
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