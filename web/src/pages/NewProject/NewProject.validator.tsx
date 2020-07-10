import * as yup from 'yup'

const requiredMessage = (field: string) => `${field} is required`
const minMessage = (min: number) => `Must have at least ${min} characters`
const maxMessage = (max: number) => `Cannot have more than ${max} characters`

export const newProjectValidator = yup.object().shape({
  title: yup
    .string()
    .min(3, minMessage(3))
    .max(255, maxMessage(255))
    .required(requiredMessage('Project title')),
  network: yup.string()
})
