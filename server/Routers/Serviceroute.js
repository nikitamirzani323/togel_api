const express = require('express')
const router = express.Router()
const ServiceController = require('../Controllers/ServiceController')

router.post('/serviceconfigtogel', ServiceController.serviceconfigtogel)
router.post('/servicelimittogel', ServiceController.servicelimittogel)
router.post('/serviceinvoicetogel', ServiceController.serviceinvoicetogel)
router.post('/servicesavetransaksi', ServiceController.servicesavetransaksi)

module.exports = router