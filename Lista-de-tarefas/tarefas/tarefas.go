package tarefas

import "fmt"

type dados struct {
	Tarefa string
	Data   string
	Hora   string
}

func InserirTarefas() []dados {
	var d []dados
	fmt.Print("Digite a tarefa, a seguir digite a data da tarefa e por último digite a hora a ser realizada: ")
	var tarefa string
	fmt.Scan(&tarefa)
	fmt.Print("a data deve conter os seguintes parâmetros - dia/mês/ano - :")
	var data string
	fmt.Scan(&data)
	fmt.Print("a hora deve conter os seguintes parâmetros - horas:minutos - :")
	var hora string
	fmt.Scan(&hora)
	d = append(d, dados{tarefa, data, hora})
	return d
}
