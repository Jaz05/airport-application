# airport-application

## Documentacion

http://localhost:8080/swagger/index.html

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

Ejecutar el script "create_insert_tables.sql" para cargar las tablas donde "/path/to" es el path absoluto al archivo.

Tambien podes usar una herramienta como mysql-workbench para ejecutar el script

``mysql>  source /path/to/create_insert_tables.sql``

**¡Importante!**

Se recomienda hacer este paso antes de cada nueva ejecucion para volver las tablas a su estado inicial

### Levantar server

Dentro del directorio cmd/app donde se encuentra main.go ejecutar

``$ go run .``

## Endpoints

### Realizar reserva de un pasaje

`POST http://localhost:8080/sales`

<details>
<summary>Request Body</summary>
<pre>
{
   "name":"Juan",
   "surname":"adsd",
   "dni":1783612,
   "seat_id":4
}
</pre>
</details>


<details>
<summary>Response</summary>
<pre>
{
	"sale_id":
}
</pre>
</details>
