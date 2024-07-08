<script setup lang="ts">
import { DarkModeToggle } from 'dys-ui';
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
  <div>
    <header class="flex items-center justify-between px-4 py-2 text-sm sm:px-6">
      <nav class="flex gap-4 text-muted-foreground">
        <RouterLink to="/" active-class="text-foreground">Todo</RouterLink>
        <RouterLink to="/about" active-class="text-foreground">Other</RouterLink>
      </nav>
      <DarkModeToggle />
    </header>
    <main class="container px-4 sm:px-6">
      <RouterView />
    </main>
  </div>
</template>
