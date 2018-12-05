<template>
  <div id="truorarecipe">
    <div class="container">
      <div class="row col-md-10 offset-md-3">
        <recipes></recipes>
        <create-recipe></create-recipe>
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