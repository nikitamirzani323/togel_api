const model = require('../Models/index')
const createError = require('http-errors')
const moment = require('moment')
const db = require('../Helpers/init_mysql_orm')
const { Op } = require('sequelize')

module.exports = {
    serviceconfigtogel: async (req,res,next) =>{
        try {
            const {member_company, pasaran_code, permainan} = req.body
            if(!member_company) throw createError.BadRequest() 
            if(!pasaran_code) throw createError.BadRequest() 
            if(!permainan) throw createError.BadRequest() 
            let result = []
            switch(permainan){
                case '4-3-2':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['1_minbet','min_bet'],
                            ['1_maxbet4d','max4d_bet'],
                            ['1_maxbet3d','max3d_bet'],
                            ['1_maxbet2d','max2d_bet'],
                            ['1_maxbet2dd','max2dd_bet'],
                            ['1_maxbet2dt','max2dt_bet'],
                            ['1_disc4d','disc4d_bet'],
                            ['1_disc3d','disc3d_bet'],
                            ['1_disc2d','disc2d_bet'],
                            ['1_disc2dd','disc2dd_bet'],
                            ['1_disc2dt','disc2dt_bet'],
                            ['1_win4d','win4d_bet'],
                            ['1_win3d','win3d_bet'],
                            ['1_win2d','win2d_bet'],
                            ['1_win2dd','win2dd_bet'],
                            ['1_win2dt','win2dt_bet'],
                            ['1_limittotal4d','limittotal4d_bet'],
                            ['1_limittotal3d','limittotal3d_bet'],
                            ['1_limittotal2d','limittotal2d_bet'],
                            ['1_limittotal2dd','limittotal2dd_bet'],
                            ['1_limittotal2dt','limittotal2dt_bet'],
                            ['limitline_4d','limitline_4d'],
                            ['limitline_3d','limitline_3d'],
                            ['limitline_2d','limitline_2d'],
                            ['limitline_2dd','limitline_2dd'],
                            ['limitline_2dt','limitline_2dt']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'colokbebas':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['2_minbet','min_bet_colokbebas'],
                            ['2_maxbet','max_bet_colokbebas'],
                            ['2_disc','disc_bet_colokbebas'],
                            ['2_win','win_bet_colokbebas'],
                            ['2_limitotal','limittotal_bet_colokbebas']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'colokmacau':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['3_minbet','min_bet_colokmacau'],
                            ['3_maxbet','max_bet_colokmacau'],
                            ['3_disc','disc_bet_colokmacau'],
                            ['3_win2digit','win_bet_colokmacau'],
                            ['3_win3digit','win3_bet_colokmacau'],
                            ['3_win4digit','win4_bet_colokmacau'],
                            ['3_limittotal','limittotal_bet_colokmacau']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'coloknaga':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['4_minbet','min_bet_coloknaga'],
                            ['4_maxbet','max_bet_coloknaga'],
                            ['4_disc','disc_bet_coloknaga'],
                            ['4_win3digit','win_bet_coloknaga'],
                            ['4_win4digit','win4_bet_coloknaga'],
                            ['4_limittotal','limittotal_bet_coloknaga']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'colokjitu':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['5_minbet','min_bet_colokjitu'],
                            ['5_maxbet','max_bet_colokjitu'],
                            ['5_desic','disc_bet_colokjitu'],
                            ['5_winas','winas_bet_colokjitu'],
                            ['5_winkop','winkop_bet_colokjitu'],
                            ['5_winkepala','winkepala_bet_colokjitu'],
                            ['5_winekor','winekor_bet_colokjitu'],
                            ['5_limitotal','limittotal_bet_colokjitu']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case '5050umum':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['6_minbet','min_bet_5050umum'],
                            ['6_maxbet','max_bet_5050umum'],
                            ['6_keibesar','keibesar_bet_5050umum'],
                            ['6_keikecil','keikecil_bet_5050umum'],
                            ['6_keigenap','keigenap_bet_5050umum'],
                            ['6_keiganjil','keiganjil_bet_5050umum'],
                            ['6_keitengah','keitengah_bet_5050umum'],
                            ['6_keitepi','keitepi_bet_5050umum'],
                            ['6_discbesar','discbesar_bet_5050umum'],
                            ['6_disckecil','disckecil_bet_5050umum'],
                            ['6_discgenap','discgenap_bet_5050umum'],
                            ['6_discganjil','discganjil_bet_5050umum'],
                            ['6_disctengah','disctengah_bet_5050umum'],
                            ['6_disctepi','disctepi_bet_5050umum'],
                            ['6_limittotal','limittotal_bet_5050umum']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case '5050special':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['7_minbet','min_bet_5050special'],
                            ['7_maxbet','max_bet_5050special'],
                            ['7_keiasganjil','keiasganjil_bet_5050special'],
                            ['7_keiasgenap','keiasgenap_bet_5050special'],
                            ['7_keiasbesar','keiasbesar_bet_5050special'],
                            ['7_keiaskecil','keiaskecil_bet_5050special'],
                            ['7_keikopganjil','keikopganjil_bet_5050special'],
                            ['7_keikopgenap','keikopgenap_bet_5050special'],
                            ['7_keikopbesar','keikopbesar_bet_5050special'],
                            ['7_keikopkecil','keikopkecil_bet_5050special'],
                            ['7_keikepalaganjil','keikepalaganjil_bet_5050special'],
                            ['7_keikepalagenap','keikepalagenap_bet_5050special'],
                            ['7_keikepalabesar','keikepalabesar_bet_5050special'],
                            ['7_keikepalakecil','keikepalakecil_bet_5050special'],
                            ['7_keiekorganjil','keiekorganjil_bet_5050special'],
                            ['7_keiekorgenap','keiekorgenap_bet_5050special'],
                            ['7_keiekorbesar','keiekorbesar_bet_5050special'],
                            ['7_keiekorkecil','keiekorkecil_bet_5050special'],
                            ['7_discasganjil','discasganjil_bet_5050special'],
                            ['7_discasgenap','discasgenap_bet_5050special'],
                            ['7_discasbesar','discasbesar_bet_5050special'],
                            ['7_discaskecil','discaskecil_bet_5050special'],
                            ['7_disckopganjil','disckopganjil_bet_5050special'],
                            ['7_disckopgenap','disckopgenap_bet_5050special'],
                            ['7_disckopbesar','disckopbesar_bet_5050special'],
                            ['7_disckopkecil','disckopkecil_bet_5050special'],
                            ['7_disckepalaganjil','disckepalaganjil_bet_5050special'],
                            ['7_disckepalagenap','disckepalagenap_bet_5050special'],
                            ['7_disckepalabesar','disckepalabesar_bet_5050special'],
                            ['7_disckepalakecil','disckepalakecil_bet_5050special'],
                            ['7_discekorganjil','discekorganjil_bet_5050special'],
                            ['7_discekorgenap','discekorgenap_bet_5050special'],
                            ['7_discekorbesar','discekorbesar_bet_5050special'],
                            ['7_discekorkecil','discekorkecil_bet_5050special'],
                            ['7_limittotal','limittotal_bet_5050special']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case '5050kombinasi':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['8_minbet','min_bet_5050kombinasi'],
                            ['8_maxbet','max_bet_5050kombinasi'],
                            ['8_belakangkeimono','kei_belakangmono_bet_5050kombinasi'],
                            ['8_belakangkeistereo','kei_belakangstereo_bet_5050kombinasi'],
                            ['8_belakangkeikembang','kei_belakangkembang_bet_5050kombinasi'],
                            ['8_belakangkeikempis','kei_belakangkempis_bet_5050kombinasi'],
                            ['8_belakangkeikembar','kei_belakangkembar_bet_5050kombinasi'],
                            ['8_tengahkeimono','kei_tengahmono_bet_5050kombinasi'],
                            ['8_tengahkeistereo','kei_tengahstereo_bet_5050kombinasi'],
                            ['8_tengahkeikembang','kei_tengahkembang_bet_5050kombinasi'],
                            ['8_tengahkeikempis','kei_tengahkempis_bet_5050kombinasi'],
                            ['8_tengahkeikembar','kei_tengahkembar_bet_5050kombinasi'],
                            ['8_depankeimono','kei_depanmono_bet_5050kombinasi'],
                            ['8_depankeistereo','kei_depanstereo_bet_5050kombinasi'],
                            ['8_depankeikembang','kei_depankembang_bet_5050kombinasi'],
                            ['8_depankeikempis','kei_depankempis_bet_5050kombinasi'],
                            ['8_depankeikembar','kei_depankembar_bet_5050kombinasi'],
                            ['8_belakangdiscmono','disc_belakangmono_bet_5050kombinasi'],
                            ['8_belakangdiscstereo','disc_belakangstereo_bet_5050kombinasi'],
                            ['8_belakangdisckembang','disc_belakangkembang_bet_5050kombinasi'],
                            ['8_belakangdisckempis','disc_belakangkempis_bet_5050kombinasi'],
                            ['8_belakangdisckembar','disc_belakangkembar_bet_5050kombinasi'],
                            ['8_tengahdiscmono','disc_tengahmono_bet_5050kombinasi'],
                            ['8_tengahdiscstereo','disc_tengahstereo_bet_5050kombinasi'],
                            ['8_tengahdisckembang','disc_tengahkembang_bet_5050kombinasi'],
                            ['8_tengahdisckempis','disc_tengahkempis_bet_5050kombinasi'],
                            ['8_tengahdisckembar','disc_tengahkembar_bet_5050kombinasi'],
                            ['8_depandiscmono','disc_depanmono_bet_5050kombinasi'],
                            ['8_depandiscstereo','disc_depanstereo_bet_5050kombinasi'],
                            ['8_depandisckembang','disc_depankembang_bet_5050kombinasi'],
                            ['8_depandisckempis','disc_depankempis_bet_5050kombinasi'],
                            ['8_depandisckembar','disc_depankembar_bet_5050kombinasi'],
                            ['8_limittotal','limittotal_bet_5050kombinasi']
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'macaukombinasi':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['9_minbet','min_bet'],
                            ['9_maxbet','max_bet'],
                            ['9_win','win_bet'],
                            ['9_discount','diskon_bet'],
                            ['9_limittotal','limit_total']
                            
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'dasar':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['10_minbet','min_bet'],
                            ['10_maxbet','max_bet'],
                            ['10_keibesar','kei_besar_bet'],
                            ['10_keikecil','kei_kecil_bet'],
                            ['10_keigenap','kei_genap_bet'],
                            ['10_keiganjil','kei_ganjil_bet'],
                            ['10_discbesar','disc_besar_bet'],
                            ['10_disckecil','disc_kecil_bet'],
                            ['10_discigenap','disc_genap_bet'],
                            ['10_discganjil','disc_ganjil_bet'],
                            ['10_limittotal','limit_total']
                            
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
                case 'shio':
                    result = await model.Mcompanypasaran.findAll({
                        attributes: [
                            ['11_minbet','min_bet'],
                            ['11_maxbet','max_bet'],
                            ['11_win','win_bet'],
                            ['11_disc','diskon_bet'],
                            ['11_limittotal','limit_total']
                            
                        ],
                        where: {
                            idcompany: member_company,
                            idpasarantogel: pasaran_code
                        }
                    })
                    break;
            }
            
            if(result.length > 0){
                res.send({
                    status: 200,
                    record: result
                });
            }else{
                throw createError.NotFound()
            }
        } catch (error) {
            next(error)
        }
    },
    servicelimittogel: async (req,res,next) =>{
        try {
            const {member_username, member_company, pasaran_code, pasaran_periode, tipe_game} = req.body
            if(!member_username) throw createError.BadRequest() 
            if(!member_company) throw createError.BadRequest() 
            if(!pasaran_code) throw createError.BadRequest() 
            if(!pasaran_periode) throw createError.BadRequest() 
            if(!tipe_game) throw createError.BadRequest() 

            let result = await model.Minvoice.findAll({
                attributes: ['typegame'],
                where: {
                    idcompany: member_company,
                    username: member_username,
                    keluaranperiode: pasaran_periode,
                    idpasarantogel: pasaran_code
                }
            })
            const newrecord = []
            let total_4d = 0;
            let total_3d = 0;
            let total_2d = 0;
            let total_2dd = 0;
            let total_2dt = 0;
            for await (const rec of result){
                let typegame = rec.dataValues.typegame
                
                if(tipe_game == '4-3-2'){
                    if(typegame == '4D'){
                        total_4d = total_4d + 1
                    }
                    if(typegame == '3D'){
                        total_3d = total_3d + 1
                    }
                    if(typegame == '2D'){
                        total_2d = total_2d + 1
                    }
                    if(typegame == '2DD'){
                        total_2dd = total_2dd + 1
                    }
                    if(typegame == '2DT'){
                        total_2dt = total_2dt + 1
                    }
                }
            }
           
            if(result.length > 0){
                res.send({
                    status: 200,
                    total_4d: total_4d,
                    total_3d: total_3d,
                    total_2d: total_2d,
                    total_2dd: total_2dd,
                    total_2dt: total_2dt
                });
            }else{
                throw createError.NotFound()
            }
        } catch (error) {
            next(error)
        }
    },
    servicesavetransaksi: async (req,res,next) =>{
        try {
            const {member_username, member_company, idtrxkeluaran, idcomppasaran, devicemember, formipaddress, totalbayarbet, list4d} = req.body
            if(!member_username) throw createError.BadRequest() 
            if(!member_company) throw createError.BadRequest() 
            if(!idtrxkeluaran) throw createError.BadRequest() 
            if(!idcomppasaran) throw createError.BadRequest() 
            if(!devicemember) throw createError.BadRequest() 
            if(!formipaddress) throw createError.BadRequest() 
            if(!totalbayarbet) throw createError.BadRequest() 
            if(!list4d) throw createError.BadRequest() 

            let date = new Date()
            let thisDay = date.getDay()
            let myDays = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']
            const dateTime_now = moment().format('YYYY-MM-DD HH:mm:ss')
            const day_now = myDays[thisDay]

            let status = ""
            let permainan = ""
            let totalbayar = 0
            let totalbet_all = 0
            let bet = 0
            let diskon = 0
            let kei = 0
            let bayar = 0
            let limit_global_togel = 0
            let status = ""
            let msg = ""
            let statuspasaran = 'ONLINE'
            let flag_next = false
            let flag_loop = false
            let flag_save = false
            let dompet = 5000000
            let pasaranjamtutup = ''
            let pasaranjamopen = ''
            let limit_sum = 0
            let limit_togel_4d = 0
            let limit_togel_3d = 0
            let limit_togel_2d = 0
            let limit_togel_2dd = 0
            let limit_togel_2dt = 0
            let limit_togel_colokbebas = 0
            let limit_togel_colokmacau = 0
            let limit_togel_coloknaga = 0
            let limit_togel_colokjitu = 0
            let limit_togel_5050umum = 0
            let limit_togel_5050special = 0
            let limit_togel_5050kombinasi = 0
            let limit_togel_kombinasi = 0
            let limit_togel_dasar = 0
            let limit_togel_shio = 0
            if(parseInt(dompet) < parseInt(totalbayarbet)){
                status = "2";
                msg = "Balance Anda Tidak Cukup";
                flag_loop = TRUE;
            }

            let result = await model.Mpasaran.findAll({
                attributes: [
                        'jamtutup','jamopen',
                        'limit_togel_4d','limit_togel_3d','limit_togel_2d','limit_togel_2dd','limit_togel_2dt',
                        'limit_togel_colokbebas','limit_togel_colokmacau','limit_togel_coloknaga','limit_togel_colokjitu',
                        'limit_togel_5050umum','limit_togel_5050special','limit_togel_5050kombinasi',
                        'limit_togel_kombinasi','limit_togel_dasar','limit_togel_shio'
                    ],
                where: {
                    [Op.and]: [
                        {
                            idcompany:member_company,
                            idcomppasaran:idcomppasaran
                        }
                    ]
                },
            })
            let pasaranOffline = await model.Mpasaranoffline.findAll({
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
            for await (const rec of result){
                pasaranjamtutup = moment().format('YYYY-MM-DD')+' '+rec.dataValues.jamtutup
                pasaranjamopen = moment().format('YYYY-MM-DD')+' '+rec.dataValues.jamopen
                limit_togel_4d = rec.dataValues.limit_togel_4d
                limit_togel_3d = rec.dataValues.limit_togel_3d
                limit_togel_2d = rec.dataValues.limit_togel_2d
                limit_togel_2dd = rec.dataValues.limit_togel_2dd
                limit_togel_2dt = rec.dataValues.limit_togel_2dt
                limit_togel_colokbebas = rec.dataValues.limit_togel_colokbebas
                limit_togel_colokmacau = rec.dataValues.limit_togel_colokmacau
                limit_togel_coloknaga = rec.dataValues.limit_togel_coloknaga
                limit_togel_colokjitu = rec.dataValues.limit_togel_colokjitu
                limit_togel_5050umum = rec.dataValues.limit_togel_5050umum
                limit_togel_5050special = rec.dataValues.limit_togel_5050special
                limit_togel_5050kombinasi = rec.dataValues.limit_togel_5050kombinasi
                limit_togel_kombinasi = rec.dataValues.limit_togel_kombinasi
                limit_togel_dasar = rec.dataValues.limit_togel_dasar
                limit_togel_shio = rec.dataValues.limit_togel_shio
            }
            if(pasaranOffline.length > 0){
                if(parseInt(moment(dateTime_now).format('x')) >= parseInt(moment(pasaranjamtutup).format('x')) && parseInt(moment(dateTime_now).format('x')) <= parseInt(moment(pasaranjamopen).format('x'))){
                    status = '2'
                    msg = 'Pasaran Sudah Tutup'
                    flag_loop = true
                }
            }
            if(flag_loop == false){
                if(list4d.length > 0){
                    for(let i=0;i<list4d.length;i++){
                        switch(list4d[i]['permainan']){
                            case "4D":
                                permainan = "4D/3D/2D"
                                limit_global_togel = limit_togel_4d
                                break;
                            case "3D":
                                permainan = "4D/3D/2D"
                                limit_global_togel = limit_togel_3d
                                break;
                            case "2D":
                                permainan = "4D/3D/2D"
                                limit_global_togel = limit_togel_2d
                                break;
                            case "2DD":
                                permainan = "4D/3D/2D"
                                limit_global_togel = limit_togel_2dd
                                break;
                            case "2DT":
                                permainan = "4D/3D/2D"
                                limit_global_togel = limit_togel_2dt
                                break;
                            case "COLOK_BEBAS":
                                permainan = "COLOK BEBAS"
                                limit_global_togel = limit_togel_colokbebas
                                break;
                            case "COLOK_MACAU":
                                permainan = "COLOK MACAU"
                                limit_global_togel = limit_togel_colokmacau
                                break;
                            case "COLOK_NAGA":
                                permainan = "COLOK NAGA"
                                limit_global_togel = limit_togel_coloknaga
                                break;
                            case "COLOK_JITU":
                                permainan = "COLOK JITU"
                                limit_global_togel = limit_togel_colokjitu
                                break;
                            case "50_50_UMUM":
                                permainan = "50 - 50 UMUM"
                                limit_global_togel = limit_togel_5050umum
                                break;
                            case "50_50_SPECIAL":
                                permainan = "50 - 50 SPECIAL"
                                limit_global_togel = limit_togel_5050special
                                break;
                            case "50_50_KOMBINASI":
                                permainan = "50 - 50 KOMBINASI"
                                limit_global_togel = limit_togel_5050kombinasi
                                break;
                            case "MACAU_KOMBINASI":
                                permainan = "MACAU / KOMBINASI"
                                limit_global_togel = limit_togel_5050kombinasi
                                break;
                            case "DASAR":
                                permainan = "DASAR"
                                limit_global_togel = limit_togel_5050kombinasi
                                break;
                            case "SHIO":
                                permainan = "SHIO"
                                limit_global_togel = limit_togel_5050kombinasi
                                break;
                        }
                        if(member_username != "" && member_company != ""){
                            bet = list4d[i]['bet']
                            diskon = list4d[i]['diskonpercen']
                            kei = list4d[i]['kei_percen']
                            bayar = bet - (bet*diskon) - (bet*kei);
                            totalbayar = totalbayar + bayar;
                            
                            let invoice = await model.Minvoice.findAll({
                                attributes: [
                                    [sequelize.fn('sum', sequelize.col('bet')), 'total']
                                ],
                                where: {
                                    [Op.and]: [
                                        {
                                            idtrxkeluaran:idtrxkeluaran,
                                            typegame:list4d[i]['permainan'],
                                            nomortogel:list4d[i]['nomor']
                                        }
                                    ]
                                },
                            })
                            for await (const rec of invoice){
                                limit_sum = rec.dataValues.total
                            }
                            totalbet_all = limit_sum + bet

                            if(parseInt(limit_global_togel) < parseInt(totalbet_all)){
                                flag_save = true
                                status = "1"
                                msg += list4d[i]['nomor']
                            }

                            if(flag_save == false){

                                let keluarandetail = await model.Mkeluarandetail.create({
                                    'idtrxkeluarandetail': idrecord,
                                    'idtrxkeluaran': idtrxkeluaran,
                                    'datetimedetail': SERVERTIME(),
                                    'ipaddress': formipaddress,
                                    'idcompany': member_company,
                                    'username': member_username,
                                    'typegame': list4d[i]['permainan'],
                                    'nomortogel': list4d[i]['nomor'],
                                    'bet': list4d[i]['bet'],
                                    'diskon': list4d[i]['diskonpercen'],
                                    'win': list4d[i]['win'],
                                    'kei': list4d[i]['kei_percen'],
                                    'browsertogel': "",
                                    'devicetogel': devicemember,
                                    'statuskeluarandetail': 'RUNNING',
                                    'createkeluarandetail': member_username,
                                    'createdatekeluarandetail': SERVERTIME(),
                                })
                                msg = 'Success'
                                flag_next = true
                            }
                        }
                    }
                }
            }

            if(flag_next == true){
                res.send({
                    status: 200,
                    message: msg
                });
            }else{
                throw createError.NotFound()
            }
        } catch (error) {
            next(error)
        }
    },
}