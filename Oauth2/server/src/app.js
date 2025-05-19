import express from 'express'

const app = express()
const port = process.env.PORT || 8000

app.listen(port, () => {
  console.log(`Server is running on port ${port} 🔥🔥🔥`)
})