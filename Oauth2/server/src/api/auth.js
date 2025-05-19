import { Router } from 'express'
import * as bcrypt from 'bcrypt'
import prisma from '#libs/prisma'
import { SALT_ROUND, SECRET_KEY } from '#config'

const router = Router()

router.post('/register', async (req, res) => {
  const { name, email, password } = req.body

  const user = await prisma.user.findUnique({
    where: { email }
  })

  if (user) {
    res.status(422).send({
      message: 'User with this email already exists'
    })
  } else {
    const salt = bcrypt.genSaltSync(Number(SALT_ROUND))
    const hashPassword = bcrypt.hashSync(password, salt)

    await prisma.user.create({
      data: {
        name,
        email,
        password: hashPassword
      }
    })

    res.status(201).send({
      message: 'The user has been created'
    })
  }
})

export default router
