<script setup lang="ts">
import { useRouter } from "vue-router";

const router = useRouter();

function goToStocks() {
  router.push({ name: "stock-ratings" });
}
</script>

<template>
  <section class="max-w-6xl mx-auto px-6 py-20">
    <div class="space-y-6">
      <h1 class="text-4xl font-semibold text-gray-900">
        Stock Ratings Service
      </h1>

      <p class="text-lg text-gray-600 max-w-3xl">
        Plataforma para consultar recomendaciones de analistas financieros,
        precios objetivo y métricas clave de acciones del mercado.
        Los datos se sincronizan desde servicios externos, se procesan y se
        almacenan localmente para ofrecer consultas rápidas y confiables.
      </p>

      <div class="flex gap-4 pt-4">
        <button
          @click="goToStocks"
          class="px-6 py-3 rounded-md bg-gray-900 text-white hover:bg-gray-800 transition"
        >
          Ver acciones
        </button>

        <a
          href="https://github.com/jabella1/stock-ratings-service"
          target="_blank"
          class="px-6 py-3 rounded-md border border-gray-300 text-gray-700 hover:bg-gray-100 transition"
        >
          Ver repositorio
        </a>
      </div>
    </div>
  </section>
</template>