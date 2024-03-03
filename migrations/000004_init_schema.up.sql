-- Tabela contatos
CREATE TABLE contatos (
    id SERIAL PRIMARY KEY,
    numero_telefone VARCHAR(20),
    fk_id_conta INTEGER REFERENCES contas(id)
);