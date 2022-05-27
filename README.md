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


### Levantar base de datos

Abrir consola mysql

``sudo mysql ``

Crear base de datos 

``mysql>  create database airport``

## Usar la aplicacion

Correr la aplicacion main.go para levantar las tablas.

Dentro del directorio cmd/app

``go run .``

Ejecutar el script "load_data.sql" para cargar las tablas donde "/path/to" es el path absoluto al archivo. 

Tambien podes usar una herramienta como mysql-workbench para ejecutar el script

``mysql>  source /path/to/load_data.sql``

Al terminar con el programa dropear todas las tablas con el script "drop_tables.sql"

``mysql>  source /path/to/drop_tables.sql``

## Endpoints

### Titulo de un endpoint

`POST http://localhost:8080/api/v1/booking`

<details>
<summary>Request Body</summary>
<pre>
{
    "nombreUsuario": "juan",
    "booking": {
        "desde": "2022-02-20",
        "personas": [
            {
                "dni": 38998262,
                "nombre": "juan",
            }
        ],
        "metodoPago":{
            "tipo": "CREDIT",
        }
    }
}
</pre>
</details>


<details>
<summary>Response</summary>
<pre>
{
	"id":1
    "nombreUsuario": "juan",
    "booking": {
        "desde": "2022-02-20",
        "personas": [
            {
                "dni": 38998262,
                "nombre": "juan",
            }
        ],
        "metodoPago":{
            "tipo": "CREDIT",
        }
    }
}
</pre>
</details>
