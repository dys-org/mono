<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
import { useForm } from 'vee-validate';
import { ref } from 'vue';
import * as z from 'zod';

import { Combobox } from '@/components/combobox';
// prettier-ignore
import { AlertDialog, AlertDialogAction, AlertDialogCancel, AlertDialogContent, AlertDialogDescription, AlertDialogFooter, AlertDialogHeader, AlertDialogTitle, AlertDialogTrigger } from '@/components/ui/alert-dialog';
import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
// prettier-ignore
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
// prettier-ignore
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuLabel, DropdownMenuPortal, DropdownMenuSeparator, DropdownMenuShortcut, DropdownMenuSub, DropdownMenuSubContent, DropdownMenuSubTrigger, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
// prettier-ignore
import { FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
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
const isModal = ref(true);
const selectedMultiple = ref([]);

const formSchema = toTypedSchema(
  z.object({
    username: z.string().min(2).max(50),
  }),
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
  alert(JSON.stringify(values));
});
</script>

<template>
  <main class="mb-96">
    <h1 class="mb-8 text-base font-semibold uppercase">Showcase</h1>
    <div class="grid gap-12">
      <!-- BUTTON -->
      <div class="grid gap-y-4">
        <h2 class="mb-4 text-2xl font-semibold">Button</h2>

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
        <Combobox v-model="selected" :options="frameworkOptions" width="w-80" />
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
            <Button variant="outline"> Show Alert Dialog </Button>
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

      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Dialog</h2>

        <Dialog :modal="isModal">
          <div class="flex items-center gap-4">
            <DialogTrigger as-child>
              <Button variant="outline"> Show Dialog </Button>
            </DialogTrigger>
            <!-- <div class="flex items-center space-x-2">
              <input
                id="isModal"
                v-model="isModal"
                type="checkbox"
                class="ring-offset-background focus-visible:ring-ring text-primary hover:text-primary size-4 shrink-0 rounded-md accent-current focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              />
              <Label for="isModal"> Modal </Label>
            </div> -->
            <div class="flex items-center space-x-2">
              <Checkbox id="terms" :checked="isModal" @update:checked="isModal = !isModal" />
              <Label for="terms"> Modal </Label>
            </div>
          </div>
          <DialogContent>
            <DialogHeader>
              <DialogTitle>Edit profile</DialogTitle>
              <DialogDescription>
                Make changes to your profile here. Click save when you're done.
              </DialogDescription>
            </DialogHeader>
            <DialogFooter class="sm:justify-start">
              <DialogClose as-child>
                <Button type="button" variant="secondary"> Close </Button>
              </DialogClose>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>

      <div class="max-w-96">
        <h2 class="mb-4 text-2xl font-semibold">Form</h2>
        <form class="w-2/3 space-y-6" @submit="onSubmit">
          <FormField v-slot="{ componentField }" name="username">
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input type="text" placeholder="dys" v-bind="componentField" />
              </FormControl>
              <FormDescription> This is your public display name. </FormDescription>
              <FormMessage />
            </FormItem>
          </FormField>
          <Button type="submit"> Submit </Button>
        </form>
      </div>
    </div>
  </main>
</template>
