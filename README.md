# airport-application

## Prerequisitos

### Instalar GO

https://go.dev/doc/install

En Ubuntu 20.04

``$ sudo apt update``

``$ sudo apt install golang``

### Instalar MySQL

https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/

En Ubuntu 20.04

``$ sudo apt update``

``$ sudo apt install mysql-server``

## Usar la aplicacion

### Levantar base de datos

Abrir consola mysql

``sudo mysql ``

Ejecutar el script "create_insert_tables.sql" para cargar las tablas, donde "/path/to" es el path absoluto al archivo.

Tambien se puede usar una herramienta como mysql-workbench para ejecutar el script

``mysql>  source /path/to/create_insert_tables.sql``

**¡Importante!**

Se recomienda hacer este paso antes de cada nueva ejecucion para volver las tablas a su estado inicial

### Levantar server

Dentro del directorio cmd/app donde se encuentra main.go ejecutar

``$ go run .``

## Endpoints

Documentación de los endpoints

http://localhost:8080/swagger/index.html

## Casos de Uso

### Obtener vuelos
Utilizando `GET /flights` se obtienen los vuelos.

### Obtener asientos
Utilizando `GET /seats` se obtienen los asientos correspondientes al los vuelos entre destinos.

### Reservar asientos
Utilizando `POST /sales` se reservan multiples asientos.
Los asientos quedan en estado **reservado** por un tiempo determinado (actualmente 5 minutos).

### Pagar reserva
Para completar la reserva, se debe realizar el pago de la misma.
Se utiliza el Token de la reserva en un request `POST /payments`

### Obtener reservas
Utilizando `GET /sales` se obtienen las reservas realizadas que esten asociadas al Token.


## Pruebas con Postman

Para poder hacer una prueba rapida se recomienda usar el Runner de la colección de Postman

Importar la colección `airport-applicaton.postman_collection.json` y
seleccionar la opcion Run. Esto hará que se ejecuten las request a todos los 
endpoints en el mismo orden que aparece en la sección Casos De Uso.
