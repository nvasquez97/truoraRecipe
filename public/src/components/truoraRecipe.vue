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
          <input class="form-control" type="search" placeholder="Busca tu receta" aria-label="Search" v-model="search">
          <button class="btn btn-outline-info my-2 my-sm-0" type="submit" v-on:click="searchRecipe">TruoraRecipe</button>
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
  import bus from "./../bus.js";
  export default {
    name: 'truoraRecipe',
    data() {
      return {
        search:'',
      }
    },
    components: {createRecipe,  recipes},
    methods: {
      searchRecipe:function(event){
        //Evita que se realice el submit del form
        event.preventDefault();
        if(this.search!=="")
        {
          bus.$emit("search",this.search);
          this.search="";
        }
      },
    }
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