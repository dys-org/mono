<script setup lang="ts">
import { ref } from 'vue';

import { DarkModeToggle } from '@/components/dark-mode-toggle';
import { Button } from '@/components/ui/button';
// prettier-ignore
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from '@/components/ui/command';
// prettier-ignore
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuLabel, DropdownMenuPortal, DropdownMenuSeparator, DropdownMenuShortcut, DropdownMenuSub, DropdownMenuSubContent, DropdownMenuSubTrigger, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
// prettier-ignore
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';
import { cn } from '@/lib/utils';

const frameworkOptions = [
  { value: 'next.js', label: 'Next.js' },
  { value: 'sveltekit', label: 'SvelteKit' },
  { value: 'nuxt', label: 'Nuxt' },
  { value: 'remix', label: 'Remix' },
  { value: 'astro', label: 'Astro' },
];

const open1 = ref(false);
const framework = ref('');

const open2 = ref(false);
const selectedFrameworks = ref(['next.js', 'sveltekit']);
</script>

<template>
  <main class="container p-6">
    <h1 class="mb-8 text-base font-semibold uppercase">Showcase</h1>
    <div class="grid gap-12">
      <!-- DARK MODE TOGGLE -->
      <div>
        <h2 class="mb-4 text-2xl font-semibold">Dark Mode Toggle</h2>
        <DarkModeToggle />
      </div>

      <!-- BUTTON -->
      <div class="grid gap-y-4">
        <h2 class="mb-4 text-2xl font-semibold">Button</h2>
        <div class="flex gap-4">
          <Button>Click Me</Button>
          <Button variant="destructive">Click Me</Button>
          <Button variant="outline">Click Me</Button>
          <Button variant="secondary">Click Me</Button>
          <Button variant="ghost">Click Me</Button>
          <Button variant="link">Click Me</Button>
        </div>
        <div class="flex gap-4">
          <Button size="sm">Click Me</Button>
          <Button size="sm" variant="destructive">Click Me</Button>
          <Button size="sm" variant="outline">Click Me</Button>
          <Button size="sm" variant="secondary">Click Me</Button>
          <Button size="sm" variant="ghost">Click Me</Button>
          <Button size="sm" variant="link">Click Me</Button>
        </div>
        <div class="flex gap-4">
          <Button size="lg">Click Me</Button>
          <Button size="lg" variant="destructive">Click Me</Button>
          <Button size="lg" variant="outline">Click Me</Button>
          <Button size="lg" variant="secondary">Click Me</Button>
          <Button size="lg" variant="ghost">Click Me</Button>
          <Button size="lg" variant="link">Click Me</Button>
        </div>
        <div class="flex gap-4">
          <Button size="icon"><span class="i-lucide-alarm-clock size-5" /></Button>
          <Button size="icon" variant="destructive"
            ><span class="i-lucide-alarm-clock size-5"
          /></Button>
          <Button size="icon" variant="outline"
            ><span class="i-lucide-alarm-clock size-5"
          /></Button>
          <Button size="icon" variant="secondary"
            ><span class="i-lucide-alarm-clock size-5"
          /></Button>
          <Button size="icon" variant="ghost"><span class="i-lucide-alarm-clock size-5" /></Button>
          <Button size="icon" variant="link"><span class="i-lucide-alarm-clock size-5" /></Button>
        </div>
      </div>

      <!-- DROPDOWN MENU -->
      <div>
        <h2 class="mb-4 text-2xl font-semibold">DropdownMenu</h2>
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button variant="outline"> Open </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent class="w-56">
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DropdownMenuItem>
                <span class="i-lucide-user mr-2 size-4" />
                <span>Profile</span>
                <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
              </DropdownMenuItem>
              <DropdownMenuItem>
                <span class="i-lucide-credit-card mr-2 size-4" />
                <span>Billing</span>
                <DropdownMenuShortcut>⌘B</DropdownMenuShortcut>
              </DropdownMenuItem>
              <DropdownMenuItem>
                <span class="i-lucide-settings mr-2 size-4" />
                <span>Settings</span>
                <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
              </DropdownMenuItem>
              <DropdownMenuItem>
                <span class="i-lucide-keyboard mr-2 size-4" />
                <span>Keyboard shortcuts</span>
                <DropdownMenuShortcut>⌘K</DropdownMenuShortcut>
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DropdownMenuItem>
                <span class="i-lucide-users mr-2 size-4" />
                <span>Team</span>
              </DropdownMenuItem>
              <DropdownMenuSub>
                <DropdownMenuSubTrigger>
                  <span class="i-lucide-user-plus mr-2 size-4" />
                  <span>Invite users</span>
                </DropdownMenuSubTrigger>
                <DropdownMenuPortal>
                  <DropdownMenuSubContent>
                    <DropdownMenuItem>
                      <span class="i-lucide-mail mr-2 size-4" />
                      <span>Email</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem>
                      <span class="i-lucide-message-square mr-2 size-4" />
                      <span>Message</span>
                    </DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem>
                      <span class="i-lucide-plus-circle mr-2 size-4" />
                      <span>More...</span>
                    </DropdownMenuItem>
                  </DropdownMenuSubContent>
                </DropdownMenuPortal>
              </DropdownMenuSub>
              <DropdownMenuItem>
                <span class="i-plus mr-2 size-4" />
                <span>New Team</span>
                <DropdownMenuShortcut>⌘+T</DropdownMenuShortcut>
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem>
              <span class="i-lucide-github mr-2 size-4" />
              <span>GitHub</span>
            </DropdownMenuItem>
            <DropdownMenuItem>
              <span class="i-lucide-life-buoy mr-2 size-4" />
              <span>Support</span>
            </DropdownMenuItem>
            <DropdownMenuItem disabled>
              <span class="i-lucide-cloud mr-2 size-4" />
              <span>API</span>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>
              <span class="i-lucide-log-out mr-2 size-4" />
              <span>Log out</span>
              <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <!-- SELECT -->
      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Select</h2>
        <Select>
          <SelectTrigger>
            <SelectValue placeholder="Select a fruit" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel>Fruits</SelectLabel>
              <SelectItem value="apple"> Apple </SelectItem>
              <SelectItem value="banana"> Banana </SelectItem>
              <SelectItem value="peach"> Peach </SelectItem>
              <SelectItem value="blueberry"> Blueberry </SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <!-- COMBOBOX -->
      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Combobox</h2>
        <Popover v-model:open="open1">
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              role="combobox"
              :aria-expanded="open1"
              class="w-[200px] justify-between"
            >
              {{
                framework
                  ? frameworkOptions.find((fw) => fw.value === framework)?.label
                  : 'Select framework...'
              }}

              <span class="i-lucide-chevrons-up-down ml-2 size-4 shrink-0 opacity-50" />
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-[200px] p-0">
            <Command v-model="framework">
              <CommandInput placeholder="Search framework..." />
              <CommandEmpty>No framework found.</CommandEmpty>
              <CommandList>
                <CommandGroup>
                  <CommandItem
                    v-for="fw in frameworkOptions"
                    :key="fw.value"
                    :value="fw.value"
                    @select="
                      (ev) => {
                        if (typeof ev.detail.value === 'string') {
                          framework = ev.detail.value;
                        }
                        open1 = false;
                      }
                    "
                  >
                    <span
                      :class="
                        cn(
                          'i-lucide-check mr-2 size-4',
                          framework === fw.value ? 'opacity-100' : 'opacity-0',
                        )
                      "
                    />
                    {{ fw.label }}
                  </CommandItem>
                </CommandGroup>
              </CommandList>
            </Command>
          </PopoverContent>
        </Popover>
      </div>

      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Combobox Multiple</h2>
        <Popover v-model:open="open2">
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              role="combobox"
              :aria-expanded="open2"
              class="w-full justify-between"
            >
              {{
                selectedFrameworks.length
                  ? frameworkOptions
                      .filter((fw) => selectedFrameworks.includes(fw.value))
                      .map((fw) => fw.label)
                      .join(', ')
                  : 'Select framework...'
              }}

              <span class="i-lucide-chevrons-up-down ml-2 size-4 shrink-0 opacity-50" />
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-[200px] p-0">
            <Command v-model="selectedFrameworks" multiple>
              <CommandInput placeholder="Search framework..." />
              <CommandEmpty>No framework found.</CommandEmpty>
              <CommandList>
                <CommandGroup>
                  <CommandItem v-for="fw in frameworkOptions" :key="fw.value" :value="fw.value">
                    <span
                      :class="
                        cn(
                          'i-lucide-check mr-2 size-4',
                          selectedFrameworks.includes(fw.value) ? 'opacity-100' : 'opacity-0',
                        )
                      "
                    />
                    {{ fw.label }}
                  </CommandItem>
                </CommandGroup>
              </CommandList>
            </Command>
          </PopoverContent>
        </Popover>
      </div>
    </div>
  </main>
</template>
