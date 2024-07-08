<script setup lang="ts">
import { toTypedSchema } from '@vee-validate/zod';
// prettier-ignore
import { Button, FormControl, FormField, FormItem, FormMessage, Input } from 'dys-ui';
import { useForm } from 'vee-validate';
import * as z from 'zod';

const formSchema = toTypedSchema(
  z.object({
    todo: z.string().min(1),
  }),
);

const { handleSubmit } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit((values) => {
  console.log(values);
});
</script>

<template>
  <form class="space-y-6" @submit="onSubmit">
    <div class="flex gap-1.5">
      <FormField v-slot="{ componentField }" name="todo">
        <FormItem class="w-full">
          <FormControl>
            <Input type="text" v-bind="componentField" aria-label="Todo Description" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <Button type="submit"> Add </Button>
    </div>
  </form>
</template>
