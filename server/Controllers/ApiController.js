const model = require('../Models/index')
const createError = require('http-errors')
const moment = require('moment')
const db = require('../Helpers/init_mysql_orm')
const { Op } = require('sequelize')
module.exports = {
    servicetoken: async (req,res,next) =>{
        try {
            const {token} = req.body
            if(!token) throw createError.BadRequest() 
            res.send({
                status: 200,
                token:token,
                member_username:'developer',
                member_company:'developer',
                member_credit:5000000,
            });
        } catch (error) {
            next(error)
        }
    },
    serviceinit: async (req,res,next) =>{
        try {
            const {member_company} = req.body
            if(!member_company) throw createError.BadRequest() 
            let result = await model.Mpasaran.findAll({
                where: {
                    idcompany: member_company
                }
            })
            const newrecord = []
            let date = new Date()
            let thisDay = date.getDay()
            let myDays = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']
            const dateTime_now = moment().format('YYYY-MM-DD HH:mm:ss')
            const day_now = myDays[thisDay]
            let idcomppasaran = ''
            let idpasarantogel = ''
            let pasarandiundi = ''
            let pasaranurl = ''
            let statuspasaran = ''
            let nmpasarantogel = ''
            let datekeluaran = ''
			let periode = ''
            let jamtutup_format = ''
            let jamopen_format = ''
            let keluaran
            let pasaranOffline
            for await (const rec of result){
                idcomppasaran = rec.dataValues.idcomppasaran
                idpasarantogel = rec.dataValues.idpasarantogel
                nmpasarantogel = rec.dataValues.nmpasarantogel
                pasarandiundi = rec.dataValues.pasarandiundi
                pasaranjamtutup = rec.dataValues.jamtutup
                pasaranjamjadwal = rec.dataValues.jamjadwal
                pasaranjamopen = rec.dataValues.jamopen
                statuspasaran = rec.dataValues.statuspasaran
                
                keluaran = await model.Mkeluaran.findAll({
                    attributes: ['keluaranperiode','keluarantogel','datekeluaran'],
                    where: {
                        [Op.and]: [
                            {
                                idcompany:member_company,
                                idcomppasaran:idcomppasaran
                            }
                        ]
                    },
                    order: [['datekeluaran','desc']],
                    limit: 1
                })
                for await (const recdetail of keluaran){
                    datekeluaran = recdetail.dataValues.datekeluaran
                    periode = recdetail.dataValues.keluaranperiode
                    keluaran = recdetail.dataValues.keluarantogel
                }
                
                pasaranOffline = await model.Mpasaranoffline.findAll({
                    attributes: ['haripasaran'],
                    where: {
                        [Op.and]: [
                            {
                                idcompany:member_company,
                                idcomppasaran:idcomppasaran,
                                haripasaran:day_now.toLowerCase()
                            }
                        ]
                    },
                })
                if(pasaranOffline.length > 0){
                    jamtutup_format = moment().format('YYYY-MM-DD')+' '+pasaranjamtutup;
                    jamopen_format = moment().format('YYYY-MM-DD')+' '+pasaranjamopen;
                    if(parseInt(moment(dateTime_now).format('x')) >= parseInt(moment(jamtutup_format).format('x')) && parseInt(moment(dateTime_now).format('x')) <= parseInt(moment(jamopen_format).format('x'))){
						statuspasaran = 'OFFLINE';
					}
                }
                
                let data2 = {}
                data2.pasaran_id = idpasarantogel
                data2.pasaran_togel = nmpasarantogel
                data2.pasaran_url = pasaranurl
                data2.pasaran_tglkeluaran = datekeluaran
                data2.pasaran_periode = periode
                data2.pasaran_undi = pasarandiundi
                data2.pasaran_marketclose = pasaranjamtutup
                data2.pasaran_marketschedule = pasaranjamjadwal
                data2.pasaran_marketopen = pasaranjamopen,
                data2.pasaran_marketclose_format = jamtutup_format,
                data2.pasaran_marketopen_format = jamopen_format,
                data2.pasaran_status = statuspasaran
                newrecord.push(data2)
            }
            if(result.length > 0){
                res.send({
                    status: 200,
                    record: newrecord
                });
            }throw createError.NotFound()
        } catch (error) {
            next(error)
        }
    },
    servicecheckpasaran: async (req,res,next) =>{
        try {
            const {member_company,pasaran_code} = req.body
            if(!member_company) throw createError.BadRequest() 
            if(!pasaran_code) throw createError.BadRequest() 
            let result = await model.Mpasaran.findAll({
                attributes: ['idcomppasaran','nmpasarantogel','jamtutup','jamopen','statuspasaran'],
                where: {
                    [Op.and]: [
                        {
                            idcompany:member_company,
                            idpasarantogel:pasaran_code
                        }
                    ]
                },
            })
            let date = new Date()
            let thisDay = date.getDay()
            let myDays = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']
            const dateTime_now = moment().format('YYYY-MM-DD HH:mm:ss')
            const day_now = myDays[thisDay]
            let idcomppasaran = ''
            let statuspasaran = ''
            let nmpasarantogel = ''
            let idtrxkeluaran = ''
            let keluaranperiode = ''
            let keluaran
            let pasaranOffline
            for await (const rec of result){
                idcomppasaran = rec.dataValues.idcomppasaran
                nmpasarantogel = rec.dataValues.nmpasarantogel
                pasaranjamtutup = rec.dataValues.jamtutup
                pasaranjamjadwal = rec.dataValues.jamjadwal
                pasaranjamopen = rec.dataValues.jamopen
                statuspasaran = rec.dataValues.statuspasaran
                
                keluaran = await model.Mkeluaran.findAll({
                    attributes: ['idtrxkeluaran','keluaranperiode'],
                    where: {
                        [Op.and]: [
                            {
                                idcompany:member_company,
                                idcomppasaran:idcomppasaran,
                                keluarantogel:''
                            }
                        ]
                    },
                    limit: 1
                })
                for await (const recdetail of keluaran){
                    idtrxkeluaran = recdetail.dataValues.idtrxkeluaran
                    keluaranperiode = recdetail.dataValues.keluaranperiode
                }

                pasaranOffline = await model.Mpasaranoffline.findAll({
                    attributes: ['haripasaran'],
                    where: {
                        [Op.and]: [
                            {
                                idcompany:member_company,
                                idcomppasaran:idcomppasaran,
                                haripasaran:day_now.toLowerCase()
                            }
                        ]
                    },
                })
                if(pasaranOffline.length > 0){
                    let jamtutup = moment().format('YYYY-MM-DD')+' '+pasaranjamtutup;
                    let jamopen = moment().format('YYYY-MM-DD')+' '+pasaranjamopen;
                    if(parseInt(moment(dateTime_now).format('x')) >= parseInt(moment(jamtutup).format('x')) && parseInt(moment(dateTime_now).format('x')) <= parseInt(moment(jamopen).format('x'))){
						statuspasaran = 'OFFLINE';
					}
                }
            }
            
            if(result.length > 0){
                res.send({
                    status: 200,
                    pasaran_idtransaction: idtrxkeluaran,
                    pasaran_name: nmpasarantogel,
                    pasaran_periode: keluaranperiode,
                    pasaran_idcomp: idcomppasaran,
                    pasaran_status: statuspasaran,
                });
            }throw createError.NotFound()
        } catch (error) {
            next(error)
        }
    },
    serviceresult: async (req, res, next) => {
        try {
            const {member_company,pasaran_code} = req.body
            if(!member_company) throw createError.BadRequest() 
            if(!pasaran_code) throw createError.BadRequest()
            let result = await model.Mkeluaran.findAll({
                attributes: ['idpasarantogel','nmpasarantogel','datekeluaran','keluaranperiode','keluarantogel'],
                where: {
                    [Op.and]: [
                        {
                            idcompany:member_company,
                            idpasarantogel:pasaran_code,
                            keluarantogel: {
                                [Op.ne]: ''
                            }
                        }
                    ]
                },
                order: [['datekeluaran','desc']],
            })
            const newrecord = []
            let no = 0
            for await (const rec of result){
                no = no + 1
                let data2 = {}
                data2.no = no
                data2.date = moment(rec.dataValues.datekeluaran).format('DD MMMM YYYY')
                data2.periode = pasaran_code+'-'+rec.dataValues.keluaranperiode
                data2.result = rec.dataValues.keluarantogel
                newrecord.push(data2)
            }
            if(result.length > 0){
                res.send({
                    status: 200,
                    record: newrecord
                });
            }throw createError.NotFound()
        } catch (error) {
            next(error)
        }
    }
}