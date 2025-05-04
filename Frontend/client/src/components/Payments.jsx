import { useEffect, useState } from 'react'
import { makePayment } from '../api/payments'
import { useCart } from '../hooks/useCart'
import { getProducts } from '../api/products'
import { useNavigate } from 'react-router'

function Payments() {
  const navigate = useNavigate()
  const { cartItems, emptyCart } = useCart()
  const [amount, setAmount] = useState(null)

  useEffect(() => {
    const calcAmount = async () => {
      const products = await getProducts()

      return cartItems.reduce((acc, current) => {
        const product = products.find((p) => p.ID === current)
        return product ? (acc += product.price) : acc
      }, 0)
    }

    calcAmount().then((finalAmount) => setAmount(finalAmount))
  }, [cartItems])

  async function payment(e) {
    e.preventDefault()
    try {
      await makePayment({ amount: +amount })
      alert(`Success payment $${amount}`)
      emptyCart()
      navigate('/')
    } catch (err) {
      //
    }
  }

  if (!cartItems.length) return <>Cart is empty!</>

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
