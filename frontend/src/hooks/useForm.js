import { useState } from 'react'

export const useForm = initVals => {
  const [values, setValues] = useState(initVals)

  const handleChange = (e) => {
    setValues({
      ...values,
      [e.target.name]: e.target.value
    })
  }

  const reset = () => {
    setValues(
      Object.fromEntries(Object.keys(values).map((key) => [key, '']))
    )
  }

  return [
    values,
    handleChange,
    reset
  ]
}
