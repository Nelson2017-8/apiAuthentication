# Imagen base con Go
FROM golang:latest

LABEL authors="Nelson"

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos necesarios al contenedor
COPY . .

# Instalar las dependencias de tu aplicación (si las tienes)

# Compilar tu aplicación
RUN go build -o serve .

# Exponer el puerto en el que se ejecuta tu aplicación
EXPOSE 8080

# Comando para ejecutar tu aplicación
CMD ["./serve"]

# docker build -t nombre_imagen .
# docker run -p 8080:8080 --name nombre_contenedor nombre_imagen
