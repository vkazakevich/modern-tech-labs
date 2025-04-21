import { useState } from 'react'

function Payments() {
  const [amount, setAmount] = useState(0)

  async function payment(e) {
    e.preventDefault()

    const response = await fetch('http://localhost:8000/payments/', {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ amount: +amount }),
    })

    if (response.ok) {
      alert('Success!')
    }
  }

  return (
    <>
      <h1>Payments</h1>
      <div>
        <form onSubmit={payment}>
          <input
            type="number"
            min={0}
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
          />
          <button type="submit">Submit</button>
        </form>
      </div>
    </>
  )
}

export default Payments
