import axios from 'axios'

export const makePayment = async function (payload) {
  const { data } = await axios.post('/payments', payload)
  return data
}
