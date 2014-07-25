#goapi
=====

Primer mini api con Go

### Peticiones

Consiste en 8 peticiones que responden en formato json:

| URL | Descripción          |
| ------------- | ----------- |
| / | index |
| /login | login |
| /logout | logout |
| /add | agregar usuario* |
| /edit | editar usuario* |
| /list | lista de usuarios* |
| /get | ver usuario* |

*funciones solo habilitadas para usuarios logueados

### Respuesta

El formato de la respuesta es el siguiente:

```json
{
	"code"	:	1, 
	"msg"	:	"Mensaje",
	"data"	:	["algo", "otro"]
}
```

El código *code* que se responde puede ser 0 o 1, donde cero índica que no hubo éxito en la petición y uno índica que se realizo correctamente.

El mensaje *msj* está relacionado a la petición.

El campo *data* es opcional y varia su formato dependiendo de la petición que se haga.
