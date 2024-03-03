package utils_date

import (
	"errors"
	"time"
)

// formatando data de nascimento para o padrão do banco de dados postgres
func Formatar_Data_Para_Padrao_Postgres(data string) (string, error) {
	dataNascimento, err := time.Parse("02-01-2006", data)
	if err != nil {
		return "", errors.New("erro ao tentar converter string em data")
	}
	// Formato YYYY-MM-DD
	dataNascimentoFormatada := dataNascimento.Format("2006-01-02")
	return dataNascimentoFormatada, nil

}
