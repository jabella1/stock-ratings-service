<template>
  <div class="flex min-h-screen bg-gray-50">

    <!-- <aside class="w-64 bg-gray-800 text-white p-4 hidden md:block">
       <h2 class="text-xl font-bold mb-6">Stock Ratings</h2>
      
    </aside> -->

    <div class="flex-1 flex flex-col">
      <header class="bg-white shadow p-4 flex justify-between items-center">
        <h1 class="text-lg font-semibold">Stock Ratings</h1>
        <nav class="space-y-2 flex">
        <router-link
          to="/"
          class="block px-4 py-2 rounded hover:bg-gray-500"
          active-class="bg-gray-500"
        >
          Inicio
        </router-link>
        <router-link
          to="/list-stock-rating"
          class="block px-4 py-2 rounded hover:bg-gray-500"
          active-class="bg-gray-500"
        >
          Listado de acciones
        </router-link>
      </nav>
      </header>
      <main class="p-6 flex-1 bg-gray-100">
        <slot />
      </main>
      <footer class="bg-gray-800 text-white text-center py-3 text-sm">
      © {{ currentYear }} Stock Ratings App. Todos los derechos reservados.
    </footer>
    </div>
  </div>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear()
</script>