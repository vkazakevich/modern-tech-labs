import axios from 'axios'

if (process.env.NODE_ENV === 'development') {
  axios.defaults.baseURL = 'http://localhost:8001'
}

if (process.env.NODE_ENV === 'production') {
  axios.defaults.baseURL = 'https://ebiznes-back-end-e2gfc5ducudeg8fv.polandcentral-01.azurewebsites.net'
}
