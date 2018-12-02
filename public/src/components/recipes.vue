<template>
    <div>
        <div class="col-md-12" v-show="recipes.length>0">
            <h3>Mira las recetas:</h3>
            <div id="accordion">
            <div class="row mrb-10" v-for="todo in todos">
        <div class="card">
        <div class="card-header" id="headingOne" data-toggle="collapse" data-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne" >
        <h5 class="mb-0">
            <p><a href="#" style="color:black">
            Arroz con Pollo:
            </a>
            Arroz con pollo y vegateles con color natural gracias a la cúrcuma.
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
                        Arroz, Pollo, Vegetales
                        <br>
                        <br>
                        <h6>
                        Instrucciones:
                        </h6>
                        Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. 3 wolf moon officia aute, non cupidatat skateboard dolor brunch. Food truck quinoa nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua put a bird on it squid single-origin coffee nulla assumenda shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident. Ad vegan excepteur butcher vice lomo. Leggings occaecat craft beer farm-to-table, raw denim aesthetic synth nesciunt you probably haven't heard of them accusamus labore sustainable VHS.
                    </div>
                    <div class="col-md-4">
                        <img class="img-fluid img-thumbnail" src="https://www.divinacocina.es/wp-content/uploads/ARROZ-CON-POLLO-2.jpg"/>
                    </div>
                </div>
            </div>
        </div>
    </div>
    </div>
        </div>
        <div class="row alert alert-info text-center" v-show="todos.length==0">
            <p class="alert alert-info">
              <strong>All Caught Up</strong>
            <br/>
            No hay ninguna receta aún</p>
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
                todos: []
            }
        },
        created: function() { // get todo items and start listening to events once component is created
            this.fetchTodo();
            this.listenToEvents();
        },
        methods: {
            fetchTodo() {
                let uri = 'http://localhost:4000/api/all';
                axios.get(uri).then((response) => {
                    this.todos = response.data;
                });
            },
            updateTodo(todo) {
                let id = todo._id;
                let uri = 'http://localhost:4000/api/update/' + id;
                todo.editing = false;
                axios.post(uri, todo).then((response) => {
                    console.log(response);
                }).catch((error) => {
                    console.log(error);
                })
            },
            deleteTodo(id) { //delete todo item
                let uri = 'http://localhost:4000/api/delete/' + id;
                axios.get(uri);
                this.fetchTodo();
            },
            listenToEvents() {
                bus.$on('refreshTodo', ($event) => {
                    this.fetchTodo(); // referesh or update todo list on the page
                })
            }
        }
    }
</script>    

<style scoped>
    .delete__icon {}
    .todo__done {
        text-decoration: line-through !important
    }
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