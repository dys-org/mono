<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

import DarkModeToggle from '@/components/DarkModeToggle.vue';
</script>

<template>
  <div>
    <header class="flex items-center justify-between px-4 py-2 text-sm sm:px-6">
      <nav class="text-muted-foreground flex gap-4">
        <RouterLink to="/" active-class="text-foreground font-medium" class="hover:text-foreground">
          Todos
        </RouterLink>
        <RouterLink
          to="/done"
          active-class="text-foreground font-medium"
          class="hover:text-foreground"
        >
          Done
        </RouterLink>
      </nav>
      <DarkModeToggle />
    </header>
    <main class="container px-4 sm:px-6">
      <RouterView />
    </main>
  </div>
</template>
