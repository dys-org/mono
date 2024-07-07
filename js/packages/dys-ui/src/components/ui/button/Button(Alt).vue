<script setup lang="ts">
import type { HTMLAttributes } from 'vue';

import { cn } from '@/lib/utils';

interface Props {
  as?: string;
  variant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link';
  size?: 'default' | 'sm' | 'lg' | 'icon' | null;
  class?: HTMLAttributes['class'];
}

const props = withDefaults(defineProps<Props>(), {
  as: 'button',
  variant: 'default',
  size: 'default',
  class: '',
});
</script>

<template>
  <component
    :is="props.as"
    :class="
      cn(
        'btn',
        // VARIANT
        'destructive' === props.variant && 'btn-destructive',
        'outline' === props.variant && 'btn-outline',
        'secondary' === props.variant && 'btn-secondary',
        'ghost' === props.variant && 'btn-ghost',
        'link' === props.variant && 'btn-link',
        // SIZE
        'sm' === props.size && 'btn-sm',
        'lg' === props.size && 'btn-lg',
        'icon' === props.size && 'btn-icon',

        props.class,
      )
    "
  >
    <slot />
  </component>
</template>

<style>
.btn {
  @apply inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50;
  & {
    @apply h-10 bg-primary px-4 py-2 text-primary-foreground hover:bg-primary/90;
  }
  &-destructive {
    @apply bg-destructive text-destructive-foreground hover:bg-destructive/90;
  }
  &-outline {
    @apply border border-input bg-background text-foreground hover:bg-accent hover:text-accent-foreground;
  }
  &-secondary {
    @apply bg-secondary text-secondary-foreground hover:bg-secondary/80;
  }
  &-ghost {
    @apply bg-transparent text-foreground hover:bg-accent hover:text-accent-foreground;
  }
  &-link {
    @apply bg-transparent text-primary underline-offset-4 hover:bg-transparent hover:underline;
  }
  &-sm {
    @apply h-9 rounded-md px-3;
  }
  &-lg {
    @apply h-11 rounded-md px-8;
  }
  &-icon {
    @apply h-10 w-10 p-0;
  }
}
</style>
