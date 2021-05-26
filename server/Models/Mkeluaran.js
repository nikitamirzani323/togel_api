const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_keluaran_togel = con.define('agen_view_togel_keluaran', {
    idtrxkeluaran:sequilize.STRING,
    idcomppasaran:sequilize.STRING,
    idcompany:sequilize.STRING,
    keluaranperiode:sequilize.STRING,
    datekeluaran:sequilize.STRING,
    keluarantogel:sequilize.STRING
},{
  freezeTableName:true,
  timestamps:false
})
tbl_keluaran_togel.removeAttribute('id')
module.exports = tbl_keluaran_togel