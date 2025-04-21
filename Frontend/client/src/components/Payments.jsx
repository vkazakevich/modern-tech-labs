import { useState } from 'react'
import { makePayment } from '../api/payments'

function Payments() {
  const [amount, setAmount] = useState(0)

  async function payment(e) {
    e.preventDefault()
    try {
      await makePayment({ amount: +amount })
      alert('Success!')
    } catch (err) {
      //
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
