<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import { DarkModeToggle } from '@/components/dark-mode-toggle';
</script>

<template>
  <header class="flex items-center justify-between px-4 py-2 text-sm sm:px-6">
    <nav class="flex gap-4 text-muted-foreground">
      <RouterLink to="/" active-class="text-foreground">Todo</RouterLink>
      <RouterLink to="/about" active-class="text-foreground">Playground</RouterLink>
    </nav>
    <DarkModeToggle />
  </header>

  <main class="container px-4 sm:px-6">
    <RouterView />
  </main>
</template>
