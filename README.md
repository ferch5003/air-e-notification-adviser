# air-e-notification-adviser

## Proyecto en ALPHA

El dicho descansas como si chambearas se lo tomo muy en serio Air-E.

El proyecto es un notificador de reclamos para enviar la info al correo. Tiene algunas variables de
entorno para configurar del lado del servidor, habria que crear un archivo `.env` en la raiz del
proyecto con las variables del `.env.local` para probar en local.

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