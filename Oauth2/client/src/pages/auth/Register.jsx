import { useForm } from 'react-hook-form'
import authApi from '../../api/auth'
import { useState } from 'react'

function Register() {
  const [success, setSuccess] = useState(false)
  const [errors, setErrors] = useState(null)
  const { register, handleSubmit } = useForm()

  const onSubmit = async (data) => {
    setErrors(null)

    try {
      await authApi.register(data)
      setSuccess(true)
    } catch (err) {
      setErrors(err.response?.data?.message)
    }
  }

  if (success)
    return (
      <>
        <div className="mx-auto w-[400px] mt-10 text-center">
          The user has been created!
        </div>
      </>
    )

  return (
    <>
      <header className="mx-auto w-[400px] mt-10">
        <h1>Register</h1>
      </header>

      <form className="mt-8" onSubmit={handleSubmit(onSubmit)}>
        <div className="mx-auto w-[400px]">
          {errors ? (
            <div className="text-red-500 border-1 border-red-500 rounded-lg px-3 py-2 mb-5">
              {errors}
            </div>
          ) : (
            ''
          )}

          <div className="py-2">
            <span className="px-1 text-sm text-gray-600">Full Name</span>
            <input
              {...register('fullName', { required: true })}
              className="text-md block px-3 py-2  rounded-lg w-full bg-white border-2 border-gray-300 placeholder-gray-600 shadow-md focus:placeholder-gray-500 focus:bg-white focus:border-gray-600 focus:outline-none"
            />
          </div>
          <div className="py-2">
            <span className="px-1 text-sm text-gray-600">Email</span>
            <input
              {...register('email', { required: true })}
              type="email"
              className="text-md block px-3 py-2  rounded-lg w-full bg-white border-2 border-gray-300 placeholder-gray-600 shadow-md focus:placeholder-gray-500 focus:bg-white focus:border-gray-600 focus:outline-none"
            />
          </div>
          <div className="py-2">
            <span className="px-1 text-sm text-gray-600">Password</span>
            <div className="relative">
              <input
                {...register('password', { required: true })}
                type="password"
                className="text-md block px-3 py-2  rounded-lg w-full bg-white border-2 border-gray-300 placeholder-gray-600 shadow-md focus:placeholder-gray-500 focus:bg-white focus:border-gray-600 focus:outline-none"
              />
            </div>
          </div>

          <button
            type="submit"
            className="mt-3 text-lg font-semibold bg-gray-800 w-full text-white rounded-lg px-6 py-3 block shadow-xl hover:text-white hover:bg-black"
          >
            Register
          </button>
        </div>
      </form>
    </>
  )
}

export default Register
