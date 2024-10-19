<script setup lang="ts">
import { ref } from 'vue';

import Button from '@/components/ui/button/Button.vue';
import Checkbox from '@/components/ui/checkbox/Checkbox.vue';
// prettier-ignore
import { Dialog, DialogClose, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import Label from '@/components/ui/label/Label.vue';

const isModal = ref(true);
</script>

<template>
  <Dialog :modal="isModal">
    <div class="flex items-center gap-4">
      <DialogTrigger as-child>
        <Button variant="outline"> Show Dialog </Button>
      </DialogTrigger>
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
</template>
