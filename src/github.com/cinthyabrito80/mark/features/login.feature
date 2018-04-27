#language: pt

Funcionalidade: Login
    Para que eu possa ter acesso as minhas tarefas #Para quem estou fazendo
    Sendo um usuario #O que eu vou fazer
    Posso meu autenticar com os meus dados previamente cadastrados #Qual o meu valor de negocio que agrega

Contexto: Formulario de Login
    Dado que eu acessei um pagina principal

Cenario: Login do usuario
    Quando faço login com "eu@papito.io" e "12345"
    Então sou autenticado com sucesso

Cenario: Senha incorreta
    Quando faço login com "eu@papito.io" e "aaaaa"
    Então devo ver a seguinte mensagem "Senha inválida."

Cenario: Email incorreta
    Quando faço login com "eu@papito.com" e "12345"
    Então devo ver a seguinte mensagem "Email inválida."