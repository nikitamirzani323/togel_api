const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_keluarandetail_togel = con.define('tbl_trx_keluarantogel_detail', {
    idtrxkeluarandetail :sequilize.STRING,
    idtrxkeluaran:sequilize.STRING,
    idcompany:sequilize.STRING,
    datetimedetail:sequilize.DATE,
    ipaddress:sequilize.STRING,
    username:sequilize.STRING,
    typegame:sequilize.STRING,
    nomortogel:sequilize.STRING,
    posisitogel:sequilize.STRING,
    bet:sequilize.DOUBLE,
    diskon:sequilize.DOUBLE,
    win:sequilize.DOUBLE,
    kei:sequilize.DOUBLE,
    upline:sequilize.STRING,
    upline_ref:sequilize.DOUBLE,
    type_ref:sequilize.STRING,
    browsertogel:sequilize.STRING,
    devicetogel:sequilize.STRING,
    statuskeluarandetail:sequilize.STRING,
    createkeluarandetail:sequilize.STRING,
    createdatekeluarandetail:sequilize.DATE,
    updatekeluarandetail:sequilize.STRING,
    updatedatekeluarandetail:sequilize.DATE
},{
  freezeTableName:true,
  timestamps:false
})
tbl_keluarandetail_togel.removeAttribute('id')
module.exports = tbl_keluarandetail_togel