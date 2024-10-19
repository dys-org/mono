<script setup lang="ts">
import { Button } from '@/components/ui/button';
</script>

<template>
  <div class="grid gap-y-4">
    <h3 class="-mb-1 mt-2 font-semibold">default</h3>
    <div class="flex gap-4">
      <Button>Click Me</Button>
      <Button variant="destructive">Destructive</Button>
      <Button variant="outline">Outline</Button>
      <Button variant="secondary">Secondary</Button>
      <Button variant="ghost">Ghost</Button>
      <Button variant="link">Link</Button>
    </div>
    <h3 class="-mb-1 mt-2 font-semibold">sm</h3>
    <div class="flex gap-4">
      <Button size="sm">Click Me</Button>
      <Button size="sm" variant="destructive">Destructive</Button>
      <Button size="sm" variant="outline">Outline</Button>
      <Button size="sm" variant="secondary">Secondary</Button>
      <Button size="sm" variant="ghost">Ghost</Button>
      <Button size="sm" variant="link">Link</Button>
    </div>
    <h3 class="-mb-1 mt-2 font-semibold">lg</h3>
    <div class="flex gap-4">
      <Button size="lg">Click Me</Button>
      <Button size="lg" variant="destructive">Destructive</Button>
      <Button size="lg" variant="outline">Outline</Button>
      <Button size="lg" variant="secondary">Secondary</Button>
      <Button size="lg" variant="ghost">Ghost</Button>
      <Button size="lg" variant="link">Link</Button>
    </div>
    <h3 class="-mb-1 mt-2 font-semibold">icon</h3>
    <div class="flex gap-4">
      <Button size="icon"><span class="i-lucide-alarm-clock size-5" /></Button>
      <Button size="icon" variant="destructive">
        <span class="i-lucide-alarm-clock size-5" />
      </Button>
      <Button size="icon" variant="outline">
        <span class="i-lucide-alarm-clock size-5" />
      </Button>
      <Button size="icon" variant="secondary">
        <span class="i-lucide-alarm-clock size-5" />
      </Button>
      <Button size="icon" variant="ghost"><span class="i-lucide-alarm-clock size-5" /></Button>
      <Button size="icon" variant="link"><span class="i-lucide-alarm-clock size-5" /></Button>
    </div>
  </div>
</template>
