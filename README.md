# Documentación

[Documentación de la API de Autenticación](https://unique-billboard-24f.notion.site/API-Authentication-eddee6ad41d84e109c6bcd1d5b53e6fc?pvs=4)


# Descripción general del proyecto

La API tiene como objetivo gestionar usuarios y permitir la funcionalidad de inicio de sesión. Está desarrollada en lenguaje Go (Golang) y utiliza PostgreSQL como base de datos. Además, la aplicación se ejecuta utilizando Air para facilitar el desarrollo.

La API proporciona diferentes endpoints para registrar usuarios, iniciar sesión, eliminar usuarios, consultar usuarios y actualizar usuarios. Además, implementa medidas de seguridad como el cifrado de contraseñas y la autenticación basada en tokens.

Para el desarrollo de la API se ha seguido una arquitectura de microservicios, lo que permite una mayor modularidad y escalabilidad. Se han utilizado patrones de diseño como MVC (Modelo-Vista-Controlador) para organizar y estructurar el código de manera eficiente.

En cuanto a la base de datos, PostgreSQL ofrece una gran robustez y rendimiento, lo que garantiza la persistencia de los datos de manera segura y confiable.

En resumen, la API brinda una solución completa y segura para la gestión de usuarios y la funcionalidad de inicio de sesión.

# Instalación

## Instalación

La instalación es el proceso de configurar y preparar un sistema o programa para su uso. Puede implicar la descarga de archivos, la configuración de ajustes y la realización de pasos específicos según el software o el sistema que se esté instalando.

Durante el proceso de instalación, es importante seguir las instrucciones proporcionadas por el proveedor del software o del sistema. Esto asegurará que la instalación se realice correctamente y que el software o el sistema funcione de manera óptima.

Además, es necesario tener en cuenta los requisitos del sistema para asegurarse de que el hardware y el software sean compatibles. Esto garantizará un rendimiento adecuado y evitará posibles problemas durante la instalación y el uso posterior.

Recuerda siempre realizar copias de seguridad de tus datos antes de realizar cualquier instalación importante. Esto te ayudará a prevenir la pérdida de información en caso de algún problema durante el proceso de instalación.

¡Asegúrate de seguir las instrucciones específicas de instalación proporcionadas para el software o el sistema que estás utilizando!

## ¿Cómo instalar?

Para instalar y configurar el proyecto, sigue estos pasos:

- Asegúrate de tener instalado Go y PostgreSQL en tu sistema.
- Clona el repositorio de tu proyecto desde GitHub.

    ```bash
    git clone https://github.com/Nelson2017-8/apiAuthentication.git
    ```

- Abre una terminal y navega hasta el directorio raíz del proyecto.
- Configurar el Archivo `.env` con sus variables de base de datos
- Crear una base de datos en PostgreSQL y cargar el script que se encuentra en la raíz `db.sql`
- Ejecutar el comando:

```bash
go run serve.go
```

- Expondrá la API en el puerto 8080 de tu máquina local, ajustar archivo `.env` para cambiar la configuración del puerto según sea requerido.

Si necesita ejecutar el proyecto en `Live Reaload` use air con el comando:

# Configuración

## Configuración

Es importante familiarizarse con las opciones de configuración disponibles y realizar los ajustes necesarios para optimizar la experiencia de uso. La configuración puede ser una herramienta poderosa para adaptar el sistema a nuestras preferencias y necesidades específicas.

Recuerda que algunos cambios de configuración pueden requerir conocimientos más avanzados o tener consecuencias importantes en el funcionamiento del sistema. Siempre es recomendable tener precaución al realizar cambios y, en caso de duda, buscar información adicional o consultar con un experto.

¡Explora las opciones de configuración y personaliza tu sistema de acuerdo a tus necesidades y preferencias!

## ¿Cómo configurar?

Antes de utilizar la API, necesitarás configurar la conexión a la base de datos PostgreSQL. Sigue estos pasos:

1. Abre el archivo de configuración  `.env`  en tu editor de código.
2. Busca la sección de configuración de la base de datos y proporciona los detalles de conexión, como el nombre de host, puerto, nombre de la base de datos, nombre de usuario y contraseña.
3. Establece tus variables de configurar la url y puerto donde se conectan la API
4. Establece tu clave secreta para codificar los tokens JWT.
5. Guarda los cambios en el archivo de configuración.

<aside>
💡 **Nota:** Es importante tener en cuenta que en nuestro entorno de trabajo contamos con dos ambientes, uno de desarrollo (`dev`) y otro de producción (`prod`). Dependiendo de las necesidades y configuraciones requeridas, se deberá utilizar el entorno adecuado. Esto nos permite tener un control y seguimiento más efectivo durante el proceso de desarrollo y asegurar que las configuraciones se apliquen correctamente en el entorno de producción.

</aside>

# Uso

## Uso

Una vez que hayas instalado y configurado el proyecto, puedes comenzar a utilizar la API, para uso básico se utiliza **Postman**. El uso de este API es muy sencillo. Simplemente sigue las instrucciones a continuación para comenzar a disfrutar de sus beneficios:

1. Abre **Postman** en tu dispositivo.
2. Crea una nueva solicitud haciendo clic en el botón "Nueva solicitud".
3. Ingresa la URL de la API que deseas consumir en el campo de la URL.
4. Selecciona el método de solicitud adecuado, como GET, POST, PUT, o DELETE.
5. Agrega los parámetros necesarios a la solicitud, si es necesario.
6. Haz clic en el botón "Enviar" para enviar la solicitud a la API.
7. Verás la respuesta de la API en la sección de respuesta de Postman.

¡Disfruta de consumir la API y aprovecha sus beneficios!


# Estructura del proyecto

## Estructura del proyecto

La API en Golang y Postgres siguiendo el patrón de diseño MVC (Modelo-Vista-Controlador), se puede establecer una estructura de proyecto similar a la siguiente:

- **serve.go**: Archivo principal del proyecto que inicializa el servidor y maneja las rutas.
- **.env**: Archivo que contiene la configuración del proyecto, como la conexión a la base de datos y variables de entorno.
- **models**: Carpeta que alberga los modelos de datos que representan las tablas de la base de datos.
- **controllers**: Carpeta que contiene los controladores, encargados de recibir las solicitudes HTTP, procesar la lógica de negocio y enviar las respuestas adecuadas.
- **routes**: Carpeta que define las rutas y los controladores asociados a cada una de ellas.
- **validations**: Carpeta que contiene funciones para realizar validaciones de datos.
- **utils**: Carpeta para almacenar funciones de utilidad o bibliotecas personalizadas.

Esta estructura de proyecto MVC proporciona una separación clara de responsabilidades y facilita el mantenimiento y la escalabilidad del sistema.


# API y endpoints

## API y endpoints

Una API (Interfaz de Programación de Aplicaciones) es un conjunto de reglas y protocolos que permiten la comunicación entre diferentes aplicaciones. Los endpoints son las URL específicas a las que se envían las solicitudes para interactuar con una API.

A continuación se presentan los endpoints:

- **Iniciar Sesión**
    - URL: `/api/v1/users/login`
    - Método: POST
    - Headers:
        - Content-Type: application/json
    - Body (Raw):

        ```json
        {
        	"phone": "prueba",
        	"password": "prueba"
        }
        ```

    - Response 200 OK:
        - Content-Type: application/json
        - Body:

        ```json
        {
            "user": {
                "id": 1,
                "email": "prueba@fakemail.com",
                "name": "prueba",
                "phone": "prueba",
                "address": "prueba",
                "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2OTkyNDQ0NTN9.HOWVOnGDWJa1isKJsi3FspU3knPoGZtbxX8HgAylka0",
                "session_active": true,
                "token_type": "bearer"
            }
        }
        
        ```

- **Obtener usuarios**
    - URL: `/api/v1/users`
    - Método: GET
    - Headers:
        - Content-Type: application/json
    - Body (Raw):
        - No se requiere
    - Response 200 OK:
        - Content-Type: application/json
        - Body:

        ```json
        [
            {
                "id": 1,
                "email": "prueba@fakemail.com",
                "name": "prueba",
                "phone": "prueba",
                "address": "prueba",
                "session_active": true
            }
        ]
        
        ```

- **Crear usuario**
    - URL: `/api/v1/users`
    - Método: POST
    - Headers:
        - Content-Type: application/json
    - Authorization
        - Authorization: Bearer `{token}`
    - Body (Raw):

        ```json
        {
            "name": "Nelson",
            "phone": "+584162102539",
            "email": "nelson10@gmail.com",
            "password": "12345678",
            "address": "Zamora, Guarenas, Venezuela"
        }
        ```

    - Response 200 OK:
        - Content-Type: application/json
        - Body:

        ```json
        [
            {
                "id": 1,
                "email": "prueba@fakemail.com",
                "name": "prueba",
                "phone": "prueba",
                "address": "prueba",
                "session_active": true
            }
        ]
        
        ```

        <aside>
        💡 Nota: Se necesita un usuario activo para este `end-point` si el usuario esta eliminado o su token no es válido enviará un error `401 Unauthorized`

        </aside>

- **Actualizar** **usuario**
    - URL: `/api/v1/users/{id_user}`
    - Método: PUT
    - Headers:
        - Content-Type: application/json
    - Authorization
        - Authorization: Bearer `{token}`
    - Body (Raw):

        ```json
        {
            "name": "Nelson Poritllo",
            "phone": "+584242102539",
            "email": "nelson@gmail.com",
            "password": "87654321",
            "address": "Caracas, Distro Capital, Venezuela"
        }
        ```

    - Response 200 OK:
        - Content-Type: application/json
        - Body:

        ```json
        {
            "email": "nelson@gmail.com",
            "name": "Nelson Poritllo",
            "phone": "+584242102539",
            "address": "Caracas, Distro Capital, Venezuela",
            "password": "87654321"
        }
        
        ```

        <aside>
        💡 Nota: Se necesita un usuario activo para este `end-point` si el usuario esta eliminado o su token no es válido enviará un error `401 Unauthorized`

        </aside>

- **Consultar usuario**
    - URL: `api/v1/users/{id_user}`
    - Método: GET
    - Headers:
        - Content-Type: application/json
    - Authorization
        - Authorization: Bearer `{token}`
    - Body (Raw):
        - No se requiere
    - Response 200 OK:
        - Content-Type: application/json
        - Body:

        ```json
        {
            "id": 1,
            "email": "prueba@fakemail.com",
            "name": "prueba",
            "phone": "prueba",
            "address": "prueba",
            "password": "",
            "session_active": true,
            "document_type_id": "",
            "document_number": ""
        }
        ```

        <aside>
        💡 Nota: Se necesita un usuario activo para este `end-point` si el usuario esta eliminado o su token no es válido enviará un error `401 Unauthorized`

        </aside>

- **Eliminar usuario**
    - URL: `/api/v1/users/{id_user}`
    - Método: DELETE
    - Headers:
        - Content-Type: application/json
    - Authorization
        - Authorization: Bearer `{token}`
    - Body (Raw):
        - No se requiere
    - Response 204 No content:

<aside>
💡 Recuerda reemplazar `{id_user}` con el ID real del usuario al realizar las solicitudes a estos endpoints y `{token}` el token que te devuelve el end-point de incio de sesión

</aside>


