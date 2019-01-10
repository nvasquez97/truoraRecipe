<template>
  <div>
    <hr>
    <b-button variant="success" v-b-toggle.anadirReceta>
      Añadir receta
    </b-button>
    <br>
    <br>
    <b-collapse id="anadirReceta">
      <h2>
        Añade tu propia receta:
      </h2>
      <b-form @submit.prevent>
        <b-container>
          <b-form-group>
            <b-row>
                <b-col md="4">
                  <label for="inputName">
                    Nombre
                  </label>
                  <b-input id="inputName" type="text" placeholder="Escribe el nombre de la receta" v-model="name"/>
                </b-col>
                <b-col md="8">
                  <label for="inputURL">
                    Imagen
                  </label>
                  <b-input id="inputURL" type="text" placeholder="Pega acá el URL de una imagen que te guste" v-model="imgURL"/>
                </b-col>
            </b-row>
            <b-form-row>
              <b-col md="12">
                <label for="inputName">
                  Describe tu plato
                </label>
                <b-input id="inputName" type="text" placeholder="Escribe una descripción del plato que se prepara en tu receta" v-model="description"/>
              </b-col>
            </b-form-row>
            <b-form-row>
                <b-col md="6">
                  <label for="inputInstructions">
                    Instrucciones
                  </label>
                  <b-input id="inputInstructions" type="text" class="form-control" placeholder="Escribe cómo se prepara tu receta" v-model="instructions"/>
                </b-col>
                <b-col md="6">
                  <label for="inputIngredients">
                    Ingredientes
                  </label>
                  <b-input id="inputIngredients" type="text" class="form-control" placeholder="Escribe los ingredientes separados por comas" v-model="ingredients"/>
                </b-col>
            </b-form-row>
          </b-form-group>
          <button type="button" class="btn btn-primary" data-toggle="collapse" href="#anadirReceta" v-on:click="addRecipe()">
            Agregar
          </button>
        </b-container>
      </b-form>
    <hr>
    </b-collapse>
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
          ingredients:this.ingredients
      };
        axios.put(url, param, {
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
        this.clear();
        this.refresh();
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