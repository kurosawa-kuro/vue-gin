<template>
  <div>
    <!-- Page header -->
    <div class="mb-8">
      <h1 class="text-2xl md:text-3xl text-gray-800 dark:text-gray-100 font-bold">Microposts</h1>
      <p class="text-gray-600 dark:text-gray-400">Create and view microposts</p>
    </div>

    <!-- Create post form -->
    <div class="card mb-8">
      <div class="p-6 border-b border-gray-200 dark:border-gray-700/60">
        <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-100">Create New Post</h2>
      </div>
      <div class="p-6">
        <form @submit.prevent="createPost" class="space-y-4">
          <div>
            <label for="title" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              Post Title
            </label>
            <input
              id="title"
              v-model="newPost.title"
              type="text"
              class="form-input"
              :class="{ 'border-red-500': hasFieldError('title') }"
              placeholder="What's on your mind?"
              required
            />
            <p v-if="hasFieldError('title')" class="mt-1 text-sm text-red-600">
              {{ getFieldError('title') }}
            </p>
          </div>
          <div>
            <label for="user" class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
              Author
            </label>
            <select
              id="user"
              v-model="newPost.userId"
              class="form-input"
              :class="{ 'border-red-500': hasFieldError('userId') }"
              required
            >
              <option value="">Select a user</option>
              <option v-for="user in users" :key="user.id" :value="user.id">
                {{ user.name }}
              </option>
            </select>
            <p v-if="hasFieldError('userId')" class="mt-1 text-sm text-red-600">
              {{ getFieldError('userId') }}
            </p>
          </div>
          <div class="flex justify-end">
            <button
              type="submit"
              :disabled="loading || !isValid"
              class="btn btn-primary disabled:opacity-50 disabled:cursor-not-allowed"
            >
              <span v-if="loading" class="inline-block animate-spin rounded-full h-4 w-4 border-b-2 border-white mr-2"></span>
              {{ loading ? 'Creating...' : 'Create Post' }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Posts list -->
    <div class="card">
      <div class="p-6 border-b border-gray-200 dark:border-gray-700/60">
        <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-100">All Posts</h2>
      </div>
      <div class="p-6">
        <div v-if="loading && microposts.length === 0" class="text-center py-8">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-violet-500"></div>
          <p class="mt-2 text-gray-600 dark:text-gray-400">Loading posts...</p>
        </div>
        <div v-else-if="error" class="text-center py-8 text-red-500">
          <p>Error: {{ error }}</p>
          <button @click="fetchData" class="btn btn-secondary mt-4">Try Again</button>
        </div>
        <div v-else-if="microposts.length === 0" class="text-center py-8 text-gray-500 dark:text-gray-400">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
          <p class="mt-2">No posts yet</p>
          <p class="text-sm">Create your first post using the form above</p>
        </div>
        <div v-else class="space-y-4">
          <div
            v-for="post in microposts"
            :key="post.id"
            class="border border-gray-200 dark:border-gray-700/60 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex items-start space-x-4">
              <div class="flex-shrink-0">
                <div class="w-10 h-10 bg-violet-500 rounded-full flex items-center justify-center">
                  <span class="text-white font-medium">{{ post.user.name.charAt(0) }}</span>
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <div class="flex items-center space-x-2">
                  <h3 class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ post.user.name }}</h3>
                  <span class="text-xs text-gray-500 dark:text-gray-500">•</span>
                  <span class="text-xs text-gray-500 dark:text-gray-500">{{ formatDate(post.createdAt) }}</span>
                </div>
                <p class="mt-2 text-gray-800 dark:text-gray-200">{{ post.title }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useApiStore } from '../stores/api'
import { computed, onMounted, reactive, watch } from 'vue'
import { useFormValidation } from '../composables/useFormValidation'
import { CreateMicropostFormSchema } from '../schemas/zod'

export default {
  name: 'Microposts',
  setup() {
    const apiStore = useApiStore()
    const { validate, isValid, getFieldError, hasFieldError, clearErrors } = useFormValidation(CreateMicropostFormSchema)

    const newPost = reactive({
      title: '',
      userId: ''
    })

    const users = computed(() => apiStore.users)
    const microposts = computed(() => apiStore.microposts)
    const loading = computed(() => apiStore.loading)
    const error = computed(() => apiStore.error)

    // Watch form changes and validate
    watch(newPost, () => {
      if (newPost.title || newPost.userId) {
        validate({
          title: newPost.title,
          userId: newPost.userId ? parseInt(newPost.userId) : null
        })
      }
    }, { deep: true })

    const createPost = async () => {
      const formData = {
        title: newPost.title,
        userId: newPost.userId ? parseInt(newPost.userId) : null
      }
      
      const validation = validate(formData)
      if (!validation.success) {
        return
      }

      try {
        await apiStore.createMicropost(formData.title, formData.userId)
        newPost.title = ''
        newPost.userId = ''
        clearErrors()
      } catch (error) {
        console.error('Failed to create post:', error)
      }
    }

    const fetchData = async () => {
      await Promise.all([
        apiStore.fetchUsers(),
        apiStore.fetchMicroposts()
      ])
    }

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleDateString() + ' ' + date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    }

    onMounted(fetchData)

    return {
      newPost,
      users,
      microposts,
      loading,
      error,
      isValid,
      getFieldError,
      hasFieldError,
      createPost,
      fetchData,
      formatDate
    }
  }
}
</script>
