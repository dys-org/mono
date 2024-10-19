<script setup lang="ts">
import { type HTMLAttributes, computed, ref } from 'vue';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from '@/components/ui/command';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
import { cn } from '@/lib/utils';

const props = withDefaults(
  defineProps<{
    placeholder?: string;
    triggerText?: string;
    options: { value: string; label: string }[];
    /** tailwindcss width class */
    width?: HTMLAttributes['class'];
  }>(),
  {
    placeholder: 'Search...',
    triggerText: 'Select...',
    width: 'w-full',
  },
);

const model = defineModel<string | string[]>({ default: '' });

const open = ref(false);
const search = ref('');

const isMultiple = computed(() => Array.isArray(model.value));

function toggleAll() {
  if (model.value.length === props.options.length) {
    model.value = [];
  } else {
    const allValues = props.options.map((option) => option.value);
    model.value = allValues;
  }
}
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button
        variant="outline"
        role="combobox"
        :aria-expanded="open"
        :class="['h-auto min-h-10 justify-between hover:bg-transparent', props.width]"
      >
        <span v-if="isMultiple && model.length" class="flex flex-wrap gap-1.5">
          <Badge
            v-for="opt in props.options.filter((opt) => model.includes(opt.value))"
            :key="opt.value"
            variant="secondary"
            >{{ opt.value }}</Badge
          >
        </span>

        <template v-else>
          {{
            typeof model === 'string' && model
              ? props.options.find((opt) => opt.value === model)?.label
              : props.triggerText
          }}
        </template>

        <span class="i-lucide-chevrons-up-down ml-2 size-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent :class="['p-0', props.width]">
      <Command :multiple="isMultiple">
        <CommandInput v-model="search" :placeholder="props.placeholder" />
        <CommandEmpty>Nothing found.</CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CommandItem v-if="isMultiple" value="Select all" @select="toggleAll">
              <span
                :class="
                  cn(
                    'i-lucide-check mr-2 size-4',
                    model.length === props.options.length ? 'opacity-100' : 'opacity-0',
                  )
                "
              />
              Select all
            </CommandItem>
            <CommandItem
              v-for="opt in props.options"
              :key="opt.value"
              :value="opt.value"
              @select="
                (e: any) => {
                  // this is used instead of v-model on Command
                  // to keep the search value from changing on select
                  if (Array.isArray(model)) {
                    // same check as isMultiple
                    model = model.includes(opt.value)
                      ? model.filter((v: string) => v !== opt.value)
                      : [...model, opt.value];
                  } else {
                    if (typeof e.detail.value === 'string') {
                      model = e.detail.value;
                    }
                    open = false;
                  }
                }
              "
            >
              <span
                :class="
                  cn(
                    'i-lucide-check mr-2 size-4',
                    // prettier-ignore
                    isMultiple && model.includes(opt.value) ? 'opacity-100' :
                    model === opt.value ? 'opacity-100' :
                    'opacity-0',
                  )
                "
              />
              {{ opt.label }}
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </PopoverContent>
  </Popover>
</template>
