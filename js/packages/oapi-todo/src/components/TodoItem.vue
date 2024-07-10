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
    class="border-gray-7 flex items-center gap-3 rounded-lg border border-solid bg-white p-4 dark:bg-transparent"
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
      <Button v-if="isEditing" variant="outline" class="w-16" @click="handleSave"> Save </Button>
      <Button v-else variant="outline" class="w-16" @click="isEditing = true"> Edit </Button>

      <Button v-if="isEditing" variant="link" class="-mr-2 w-16 px-2" @click="handleCancel">
        Cancel
      </Button>
      <Button
        v-else
        variant="link"
        class="-mr-2 w-16 px-2 text-muted-foreground hover:text-destructive hover:no-underline"
        @click="handleDelete"
      >
        Delete
      </Button>
    </div>
  </div>
</template>
