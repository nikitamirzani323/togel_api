const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_counter = con.define('tbl_counter', {
    idcounter  :sequilize.INTEGER,
    nmcounter:sequilize.STRING,
    counter:sequilize.BIGINT
},{
  freezeTableName:true,
  timestamps:false
})
tbl_counter.removeAttribute('id')
module.exports = tbl_counter