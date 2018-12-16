# truoraRecipe
CRUD de recetas utilizando Golang, CrockroachDB para el backend y VueJS con bootstrap-vue para el frontend

## Instalar y correr el Backend:

#### CockRoachDB. 

Iniciar un cluster de CockRoachDB:

```console
cockroach start --insecure --listen-addr=localhost
```

Luego, conectarse al cluster desde un terminal:
```sh
cockroach sql --insecure --host=localhost:26257
```
Crear una base de datos llamada cook, un usuario llamado truora y darle acceso al usuario a la base de datos:

```sql
CREATE DATABASE COOK;
CREATE USER TRUORA;
GRANT ALL ON DATABASE COOK TO TRUORA;
```

#### Ejecutar el backend desarrollado en go
```sh
go run ./backend/main.go
```

## Instalar y correr el Frontend

#### Instalar los paquetes de NodeJS y correr node

Instalar los paquetes:
```sh
npm install
```

Hacer build con webpack
```sh
npm run build
```

Ejecutar node
```sh
npm start
```
