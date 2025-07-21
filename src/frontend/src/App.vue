<template>
  <div class="font-inter antialiased bg-gray-100 dark:bg-gray-900 text-gray-600 dark:text-gray-400 min-h-screen">
    <div class="flex h-screen overflow-hidden">
      <!-- Sidebar -->
      <div class="min-w-fit">
        <div class="flex flex-col w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700/60 p-4 h-full">
          <!-- Logo -->
          <div class="flex justify-center mb-10">
            <div class="flex items-center">
              <svg class="fill-violet-500" xmlns="http://www.w3.org/2000/svg" width="32" height="32">
                <path d="M31.956 14.8C31.372 6.92 25.08.628 17.2.044V5.76a9.04 9.04 0 0 0 9.04 9.04h5.716ZM14.8 26.24v5.716C6.92 31.372.63 25.08.044 17.2H5.76a9.04 9.04 0 0 1 9.04 9.04Zm11.44-9.04h5.716c-.584 7.88-6.876 14.172-14.756 14.756V26.24a9.04 9.04 0 0 1 9.04-9.04ZM.044 14.8C.63 6.92 6.92.628 14.8.044V5.76a9.04 9.04 0 0 1-9.04 9.04H.044Z" />
              </svg>
              <span class="ml-3 text-lg font-semibold text-gray-800 dark:text-gray-100">Vue Express</span>
            </div>
          </div>
          
          <!-- Navigation -->
          <nav class="space-y-2">
            <router-link 
              to="/" 
              class="flex items-center px-4 py-2 text-sm font-medium rounded-lg transition-colors"
              :class="$route.path === '/' ? 'bg-violet-500/10 text-violet-500' : 'text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/50'"
            >
              <svg class="w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="M10.707 2.293a1 1 0 00-1.414 0l-7 7a1 1 0 001.414 1.414L4 10.414V17a1 1 0 001 1h2a1 1 0 001-1v-2a1 1 0 011-1h2a1 1 0 011 1v2a1 1 0 001 1h2a1 1 0 001-1v-6.586l.293.293a1 1 0 001.414-1.414l-7-7z" />
              </svg>
              Dashboard
            </router-link>
            
            <router-link 
              to="/microposts" 
              class="flex items-center px-4 py-2 text-sm font-medium rounded-lg transition-colors"
              :class="$route.path === '/microposts' ? 'bg-violet-500/10 text-violet-500' : 'text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/50'"
            >
              <svg class="w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M2 5a2 2 0 012-2h8a2 2 0 012 2v10a2 2 0 002 2H4a2 2 0 01-2-2V5zm3 1h6v4H5V6zm6 6H5v2h6v-2z" clip-rule="evenodd" />
              </svg>
              Microposts
            </router-link>
            
            <router-link 
              to="/users" 
              class="flex items-center px-4 py-2 text-sm font-medium rounded-lg transition-colors"
              :class="$route.path === '/users' ? 'bg-violet-500/10 text-violet-500' : 'text-gray-600 dark:text-gray-400 hover:text-gray-800 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/50'"
            >
              <svg class="w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20">
                <path d="M9 6a3 3 0 11-6 0 3 3 0 016 0zM17 6a3 3 0 11-6 0 3 3 0 016 0zM12.93 17c.046-.327.07-.66.07-1a6.97 6.97 0 00-1.5-4.33A5 5 0 0119 16v1h-6.07zM6 11a5 5 0 015 5v1H1v-1a5 5 0 015-5z" />
              </svg>
              Users
            </router-link>
          </nav>
        </div>
      </div>
      
      <!-- Main content -->
      <div class="flex-1 flex flex-col overflow-hidden">
        <main class="flex-1 overflow-y-auto">
          <div class="px-4 sm:px-6 lg:px-8 py-8 w-full max-w-[96rem] mx-auto">
            <router-view />
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
