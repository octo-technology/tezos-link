import * as yup from 'yup'

const requiredMessage = (field: string) => `${field} is required`
const uuidMessage = `Must be a valid UUID`

export const signInProjectValidator = yup.object().shape({
  uuid: yup
    .string()
    .matches(RegExp('^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$'), {
      message: uuidMessage,
      excludeEmptyString: true
    })
    .required(requiredMessage('Project ID'))
})
