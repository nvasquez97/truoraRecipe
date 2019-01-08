<template>
    <div>
        <div v-show="noBusqueda">
        <hr>
        <h3 v-show="recipes.length>0">Mira las recetas:</h3>
            <div class="row alert alert-info text-center" v-show="recipes.length==0">
                <h3 class="alert alert-info">
                    No hay recetas aún, ¡agrega una!
                </h3>
            </div>
        </div>

        <div v-show="!noBusqueda">
            <hr>
            <div class="row text-center">
                <h3 v-show="recipes.length==0" class="alert alert-info">
                    No hay recetas que coincidan con "{{search}}"
                </h3>
                <h3 v-show="recipes.length>0" class="alert alert-info">
                    Mira las recetas que coinciden con "{{search}}"
                </h3>
            </div>
            <button class="btn btn-dark" v-on:click="fetchRecipes()">
                    Mostrar todas
            </button>
        </div>

        <div v-show="recipes.length>0">
            
            <div id="accordion" class="row">
            <div class="col-md-6" v-for="recipe in recipes">
        <div class="card">
        <div class="card-header" id="headingOne" data-toggle="collapse" v-bind:data-target="'#'+recipe.id" aria-expanded="true" aria-controls="collapseOne" >
        <h5 class="mb-0">
            <p><a href="#" style="color:black">
            {{recipe.name}}: 
            </a>
            {{recipe.description}}
            </p>
        </h5>
        </div>

        <div v-bind:id="recipe.id" class="collapse show" aria-labelledby="headingOne" data-parent="#accordion">
            <div class="card-body">
                <div class="row">
                    <div class="col-md-8">
                        <h6>
                        Ingredientes:
                        </h6>
                        {{recipe.ingredients}}
                        <br>
                        <br>
                        <h6>
                        Instrucciones:
                        </h6>
                        {{recipe.instructions}}
                        <hr>
                    </div>
                    <div class="col-md-4">
                        <img class="img-fluid img-thumbnail" :src="recipe.imgURL"/>
                    </div>
                    <div class="container">
                        <button class="btn btn-info" v-on:click="updateRecipe(recipe)">
                            Editar receta
                        </button>
                        <button class="btn btn-danger" v-on:click="deleteRecipe(recipe.id)">
                            Eliminar receta
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
    </div>
        </div>
    </div>
    
</div>
</template>

<script>
    import axios from 'axios';
    import bus from './../bus.js';

    export default {
        data() {
            return {
                recipes: [],
                noBusqueda:true,
                search:''
            }
        },
        created: function() {
            this.fetchRecipes();
            this.listenToEvents();
        },
        methods: {
            fetchRecipes() {
                let uri = 'http://localhost:8000/recipes/';
                this.noBusqueda=true;
                axios.get(uri).then((response) => {
                    if(response.data !== "No hay recetas en el momento")
                    {
                        this.recipes = response.data;    
                    }
                    else
                    {
                        this.recipes = [];
                    }
                });
            },
            searchRecipes(name){
                let uri = 'http://localhost:8000/search/'+name;
                console.log(name);
                this.noBusqueda=false;
                this.search=name;
                axios.get(uri).then((response)=>{
                    if(response.data !== "No hay recetas con ese nombre")
                    {
                        this.recipes = response.data;    
                    }
                    else
                    {
                        this.recipes = [];
                    }
                });
            },
            updateRecipe(recipe) {
                this.noBusqueda=true;
                let id = recipe._id;
                let uri = 'http://localhost:8000/recipe/' + id;
                todo.editing = false;
                axios.post(uri, recipe).then((response) => {
                    console.log(response);
                    this.fetchRecipes();
                }).catch((error) => {
                    console.log(error);
                    this.fetchRecipes();
                });
            },
            deleteRecipe(id) {
                this.noBusqueda=true;
                let uri = 'http://localhost:8000/recipe/' + id;
                axios.delete(uri).then((response)=>{
                    this.fetchRecipes();
                });
            },
            listenToEvents() {
                bus.$on('reloadRecipes', ($event) => {
                    this.fetchRecipes(); 
                });
                bus.$on('search', (name)=> {
                    console.log("Pasa");
                    this.searchRecipes(name);
                });
            }
        }
    }
</script>    

<style scoped>
    .delete__icon {}
    .no_border_left_right {
        border-left: 0px;
        border-right: 0px;
    }
    .flat_form {
        border-radius: 0px;
    }
    .mrb-10 {
        margin-bottom: 10px;
    }
    .addon-left {
        background-color: none !important;
        border-left: 0px !important;
        cursor: pointer !important;
    }
    .addon-right {
        background-color: none !important;
        border-right: 0px !important;
    }
</style> 