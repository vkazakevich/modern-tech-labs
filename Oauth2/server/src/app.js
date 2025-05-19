import express from 'express'
import cors from 'cors'

import routes from '#api'

const app = express()
const port = process.env.PORT || 8000

app.use(cors())
app.use(express.json())
app.use(express.urlencoded())

app.use('/api', routes)

app.listen(port, () => {
  console.log(`Server is running on port ${port} 🔥🔥🔥`)
})
