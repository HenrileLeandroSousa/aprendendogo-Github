package tarefas

type dados struct {
	Tarefa string
	Data   string
	Hora   string
	ID     int
}

func Dados() []dados {
	id := 1

	data := []dados{
		{Tarefa: "Acordar", Data: "14/05/23", Hora: "08:00", ID: id},
		{Tarefa: "Lanchar", Data: "14/05/23", Hora: "09:00", ID: id},
		{Tarefa: "Almoço", Data: "14/05/23", Hora: "12:00", ID: id},
		{Tarefa: "Estudar", Data: "14/05/23", Hora: "13:00", ID: id},
		{Tarefa: "café", Data: "14/05/23", Hora: "17:00", ID: id},
		{Tarefa: "Jogar", Data: "14/05/23", Hora: "18:00", ID: id},
		{Tarefa: "jantar", Data: "14/05/23", Hora: "20:00", ID: id},
		{Tarefa: "dormir", Data: "14/05/23", Hora: "23:00", ID: id},
	}

	return data

}
