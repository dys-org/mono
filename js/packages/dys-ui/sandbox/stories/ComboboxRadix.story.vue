<script setup lang="ts">
// prettier-ignore
import { ComboboxAnchor, ComboboxContent, ComboboxEmpty, ComboboxInput, ComboboxItem, ComboboxItemIndicator, ComboboxRoot, ComboboxTrigger } from 'radix-vue';
import { computed, ref } from 'vue';

import { Label } from '../../src/components/ui/label';

const frameworkOptions = [
  { value: 'next.js', label: 'Next.js' },
  { value: 'sveltekit', label: 'SvelteKit' },
  { value: 'nuxt', label: 'Nuxt' },
  { value: 'remix', label: 'Remix' },
  { value: 'astro', label: 'Astro' },
];

export interface ComboboxOption {
  value: string;
  label: string;
}

const selected = ref<ComboboxOption>();
const open = ref(false);
const searchTerm = ref('');

const filteredItems = computed(() => {
  return frameworkOptions.filter((item) =>
    item.label.toLowerCase().includes(searchTerm.value.toLowerCase()),
  );
});

const addNewItem = () => {
  if (searchTerm.value && !frameworkOptions.some((item) => item.label === searchTerm.value)) {
    const newItem: ComboboxOption = {
      value: searchTerm.value.toLowerCase(),
      label: searchTerm.value,
    };
    frameworkOptions.push(newItem);
    selected.value = newItem;
    searchTerm.value = '';
  }
};
</script>

<template>
  <div class="mx-auto w-full max-w-md">
    <ComboboxRoot
      v-model="selected"
      v-model:open="open"
      v-model:search-term="searchTerm"
      :display-value="(item) => item.label"
      class="relative"
    >
      <Label for="framework" class="mb-2 block">Framework</Label>
      <ComboboxAnchor class="relative">
        <ComboboxInput
          id="framework"
          class="focus:border-primary focus:ring-primary border-border bg-background w-full rounded-md border py-2 pl-3 pr-10 shadow-sm focus:outline-none focus:ring-1 sm:text-sm"
          placeholder="Choose a framework"
          @focus="open = true"
        />
        <ComboboxTrigger class="absolute inset-y-0 right-0 flex items-center pr-2">
          <svg class="size-5 text-gray-400" viewBox="0 0 20 20" fill="none" stroke="currentColor">
            <path
              d="M7 7l3-3 3 3m0 6l-3 3-3-3"
              strokeWidth="1.5"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </ComboboxTrigger>
      </ComboboxAnchor>
      <ComboboxContent
        class="bg-background border-border absolute z-10 mt-1.5 max-h-60 w-full overflow-auto rounded-md border py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm"
      >
        <ComboboxEmpty class="relative cursor-default select-none px-4 py-2 text-gray-700">
          <span>No results found</span>
          <button
            class="bg-primary hover:bg-primary/90 mt-2 w-full truncate rounded-md px-3 py-2 text-sm font-semibold text-white shadow-sm focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
            @click="addNewItem"
          >
            Add "{{ searchTerm }}"
          </button>
        </ComboboxEmpty>
        <ComboboxItem
          v-for="item in filteredItems"
          :key="item.value"
          :value="item"
          class="data-[highlighted]:bg-muted text-foreground relative cursor-default select-none py-2 pl-3 pr-9"
        >
          <span :class="['block truncate']">
            {{ item.label }}
          </span>
          <ComboboxItemIndicator class="absolute inset-y-0 right-0 flex items-center pr-4">
            <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path
                fillRule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clipRule="evenodd"
              />
            </svg>
          </ComboboxItemIndicator>
        </ComboboxItem>
      </ComboboxContent>
    </ComboboxRoot>
    <p class="mt-4 text-sm text-gray-600">Selected framework: {{ selected ?? 'None' }}</p>
  </div>
</template>
