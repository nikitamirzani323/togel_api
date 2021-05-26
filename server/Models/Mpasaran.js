const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_pasaran = con.define('client_view_pasaran', {
    idcomppasaran:sequilize.STRING,
    idcompany:sequilize.STRING,
    idpasarantogel:sequilize.STRING,
    pasarandiundi:sequilize.STRING,
    pasaranlibur:sequilize.STRING,
    pasaranurl:sequilize.STRING,
    jamtutup:sequilize.STRING,
    jamjadwal:sequilize.STRING,
    jamopen:sequilize.STRING,
    statuspasaran:sequilize.STRING,
    nmpasarantogel:sequilize.STRING
},{
  freezeTableName:true,
  timestamps:false
})
tbl_pasaran.removeAttribute('id')
module.exports = tbl_pasaran