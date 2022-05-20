# airport-application

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
