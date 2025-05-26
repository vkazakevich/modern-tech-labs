import { Router } from 'express'

import auth from '#api/auth'
import authGoogle from '#api/auth/google'

const router = Router()

router.use('/api', auth)

router.use('/', authGoogle)

export default router
