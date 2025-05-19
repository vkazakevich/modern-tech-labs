import PropTypes from 'prop-types'
import { useState } from 'react'

import { AuthContext } from '../contexts/AuthContext'
import { useEffect } from 'react'
import axios from 'axios'

const TOKEN_KEY = 'token'

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(null)

  const auth = async ({ token }) => {
    setToken(token)
  }

  const setupToken = (token) => {
    if (token) {
      localStorage.setItem(TOKEN_KEY, token)
      axios.defaults.headers.common['Authorization'] = `Bearer ${token}`
    } else {
      localStorage.removeItem(TOKEN_KEY)
      axios.defaults.headers.common['Authorization'] = undefined
    }

    setToken(token)
  }

  useEffect(() => {
    setupToken(localStorage.getItem(TOKEN_KEY))
  }, [])

  const value = {
    token,
    auth,
    isAuth: !!token
  }

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

AuthProvider.propTypes = {
  children: PropTypes.node.isRequired
}
