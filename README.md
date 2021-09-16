# GoLangTP

##Trabajo Entregable GoLang
El trabajo consiste en desestructurar una entrada de tipo String, y a partir de ella generar una estructura (utilizando Go)

##Implementacion
Para el desarrollo se separa la entrada en las distintas partes de la estructura (caracteres 0 y 1 son el tipo, 2 y 3 son el tamaño, y el resto es el contenido o valor de la estructura)
Se realizan validaciones de tipo INT para los campos del tamaño, y para el valor/contenido sólo si el tipo es numérico (NN), y se valida que el tamaño del contenido sea el especificado.

##Ejecucion
Para la ejecucion, se debe clonar el repositorio en algun destino que permita ejecutar Go.
Luego abrir un terminal dentro de la carpeta clonada y ejecutar el archivo .go con un comando como el siguiente:

go run tp.go TX03TPF