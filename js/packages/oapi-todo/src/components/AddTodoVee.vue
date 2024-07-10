<script setup lang="ts">
import { useQueryClient } from '@tanstack/vue-query';
import { toTypedSchema } from '@vee-validate/zod';
// prettier-ignore
import { Button, FormControl, FormField, FormItem, FormMessage, Input } from 'dys-ui';
import { useForm } from 'vee-validate';
import * as z from 'zod';

import { usePostTodos } from '@/api';
import { TODOS_KEY } from '@/lib/constants';

const queryClient = useQueryClient();

const formSchema = toTypedSchema(
  z.object({
    todo: z.string().min(1, { message: 'Field is required' }),
  }),
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const createTodo = usePostTodos({
  mutation: { onSuccess: () => queryClient.invalidateQueries({ queryKey: [TODOS_KEY] }) },
});

const onSubmit = handleSubmit((values, { resetForm }) => {
  // console.log(values);
  createTodo.mutate({ data: { text: values.todo, done: false } });
  resetForm();
});
</script>

<template>
  <form class="space-y-6" @submit="onSubmit">
    <div class="flex min-h-[68px] gap-1.5">
      <FormField v-slot="{ componentField }" name="todo">
        <FormItem class="w-full">
          <FormControl>
            <Input type="text" v-bind="componentField" aria-label="Todo Description" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <Button type="submit" variant="default"> Add </Button>
    </div>
  </form>
</template>
