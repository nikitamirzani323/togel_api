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
    nmpasarantogel:sequilize.STRING,
    limit_togel_4d:sequilize.DOUBLE,
    limit_togel_3d:sequilize.DOUBLE,
    limit_togel_2d:sequilize.DOUBLE,
    limit_togel_2dd:sequilize.DOUBLE,
    limit_togel_2dt:sequilize.DOUBLE,
    limit_togel_colokbebas:sequilize.DOUBLE,
    limit_togel_colokmacau:sequilize.DOUBLE,
    limit_togel_coloknaga:sequilize.DOUBLE,
    limit_togel_colokjitu:sequilize.DOUBLE,
    limit_togel_5050umum:sequilize.DOUBLE,
    limit_togel_5050special:sequilize.DOUBLE,
    limit_togel_5050kombinasi:sequilize.DOUBLE,
    limit_togel_kombinasi:sequilize.DOUBLE,
    limit_togel_dasar:sequilize.DOUBLE,
    limit_togel_shio:sequilize.DOUBLE
},{
  freezeTableName:true,
  timestamps:false
})
tbl_pasaran.removeAttribute('id')
module.exports = tbl_pasaran