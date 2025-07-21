<template>
  <div>
    <!-- Page header -->
    <div class="mb-8">
      <h1 class="text-2xl md:text-3xl text-gray-800 dark:text-gray-100 font-bold">Users</h1>
      <p class="text-gray-600 dark:text-gray-400">View all users and their posts</p>
    </div>

    <!-- Users list -->
    <div class="card">
      <div class="p-6 border-b border-gray-200 dark:border-gray-700/60">
        <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-100">All Users</h2>
      </div>
      <div class="p-6">
        <div v-if="loading" class="text-center py-8">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-violet-500"></div>
          <p class="mt-2 text-gray-600 dark:text-gray-400">Loading users...</p>
        </div>
        <div v-else-if="error" class="text-center py-8 text-red-500">
          <p>Error: {{ error }}</p>
          <button @click="fetchUsers" class="btn btn-secondary mt-4">Try Again</button>
        </div>
        <div v-else-if="users.length === 0" class="text-center py-8 text-gray-500 dark:text-gray-400">
          <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
          </svg>
          <p class="mt-2">No users found</p>
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="user in users"
            :key="user.id"
            class="border border-gray-200 dark:border-gray-700/60 rounded-lg p-6 hover:shadow-md transition-shadow"
          >
            <div class="flex items-center space-x-4">
              <div class="flex-shrink-0">
                <div class="w-12 h-12 bg-violet-500 rounded-full flex items-center justify-center">
                  <span class="text-white text-lg font-medium">{{ user.name.charAt(0) }}</span>
                </div>
              </div>
              <div class="flex-1 min-w-0">
                <h3 class="text-lg font-medium text-gray-900 dark:text-gray-100 truncate">{{ user.name }}</h3>
                <p class="text-sm text-gray-500 dark:text-gray-400">
                  {{ user.microposts.length }} {{ user.microposts.length === 1 ? 'post' : 'posts' }}
                </p>
                <p class="text-xs text-gray-400 dark:text-gray-500 mt-1">
                  Joined {{ formatDate(user.createdAt) }}
                </p>
              </div>
            </div>
            
            <!-- User's recent posts -->
            <div v-if="user.microposts.length > 0" class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700/60">
              <h4 class="text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">Recent Posts</h4>
              <div class="space-y-2">
                <div
                  v-for="post in user.microposts.slice(0, 3)"
                  :key="post.id"
                  class="text-sm text-gray-600 dark:text-gray-400 bg-gray-50 dark:bg-gray-700/50 rounded p-2"
                >
                  <p class="truncate">{{ post.title }}</p>
                  <p class="text-xs text-gray-500 dark:text-gray-500 mt-1">{{ formatDate(post.createdAt) }}</p>
                </div>
                <div v-if="user.microposts.length > 3" class="text-xs text-gray-500 dark:text-gray-500">
                  +{{ user.microposts.length - 3 }} more posts
                </div>
              </div>
            </div>
            <div v-else class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700/60">
              <p class="text-sm text-gray-500 dark:text-gray-400 italic">No posts yet</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { useApiStore } from '../stores/api'
import { computed, onMounted } from 'vue'

export default {
  name: 'Users',
  setup() {
    const apiStore = useApiStore()

    const users = computed(() => apiStore.users)
    const loading = computed(() => apiStore.loading)
    const error = computed(() => apiStore.error)

    const fetchUsers = async () => {
      await apiStore.fetchUsers()
    }

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleDateString()
    }

    onMounted(fetchUsers)

    return {
      users,
      loading,
      error,
      fetchUsers,
      formatDate
    }
  }
}
</script>
