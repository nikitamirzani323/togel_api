const sequilize = require('sequelize')
const con = require('../Helpers/init_mysql_orm')

const tbl_admin = con.define('tbl_admin', {
  username:sequilize.STRING,
  password:sequilize.STRING
},{
  freezeTableName:true,
  timestamps:false
})
tbl_admin.removeAttribute('id')
module.exports = tbl_admin