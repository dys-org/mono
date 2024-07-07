<script lang="ts" setup>
import { CalendarPrev, type CalendarPrevProps, useForwardProps } from 'radix-vue';
import { type HTMLAttributes, computed } from 'vue';

import { cn } from '@/lib/utils';

const props = defineProps<CalendarPrevProps & { class?: HTMLAttributes['class'] }>();

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props;

  return delegated;
});

const forwardedProps = useForwardProps(delegatedProps);
</script>

<template>
  <CalendarPrev
    :class="
      cn(
        'border border-input hover:bg-accent hover:text-accent-foreground', // variant outline
        'size-7 bg-transparent p-0 opacity-50 hover:opacity-100',
        props.class,
      )
    "
    v-bind="forwardedProps"
  >
    <slot>
      <span class="i-lucide-chevron-left size-4" />
    </slot>
  </CalendarPrev>
</template>
