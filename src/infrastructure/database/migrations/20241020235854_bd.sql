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


insert into Tenant (TenantID, Nome, InformacoesDeContato) values (1, 'PUCRS', 'pucrs@pucrs.com');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (2, 'UFRGS', 'ufrgs@ufrgs.com');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (3, 'Fanta', 'fanta@fanta.com');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (4, 'Muller', 'muller@muller.com');

INSERT INTO Tenant (TenantID, Nome, InformacoesDeContato) VALUES
    (5, 'Empresa Alpha', 'contato@alpha.com'),
    (6, 'Organização Beta', 'suporte@beta.org');

insert into Usuario (UserID, TenantID, Nome, Email) values (1, 1, 'Francisco', 'francisco@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (2, 1, 'Luis', 'luis@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (3, 2, 'Lucas', 'lucas@emai.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (4, 2, 'Joao', 'joao@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (5, 3, 'Maria', 'maria@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (6, 3, 'Ana', 'ana@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (7, 4, 'Pedro', 'pedro@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (8, 4, 'Paulo', 'paulo@email.com');

INSERT INTO Usuario (UserID, TenantID, Nome, Email) VALUES
    (9, 5, 'Alice Silva', 'alice.silva@alpha.com'),
    (10, 5, 'Bruno Souza', 'bruno.souza@alpha.com'),
    (11, 6, 'Carlos Lima', 'carlos.lima@beta.org'),
    (12, 6, 'Daniela Santos', 'daniela.santos@beta.org');


insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (1, 3, 'Festa do refri', 'Festa', 'Av da Festa 11', '2024-10-20 23:58:54');
insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (2, 4, 'Seminário de Negócios', 'Palestra', 'Rua amarela 24', '2024-10-29 11:58:54');
insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (3, 4, 'Festa do vinho', 'Festa', 'Rua do vinho 11', '2024-11-20 20:58:54');

INSERT INTO Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) VALUES
    (4, 5, 'Concerto de Rock', 'Show', 'Arena Alpha', '2023-12-15 20:00:00'),
    (5, 5, 'Peça de Teatro', 'Teatro', 'Teatro Central', '2024-01-10 19:30:00'),
    (6, 6, 'Exposição de Arte', 'Exposição', 'Galeria Beta', '2023-11-20 10:00:00');

INSERT INTO Ticket (TicketID, EventoID, TenantID, PrecoOriginal, IDDoVendedor, CodigoUnicoDeVerificacao, Status, Usado) VALUES
    (1, 4, 5, 150.00, 9, 'VERIFCODE123', 'disponivel', FALSE),
    (2, 4, 5, 150.00, 10, 'VERIFCODE124', 'disponivel', FALSE),
    (3, 5, 5, 80.00, 9, 'VERIFCODE125', 'disponivel', FALSE),
    (4, 6, 6, 50.00, 11, 'VERIFCODE126', 'disponivel', FALSE);


INSERT INTO Transacao (TransacaoID, TenantID, IDDoComprador, IDDoTicket, PrecoDeVenda, DataDaTransacao, StatusDaTransacao) VALUES
    (1, 5, 2, 1, 150.00, '2024-10-22 09:00:00', 'concluida'),
    (2, 6, 4, 4, 50.00, '2024-10-22 09:30:00', 'concluida');

INSERT INTO PreferenciasDeNotificacao (PreferenciasID, UserID, ReceberEmails) VALUES
    (1, 1, TRUE),
    (2, 2, FALSE),
    (3, 3, TRUE),
    (4, 4, FALSE),
    (5, 5, TRUE),
    (6, 6, FALSE),
    (7, 7, TRUE),
    (8, 8, TRUE),
    (9, 9, TRUE),
    (10, 10, FALSE),
    (11, 11, FALSE),
    (12, 12, TRUE);

INSERT INTO Avaliacao (AvaliacaoID, CompradorID, VendedorID, TransacaoID, Nota, Comentario, Data) VALUES
    (1, 2, 9, 1, 5, 'Excelente vendedor! A transação foi rápida e tranquila.', '2024-10-22 10:00:00'),
    (2, 4, 11, 2, 4, 'Bom vendedor, mas a comunicação poderia ser melhor.', '2024-10-22 10:15:00');


INSERT INTO MovimentoFinanceiro (MovimentoID, UserID, Valor, DataMovimento, TipoMovimento, Descricao) VALUES
    (1, 9, 150.00, '2024-10-22 09:05:00', 'credito', 'Venda do ingresso 1'),
    (2, 2, -150.00, '2024-10-22 09:05:00', 'debito', 'Compra do ingresso 1'),
    (3, 11, 50.00, '2024-10-22 09:35:00', 'credito', 'Venda do ingresso 4'),
    (4, 4, -50.00, '2024-10-22 09:35:00', 'debito', 'Compra do ingresso 4');


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