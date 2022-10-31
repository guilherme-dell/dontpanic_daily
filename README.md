# Como jogar?

## Objetivo do jogo
Acerte a equação oculta do dia que tem como resposta o número 42. Para encontrar a solução misteriosa, você poderá utilizar números de 0 a 9 e operadores matemáticos(+, -, * e /) em cada campo.

`Atenção:`
  - Um exemplo de equação é:  1 6 8 / 0 4;
  - não entenderemos como equação operadores na primeira e/ou última posição, exemplo: + 6 8 /4 -;
  - ou dois operadores em seguida, exemplo: 1 6 / / 0 4.
  - Utilize todos os campos!

A cada tentativa válida, isto é, a cada tentativa que a equação resulte em 42, o jogo retornará dicas para te ajudar!

Mostre que você sabe matemática e ganhe este jogo!

## Dicas
Caso sua equação tenha como resultado 42, mas não seja a equação misteriosa, as seguintes dicas aparecerão na tela:

`C` = significa que o número ou operador está na posição correta.

`T` = significa que o caractere existe na equação oculta, mas não está na posição correta.

`X` = significa que o caractere não está na equação oculta.

Tente até acertar! Boa sorte!

## Uso
Para jogar o game primeiro execute o comando `docker-compose build` após o build ocorrer execute o `docker-compose up`.
Abra uma aba em seu navegador com o caminho: `localhost:5010`.

## Fluxograma para criação do jogo

![Fluxograma](https://user-images.githubusercontent.com/98241492/192639912-6f7f6c8b-c463-44ac-9068-a1eae5ca220a.png)

## Documentação da API
Acesse `localhost:5010/swagger/index.html`.
