const sqlite = require("sqlite3").verbose();
const db = new sqlite.Database("mailingList.db");

db.serialize()