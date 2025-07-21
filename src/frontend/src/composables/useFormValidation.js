import { ref, reactive } from 'vue'
import { ZodError } from 'zod'

export function useFormValidation(schema) {
  const errors = reactive({})
  const isValid = ref(true)

  const validate = (data) => {
    // Clear previous errors
    Object.keys(errors).forEach(key => delete errors[key])
    isValid.value = true

    try {
      const result = schema.parse(data)
      return { success: true, data: result }
    } catch (error) {
      if (error instanceof ZodError) {
        isValid.value = false
        error.errors.forEach(err => {
          const path = err.path.join('.')
          if (!errors[path]) {
            errors[path] = []
          }
          errors[path].push(err.message)
        })
      }
      return { success: false, errors: error.errors }
    }
  }

  const getFieldError = (fieldName) => {
    return errors[fieldName]?.[0] || null
  }

  const hasFieldError = (fieldName) => {
    return Boolean(errors[fieldName]?.length)
  }

  const clearErrors = () => {
    Object.keys(errors).forEach(key => delete errors[key])
    isValid.value = true
  }

  return {
    errors,
    isValid,
    validate,
    getFieldError,
    hasFieldError,
    clearErrors
  }
}