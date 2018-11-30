<template>
  <div>
    <h2>Create a Recipe</h2>
    <form @submit.prevent>
  <div class="container">
      <div class="form-group">
        <div class="form-row">
          <div class="form-group col-md-4">
            <label for="inputName">Nombre</label>
            <input id="inputName" type="text" class="form-control" placeholder="Escribe el nombre de la receta">
          </div>
          <div class="form-group col-md-8">
            <label for="inputURL">Imagen</label>
            <input id="inputURL" type="text" class="form-control" placeholder="Pega acá el URL de una imagen que te guste">
          </div>
        </div>
        <div class="form-row">
          <div class="col-md-12">
              <label for="inputName">Describe tu plato</label>
            <input id="inputName" type="text" class="form-control" placeholder="Escribe una descripción del plato que se prepara en tu receta">
            </div>
        </div>
        <div class="form-row">
          <div class="form-group col-md-6">
            <label for="inputInstructions">Instrucciones</label>
            <input id="inputInstructions" type="text" class="form-control" placeholder="Escribe cómo se prepara tu receta">
          </div>
          <div class="form-group col-md-6">
            <label for="inputIngredients">Ingredientes</label>
            <input id="inputIngredients" type="text" class="form-control" placeholder="Escribe los ingredientes separados por comas">
          </div>
        </div>
      </div>
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>

<script>
  import axios from 'axios';  
  import bus from "./../bus.js";

  export default {
    data() {
      return {
        todo: '',
        typing: false,
      }
    },
    methods: {
      addRecipe(event) {
        if (event) event.preventDefault();
        let url = 'http://localhost:9000/api/add';
        let param = {
          name: this.todo,
          done: 0
      };
        axios.post(url, param).then((response) => {
          console.log(response);
          this.clearTodo();
          this.refreshTodo();
          this.typing = false;
        }).catch((error) => {
          console.log(error);
        })
      },
      clearTodo() {
        this.todo = '';
      },
      refreshTodo() {
        bus.$emit("refreshTodo")
      }
    }
  }
</script>