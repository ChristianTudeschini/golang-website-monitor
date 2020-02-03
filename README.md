## Monitorador de Sites

###### Monitorador feito em Go que verifica se determinado site está ou não online.
---
## Instalação
Após clonar o repositório para seu dispositivo, vá na pasta ```/src``` do seu diretório Go e jogue a pasta ```/monitora-site``` dentro dela.

> ```Disco Local/Go/src/monitora-site```

## Executar
Para executar o projeto, abra o terminal no diretório onde o mesmo está localizado e dê o comando ```go run monitora-site.go```.

## Informações Adicionais
Para que a verificação dos sites seja realizada, é necessário escrever os sites que deseja monitorar no arquivo ```sites.txt```, da seguinte forma:
```
http(s)://www.nomedosite.com/org/etc
```
Se quiser verificar mais de um site, pule uma linha e coloque o outro desejado, Exemplo:
```
https://www.unionu.com.br
http://www.univoice.hol.es
```
