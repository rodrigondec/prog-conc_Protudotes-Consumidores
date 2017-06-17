package etapa2

import (
	"sync"
	"time"
	"fmt"
	"strconv"
	"os"
)

const SEPARADOR = " / "

//estrutura que representa um pedido
type Pedido struct {
	id int //identificador
	dados string
}

var wg sync.WaitGroup //cria grupo de espera

/*gorotina consumidora que consumira um canal
bufferizado com 5000 pedidos*/
func consumidor (ch chan Pedido, n int) {
	for p := range ch {
		horario_inicio := time.Now()
		time.Sleep(500 * time.Millisecond)
		horario_termino := time.Now()

		fmt.Println("Agente: " + strconv.Itoa(n) + SEPARADOR +
			"Pedido: " + strconv.Itoa(p.id) + SEPARADOR +
			"Inicio proc: " + horario_inicio.String() + SEPARADOR +
			"Termino proc: " + horario_termino.String() + SEPARADOR +
			"Duracao: " + horario_termino.Sub(horario_inicio).String())
	}
	wg.Done()
}

func main() {
	if len(os.Args) == 3 {
		TAMANHO_BUFFER, _ := strconv.Atoi(os.Args[1])
		QTD_CONSUMIDORES, _ := strconv.Atoi(os.Args[2])
		var p Pedido
		ch := make(chan Pedido, TAMANHO_BUFFER) //cria canal

		//loop adiciona pedidos no canal
		for i := 1; i <= TAMANHO_BUFFER; i++ {
			p = Pedido{i, "Dados do pedido #" + strconv.Itoa(i)}
			ch <- p
		}
		close(ch) //fecha o canal

		//executa todos os consumidores
		for i := 1; i <= QTD_CONSUMIDORES; i++ {
			wg.Add(1)
			go consumidor(ch, i)
		}

		//espera termino de execucao de todos os consumidores
		wg.Wait()
	} else {
		fmt.Println("Numero invalido de argumentos. Requer exatamente 2 parametros enviados:\n")
		fmt.Println("1 - Tamanho do buffer / quantidade de pedidos\n2 - Quantidade de gorotinas")
	}
}
