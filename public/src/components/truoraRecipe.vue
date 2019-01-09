<template>
  <div id="truorarecipe">
    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-container>
        <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
        <b-navbar-brand href="#">Nicolás Vásquez Murcia</b-navbar-brand>
        <b-collapse is-nav id="nav_collapse">
          <!-- Right aligned nav items -->
          <b-navbar-nav class="ml-auto">
            <b-nav-form>
              <b-form-input size="sm" class="mr-sm-2" type="text" v-model="search" placeholder="Busca tu receta"/>
              <b-button size="sm" class="my-2 my-sm-0" type="submit" variant="outline-info" v-on:click="searchRecipe">TruoraRecipe</b-button>
            </b-nav-form>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
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