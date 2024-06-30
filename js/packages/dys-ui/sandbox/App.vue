<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
  <header class="w-full px-6 py-4">
    <nav class="flex gap-4 text-sm">
      <RouterLink to="/">Showcase</RouterLink>
      <RouterLink to="/about">About</RouterLink>
    </nav>
  </header>

  <RouterView />
</template>
