<template>
  <div>
    <!-- Page header -->
    <div class="mb-8">
      <h1 class="text-2xl md:text-3xl text-gray-800 dark:text-gray-100 font-bold">Dashboard</h1>
      <p class="text-gray-600 dark:text-gray-400">Welcome to the Vue Express sample application</p>
    </div>

    <!-- Stats cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
      <!-- Users card -->
      <div class="card p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-blue-500 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <div class="text-sm font-medium text-gray-500 dark:text-gray-400">Total Users</div>
            <div class="text-2xl font-bold text-gray-900 dark:text-gray-100">{{ users.length }}</div>
          </div>
        </div>
      </div>

      <!-- Microposts card -->
      <div class="card p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-green-500 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z" clip-rule="evenodd" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <div class="text-sm font-medium text-gray-500 dark:text-gray-400">Total Posts</div>
            <div class="text-2xl font-bold text-gray-900 dark:text-gray-100">{{ microposts.length }}</div>
          </div>
        </div>
      </div>

      <!-- Recent activity card -->
      <div class="card p-6">
        <div class="flex items-center">
          <div class="flex-shrink-0">
            <div class="w-8 h-8 bg-purple-500 rounded-lg flex items-center justify-center">
              <svg class="w-5 h-5 text-white" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clip-rule="evenodd" />
              </svg>
            </div>
          </div>
          <div class="ml-4">
            <div class="text-sm font-medium text-gray-500 dark:text-gray-400">Recent Activity</div>
            <div class="text-2xl font-bold text-gray-900 dark:text-gray-100">{{ recentPostsCount }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Recent posts -->
    <div class="card">
      <div class="p-6 border-b border-gray-200 dark:border-gray-700/60">
        <h2 class="text-lg font-semibold text-gray-800 dark:text-gray-100">Recent Posts</h2>
      </div>
      <div class="p-6">
        <div v-if="loading" class="text-center py-4">
          <div class="inline-block animate-spin rounded-full h-6 w-6 border-b-2 border-violet-500"></div>
          <span class="ml-2 text-gray-600 dark:text-gray-400">Loading...</span>
        </div>
        <div v-else-if="error" class="text-center py-4 text-red-500">
          Error: {{ error }}
        </div>
        <div v-else-if="recentPosts.length === 0" class="text-center py-8 text-gray-500 dark:text-gray-400">
          No posts yet. <router-link to="/microposts" class="text-violet-500 hover:text-violet-600">Create your first post</router-link>
        </div>
        <div v-else class="space-y-4">
          <div v-for="post in recentPosts" :key="post.id" class="flex items-start space-x-3 p-4 bg-gray-50 dark:bg-gray-700/50 rounded-lg">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-violet-500 rounded-full flex items-center justify-center">
                <span class="text-white text-sm font-medium">{{ post.user.name.charAt(0) }}</span>
              </div>
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-gray-900 dark:text-gray-100">{{ post.user.name }}</p>
              <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">{{ post.title }}</p>
              <p class="text-xs text-gray-500 dark:text-gray-500 mt-1">{{ formatDate(post.createdAt) }}</p>
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
  name: 'Dashboard',
  setup() {
    const apiStore = useApiStore()

    const users = computed(() => apiStore.users)
    const microposts = computed(() => apiStore.microposts)
    const loading = computed(() => apiStore.loading)
    const error = computed(() => apiStore.error)

    const recentPosts = computed(() => {
      return microposts.value.slice(0, 5)
    })

    const recentPostsCount = computed(() => {
      const oneDayAgo = new Date()
      oneDayAgo.setDate(oneDayAgo.getDate() - 1)
      return microposts.value.filter(post => new Date(post.createdAt) > oneDayAgo).length
    })

    const formatDate = (dateString) => {
      const date = new Date(dateString)
      return date.toLocaleDateString() + ' ' + date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    }

    onMounted(async () => {
      await Promise.all([
        apiStore.fetchUsers(),
        apiStore.fetchMicroposts()
      ])
    })

    return {
      users,
      microposts,
      loading,
      error,
      recentPosts,
      recentPostsCount,
      formatDate
    }
  }
}
</script>
