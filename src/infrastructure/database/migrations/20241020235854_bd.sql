-- +goose Up
-- +goose StatementBegin

CREATE TABLE Tenant (
    TenantID SERIAL PRIMARY KEY,
    Nome VARCHAR(255) NOT NULL,
    InformacoesDeContato TEXT
);

CREATE TABLE Usuario (
    UserID SERIAL PRIMARY KEY,
    TenantID INTEGER NOT NULL REFERENCES Tenant(TenantID),
    Nome VARCHAR(255) NOT NULL,
    Email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE Evento (
    EventoID SERIAL PRIMARY KEY,
    TenantID INTEGER NOT NULL REFERENCES Tenant(TenantID),
    NomeDoEvento VARCHAR(255) NOT NULL,
    Tipo VARCHAR(100),
    Localizacao VARCHAR(255),
    DataEHora TIMESTAMP NOT NULL
);

CREATE TABLE Ticket (
    TicketID SERIAL PRIMARY KEY,
    EventoID INTEGER NOT NULL REFERENCES Evento(EventoID),
    TenantID INTEGER NOT NULL REFERENCES Tenant(TenantID),
    PrecoOriginal NUMERIC(10,2) NOT NULL,
    IDDoVendedor INTEGER NOT NULL REFERENCES Usuario(UserID),
    CodigoUnicoDeVerificacao VARCHAR(255) UNIQUE NOT NULL,
    Status VARCHAR(20) NOT NULL CHECK (Status IN ('disponivel', 'reservado', 'vendido', 'usado')),
    Usado BOOLEAN DEFAULT FALSE
);

CREATE TABLE Transacao (
    TransacaoID SERIAL PRIMARY KEY,
    TenantID INTEGER NOT NULL REFERENCES Tenant(TenantID),
    IDDoComprador INTEGER NOT NULL REFERENCES Usuario(UserID),
    IDDoTicket INTEGER UNIQUE NOT NULL REFERENCES Ticket(TicketID),
    PrecoDeVenda NUMERIC(10,2) NOT NULL,
    DataDaTransacao TIMESTAMP DEFAULT NOW(),
    StatusDaTransacao VARCHAR(20) NOT NULL CHECK (StatusDaTransacao IN ('pendente', 'concluida', 'cancelada', 'reembolsada'))
);

CREATE TABLE PreferenciasDeNotificacao (
    PreferenciasID SERIAL PRIMARY KEY,
    UserID INTEGER UNIQUE NOT NULL REFERENCES Usuario(UserID),
    ReceberEmails BOOLEAN DEFAULT TRUE
);

CREATE TABLE Avaliacao (
    AvaliacaoID SERIAL PRIMARY KEY,
    CompradorID INTEGER NOT NULL REFERENCES Usuario(UserID),
    VendedorID INTEGER NOT NULL REFERENCES Usuario(UserID),
    TransacaoID INTEGER UNIQUE NOT NULL REFERENCES Transacao(TransacaoID),
    Nota INTEGER NOT NULL CHECK (Nota BETWEEN 1 AND 5),
    Comentario TEXT,
    Data TIMESTAMP DEFAULT NOW()
);

CREATE TABLE MovimentoFinanceiro (
    MovimentoID SERIAL PRIMARY KEY,
    UserID INTEGER NOT NULL REFERENCES Usuario(UserID),
    Valor NUMERIC(10,2) NOT NULL,
    DataMovimento TIMESTAMP DEFAULT NOW(),
    TipoMovimento VARCHAR(20) NOT NULL CHECK (TipoMovimento IN ('credito', 'debito')),
    Descricao TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS MovimentoFinanceiro;
DROP TABLE IF EXISTS Avaliacao;
DROP TABLE IF EXISTS PreferenciasDeNotificacao;
DROP TABLE IF EXISTS Transacao;
DROP TABLE IF EXISTS Ticket;
DROP TABLE IF EXISTS Evento;
DROP TABLE IF EXISTS Usuario;
DROP TABLE IF EXISTS Tenant;

-- +goose StatementEnd
