# airport-application


## Pasos para levantar la aplicacion

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


### Levantar base de datos con los vuelos precargados

Abrir consola mysql

``sudo mysql ``

Crear base de datos 

``mysql>  create database airport``

Permitir que agrege tablas

``mysql>  use airport``

Correr script "TODO.sql" para cargar las tablas de la base de datos

``mysql>  source /path/to/TODO.sql``


## Endpoint de ejemplo

Descripcion del endpoint

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
