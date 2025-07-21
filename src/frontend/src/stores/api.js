import { defineStore } from 'pinia'
import axios from 'axios'
import { UsersResponseSchema, MicropostsResponseSchema, MicropostSchema, ApiErrorSchema } from '../schemas/zod.js'

const API_BASE_URL = 'http://localhost:3001'

export const useApiStore = defineStore('api', {
  state: () => ({
    users: [],
    microposts: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchUsers() {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get(`${API_BASE_URL}/users`)
        const validatedUsers = UsersResponseSchema.parse(response.data)
        this.users = validatedUsers
      } catch (error) {
        if (error.name === 'ZodError') {
          this.error = 'Invalid data format received from server'
          console.error('User data validation failed:', error.errors)
        } else {
          this.error = error.message
          console.error('Error fetching users:', error)
        }
      } finally {
        this.loading = false
      }
    },

    async fetchMicroposts() {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get(`${API_BASE_URL}/microposts`)
        const validatedMicroposts = MicropostsResponseSchema.parse(response.data)
        this.microposts = validatedMicroposts
      } catch (error) {
        if (error.name === 'ZodError') {
          this.error = 'Invalid data format received from server'
          console.error('Micropost data validation failed:', error.errors)
        } else {
          this.error = error.message
          console.error('Error fetching microposts:', error)
        }
      } finally {
        this.loading = false
      }
    },

    async createMicropost(title, userId) {
      this.loading = true
      this.error = null
      try {
        const response = await axios.post(`${API_BASE_URL}/microposts`, {
          title,
          userId
        })
        const validatedMicropost = MicropostSchema.parse(response.data)
        this.microposts.unshift(validatedMicropost)
        return validatedMicropost
      } catch (error) {
        if (error.name === 'ZodError') {
          this.error = 'Invalid data format received from server'
          console.error('Created micropost validation failed:', error.errors)
        } else if (error.response?.data) {
          const apiError = ApiErrorSchema.safeParse(error.response.data)
          if (apiError.success) {
            this.error = apiError.data.error
          } else {
            this.error = error.message
          }
        } else {
          this.error = error.message
        }
        console.error('Error creating micropost:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    async fetchMicropostById(id) {
      this.loading = true
      this.error = null
      try {
        const response = await axios.get(`${API_BASE_URL}/microposts/${id}`)
        const validatedMicropost = MicropostSchema.parse(response.data)
        return validatedMicropost
      } catch (error) {
        if (error.name === 'ZodError') {
          this.error = 'Invalid data format received from server'
          console.error('Micropost by ID validation failed:', error.errors)
        } else {
          this.error = error.message
          console.error('Error fetching micropost:', error)
        }
        throw error
      } finally {
        this.loading = false
      }
    }
  }
})
