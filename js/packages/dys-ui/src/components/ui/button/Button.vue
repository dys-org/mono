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
        'inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50',

        // VARIANT
        'default' === props.variant && 'bg-primary text-primary-foreground hover:bg-primary/90',
        'destructive' === props.variant &&
          'bg-destructive text-destructive-foreground hover:bg-destructive/90',
        'outline' === props.variant &&
          'border border-input bg-background hover:bg-accent hover:text-accent-foreground',
        'secondary' === props.variant &&
          'bg-secondary text-secondary-foreground hover:bg-secondary/80',
        'ghost' === props.variant && 'hover:bg-accent hover:text-accent-foreground',
        'link' === props.variant && 'text-primary underline-offset-4 hover:underline',

        // SIZE
        'default' === props.size && 'h-10 px-4 py-2',
        'sm' === props.size && 'h-9 rounded-md px-3',
        'lg' === props.size && 'h-11 rounded-md px-8',
        'icon' === props.size && 'h-10 w-10',

        props.class,
      )
    "
  >
    <slot />
  </component>
</template>
