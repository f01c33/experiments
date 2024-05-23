#### Diagrama de Alocação
| pessoa\tarefa| T1 | T2 | T3 | T4 | T5 | T6 |
|:-------------|:--:|:--:|:--:|:--:|:--:|:--:|
| Caue         | x  | x  | x  | x  | x  | x  | 
| Maria        | x  | x  | x  | x  | x  | x  | 
| Walter       | x  | x  | x  | x  | x  | x  | 

#### Rede de Atividades
| tarefa | dependencia | duração |
|:------:|:-----------:|:-------:|
| T1     |             | 3       |
| T2     |             | 3       |
| T3     | T1,T2(M1)   | 2       |
| T4     | T3(M2)      | 1       |
| T5     | T4(M3)      | 6       |
| T6     | T5(M4)      | 1       |

T1 - Questionario, 

T2 - Protótipos, 

T3 - Brainstorm, 

T4 - Arquitetura, 

T5 - Implementação, 

T6 - Teste de Aceitação