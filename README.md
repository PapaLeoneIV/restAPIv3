# Network 42 Historical Tracker

Questo è un progetto personale nato con l'intento di essere uno strumento utile all'interno del network della 42. L'idea è quella di avere uno storico di fatti, espressioni, e situazioni registrate all'interno delle diverse scuole.

## Struttura del Progetto

Il progetto è diviso in due parti: **Backend** e **Frontend**.

### Backend

Per il backend, ho utilizzato principalmente la standard library di Go, eccetto per il driver di PostgreSQL. Le feature implementate includono:

- Acquisizione dei dati necessari all'avviamento dei servizi tramite `.env`
- Creazione e gestione del database (PostgreSQL) tramite `sqlc`
- Creazione di un multiplexer per la gestione delle richieste in arrivo
- Creazione di endpoint CRUD
- Rate Limiting basato sull'IP
- Richiesta dei certificati per HTTPS

**Prossimi upgrade backend:**

- Altre misure di sicurezza (JWT, Autenticazione, Autorizzazione)

### Frontend

Per il frontend, ho utilizzato Next.js e Tailwind CSS per migliorare le mie competenze. Le feature implementate includono:

- Implementazione di alcuni componenti riutilizzabili
- Installazione di un unico modulo esterno: `terminal.css`
- Invio di richieste POST al backend per l'inserimento nel database

### Porte Utilizzate

- **Frontend:** `3000`
- **Backend:** `8443`

### Problemi Noti

- Il database viene inizializzato con una entry durante la prima creazione del container. È necessario chiuderlo e rilanciare di nuovo il container.

### Come Usarlo

## Dockerizzazione

L'intero progetto è stato dockerizzato. Per avviare il progetto, utilizzare il comando:

```bash
docker compose up
```



### Contribuire

Se desideri contribuire a questo progetto, segui questi passaggi:

1. Fai un fork del repository
2. Crea un branch per la tua feature (`git checkout -b feature/nome-feature`)
3. Fai commit delle tue modifiche (`git commit -m 'Aggiungi feature'`)
4. Pusha il tuo branch (`git push origin feature/nome-feature`)
5. Apri una Pull Request

Grazie per aver scelto di contribuire al progetto Network 42 Historical Tracker!

---
