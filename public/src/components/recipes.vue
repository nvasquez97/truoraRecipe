<template>
    <div>
        <div v-show="recipes.length>0">
            <hr>
            <h3>Mira las recetas:</h3>
            <div id="accordion">
            <div class="col-md-6" v-for="recipe in recipes">
        <div class="card">
        <div class="card-header" id="headingOne" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne" >
        <h5 class="mb-0">
            <p><a href="#" style="color:black">
            {{recipe.name}}
            </a>
            {{recipe.description}}
            </p>
        </h5>
        </div>

        <div id="collapseOne" class="collapse show" aria-labelledby="headingOne" data-parent="#accordion">
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
                </div>
            </div>
        </div>
    </div>
    </div>
        </div>
        <div class="row alert alert-info text-center" v-show="recipes.length==0">
            <p class="alert alert-info">
              <strong>No tienes recetas aún</strong>
            <br/>
            ¡Agrega una!</p>
        </div>
    </div>
</div>
</template>

<script>
    import axios from 'axios';
    import bus from './../bus.js'

    export default {
        data() {
            return {
                recipes: []
            }
        },
        created: function() {
            this.fetchRecipes();
            this.listenToEvents();
        },
        methods: {
            fetchRecipes() {
                let uri = 'http://localhost:8000/recipes/';
                axios.get(uri).then((response) => {
                    this.recipes = response.data;
                });
            },
            updateRecipe(recipe) {
                let id = recipe._id;
                let uri = 'http://localhost:8000/recipe/' + id;
                todo.editing = false;
                axios.post(uri, recipe).then((response) => {
                    console.log(response);
                }).catch((error) => {
                    console.log(error);
                });
                this.fetchRecipes();
            },
            deleteRecipe(id) {
                let uri = 'http://localhost:8000/recipe/' + id;
                axios.delete(uri);
                this.fetchRecipes();
            },
            listenToEvents() {
                bus.$on('reloadRecipes', ($event) => {
                    this.fetchTodo(); 
                })
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