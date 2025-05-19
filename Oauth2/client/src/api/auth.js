import axios from 'axios'

const BASE_URL = 'http://localhost:8000'

const register = async (payload) => {
  return await axios.post(`${BASE_URL}/api/register`, payload)
}

const login = async (payload) => {
  return await axios.post(`${BASE_URL}/api/login`, payload)
}

export default {
  register,
  login
}
