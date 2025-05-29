import prisma from '#libs/prisma'
import * as bcrypt from 'bcrypt'
import jwt from 'jsonwebtoken'
import { SALT_ROUND, SECRET_KEY } from '#config'

export const getJwtTokenForUser = async (user) => {
  const userPayload = {
    id: user.id
  }

  const token = jwt.sign(userPayload, SECRET_KEY, {
    expiresIn: 3600
  })

  return token
}

export const createUserIfNotExist = async ({ fullName, email, password }) => {
  const user = await prisma.user.findUnique({
    where: { email: email }
  })

  if (user) return false

  const salt = bcrypt.genSaltSync(Number(SALT_ROUND))
  const hashPassword = password ? bcrypt.hashSync(password, salt) : 'password'

  await prisma.user.create({
    data: {
      fullName,
      email,
      password: hashPassword
    }
  })
}
