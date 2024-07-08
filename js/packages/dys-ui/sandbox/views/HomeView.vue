<script setup lang="ts">
import { ref } from 'vue';

import { Combobox } from '@/components/combobox';
import { DarkModeToggle } from '@/components/dark-mode-toggle';
// prettier-ignore
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from '@/components/ui/alert-dialog';
import { Button } from '@/components/ui/button';
// prettier-ignore
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuLabel, DropdownMenuPortal, DropdownMenuSeparator, DropdownMenuShortcut, DropdownMenuSub, DropdownMenuSubContent, DropdownMenuSubTrigger, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
// prettier-ignore
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';

const frameworkOptions = [
  { value: 'next.js', label: 'Next.js' },
  { value: 'sveltekit', label: 'SvelteKit' },
  { value: 'nuxt', label: 'Nuxt' },
  { value: 'remix', label: 'Remix' },
  { value: 'astro', label: 'Astro' },
];

const selected = ref('');
const selectedMultiple = ref([]);
</script>

<template>
  <main class="container mb-96 p-6">
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
        <Combobox v-model="selected" :options="frameworkOptions" width="w-96" />
      </div>

      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Combobox Multiple</h2>
        <Combobox v-model="selectedMultiple" :options="frameworkOptions" />
      </div>
      <!-- CONFIRM DIALOG -->
      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Alert Dialog</h2>

        <AlertDialog>
          <AlertDialogTrigger as-child>
            <Button variant="outline"> Show Dialog </Button>
          </AlertDialogTrigger>
          <AlertDialogContent>
            <AlertDialogHeader>
              <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
              <AlertDialogDescription>
                This action cannot be undone. This will permanently delete your account and remove
                your data from our servers.
              </AlertDialogDescription>
            </AlertDialogHeader>
            <AlertDialogFooter>
              <AlertDialogCancel>Cancel</AlertDialogCancel>
              <AlertDialogAction>Continue</AlertDialogAction>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialog>
      </div>
    </div>
  </main>
</template>
