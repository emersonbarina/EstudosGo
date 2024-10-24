package repositories

import (
	"api/source/models"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um nova instância de repositório publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere um publicaçao no banco de dados
func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) value (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPublicacao traz uma única publicação
func (repositorio Publicacoes) BuscarPublicacao(ID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select 
			p.*, u.nick 
		from 
			publicacoes p
			inner join usuarios u on u.id = p.autor_id
		where
		  p.id = ?`,
		ID,
	)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}
