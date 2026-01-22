<template>
  <div class="flex min-h-screen bg-gray-50">

    <!-- <aside class="w-64 bg-gray-800 text-white p-4 hidden md:block">
       <h2 class="text-xl font-bold mb-6">Stock Ratings</h2>
      
    </aside> -->

    <div class="flex-1 flex flex-col">
    <header class="bg-white border-b border-gray-200 px-6 py-4 flex justify-between items-center">
      <h1 class="text-lg font-semibold text-gray-900">
        Stock Ratings
      </h1>

      <nav class="flex gap-1">
        <router-link
          to="/"
          class="px-4 py-2 rounded-md text-gray-700
                hover:bg-gray-100
                transition"
          active-class="bg-gray-200 text-gray-900 font-medium"
        >
          Inicio
        </router-link>

        <router-link
          to="/list-stock-rating"
          class="px-4 py-2 rounded-md text-gray-700
                hover:bg-gray-100
                transition"
          active-class="bg-gray-200 text-gray-900 font-medium"
        >
          Listado de acciones
        </router-link>
      </nav>
    </header>
      <main class="p-6 flex-1 bg-gray-100">
        <slot />
      </main>
    <footer class="bg-gray-900 text-gray-300 text-center py-3 text-sm">
      © {{ currentYear }} Stock Ratings App. Todos los derechos reservados.
    </footer>
    </div>
  </div>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear()
</script>