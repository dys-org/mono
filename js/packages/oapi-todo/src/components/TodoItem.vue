<script setup lang="ts">
import { Button, Checkbox, Input, Label } from 'dys-ui';
import { ref } from 'vue';

import type { Todo } from '@/model';

const props = defineProps<{
  todo: Todo;
}>();

const emit = defineEmits<{
  save: [text: string];
  'update-done': [done: boolean];
  delete: [id: number];
}>();

const isEditing = ref(false);
const innerText = ref(props.todo.text);

function handleSave() {
  emit('save', innerText.value);
  isEditing.value = false;
}
function handleCancel() {
  innerText.value = props.todo.text;
  isEditing.value = false;
}

function handleDelete() {
  if (confirm('Are you sure you want to delete this todo?')) {
    emit('delete', props.todo.id!);
  }
}
</script>

<template>
  <div
    class="flex items-center gap-3 rounded-lg bg-white p-4 shadow dark:border dark:border-solid dark:bg-transparent"
  >
    <Checkbox
      id="done"
      :checked="props.todo.done"
      @update:checked="emit('update-done', !props.todo.done)"
    />
    <Label for="done" class="sr-only">Done</Label>

    <Input
      id="text"
      v-model="innerText"
      :readonly="!isEditing"
      :class="[
        'read-only:border-transparent focus:read-only:ring-0',
        { 'line-through': props.todo.done },
      ]"
    />
    <Label for="text" class="sr-only">Description</Label>

    <div class="ml-auto flex gap-1.5">
      <Button v-if="isEditing" variant="outline" size="icon" title="Save" @click="handleSave">
        <span class="sr-only">Save</span>
        <span class="i-lucide-save size-4" aria-hidden="true" />
      </Button>
      <Button v-else variant="outline" size="icon" title="Edit" @click="isEditing = true">
        <span class="sr-only">Edit</span>
        <span class="i-lucide-edit size-4" aria-hidden="true" />
      </Button>

      <Button
        v-if="isEditing"
        variant="link"
        size="icon"
        class="text-muted-foreground hover:text-primary -mr-2"
        title="Cancel"
        @click="handleCancel"
      >
        <span class="sr-only">Cancel</span>
        <span class="i-lucide-x size-4" aria-hidden="true" />
      </Button>
      <Button
        v-else
        variant="link"
        size="icon"
        class="text-muted-foreground hover:text-destructive -mr-2"
        title="Delete"
        @click="handleDelete"
      >
        <span class="sr-only">Delete</span>
        <span class="i-lucide-trash size-4" aria-hidden="true" />
      </Button>
    </div>
  </div>
</template>
