import axios from 'axios'

export const getProducts = async function () {
  const { data } = await axios.get('/products')
  return data
}
