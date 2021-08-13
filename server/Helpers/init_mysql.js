const mysql = require('mysql')
const con = mysql.createConnection({
    host: 'localhost',
    port: 3306,
    user: 'root',
    password: '',
    database: 'db_togel2'
  })

con.connect(function (err){
    if(err) throw err;
    console.log("Mysql connect")
})


module.exports = con