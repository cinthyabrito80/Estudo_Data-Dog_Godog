#language: pt

@dogs
Funcionalidade: Comer Godogs
  Para ser feliz
  Como um gopher faminto
  Eu preciso ser capaz de comer godogs

  Cenário: coma 5 de 12
    Dado que existem 12 Godogs
    Quando eu como 5
    Então deve restar 7 Godogs