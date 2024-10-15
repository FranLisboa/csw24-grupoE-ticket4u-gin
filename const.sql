CREATE TABLE `Tenant` (
  `TenantID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `Nome` VARCHAR(255) NOT NULL,
  `InformaçõesDeContato` TEXT,
  `ConfiguraçõesEspecíficas` JSONB
);

CREATE TABLE `Usuario` (
  `UserID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `TenantID` INT NOT NULL,
  `Nome` VARCHAR(255) NOT NULL,
  `Email` VARCHAR(255) UNIQUE NOT NULL,
  `FirebaseToken` VARCHAR(255),
  `ConfiguraçõesDePrivacidadeID` INT
);

CREATE TABLE `Evento` (
  `EventoID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `TenantID` INT NOT NULL,
  `NomeDoEvento` VARCHAR(255) NOT NULL,
  `Tipo` VARCHAR(50),
  `Localização` TEXT,
  `DataEHora` TIMESTAMP
);

CREATE TABLE `Ticket` (
  `TicketID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `EventoID` INT NOT NULL,
  `TenantID` INT NOT NULL,
  `PreçoOriginal` DECIMAL(10,2) NOT NULL,
  `IDDoVendedor` INT NOT NULL,
  `CódigoÚnicoDeVerificação` VARCHAR(255) UNIQUE,
  `Status` VARCHAR(50) NOT NULL
);

CREATE TABLE `Transacao` (
  `TransacaoID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `TenantID` INT NOT NULL,
  `IDDoComprador` INT NOT NULL,
  `IDDoTicket` INT UNIQUE NOT NULL,
  `PreçoDeVenda` DECIMAL(10,2) NOT NULL,
  `DataDaTransação` TIMESTAMP NOT NULL DEFAULT (NOW()),
  `StatusDaTransação` VARCHAR(50) NOT NULL
);

CREATE TABLE `PreferênciasDeNotificação` (
  `PreferênciasID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `UserID` INT NOT NULL,
  `ReceberEmails` BOOLEAN NOT NULL DEFAULT true,
  `ReceberPushNotifications` BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE `ConfiguraçõesDePrivacidade` (
  `ConfiguraçõesDePrivacidadeID` SERIAL PRIMARY KEY AUTO_INCREMENT,
  `UserID` INT NOT NULL,
  `PermitirCompartilhamentoDados` BOOLEAN NOT NULL DEFAULT false,
  `VisibilidadePerfil` VARCHAR(50) NOT NULL DEFAULT 'Privado',
  `HistóricoDeTransaçõesVisível` BOOLEAN NOT NULL DEFAULT false,
  `ReceberComunicaçõesMarketing` BOOLEAN NOT NULL DEFAULT true
);

ALTER TABLE `Usuario` ADD FOREIGN KEY (`TenantID`) REFERENCES `Tenant` (`TenantID`);

ALTER TABLE `Usuario` ADD FOREIGN KEY (`ConfiguraçõesDePrivacidadeID`) REFERENCES `ConfiguraçõesDePrivacidade` (`ConfiguraçõesDePrivacidadeID`);

ALTER TABLE `Evento` ADD FOREIGN KEY (`TenantID`) REFERENCES `Tenant` (`TenantID`);

ALTER TABLE `Ticket` ADD FOREIGN KEY (`EventoID`) REFERENCES `Evento` (`EventoID`);

ALTER TABLE `Ticket` ADD FOREIGN KEY (`TenantID`) REFERENCES `Tenant` (`TenantID`);

ALTER TABLE `Ticket` ADD FOREIGN KEY (`IDDoVendedor`) REFERENCES `Usuario` (`UserID`);

ALTER TABLE `Transacao` ADD FOREIGN KEY (`TenantID`) REFERENCES `Tenant` (`TenantID`);

ALTER TABLE `Transacao` ADD FOREIGN KEY (`IDDoComprador`) REFERENCES `Usuario` (`UserID`);

ALTER TABLE `Transacao` ADD FOREIGN KEY (`IDDoTicket`) REFERENCES `Ticket` (`TicketID`);

ALTER TABLE `PreferênciasDeNotificação` ADD FOREIGN KEY (`UserID`) REFERENCES `Usuario` (`UserID`);

ALTER TABLE `ConfiguraçõesDePrivacidade` ADD FOREIGN KEY (`UserID`) REFERENCES `Usuario` (`UserID`);
