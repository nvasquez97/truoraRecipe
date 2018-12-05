<template>
  <div>
    <h2>Create a Recipe</h2>
    <form @submit.prevent>
  <div class="container">
      <div class="form-group">
        <div class="form-row">
          <div class="form-group col-md-4">
            <label for="inputName">Nombre</label>
            <input id="inputName" type="text" class="form-control" placeholder="Escribe el nombre de la receta" v-model="name">
          </div>
          <div class="form-group col-md-8">
            <label for="inputURL">Imagen</label>
            <input id="inputURL" type="text" class="form-control" placeholder="Pega acá el URL de una imagen que te guste" v-model="imgURL">
          </div>
        </div>
        <div class="form-row">
          <div class="col-md-12">
              <label for="inputName">Describe tu plato</label>
            <input id="inputName" type="text" class="form-control" placeholder="Escribe una descripción del plato que se prepara en tu receta" v-model="description">
            </div>
        </div>
        <div class="form-row">
          <div class="form-group col-md-6">
            <label for="inputInstructions">Instrucciones</label>
            <input id="inputInstructions" type="text" class="form-control" placeholder="Escribe cómo se prepara tu receta" v-model="instructions">
          </div>
          <div class="form-group col-md-6">
            <label for="inputIngredients">Ingredientes</label>
            <input id="inputIngredients" type="text" class="form-control" placeholder="Escribe los ingredientes separados por comas"
            v-model="ingredients">
          </div>
        </div>
      </div>
      <button type="submit" class="btn btn-primary">Submit</button>
    </div>
    </form>
  </div>
</template>

<script>
  import axios from 'axios';  
  import bus from "./../bus.js";

  export default {
    data() {
      return {
        name: '',
        imgURL:'',
        ingredients:'',
        description:'',
        instructions:'',
      }
    },
    methods: {
      addRecipe() {
        let url = 'http://localhost:8000/recipe/';
        let param = {
          name: this.name,
          imgURL:this.imgURL,
          description:this.description,
          instructions:this.instructions,
          ingredients:this.ingredients,
          done: 0
      };
        axios.post(url, param, {
          headers: {
            'Content-Type':'application/json'
            //'Content-Type':'multipart/form-data'
          }
        }).then((response) => {
          console.log(response);
          this.clear();
          this.refresh();
          this.typing = false;
        }).catch((error) => {
          console.log(error);
        })
      },
      clear() {
        this.name = '';
        this.imgURL ='';
        this.ingredients='';
        this.description='';
        this.instructions='';
      },
      refresh() {
        bus.$emit("reloadRecipes")
      }
    }
  }
</script> 