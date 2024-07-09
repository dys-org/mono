<script setup lang="ts">
import { useQueryClient } from '@tanstack/vue-query';

import { useDeleteTodosId, useGetTodos, usePutTodosId } from '@/api';
import AddTodoVee from '@/components/AddTodoVee.vue';
import TodoItem from '@/components/TodoItem.vue';

const queryClient = useQueryClient();
const invalidateTodos = () => queryClient.invalidateQueries({ queryKey: ['todos'] });

const { isPending, error, data } = useGetTodos({ query: { queryKey: ['todos'] } });

// should we refetch the data here?
const deleteTodo = useDeleteTodosId({
  mutation: { onSuccess: invalidateTodos },
});
const updateTodo = usePutTodosId({
  mutation: { onSuccess: invalidateTodos },
});
</script>

<template>
  <div>
    <h1 class="mb-8 mt-4 text-4xl font-bold">Todos</h1>
    <AddTodoVee />
    <!-- Todo items -->
    <span v-if="isPending">Loading...</span>
    <span v-else-if="error">Error: {{ error.message }}</span>
    <div v-else class="mt-8 grid gap-4">
      <TodoItem
        v-for="todo in data?.data"
        :key="todo.id"
        :todo="todo"
        @update-done="updateTodo.mutate({ id: todo.id, data: { ...todo, done: $event } })"
        @save="updateTodo.mutate({ id: todo.id, data: { ...todo, text: $event } })"
        @delete="deleteTodo.mutate({ id: $event })"
      />
    </div>
  </div>
</template>
