# truoraRecipe
CRUD de recetas utilizando Golang, CrockroachDB para el backend y VueJS con bootstrap-vue para el frontend

## Instalar y correr el Backend:

#### CockRoachDB. 

Iniciar un cluster de CockRoachDB:

```shell
cockroach start --insecure --listen-addr=localhost
```

Luego, crear una base de datos llamada cook, un usuario llamado truora y darle acceso al usuario a la base de datos:

```sql
CREATE DATABASE COOK
CREATE USER TRUORA
GRANT ALL ON DATABASE COOK TO TRUORA
```

#### Ejecutar el backend desarrollado en go
```
go run main.go
```

## Instalar y correr el Frontend

#### Instalar los paquetes de NodeJS y correr node

Instalar los paquetes:
```
npm install
```

Hacer build con webpack
```
npm run build
```

Ejecutar node
```
npm start
```
