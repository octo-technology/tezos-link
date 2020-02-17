import * as yup from 'yup'

export const projectNameValidator: yup.StringSchema = yup
  .string()
  .min(3, 'Name must have at least 3 characters')
  .max(255)
  .required()
  .label('Name')
