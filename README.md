# Kubernetes Unapp Demonstration

Assets que utilizo para realizar as demonstrações de Kubernetes utilizando o IBM Cloud Kubernetes Services (IKS).

1. Abrir o terminal e efetuar o Login na IBM Cloud
```
ibmcloud login
```
2. Acesso ao cluster Kubernetes

ibmcloud ks cluster-config <nome_do_cluster>

3. Executar o comando export sugerido pelo comando anterior. 

4. Acessar o diretório unapp ou unapp-k8s e executar os comandos.


# Makefile

Alguns alias foram criados utilizando-se o arquivo Makefile. A execução de cada comando deverá ser realizado no diretório onde do Makefile.

1. Entrar no diretório unapp-k8s

cd unapp-k8s

2. Executar os comandos. Exemplo, comando para verificar os deployments, pods e serviços do Kubernetes (getAll):

make getAll

# Código Fonte do Aplicativo Unapp

1. Entrar no diretório unapp

2. Você poderá encontrar todo o código fonte (HTML, css, js, etc) neste diretório. 

3. Criei um arquivo DockerFile para que possa criar um Docker Container e executar a aplicação.

4. Executar o comando. Não esquecer do "ponto" no final:

docker build -t <nome_da_image> .

5. Verificar se a imagem foi criada corretamente:

docker images

6. Executar o Container:

docker run -d -p 8000:80 <nome_da_imagem>

7. Abrir um Navegador e digitar o endereço:

localhost:8000
