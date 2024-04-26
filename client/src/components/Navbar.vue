
<template>
    <nav class="bg-gray-700">
        <div class="mx-auto max-w-9xl px-0 sm:px-0 lg:px-0">
            <div class="flex items-center justify-between">
                <router-link to="/">
                    <h1 class="text-white bg-black px-3 py-2 text-xl font-bold border-radius-xl shadow-lg transition-all duration-300 hover:opacity-75">
                        <span class="text-green-400">WP</span>
                        <span class="text-blue-400">Manager</span>
                    </h1>
                </router-link>
                

                <div class="hidden sm:block">
            
              
                    <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Dashboard</a>
                    <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Team</a>
                    <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Projects</a>
                    <a href="#" class="text-gray-300 hover:bg-gray-700 hover:text-white rounded-md px-3 py-2 text-sm font-medium">Calendar</a>
                </div>
            </div>
        </div>
  
    </nav>
</template>


<script>
export default {
    name: 'Navbar'
}
</script>

<style>
/* Add your custom styles here */
</style>
