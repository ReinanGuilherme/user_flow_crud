-- Tabela pessoas
CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    sobrenome VARCHAR(100),
    genero VARCHAR(10),
    data_nascimento DATE,
    fk_id_conta INTEGER REFERENCES contas(id)
);

