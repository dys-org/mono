<script setup lang="ts">
import { type HTMLAttributes, computed, ref, watch } from 'vue';

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
    width?: HTMLAttributes['class'];
  }>(),
  {
    multiple: false,
    placeholder: 'Search...',
    triggerText: 'Select ...',
    width: 'w-80',
  },
);

const model = defineModel<string | string[]>({ default: '' });

const open = ref(false);
const search = ref('');

watch(search, (newValue, oldValue) => {
  console.log(newValue);
  if (!newValue) console.log('FUCKK');
});

watch(open, (newValue, oldValue) => {
  if (!newValue) search.value = '';
});

const isMultiple = computed(() => Array.isArray(model.value));
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button
        variant="outline"
        role="combobox"
        :aria-expanded="open"
        :class="['justify-between', props.width]"
      >
        {{
          isMultiple && model.length
            ? props.options
                .filter((opt) => model.includes(opt.value))
                .map((opt) => opt.label)
                .join(', ')
            : typeof model === 'string' && model
              ? props.options.find((opt) => opt.value === model)?.label
              : props.triggerText
        }}

        <span class="i-lucide-chevrons-up-down ml-2 size-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent :class="['p-0', props.width]">
      <Command :multiple="isMultiple">
        <CommandInput v-model="search" :placeholder="props.placeholder" />
        <CommandEmpty>No framework found.</CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CommandItem
              v-for="opt in props.options"
              :key="opt.value"
              :value="opt.value"
              @select="
                (ev) => {
                  // this is used instead of v-model on Command
                  // to keep the search value from changing on select of a single item
                  if (isMultiple) {
                    model = model.includes(opt.value)
                      ? // @ts-expect-error
                        model.filter((v) => v !== opt.value)
                      : [...model, opt.value];
                    // }
                  } else {
                    if (typeof ev.detail.value === 'string') {
                      model = ev.detail.value;
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
