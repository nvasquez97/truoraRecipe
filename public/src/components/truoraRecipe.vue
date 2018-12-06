<template>
  <div id="truorarecipe">
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <div class="collapse navbar-collapse container">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item active">
              <a class="navbar-brand" href="#">Nicolás Vásquez Murcia</a>
          </li>
        </ul>
        <br>
        <form class="form-inline">
          <input class="form-control" type="search" placeholder="Busca tu receta" aria-label="Search">
          <button class="btn btn-outline-info my-2 my-sm-0" type="submit">TruoraRecipe</button>
        </form>
      </div>
    </nav>
    <div class="container">
        <recipes></recipes>
        <create-recipe></create-recipe>
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