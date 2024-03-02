-- Tabela enderecos
CREATE TABLE enderecos (
    id SERIAL PRIMARY KEY,
    endereco VARCHAR(200),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    estado VARCHAR(50),
    cep VARCHAR(10),
    fk_id_pessoa INTEGER REFERENCES pessoas(id)
);