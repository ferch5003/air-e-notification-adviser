# air-e-notification-adviser

El proyecto es para enviar correos de Notificaciones por aviso (Proceso de Recurso de Reposición) a 
algún correo en especifico. Tiene algunas variables de entorno para configurar del lado del servidor, 
habria que crear un archivo `.env` en la raiz delproyecto con las variables del `.env.local` para 
probar en local.

En el momento de modificar el `.env` a tener en cuenta:

CRON:
```
SEARCH_NIC_CRON='* 0,12 * * *' // Para recibir el correo en dos momentos del dia (12AM y 12PM).
```

NIC:
```
NIC=unknown // Puede ser el NIC del recibo (Si se escoge TIPO=1) o Número de documento (Si se escoge TIPO=2)
TIPO=unknown // 1 = NIC, 2 = Número de documento 
```

MAIL:
```
SMTP_FROM_ADDRESS=unknown // Desde que correo se envia.
FROM_MAIL=unknown // Desde que correo se envia (Se sugiere mismo que FROM_MAIL).
TO_MAIL=unknown // A que correo enviar.
```

El correo que se envia tiene este formato:

```
Subject: Notificación de Air-E aun sin responder
```

```
Este es la notificación de las 1986-01-01 00:00:00.000000000 -0000 -00

Estado: 0
Mensaje: no hay notificaciones con este nic
Historial: []
```

PSDT: Gran laburo Air-E, segui asi👍.
