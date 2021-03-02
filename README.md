# GO-VUEJS
```
Proyecto realizado con 
```


## Tecnologías
```
- Lenguaje: Go.
- Base de Datos: DGraph.
- API Router: chi (pronunciado kai).
- Interfaz: Vue.js & Vuetify.
```
## Instalación de la base de datos
```
Para este proyecto se utilizó la base de datos DGraph. 
Se puede utilizar Docker para correr la base de datos, la informacion se encuentra en el siguiente link: <a href='https://dgraph.io/downloads#code-dgraph-sdk-go'>link</a>.

En mi caso utilice la base de datos descargada en linux: <a href='https://dgraph.io/downloads#code-dgraph-sdk-go'>link</a>.

Luego de tener la base de datos instalada en linux se corre en una terminal:
```/usr/local/bin/dgraph alpha --lru_mb 1024```

en otra:
```/usr/local/bin/dgraph zero```

y en otra:
```/usr/local/bin/dgraph-ratel```

```

### Compilacion e inicio del proyecto
```
El proyecto cuenta con dos carpetas: backend y frontend.

Ir a la carpeta backend e instalar todas las dependencias del proyecto 
```

### Compiles and minifies for production
```
npm run build
```