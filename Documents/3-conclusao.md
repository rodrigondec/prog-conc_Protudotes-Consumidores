# Conclusão

Analisando-se os tempos médios obtidos na etapa 3, a implementação da etapa 2 foi ligeiramente mais rápida. Ela é provavelmente mais fácil para manutenção uma vez que não necessita do uso de mecanismos de sincronização explícitos. Note que a análise foi um pouco prejudica pois não foi possível realizar o benchmarking da etapa 3 com 500ms \(devido ao prazo de entrega\). Entretanto, como o tempo de execução das duas etapas foi muito parecido, é muito provável que, ao utilizar 500 ms para simular o processamento, a melhor abordagem seja usando 500 gorotinas \(igual à etapa2\).

Podemos perceber que assim que produtores são adicionados para ir produzindo on-the-go, o programa para de ter ganho de desempenho no tempo acima da quantidade de 500 consumidores/produtores. Estagnando em 6~7 segundos de execução para as quantidades de 500, 1000 e 5000 nas etapas 2 e 3.

