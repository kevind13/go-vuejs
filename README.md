# GO-VUEJS

## Tecnologías

- Lenguaje: Go.
- Base de Datos: DGraph.
- API Router: chi (pronunciado kai).
- Interfaz: Vue.js & Vuetify.

## Instalación e inicio de la base de datos

Para este proyecto se utilizó la base de datos DGraph.
Se puede utilizar Docker para correr la base de datos, la información se encuentra en el siguiente link: <a href='https://dgraph.io/downloads#code-dgraph-sdk-go'>link</a>.

En el caso de este proyecto se usó la base de datos descargada en linux: <a href='https://dgraph.io/downloads#code-dgraph-sdk-go'>link</a>.

Luego de tener la base de datos instalada en linux se corre en una terminal:
 
```/usr/local/bin/dgraph alpha --lru_mb 1024```
 
en otra:
 
```/usr/local/bin/dgraph zero```
 
y en otra:
 
```/usr/local/bin/dgraph-ratel```
 
 
 
## Compilación e inicio del proyecto
 
El proyecto cuenta con dos carpetas: backend y frontend.
 
Ir a la carpeta backend e instalar todas las dependencias del proyecto:
 
```go install```
 
y se corre el proyecto:
 
```go run main.go```
 
Ir a la carpeta frontend e instalar todas las dependencias del proyecto:
 
```
npm install
```
 
y se compila e inicia el servidor:
```
npm run serve
```

## Uso del proyecto

TODO

