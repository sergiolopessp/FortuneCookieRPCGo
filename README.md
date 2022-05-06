# Hands On | Criando um RPC Server em Go

E chegamos a mais um hands on, e vamos brincar aqui com o Go Lang e a criação de RPC Server e um Client para consumir a função e mostrar mais uma opção de desenvolvimento e mais um tecnica para vocês. Então vamos lá

Mas antes de começar a por a mão no código, vamos começar falando o que é o RPC. Remote Procedure Call, ou RPC, é um jeito alternativo ao uso do HTTP, onde a gente tem menos restrições que o REST, o que acaba trazendo mais performance e que vai permitir para você criar operação, tipicamente chamadas de "procedures" que um servidor provê para clientes, que neste caso aqui são outros programas, então basicamente são funções para outros programas consumirem.

Então, vamos a mão na massa ? Então vamos revisar o que a gente vai fazer aqui:

- Criar um RPC Server que vai hospedar, executar e entregar o resultado da função.
- Criar um Client que vai consumir essa função e usar o resultado dela para alguma coisa.

Então aqui, quem me segue já algum tempo, sabe que eu adoro brincar a aplicação que eu chamo de FortuneCookie, que é basicamente uma lista de frases, onde nesta lista a gente sorteia uma aleatoriamente e devolve a frase. E vamos fazer em Go Lang, e porque ? Primeiro porque tenho estudado a linguagem e segundo porque ela tem uma boa abstração para fazer isso com facilidade para você, e você irá entender seguindo o fio aqui porque.

**Pré-requisitos**

Para reproduzir esse hands-on você ira precisar de dois pré-requisitos para seguir, que são:

Go Lang (Estou usando a versão 1.18.1 aqui)

Uma IDE que suporte o Go. (Estou usando o VSCode com a Extensão de Go for Visual Studio Code)

Para ver todos os detalhes, acesse: [https://deviniciative.wordpress.com/2022/05/06/hands-on--criando-um-rpc-server-em-go/(abrir em uma nova aba)](https://deviniciative.wordpress.com/?p=2791)