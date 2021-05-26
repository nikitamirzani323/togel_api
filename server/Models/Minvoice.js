const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_invoice = con.define('client_view_invoice', {
    idtrxkeluarandetail:sequilize.INTEGER,
    idtrxkeluaran:sequilize.INTEGER,
    idcompany:sequilize.STRING,
    datetimedetail:sequilize.DATE,
    typegame:sequilize.STRING
},{
  freezeTableName:true,
  timestamps:false
})
tbl_invoice.removeAttribute('id')
module.exports = tbl_invoice