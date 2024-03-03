-- Tabela contatos
CREATE TABLE contatos (
    id SERIAL PRIMARY KEY,
    numero_telefone VARCHAR(20),
    fk_id_pessoa INTEGER REFERENCES pessoas(id)
);