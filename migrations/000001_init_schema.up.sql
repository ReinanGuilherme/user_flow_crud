-- Tabela contas
CREATE TABLE contas (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100),
    senha VARCHAR(50),
    ativo BOOLEAN,
    token VARCHAR(100),
    foto_perfil VARCHAR(200),
    data_ultima_sessao TIMESTAMP WITH TIME ZONE
);