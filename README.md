# bienestar_crud
API de gestión general de servicios de bienestar virtual

Integración con

 - `CI`
 - `Drone 1.x`
 - `bienestar_crud master/develop`

## Requerimientos
Go version >= 1.9.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/bienestar_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `BIENESTAR_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `BIENESTAR_CRUD__PGUSER`: Usuario de la base de datos
 - `BIENESTAR_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `BIENESTAR_CRUD__PGURLS`: Host de conexión
 - `BIENESTAR_CRUD__PGDB`: Nombre de la base de datos
 - `BIENESTAR_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
BIENESTAR_HTTP_PORT=9018 BIENESTAR_CRUD__PGUSER=user BIENESTAR_CRUD__PGPASS=password BIENESTAR_CRUD__PGURLS=localhost BIENESTAR_CRUD__PGDB=bd BIENESTAR_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/bienestar_crud/blob/master/modelo_bienestar_crud.png).
