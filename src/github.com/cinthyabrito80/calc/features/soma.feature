#language: pt

Funcionalidade: Soma

    # Cenario: Soma de valores
    #     Quando faço a soma de 2 + 2
    #     Então o resultado deve ser 6

    Esquema do Cenario: Soma de valores
        Quando faço a soma de <v1> + <v2>
        Então o resultado deve ser <result>

        Exemplos:
        |v1|v2|result|
        |2 |3 |5     |
        |6 |6 |12    |
        |2 |2 |4     |