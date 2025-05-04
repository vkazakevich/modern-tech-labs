import { useEffect, useState } from 'react'
import { makePayment } from '../api/payments'
import { useCart } from '../hooks/useCart'
import { getProducts } from '../api/products'

function Payments() {
  const { cartItems } = useCart()
  const [amount, setAmount] = useState(null)

  useEffect(() => {
    const calcAmount = async () => {
      const products = await getProducts()

      return cartItems.reduce((acc, currentProductId) => {
        const product = products.find(
          (product) => product.ID == currentProductId
        )
        return acc + product?.price
      }, 0)
    }

    calcAmount().then((finalAmount) => setAmount(finalAmount))
  }, [cartItems])

  async function payment(e) {
    e.preventDefault()
    try {
      await makePayment({ amount: +amount })
      alert(`Success payment $${amount}`)
    } catch (err) {
      //
    }
  }

  return (
    <>
      <h1>Payments</h1>
      <div>
        {amount ? (
          <form onSubmit={payment}>
            <p>
              Final amount: 
              <strong> ${amount}</strong>
            </p>
            <button type="submit">Pay</button>
          </form>
        ) : (
          'Calculating...'
        )}
      </div>
    </>
  )
}

export default Payments
