<template>
    <div>
        <div v-show="noBusqueda">
            <hr>
            <h3 v-show="recipes.length>0">
                Mira las recetas:
            </h3>
            <div v-show="recipes.length==0">
                <h3>
                    No hay recetas aún, ¡agrega una!
                </h3>
            </div>
        </div>
        <div v-show="!noBusqueda">
            <b-container>
                <hr>
                <div>
                    <h3 v-show="recipes.length==0">
                        No hay recetas que coincidan con "{{search}}"
                    </h3>
                    <h3 v-show="recipes.length>0">
                        Mira las recetas que coinciden con "{{search}}"
                    </h3>
                </div>
                <b-button type="dark" v-on:click="fetchRecipes()">
                    Mostrar todas
                </b-button>
                <hr>
            </b-container>
        </div>

        <div v-show="recipes.length>0">
            <b-row id="accordion">
                <b-col md="6" v-for="recipe in recipes">
                    <b-card no-body class="mb-1">
                        <b-card-header v-show="idEditar!==recipe.id" header-tag="header" class="p-1" role="tab" v-on:click="reviewHeaderClick()">
                            <b-container>
                                <h5 href="#" v-b-toggle="'accordion'+recipe.id">
                                    {{recipe.name}}: {{recipe.description}}
                                </h5>
                            </b-container>
                        </b-card-header>
                        <b-card-header v-show="idEditar===recipe.id" header-tag="header" class="p-1" role="tab">
                            <b-container>
                                <b-row>
                                    <b-col cols="6">
                                        <label for="recipeName">
                                            Editar Nombre:
                                        </label>
                                        <b-input id="recipeName" v-bind:placeHolder="recipe.name" v-model="name"/>
                                    </b-col>
                                    <b-col cols="6">
                                        <label for="recipeDescription">
                                            Editar Descripción:
                                        </label>
                                        <b-input id="recipeDescription" v-bind:placeholder="recipe.description" v-model="description"/>
                                    </b-col>
                                </b-row>
                            </b-container>
                        </b-card-header>
                        <b-collapse v-bind:id="'accordion'+recipe.id" visible accordion="my-accordion" role="tabpanel">
                            <b-card-body v-show="idEditar!==recipe.id">
                                <b-row>
                                    <b-col md="8">
                                        <h6>
                                            Ingredientes:
                                        </h6>
                                        <p class="card-text">
                                            {{recipe.ingredients}}
                                        </p>
                                        <h6>
                                            Instrucciones:
                                        </h6>
                                        <p class="card-text">
                                        {{ recipe.instructions }}
                                        </p>
                                        <hr>
                                    </b-col>
                                    <b-col md="4">
                                        <b-img center fluid thumbnail v-bind:src="recipe.imgURL"/>
                                    </b-col>
                                    <hr>
                                    <b-container>
                                        <b-button variant="primary" v-on:click="updateSpecificRecipe(recipe)"> 
                                            Editar receta
                                        </b-button>
                                        <b-button variant="danger" v-on:click="deleteRecipe(recipe.id)">
                                            Eliminar receta
                                        </b-button>
                                    </b-container>
                                </b-row>
                            </b-card-body>
                            <b-card-body v-show="idEditar===recipe.id">
                                <b-row>
                                    <b-container>
                                        <h6>
                                            Editar ingredientes:
                                        </h6>
                                        <b-input v-bind:placeholder="recipe.ingredients" v-model="ingredients"/>
                                        <h6>
                                            Editar instrucciones:
                                        </h6>
                                        <b-input v-bind:placeholder="recipe.instructions" v-model="instructions"/>
                                        <h6>
                                            Editar imagen:
                                        </h6>
                                        <b-input v-bind:placeholder="recipe.imgURL" v-model="imgURL"/>
                                        <hr>
                                    </b-container>
                                    <hr>
                                    <b-container>
                                        <b-button variant="primary" v-on:click="updateRecipe(recipe)"> 
                                            Actualizar receta
                                        </b-button>
                                        <b-button variant="danger" v-on:click="cancelEdit()"> 
                                            Cancelar
                                        </b-button>
                                    </b-container>
                                </b-row>
                            </b-card-body>
                        </b-collapse>
                    </b-card>
                </b-col>
            </b-row>
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
                search:'',
                idEditar:0,
                name:'',
                description:'',
                instructions:'',
                imgURL:'',
                ingredients:''
            }
        },
        created: function() {
            this.fetchRecipes();
            this.listenToEvents();
        },
        methods: {
            fetchRecipes() {
                this.recipes = [];
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
                this.recipes = [];
                let uri = 'http://localhost:8000/search/'+name;
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
                let uri = 'http://localhost:8000/recipe/' + recipe.id;
                let nameU = this.name==="" ? recipe.name : this.name;
                let imgURLU = this.imgURL==="" ? recipe.imgURL : this.imgURL;
                let descriptionU = this.description==="" ? recipe.description : this.description;
                let ingredientsU = this.ingredients==="" ? recipe.ingredients : this.ingredients;
                let instructionsU = this.instructions==="" ? recipe.instructions : this.instructions;
                let recipeU = {name:nameU, imgURL:imgURLU, description:descriptionU, ingredients:ingredientsU,instructions:instructionsU};
                axios.post(uri, recipeU).then((response) => {
                    this.fetchRecipes();
                }).catch((error) => {
                    console.log(error);
                    this.fetchRecipes();
                });
                this.cancelEdit();
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
                    this.searchRecipes(name);
                });
            },
            updateSpecificRecipe(recipe)
            {
                this.idEditar = recipe.id;
            },
            cancelEdit()
            {
                this.idEditar=0;
                this.name="";
                this.description="";
                this.instructions="";
                this.imgURL="";
                this.ingredients="";
            },
            reviewHeaderClick()
            {
                if(this.idEditar !== 0)
                {
                    this.cancelEdit();
                }
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