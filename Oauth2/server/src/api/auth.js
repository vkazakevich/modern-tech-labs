import { Router } from 'express'
import * as bcrypt from 'bcrypt'
import jwt from 'jsonwebtoken'
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

router.post('/login', async (req, res) => {
  const { email, password } = req.body

  const user = await prisma.user.findUnique({
    where: { email }
  })

  if (!user) {
    res.status(422).send({ message: 'Invalid email or password' })
  } else {
    const isSuccess = bcrypt.compareSync(password, user.password)

    if (isSuccess) {
      const userPayload = {
        id: user.id
      }

      const token = jwt.sign(userPayload, SECRET_KEY, {
        expiresIn: 3600
      })
      res.status(200).send({ token })
    } else {
      res.status(422).send({ message: 'Invalid email or password' })
    }
  }
})

export default router
