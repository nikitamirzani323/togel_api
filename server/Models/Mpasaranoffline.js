const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_pasaran_offline = con.define('tbl_mst_company_game_pasaran_offline', {
    idcomppasaranoff:sequilize.INTEGER,
    idcomppasaran:sequilize.INTEGER,
    idcompany:sequilize.STRING,
    haripasaran:sequilize.STRING,
},{
  freezeTableName:true,
  timestamps:false
})
tbl_pasaran_offline.removeAttribute('id')
module.exports = tbl_pasaran_offline