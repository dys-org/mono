<script setup lang="ts">
import { ref } from 'vue';

import AlertDialogStory from '../stories/AlertDialog.story.vue';
import ButtonStory from '../stories/Button.story.vue';
import ComboboxRadixStory from '../stories/ComboboxRadix.story.vue';
import DialogStory from '../stories/Dialog.story.vue';
import DropdownStory from '../stories/Dropdown.story.vue';
import FormStory from '../stories/Form.story.vue';
import SelectStory from '../stories/Select.story.vue';

const selectedMultiple = ref([]);
</script>

<template>
  <main class="mb-96">
    <h1 class="mb-8 text-base font-semibold uppercase">Showcase</h1>
    <div class="grid gap-12">
      <!-- BUTTON -->
      <div>
        <h2 class="">Button</h2>
        <ButtonStory />
      </div>

      <!-- DROPDOWN MENU -->
      <div>
        <h2>DropdownMenu</h2>
        <DropdownStory />
      </div>

      <!-- SELECT -->
      <div class="max-w-96">
        <h2>Select</h2>
        <SelectStory />
      </div>

      <!-- COMBOBOX -->
      <div class="max-w-96">
        <h2>Combobox</h2>
        <ComboboxRadixStory />
      </div>

      <!-- <div class="max-w-96">
        <h2 >Combobox Multiple</h2>
        <Combobox v-model="selectedMultiple" :options="frameworkOptions" />
      </div> -->

      <!-- CONFIRM DIALOG -->
      <div class="max-w-96">
        <h2>Alert Dialog</h2>
        <AlertDialogStory />
      </div>

      <!-- DIALOG -->
      <div class="max-w-96">
        <h2>Dialog</h2>
        <DialogStory />
      </div>

      <!-- FORM -->
      <div class="max-w-96">
        <h2>Form</h2>
        <FormStory />
      </div>
    </div>
  </main>
</template>

<style>
h2 {
  @apply mb-4 text-2xl font-semibold;
}
</style>
