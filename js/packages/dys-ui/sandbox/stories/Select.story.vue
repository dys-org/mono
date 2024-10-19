<script setup lang="ts">
// prettier-ignore
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';
</script>

<template>
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
</template>
