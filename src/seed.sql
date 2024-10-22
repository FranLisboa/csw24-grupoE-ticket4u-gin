insert into Tenant (TenantID, Nome, InformacoesDeContato) values (1, 'PUCRS', 'Faculdade');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (2, 'UFRGS', 'Faculdade');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (3, 'Fanta', 'Empresa');
insert into Tenant (TenantID, Nome, InformacoesDeContato) values (4, 'Muller', 'Empresa');

insert into Usuario (UserID, TenantID, Nome, Email) values (1, 1, 'Francisco', 'francisco@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (2, 1, 'Luis', 'luis@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (3, 2, 'Lucas', 'lucas@emai.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (4, 2, 'Joao', 'joao@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (5, 3, 'Maria', 'maria@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (6, 3, 'Ana', 'ana@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (7, 4, 'Pedro', 'pedro@email.com');
insert into Usuario (UserID, TenantID, Nome, Email) values (8, 4, 'Paulo', 'paulo@email.com');

insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (1, 3, 'Festa do refri', 'Festa', 'Av da Festa 11', '2024-10-20 23:58:54');
insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (2, 4, 'Seminário de Negócios', 'Palestra', 'Rua amarela 24', '2024-10-29 11:58:54');
insert into Evento (EventoID, TenantID, NomeDoEvento, Tipo, Localizacao, DataEHora) values (3, 4, 'Festa do vinho', 'Festa', 'Rua do vinho 11', '2024-11-20 20:58:54');