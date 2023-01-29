# Instrucciones para lanzar App

## Iniciando Backend.

Dentro de la carpeta _API_ instalar las dependencias:
```
npm install
```

  Nota: En este caso usamos una base de datos local *PostgreSQL*. Asi es que debemos crear una Base de datos desde *psql* con el nombre: _food_.

Una vez instaladas las dependencias y haber creado nuestra base de datos, procedemos a levantar el servidor con el siguiente comando:
```
air
```
_RECUERDA dejar el servidor en una terminal y correr el cliente en una distinta_

## Creando la base de datos.

El servidor esta trabajando ahora en el puerto 3000. Haremos uso de *Postman* (o la herramienta de peticiones de su preferencia) para cargar datos.
Desde el servidor podemos crear categorias y productos... Siga los siguientes pasos:

- Crear una categoría con POST en la siguiente ruta junto con su body:

```
http://localhost:3000/categories

{
  categoName: "Fast food"
}
```
- Verificar la categoría creada con GET en la misma ruta:

```
http://localhost:3000/categories
```

- Para crear un producto usaremos POST en la ruta siguiente junto con su body:
```
http://localhost:3000/products

{
  "name": "Hot Dog",
  "img": "https://th.bing.com/th/id/OIP.dgIb8DNI04jDDWLo-XGpAwHaE8?pid=ImgDet&rs=1",
  "price": 35,
  "categoryId": 1
 }
```
  El valor de _categoryId_ debe ser el ID de la categoría que se creó anteriormente.

## Lanzar Cliente.

Dentro de la carpeta _client_ instalaremos las dependencias:
```
npm install
```

Para alojar el cliente de forma local usaremos
```
npm run serve
```

## Gracias por probar mi app :)